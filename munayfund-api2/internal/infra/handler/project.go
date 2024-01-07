package handler

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"munayfund-api2/internal/core/domain"
	"munayfund-api2/internal/core/service"
	fileinfra "munayfund-api2/internal/infra/file"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

const UploadTempDir = "./uploads/temp/"

type ProjectHandler struct {
	projectService *service.ProjectService
}

func NewProjectHandler(projectService *service.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

// @Summary Crear un nuevo proyecto
// @Description Crea un nuevo proyecto con información multimedia.
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags projects
// @ID create-project
// @Accept mpfd
// @Produce json
// @Param userId formData string true "ID del usuario creador"
// @Param title formData string true "Título del proyecto"
// @Param description formData string true "Descripción del proyecto"
// @Param media formData file false "Archivos multimedia (imágenes y videos)" format=multi
// @Success 200 {object} gin.H "Operación exitosa"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router /projects/create [post]
func (h *ProjectHandler) CreateProject(c *gin.Context) {
	if c.Request.Method != http.MethodPost {
		httputil.NewError(c, http.StatusMethodNotAllowed, errors.New("Method not allowed"))
		return
	}

	// Parsear el formulario
	err := c.Request.ParseMultipartForm(10 << 20) // 10 MB límite de tamaño del formulario
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, errors.New("Error parsing the form"))
		return
	}

	// Obtener los archivos del formulario
	formFiles := c.Request.MultipartForm.File["media"]

	// Crear un directorio temporal para almacenar los archivos
	if err := os.MkdirAll(UploadTempDir, os.ModePerm); err != nil {
		httputil.NewError(c, http.StatusInternalServerError, errors.New("Error creating tmp directory"))
		return
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var uploadedMedia []domain.Multimedia
	errorsCh := make(chan error, len(formFiles))

	for _, fileHeader := range formFiles {
		wg.Add(1)
		go func(fileHeader *multipart.FileHeader) {
			defer wg.Done()

			// Abrir el archivo
			file, err := fileHeader.Open()
			if err != nil {
				errorsCh <- err
				return
			}
			defer file.Close()

			// Crear un archivo en el servidor para almacenar el multimedia
			filename := filepath.Join(UploadTempDir, fileHeader.Filename)
			dst, err := os.Create(filename)
			if err != nil {
				errorsCh <- err
				return
			}
			defer dst.Close()

			// Copiar el contenido del archivo recibido al archivo en el servidor
			_, err = io.Copy(dst, file)
			if err != nil {
				errorsCh <- err
				return
			}

			// Subir el multimedia a IPFS
			ipfsURL, err := h.projectService.UploadMultimedia(c.Request.Context(), filename)
			if err != nil {
				errorsCh <- err
				return
			}

			mu.Lock()
			uploadedMedia = append(uploadedMedia, domain.Multimedia{URL: ipfsURL, Type: fileinfra.GetTypeByExtension(filename)})
			mu.Unlock()
		}(fileHeader)
	}

	// Esperar a que todas las goroutines terminen
	wg.Wait()
	defer close(errorsCh)

	// Verificar si hubo errores durante la ejecución de las goroutines
	select {
	case err := <-errorsCh:
		// Al menos una goroutine encontró un error
		httputil.NewError(c, http.StatusInternalServerError, errors.New(fmt.Sprintf("Failure on processing: %v", err)))
		return
	default:
		userID := c.PostForm("userId")
		title := c.PostForm("title")
		description := c.PostForm("description")
		// Todas las goroutines se ejecutaron sin errores
		// Crear un nuevo proyecto con los multimedia subidos
		project := &domain.Project{
			UserID:      userID,
			Title:       title,
			Description: description,
			Advances: []domain.ProjectAdvance{
				{
					Date:    time.Now().String(),
					Details: "created",
				},
			},
			Multimedia: uploadedMedia,
		}

		// Llamar al servicio para crear el nuevo proyecto
		createdProject, err := h.projectService.NewProject(c.Request.Context(), project)
		if err != nil {
			httputil.NewError(c, http.StatusInternalServerError, errors.New("Error creating the project"))
			return
		}

		// Responder con un mensaje de éxito y los detalles del proyecto creado
		c.JSON(http.StatusOK, gin.H{"message": "Project created successfully", "project": createdProject})
	}
}

// @Summary Obtener detalles del proyecto
// @Description Obtiene los detalles de un proyecto específico.
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "ID del proyecto"
// @Success 200 {object} domain.Project "Detalles del proyecto"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router /projects/{id} [get]
func (h *ProjectHandler) GetProjectDetails(c *gin.Context) {
	projectID := c.Param("id")

	// Llamar al servicio para obtener los detalles del proyecto
	project, err := h.projectService.GetProject(c.Request.Context(), projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los detalles del proyecto"})
		return
	}

	// Verificar si el proyecto existe
	if project == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Proyecto no encontrado"})
		return
	}

	// Responder con los detalles del proyecto
	c.JSON(http.StatusOK, project)
}

// @Summary Obtener lista de proyectos
// @Description Obtiene una lista paginada de proyectos.
// @Tags projects
// @Accept json
// @Produce json
// @Param limit query int false "Número máximo de proyectos por página" default(10)
// @Param page query int false "Número de página" default(1)
// @Success 200 {array} domain.Project "Lista de proyectos"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router /projects [get]
func (h *ProjectHandler) GetProjects(c *gin.Context) {
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	// Convertir los parámetros a enteros
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetro 'limit' inválido"})
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetro 'page' inválido"})
		return
	}

	// Llamar al servicio para obtener la lista de proyectos
	projects, err := h.projectService.GetProjects(c.Request.Context(), limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la lista de proyectos"})
		return
	}

	// Responder con la lista de proyectos
	c.JSON(http.StatusOK, projects)
}

// @Summary Eliminar proyecto
// @Description Elimina un proyecto por su ID.
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "ID del proyecto a eliminar"
// @Success 200 {object} gin.H "Proyecto eliminado exitosamente"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router /projects/{id} [delete]
func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	projectID := c.Param("id") // Supongo que el ID está en la ruta, ajusta según tu implementación

	// Llamar al servicio para eliminar el proyecto
	err := h.projectService.DeleteProject(c.Request.Context(), projectID)
	if err != nil {
		// Manejar el error y responder adecuadamente
		if errors.Is(err, domain.ErrProjectNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Proyecto no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el proyecto"})
		}
		return
	}

	// Responder con un mensaje de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Proyecto eliminado exitosamente"})
}

// @Summary Obtener proyectos por nombre similar
// @Description Obtiene la lista de proyectos cuyos nombres son similares al proporcionado.
// @Tags projects
// @Accept json
// @Produce json
// @Param partialName query string true "Nombre parcial del proyecto a buscar"
// @Success 200 {array} domain.Project "Lista de proyectos"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router /projects/search [get]
func (h *ProjectHandler) GetProjectsBySimilarName(c *gin.Context) {
	partialName := c.Query("partialName")

	// Verificar si se proporcionó el parámetro necesario
	if partialName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Se requiere el parámetro 'partialName'"})
		return
	}

	// Llamar al servicio para obtener la lista de proyectos con nombres similares
	projects, err := h.projectService.GetProjectsBySimilarName(c.Request.Context(), partialName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener la lista de proyectos"})
		return
	}

	// Responder con la lista de proyectos
	c.JSON(http.StatusOK, projects)
}

// @Summary Actualizar proyecto
// @Description Actualiza un proyecto existente con la información proporcionada.
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags projects
// @Accept json
// @Produce json
// @Param id path string true "ID del proyecto a actualizar"
// @Param project body domain.Project true "Datos del proyecto a actualizar"
// @Success 200 {object} domain.Project "Proyecto actualizado exitosamente"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router /projects/{id} [put]
func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	var updatedProject domain.Project

	// Bind the JSON request to the Project struct
	if err := c.ShouldBindJSON(&updatedProject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	updatedProject.ID = c.Param("id")

	// Llamar al servicio para actualizar el proyecto
	project, err := h.projectService.UpdateProject(c.Request.Context(), &updatedProject)
	if err != nil {
		// Manejar el error y responder adecuadamente
		if errors.Is(err, domain.ErrProjectNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Proyecto no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el proyecto"})
		}
		return
	}

	// Responder con el proyecto actualizado
	c.JSON(http.StatusOK, project)
}

// UpdateProjectMultimedia actualiza la multimedia de un proyecto.
// @Summary Actualiza la multimedia de un proyecto.
// @Description Este endpoint permite actualizar la multimedia de un proyecto existente.
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @ID updateProjectMultimedia
// @Accept mpfd
// @Tags projects
// @Produce json
// @Param id path string true "ID del proyecto"
// @Param files formData file true "Archivos multimedia a agregar"
// @Param updatedMultimedia body []domain.Multimedia true "Lista de objetos multimedia actualizados"
// @Success 200 {object} domain.Project "Proyecto actualizado exitosamente"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router /projects/{id}/media [put]
func (h *ProjectHandler) UpdateProjectMultimedia(c *gin.Context) {
	projectID := c.Param("id")

	var updatedMultimedia []domain.Multimedia

	if err := c.ShouldBindJSON(&updatedMultimedia); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Obtener los archivos del formulario
	formFiles := c.Request.MultipartForm.File["files"]

	var wg sync.WaitGroup
	var mu sync.Mutex
	var uploadedMedia []domain.Multimedia
	errorsCh := make(chan error, len(formFiles))

	for _, fileHeader := range formFiles {
		wg.Add(1)
		go func(fileHeader *multipart.FileHeader) {
			defer wg.Done()

			// Abrir el archivo
			file, err := fileHeader.Open()
			if err != nil {
				errorsCh <- err
				return
			}
			defer file.Close()

			// Crear un archivo en el servidor para almacenar el multimedia
			filename := filepath.Join(UploadTempDir, fileHeader.Filename)
			dst, err := os.Create(filename)
			if err != nil {
				errorsCh <- err
				return
			}
			defer dst.Close()

			// Copiar el contenido del archivo recibido al archivo en el servidor
			_, err = io.Copy(dst, file)
			if err != nil {
				errorsCh <- err
				return
			}

			// Subir el multimedia a IPFS
			ipfsURL, err := h.projectService.UploadMultimedia(c.Request.Context(), filename)
			if err != nil {
				errorsCh <- err
				return
			}

			mu.Lock()
			uploadedMedia = append(uploadedMedia, domain.Multimedia{URL: ipfsURL, Type: fileinfra.GetTypeByExtension(filename)})
			mu.Unlock()
		}(fileHeader)
	}

	// Esperar a que todas las goroutines terminen
	go func() {
		wg.Wait()
		close(errorsCh)
	}()

	// Verificar si hubo errores durante la ejecución de las goroutines
	select {
	case err := <-errorsCh:
		// Al menos una goroutine encontró un error
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failure on processing: %v", err)})
		return
	default:
		// Todas las goroutines se ejecutaron sin errores
		// Llamar al servicio para actualizar el proyecto con la nueva multimedia
		updatedMultimedia = append(updatedMultimedia, uploadedMedia...)

		project, err := h.projectService.UpdateProjectMultimedia(c.Request.Context(), projectID, updatedMultimedia)
		if err != nil {
			// Manejar el error y responder adecuadamente
			if errors.Is(err, domain.ErrProjectNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "Proyecto no encontrado"})
			} else if errors.Is(err, domain.ErrInvalidMediaType) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de multimedia no válido"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la multimedia del proyecto"})
			}
			return
		}

		// Responder con el proyecto actualizado
		c.JSON(http.StatusOK, project)
	}
}

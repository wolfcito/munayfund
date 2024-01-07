package handler

import (
	"errors"
	"munayfund-api2/internal/core/domain"
	"munayfund-api2/internal/core/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/swag/example/celler/httputil"
)

type HistoryHandler struct {
	historyService *service.HistoryService
}

func NewHistoryHandler(historySrv *service.HistoryService) *HistoryHandler {
	return &HistoryHandler{
		historyService: historySrv,
	}
}

// @Summary Obtener historial de proyectos
// @Description Obtiene el historial de avances de proyectos para un usuario específico.
// @Tags projects-history
// @ID get-history
// @Param id path string true "ID del proyecto"
// @Accept json
// @Produce json
// @Success 200 {array} domain.ProjectAdvance "Historial de proyectos"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router /projects/history/{id} [get]
func (h *HistoryHandler) GetHistory(c *gin.Context) {
	userID := c.Param("id")

	historyList, err := h.historyService.GetHistoryList(c.Request.Context(), userID)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, historyList)
}

// @Summary Actualizar historial de proyectos
// @Description Agrega un nuevo avance al historial de proyectos para un proyecto específico.
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Tags projects-history
// @ID update-history
// @Param id path string true "ID del proyecto"
// @Accept json
// @Produce json
// @Param request body domain.ProjectAdvance true "Datos del avance del historial a agregar"
// @Success 200 {array} domain.ProjectAdvance "Historial actualizado de proyectos"
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router /projects/history/{id} [put]
func (h *HistoryHandler) UpdateHistory(c *gin.Context) {
	projectID := c.Param("id")

	// Crear una estructura para recibir los datos del avance del historial desde el cuerpo de la solicitud
	var updatedHistory domain.ProjectAdvance

	// Decodificar el cuerpo JSON de la solicitud en la estructura del avance del historial
	if err := c.ShouldBindJSON(&updatedHistory); err != nil {
		httputil.NewError(c, http.StatusBadRequest, errors.New("Error al decodificar el cuerpo JSON"))
		return
	}

	// Llamar al servicio para agregar el avance al historial
	updatedHistories, err := h.historyService.AddHistory(c.Request.Context(), projectID, &updatedHistory)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	// Responder con la lista actualizada de avances del historial
	c.JSON(http.StatusOK, updatedHistories)
}

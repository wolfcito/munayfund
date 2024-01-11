package service

import (
	"context"
	"errors"
	"munayfund-api2/internal/core/domain"
	"munayfund-api2/internal/core/port"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProjectService struct {
	repo           port.ProjectRepository
	multimediaRepo port.MultimediaRepository
}

func NewProjectService(repo port.ProjectRepository, multimediaRepo port.MultimediaRepository) *ProjectService {
	return &ProjectService{
		repo:           repo,
		multimediaRepo: multimediaRepo,
	}
}

func (p *ProjectService) GetProject(ctx context.Context, projectID string) (*domain.Project, error) {
	projectObjectID, _ := primitive.ObjectIDFromHex(projectID)
	filter := bson.M{"_id": projectObjectID}

	// Llamar al repositorio para obtener el proyecto
	project := &domain.Project{}
	err := p.repo.FindOne(ctx, project, filter)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, errors.New("Error al obtener el proyecto desde la base de datos")
	}

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	return project, nil
}

func (p *ProjectService) GetProjects(ctx context.Context, limit, page int) ([]*domain.Project, error) {
	skip := (page - 1) * limit

	// Configurar opciones para la paginación
	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(skip))

	// Llamar al repositorio para obtener la lista de proyectos
	projects := []*domain.Project{}
	filter := bson.M{} // Puedes agregar filtros adicionales aquí si es necesario
	err := p.repo.Find(ctx, &projects, filter, findOptions)
	if err != nil {
		return nil, errors.New("Error al obtener la lista de proyectos desde la base de datos")
	} else if err != mongo.ErrNoDocuments {
		return projects, nil
	}

	return projects, nil
}

func (p *ProjectService) NewProject(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	err := p.repo.InsertOne(ctx, project)
	if err != nil {
		return nil, errors.New("Error al crear el proyecto en la base de datos")
	}

	return project, nil
}

func (p *ProjectService) UpdateProject(ctx context.Context, project *domain.Project) (*domain.Project, error) {
	existingProject := &domain.Project{}

	projectObjectID, _ := primitive.ObjectIDFromHex(project.ID)

	err := p.repo.FindOne(ctx, existingProject, bson.M{"_id": projectObjectID})
	if err != nil {
		return nil, domain.ErrProjectNotFound
	}

	// Actualizar campos necesarios
	existingProject.Title = project.Title
	existingProject.Description = project.Description
	// Agregar lógica para actualizar otros campos según sea necesario

	// Llamar al repositorio para actualizar el proyecto
	err = p.repo.UpdateOne(ctx, bson.M{"_id": project.ID}, bson.M{"$set": existingProject})
	if err != nil {
		return nil, errors.New("Error al actualizar el proyecto en la base de datos")
	}

	return existingProject, nil
}

func (p *ProjectService) DeleteProject(ctx context.Context, projectID string) error {
	existingProject := &domain.Project{}

	projectObjectID, _ := primitive.ObjectIDFromHex(projectID)

	err := p.repo.FindOne(ctx, existingProject, bson.M{"_id": projectObjectID})
	if err != nil {
		return domain.ErrProjectNotFound
	}

	// Llamar al repositorio para eliminar el proyecto
	err = p.repo.DeleteOne(ctx, bson.M{"_id": projectObjectID})
	if err != nil {
		return errors.New("Error al eliminar el proyecto en la base de datos")
	}

	return nil
}

func (p *ProjectService) GetProjectsBySimilarName(ctx context.Context, partialName string) ([]*domain.Project, error) {
	// Crear filtro para buscar proyectos con nombres similares
	filter := bson.M{"title": bson.M{"$regex": primitive.Regex{Pattern: partialName, Options: "i"}}}

	// Llamar al repositorio para obtener la lista de proyectos
	projects := []*domain.Project{}
	err := p.repo.Find(ctx, &projects, filter)
	if err != nil {
		return nil, errors.New("Error al obtener la lista de proyectos desde la base de datos")
	}

	return projects, nil
}

func (p *ProjectService) UploadMultimedia(ctx context.Context, filePath string) (string, error) {
	return p.multimediaRepo.UploadFile(filePath)
}

// UpdateProjectMultimedia actualiza la multimedia de un proyecto
func (p *ProjectService) UpdateProjectMultimedia(ctx context.Context, projectID string, updatedMultimediaList []domain.Multimedia) (*domain.Project, error) {
	existingProject := &domain.Project{}

	projectObjectID, _ := primitive.ObjectIDFromHex(projectID)

	err := p.repo.FindOne(ctx, existingProject, bson.M{"_id": projectObjectID})
	if err != nil {
		return nil, errors.New("Proyecto no encontrado")
	}

	// Lógica para agregar o quitar elementos multimedia según la operación indicada
	for _, updatedMultimedia := range updatedMultimediaList {
		switch updatedMultimedia.Type {
		case "image":
			// Agregar la imagen si no está presente
			if !containsMedia(existingProject.Multimedia, &updatedMultimedia) {
				existingProject.Multimedia = append(existingProject.Multimedia, updatedMultimedia)
			}
		case "video":
			// Quitar el video si está presente
			if containsMedia(existingProject.Multimedia, &updatedMultimedia) {
				existingProject.Multimedia = removeMedia(existingProject.Multimedia, &updatedMultimedia)
			}
		default:
			return nil, domain.ErrInvalidMediaType
		}
	}
	// Llamar al repositorio para actualizar el proyecto
	err = p.repo.UpdateOne(ctx, bson.M{"_id": projectID}, bson.M{"$set": existingProject})
	if err != nil {
		return nil, errors.New("Error al actualizar el proyecto en la base de datos")
	}

	return existingProject, nil
}

// containsMedia verifica si un elemento multimedia ya está presente en la lista
func containsMedia(mediaList []domain.Multimedia, targetMedia *domain.Multimedia) bool {
	for _, media := range mediaList {
		if media.URL == targetMedia.URL && media.Type == targetMedia.Type {
			return true
		}
	}
	return false
}

// removeMedia elimina un elemento multimedia de la lista
func removeMedia(mediaList []domain.Multimedia, targetMedia *domain.Multimedia) []domain.Multimedia {
	var updatedList []domain.Multimedia
	for _, media := range mediaList {
		if media.URL != targetMedia.URL || media.Type != targetMedia.Type {
			updatedList = append(updatedList, media)
		}
	}
	return updatedList
}

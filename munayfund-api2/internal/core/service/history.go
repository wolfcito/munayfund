package service

import (
	"context"
	"errors"
	"munayfund-api2/internal/core/domain"
	"munayfund-api2/internal/core/port"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HistoryService struct {
	repo           port.HistoryRepository
	multimediaRepo port.MultimediaRepository
}

func NewHistoryService(repo port.HistoryRepository, multimediaRepo port.MultimediaRepository) *HistoryService {
	return &HistoryService{
		repo:           repo,
		multimediaRepo: multimediaRepo,
	}
}

func (s *HistoryService) GetHistoryList(ctx context.Context, id string) ([]domain.ProjectAdvance, error) {
	var project domain.Project

	projectObjectID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": projectObjectID}

	err := s.repo.FindOne(ctx, &project, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("Historial no encontrado")
		}
		return nil, errors.New("Error al obtener el historial desde la base de datos")
	}

	return project.Advances, nil
}

func (s *HistoryService) AddHistory(ctx context.Context, projectID string, updatedHistory *domain.ProjectAdvance) ([]domain.ProjectAdvance, error) {
	// Verificar si el historial existe
	existingProject := &domain.Project{}

	projectObjectID, _ := primitive.ObjectIDFromHex(projectID)

	err := s.repo.FindOne(ctx, existingProject, bson.M{"_id": projectObjectID})
	if err != nil {
		if errors.Is(err, domain.ErrHistoryNotFound) {
			return nil, domain.ErrHistoryNotFound
		}
		return nil, errors.New("Error al obtener el historial desde la base de datos")
	}

	// Actualizar campos necesarios
	updatedHistory.Date = time.Now().String()
	existingProject.Advances = append(existingProject.Advances, *updatedHistory)

	// Agregar lógica para actualizar otros campos según sea necesario

	// Llamar al repositorio para actualizar el historial
	err = s.repo.UpdateOne(ctx, bson.M{"_id": projectObjectID},
		bson.M{
			"$set": bson.M{
				"advances": existingProject.Advances,
			},
		},
	)
	if err != nil {
		return nil, errors.New("Error al actualizar el historial en la base de datos")
	}

	// Obtener la lista actualizada de historiales
	return existingProject.Advances, nil
}

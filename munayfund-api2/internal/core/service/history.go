package service

import (
	"context"
	"errors"
	"munayfund-api2/internal/core/domain"
	"munayfund-api2/internal/core/port"

	"go.mongodb.org/mongo-driver/bson"
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

	filter := bson.M{"_id": id}

	err := s.repo.Find(ctx, &project, filter)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("Historial no encontrado")
		}
		return nil, errors.New("Error al obtener el historial desde la base de datos")
	}

	return project.Advances, nil
}

func (s *HistoryService) AddHistory(ctx context.Context, projectID string, updatedHistory *domain.ProjectAdvance) ([]*domain.ProjectAdvance, error) {
	// Verificar si el historial existe
	existingProject := &domain.Project{}
	err := s.repo.Find(ctx, existingProject, bson.M{"_id": projectID})
	if err != nil {
		if errors.Is(err, domain.ErrHistoryNotFound) {
			return nil, domain.ErrHistoryNotFound
		}
		return nil, errors.New("Error al obtener el historial desde la base de datos")
	}

	// Actualizar campos necesarios
	existingProject.Advances = append(existingProject.Advances, *updatedHistory)

	// Agregar lógica para actualizar otros campos según sea necesario

	// Llamar al repositorio para actualizar el historial
	err = s.repo.UpdateOne(ctx, bson.M{"_id": projectID}, bson.M{"$set": existingProject})
	if err != nil {
		return nil, errors.New("Error al actualizar el historial en la base de datos")
	}

	// Obtener la lista actualizada de historiales
	updatedHistories, err := s.getAllHistories(ctx)
	if err != nil {
		return nil, errors.New("Error al obtener la lista actualizada de historiales")
	}

	return updatedHistories, nil
}

func (s *HistoryService) getAllHistories(ctx context.Context) ([]*domain.ProjectAdvance, error) {
	var histories []*domain.ProjectAdvance
	err := s.repo.Find(ctx, &histories, bson.M{})
	if err != nil {
		return nil, err
	}
	return histories, nil
}

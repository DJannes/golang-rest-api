package service

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/janneseffendi/rest-api/internal/dto"
	"gitlab.com/janneseffendi/rest-api/internal/repository"
)

type PublicService struct {
	repo *repository.Repository
}

func NewPublicService() *PublicService {
	return &PublicService{
		repo: repository.NewRepository(),
	}
}

func (s *PublicService) GetPublicData(ctx context.Context) (*dto.PublicData, error) {
	persons, err := s.repo.PublicRepo.GetPersons(ctx)
	if err != nil {
		return nil, fmt.Errorf("getPersons: %w", err)
	}

	return &dto.PublicData{
		PublicName:         persons[0].Email,
		PublicDate:         time.Now(),
		PublicDescriptions: []string{"Data 1"},
	}, nil
}

func (s *PublicService) SavePublicData(req dto.SavePublicData) dto.SavePublicData {
	return req
}

package service

import (
	"time"

	"gitlab.com/janneseffendi/rest-api/internal/dto"
)

type PublicService struct {
}

func NewPublicService() *PublicService {
	return &PublicService{}
}

func (s *PublicService) GetPublicData() dto.PublicData {
	return dto.PublicData{
		PublicName:         "public name",
		PublicDate:         time.Now(),
		PublicDescriptions: []string{"Data 1"},
	}
}
func (s *PublicService) SavePublicData(req dto.SavePublicData) dto.SavePublicData {
	return req
}

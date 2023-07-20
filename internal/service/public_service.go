package service

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
	"gitlab.com/janneseffendi/rest-api/internal/dto"
	"gitlab.com/janneseffendi/rest-api/internal/internal_utils"
	"gitlab.com/janneseffendi/rest-api/internal/repository"
	"gitlab.com/janneseffendi/rest-api/internal/repository/generated"
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
	publicData, err := s.repo.PublicRepo.GetPublicDataByEmail(ctx, "test@gmail.com")
	if err != nil {
		return nil, fmt.Errorf("getPersonByEmail: %w", err)
	}

	return &dto.PublicData{
		Email:             publicData.Email,
		Name:              publicData.Name.String,
		AdditionalInfo:    publicData.AdditionalInfo,
		Birthdate:         publicData.Birthdate.Time.Format("2006-01-02"),
		AccBalanceNull:    publicData.AccBalanceNull.Decimal.StringFixed(0),
		AccBalance:        publicData.AccBalance.StringFixed(5),
		UserCredentialsID: publicData.UserCredentialsID,
		CommentsNull:      publicData.CommentsNull,
		Comments:          publicData.Comments,
	}, nil
}

func (s *PublicService) SavePublicData(ctx context.Context, uuid string, req dto.SavePublicData) error {
	birthDate, _ := time.Parse("2006-01-02", req.PublicDate)
	if err := s.repo.PublicRepo.UpsertPublicDataByUUID(ctx, generated.UpsertPublicDataByUUIDParams{
		Uuid:              pgtype.UUID{Bytes: internal_utils.GetUUIDFromString(uuid), Valid: true},
		Email:             "test@gmail.com",
		Name:              pgtype.Text{String: req.PublicName, Valid: true},
		AdditionalInfo:    req.PublicDescriptions,
		Birthdate:         pgtype.Timestamptz{Time: birthDate, Valid: true},
		AccBalance:        decimal.NewFromFloat(0.123412),
		AccBalanceNull:    decimal.NullDecimal{},
		UserCredentialsID: 1,
		Comments:          []byte("{}"),
		CommentsNull:      nil,
	}); err != nil {
		return fmt.Errorf("UpsertPublicDataByUUID: %w", err)
	}

	return nil
}

func (s *PublicService) DeletePublicData(ctx context.Context, uuid string) error {
	if err := s.repo.PublicRepo.DeletePublicDataByUUID(ctx, pgtype.UUID{
		Bytes: internal_utils.GetUUIDFromString(uuid),
		Valid: true},
	); err != nil {
		return fmt.Errorf("DeletePublicDataByUUID: %w", err)
	}

	return nil
}

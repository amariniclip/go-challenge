package services

import (
	"context"
	"go-challenge/internal/domain"
	"go-challenge/internal/repository"
	"time"

	"github.com/google/uuid"
)

// ShirtService is responsible for creating, reading, updating and deleting t-shirts
// from our store's stock.
type ShirtService struct {
	repo repository.Shirts
}

// NewShirtService instantiates a new ShirtService.
func NewShirtService(repo repository.Shirts) *ShirtService {
	return &ShirtService{
		repo: repo,
	}
}

// Create will create a new shirt entity in our database.
//
// This method also makes some validations about incoming shirt data.
func (s *ShirtService) Create(ctx context.Context, input *domain.Shirt) (*domain.Shirt, error) {
	shirt := &domain.Shirt{
		Brand:          input.Brand,
		Description:    input.Description,
		AvailableUnits: input.AvailableUnits,
		PricePerUnit:   input.PricePerUnit,
	}

	shirt.ID = uuid.NewString()
	shirt.CreatedAt = time.Now()
	shirt.UpdatedAt = time.Now()
	shirt.Status = domain.ShirtStatusNormal

	if err := shirt.Validate(); err != nil {
		return nil, err
	}

	if err := s.repo.Save(ctx, shirt); err != nil {
		return nil, err
	}

	return shirt, nil
}

func (s *ShirtService) Read(ctx context.Context, id string) (*domain.Shirt, error) {
	return s.repo.Get(ctx, id)
}

// UpdateShirtParams are the available params to update a shirt entity.
type UpdateShirtParams struct {
	Brand          string             `json:"brand"`
	Size           string             `json:"size"`
	Description    string             `json:"description"`
	Status         domain.ShirtStatus `json:"status"`
	AvailableUnits *int               `json:"available_units"`
	PricePerUnit   *float64           `json:"price_per_unit"`
}

// Update runs an update over a Shirt entity using UpdateShirtParams object.
//
// This works like a 'patch' so the null or empty attributes means no alteration.
func (s *ShirtService) Update(ctx context.Context, id string, updateParams *UpdateShirtParams) (*domain.Shirt, error) {
	shirt, err := s.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	if updateParams.Brand != "" {
		shirt.Brand = updateParams.Brand
	}

	if updateParams.Description != "" {
		shirt.Description = updateParams.Description
	}

	if updateParams.Status != "" {
		shirt.Status = updateParams.Status
	}

	if updateParams.AvailableUnits != nil {
		shirt.AvailableUnits = *updateParams.AvailableUnits
	}

	if updateParams.PricePerUnit != nil && *updateParams.PricePerUnit > 0 {
		shirt.PricePerUnit = *updateParams.PricePerUnit
	}

	shirt.UpdatedAt = time.Now()
	if err := s.repo.Save(ctx, shirt); err != nil {
		return nil, err
	}

	return shirt, nil
}

func (s *ShirtService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

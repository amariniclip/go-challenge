package services

import (
	"context"
	"go-challenge/internal/domain"
	"go-challenge/internal/repository"
	"time"

	"github.com/google/uuid"
)

// SalesService is sales manager in our system.
//
// The main responsibility of this service is to manage the lifecycle of a sale.
// This includes creation and refund.
type SalesService struct {
	shirtRepo repository.Shirts
	salesRepo repository.Sales
}

// NewSalesService instantiates a new SalesService.
func NewSalesService(salesRepo repository.Sales, shirtRepo repository.Shirts) *SalesService {
	return &SalesService{
		shirtRepo: shirtRepo,
		salesRepo: salesRepo,
	}
}

// Create will try to create a new sale.
//
// This process runs as follows:
// - Get information about each shirt included in the sale.
// - Reduce the stock of each shirt (change status if applies)
// - Save the sale.
func (s *SalesService) Create(ctx context.Context, input *domain.Sale) (*domain.Sale, error) {
	sale := &domain.Sale{
		Customer: input.Customer,
		Cart:     input.Cart,
	}

	sale.ID = uuid.NewString()
	sale.CreatedAt = time.Now()
	sale.UpdatedAt = time.Now()
	sale.Status = domain.SaleStatusApproved

	for i, cartItem := range sale.Cart {
		shirt, err := s.shirtRepo.Get(ctx, cartItem.ID)
		if err != nil {
			return nil, err
		}

		sale.Cart[i].Description = shirt.Description

		shirt.AvailableUnits -= cartItem.Units

		if err := s.shirtRepo.Save(ctx, shirt); err != nil {
			return nil, err
		}
	}

	if err := s.salesRepo.Save(ctx, sale); err != nil {
		return nil, err
	}

	return sale, nil
}

// Create will try to create a refund from a sale.
func (s *SalesService) Refund(ctx context.Context, id string) (*domain.Sale, error) {
	sale, err := s.salesRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	sale.UpdatedAt = time.Now()
	sale.Status = domain.SaleStatusRefunded

	if err := s.salesRepo.Save(ctx, sale); err != nil {
		return nil, err
	}

	return sale, nil
}

func (s *SalesService) Read(ctx context.Context, id string) (*domain.Sale, error) {
	return s.salesRepo.Get(ctx, id)
}

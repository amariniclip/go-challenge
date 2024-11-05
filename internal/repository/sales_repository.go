package repository

import (
	"context"
	"errors"
	"go-challenge/internal/domain"
)

var ErrSaleNotFound = errors.New("sale not found")

type Sales interface {
	Save(ctx context.Context, sale *domain.Sale) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*domain.Sale, error)
}

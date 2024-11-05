package repository

import (
	"context"
	"errors"
	"go-challenge/internal/domain"
)

var ErrShirtNotFound = errors.New("shirt not found")

type Shirts interface {
	Save(ctx context.Context, shirt *domain.Shirt) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*domain.Shirt, error)
}

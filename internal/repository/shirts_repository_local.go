package repository

import (
	"context"
	"errors"
	"go-challenge/internal/domain"
)

type shirtsLocal struct {
	localStorage map[string]interface{}
}

func NewShirtsLocal() Shirts {
	return &shirtsLocal{
		localStorage: map[string]interface{}{},
	}
}

func (l *shirtsLocal) Save(ctx context.Context, shirt *domain.Shirt) error {
	if shirt.ID == "" {
		return errors.New("id for new shirt is empty")
	}

	l.localStorage[shirt.ID] = shirt
	return nil
}

func (l *shirtsLocal) Delete(ctx context.Context, id string) error {
	delete(l.localStorage, id)
	return nil
}

func (l *shirtsLocal) Get(ctx context.Context, id string) (*domain.Shirt, error) {
	s, ok := l.localStorage[id].(*domain.Shirt)
	if !ok {
		return nil, ErrShirtNotFound
	}

	return s, nil
}

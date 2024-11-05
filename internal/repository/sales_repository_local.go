package repository

import (
	"context"
	"errors"
	"go-challenge/internal/domain"
)

type salesLocal struct {
	localStorage map[string]interface{}
}

func NewSalesLocal() Sales {
	return &salesLocal{
		localStorage: map[string]interface{}{},
	}
}

func (l *salesLocal) Save(ctx context.Context, sale *domain.Sale) error {
	if sale.ID == "" {
		return errors.New("id for new sale is empty")
	}

	l.localStorage[sale.ID] = sale
	return nil
}

func (l *salesLocal) Delete(ctx context.Context, id string) error {
	_, err := l.Get(ctx, id)
	if err != nil {
		return err
	}

	delete(l.localStorage, id)
	return nil
}

func (l *salesLocal) Get(ctx context.Context, id string) (*domain.Sale, error) {
	s, ok := l.localStorage[id].(*domain.Sale)
	if !ok {
		return nil, ErrSaleNotFound
	}

	return s, nil
}

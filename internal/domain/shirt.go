package domain

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

const (
	_maxSizeDescription = 100
	_maxUnits           = 100000
)

// Shirt represents the shirts available in our store.
type Shirt struct {
	ID             string      `json:"id"`
	Brand          string      `json:"brand" binding:"required"`
	Description    string      `json:"description" binding:"required"`
	Status         ShirtStatus `json:"status"`
	AvailableUnits int         `json:"available_units" binding:"required"`
	PricePerUnit   float64     `json:"price_per_unit" binding:"required"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
}

type ShirtValidationError []string

func (ss ShirtValidationError) Error() string {
	return strings.Join(ss, " AND ")
}

func (s *Shirt) Validate() ShirtValidationError {
	var toRet ShirtValidationError
	if !slices.Contains(_validStatus, string(s.Status)) {
		toRet = append(toRet, fmt.Sprintf("%s is invalid status", s.Status))
	}

	if len(s.Description) > _maxSizeDescription {
		toRet = append(toRet, "description has to be less than 100 chars")
	}

	if s.AvailableUnits > _maxUnits {
		toRet = append(toRet, "units must to be less than 100000")
	}

	return toRet
}

type ShirtStatus string

const (
	// ShirtStatusNormal indicates that the shirts are ready to sell.
	ShirtStatusNormal ShirtStatus = "NORMAL"

	// ShirtStatusDeleted indicates that the shirts was deleted logically.
	ShirtStatusDeleted ShirtStatus = "DELETED"

	// ShirtStatusWithoutStock indicates that the shirts runs out of stock.
	// This means 'available_units' is zero.
	ShirtStatusWithoutStock ShirtStatus = "WITHOUT_STOCK"
)

var _validStatus = []string{
	string(ShirtStatusDeleted),
	string(ShirtStatusNormal),
	string(ShirtStatusWithoutStock),
}

package domain

import "time"

// Sale represents a sale of a few shirts in our system.
type Sale struct {
	ID          string         `json:"id"`
	Customer    SaleCustomer   `json:"customer" binding:"required"`
	Cart        []SaleCartItem `json:"cart" binding:"required"`
	TotalAmount float64        `json:"total_amount"`
	Status      SaleStatus     `json:"status"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

// SaleCustomer is who buy the shirts.
type SaleCustomer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type SaleCartItem struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Units       int    `json:"units"`
}

type SaleStatus string

const (
	SaleStatusApproved SaleStatus = "APPROVED"
	SaleStatusRefunded SaleStatus = "REFUNDED"
)

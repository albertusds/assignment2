package params

import (
	"assignment2/models"
	"time"
)

type CreateOrder struct {
	CustomerName string       `json:"customerName"`
	Items        []CreateItem `json:"items"`
}

type UpdateOrder struct {
	ID           uint         `json:"order_id,omitempty"`
	CustomerName string       `json:"customerName"`
	Items        []CreateItem `json:"items"`
}

type ReponseGetOrder struct {
	OrderId      uint          `json:"order_id"`
	CustomerName string        `json:"customerName,omitempty"`
	OrderedAt    time.Time     `json:"OrderedAt,omitempty"`
	Items        []models.Item `json:"items"`
}

type ReponseDeleteOrder struct {
	OrderId uint `json:"order_id"`
}

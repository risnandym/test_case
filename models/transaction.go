package models

import (
	"time"
)

type (

	// transaction
	Transaction struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		Customer  string    `json:"customer"`
		Quantity  int       `json:"quantity"`
		Price     float64   `json:"price"`
		RequestID int       `json:"request_id"`
		TimeStamp time.Time `json:"timestamp"`
	}

	// request
	TransactionRequest struct {
		RequestID int           `json:"request_id"`
		Data      []Transaction `json:"data"`
	}

	InputDummy struct {
		Total int `json:"total"`
	}
)

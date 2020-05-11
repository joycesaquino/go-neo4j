package service

import (
	"time"
)

type Service struct {
	Description      string    `json:"description"`
	Value            float64   `json:"value"`
	Date             time.Time `json:"date"`
	MinSubscriptions int       `json:"minSubscriptions"`
	MaxSubscriptions int       `json:"maxSubscriptions"`
}

package service

import (
	"time"
)

type Service struct {
	Id               string    `json:"id"`
	Description      string    `json:"description"`
	Value            float64   `json:"value"`
	Date             time.Time `json:"date"`
	MinSubscriptions int       `json:"minSubscriptions"`
	MaxSubscriptions int       `json:"maxSubscriptions"`
}

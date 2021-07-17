package model

import "github.com/neo4j/neo4j-go-driver/v4/neo4j"

type Service struct {
	Id          string              `json:"id"`
	Description string              `json:"description"`
	Value       float64             `json:"value"`
	CreatedAt   neo4j.LocalDateTime `json:"createdAt"`
}
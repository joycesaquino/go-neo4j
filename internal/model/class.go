package model

import "github.com/neo4j/neo4j-go-driver/neo4j"

type Class struct {
	Id               string              `json:"id"`
	Description      string              `json:"description"`
	Value            float64             `json:"value"`
	InitialDateTime  neo4j.LocalDateTime `json:"initialDate"`
	FinalDateTime    neo4j.LocalDateTime `json:"finalDate"`
	Subscriptions    int64               `json:"subscriptions"`
	MinSubscriptions int64               `json:"minSubscriptions"`
	MaxSubscriptions int64               `json:"maxSubscriptions"`
	CreatedAt        neo4j.LocalDateTime `json:"createdAt"`
}

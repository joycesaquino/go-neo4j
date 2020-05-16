package service

import (
	"dera-services-api/internal/persistence"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"log"
)

type Dao struct {
	neo4jConnection *persistence.Neo4Go
}

type Service struct {
	Id               string              `json:"id"`
	Description      string              `json:"description"`
	Value            float64             `json:"value"`
	InitialDateTime  neo4j.LocalDateTime `json:"date"`
	FinalDateTime    neo4j.LocalDateTime `json:"date"`
	MinSubscriptions int64               `json:"minSubscriptions"`
	MaxSubscriptions int64               `json:"maxSubscriptions"`
	CreatedAt        neo4j.LocalDateTime `json:"createdAt"`
}

func (dao Dao) Insert(s Service) error {
	result, err := dao.neo4jConnection.Session.Run("CREATE (n:Service { id: $id, description: $description, value: $value, initialDateTime: $initialDateTime, finalDateTime: $finalDateTime, minSubscriptions: $minSubscriptions, maxSubscriptions: $maxSubscriptions,createdAt: $createdAt}) RETURN n.id, n.description",
		map[string]interface{}{
			"id":               s.Id,
			"description":      s.Description,
			"value":            s.Value,
			"initialDateTime":  s.InitialDateTime,
			"finalDateTime":    s.FinalDateTime,
			"minSubscriptions": s.MinSubscriptions,
			"maxSubscriptions": s.MaxSubscriptions,
			"createdAt":        s.CreatedAt,
		})
	if err != nil {
		return err
	}
	for result.Next() {
		log.Printf("Created Item with Id = '%s' and Description = '%s'\n", result.Record().GetByIndex(0).(string), result.Record().GetByIndex(1).(string))
	}

	if err = result.Err(); err != nil {
		return err // handle error
	}
	defer dao.neo4jConnection.Close()

	return nil
}

func NewDao() *Dao {
	return &Dao{neo4jConnection: persistence.NewNeo4Go()}
}

package service

import (
	"dera-services-api/internal/persistence"
	"time"
)

type Dao struct {
	neo4jConnection persistence.Neo4Go
}

type Service struct {
	Id               string    `json:"id"`
	Description      string    `json:"description"`
	Value            float64   `json:"value"`
	Date             time.Time `json:"date"`
	MinSubscriptions int       `json:"minSubscriptions"`
	MaxSubscriptions int       `json:"maxSubscriptions"`
}

func (dao Dao) insert(s Service) error {
	_, err := dao.neo4jConnection.Session.Run("CREATE (n:Service { id: $id, description: $description, value: $value, date: $date, minSubscriptions: $minSubscriptions, maxSubscriptions: $maxSubscriptions}) RETURN n.id",
		map[string]interface{}{
			"Id":               s.Id,
			"Description":      s.Description,
			"Value":            s.Value,
			"Date":             s.Date,
			"MinSubscriptions": s.MinSubscriptions,
			"MaxSubscriptions": s.MaxSubscriptions,
		})
	if err != nil {
		return err
	}
	defer dao.neo4jConnection.Close()

	return nil
}

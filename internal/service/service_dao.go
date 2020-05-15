package service

import (
	"dera-services-api/internal/persistence"
)

type Dao struct {
	neo4jConnection *persistence.Neo4Go
}

type Service struct {
	Id               string  `json:"id"`
	Description      string  `json:"description"`
	Value            float64 `json:"value"`
	Date             string  `json:"date"`
	MinSubscriptions int     `json:"minSubscriptions"`
	MaxSubscriptions int     `json:"maxSubscriptions"`
}

func (dao Dao) Insert(s Service) error {
	_, err := dao.neo4jConnection.Session.Run("CREATE (n:Service { id: $id, description: $description, value: $value, date: $date, minSubscriptions: $minSubscriptions, maxSubscriptions: $maxSubscriptions}) RETURN n.id",
		map[string]interface{}{
			"id":               s.Id,
			"description":      s.Description,
			"value":            s.Value,
			"date":             s.Date,
			"minSubscriptions": s.MinSubscriptions,
			"maxSubscriptions": s.MaxSubscriptions,
		})
	if err != nil {
		return err
	}

	defer dao.neo4jConnection.Close()

	return nil
}

func NewDao() *Dao {
	return &Dao{neo4jConnection: persistence.NewNeo4Go()}
}

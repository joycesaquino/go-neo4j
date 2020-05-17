package dao

import (
	"dera-services-api/internal/persistence"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"log"
)

type Dao struct {
	neo4jConnection *persistence.Neo4Go
}
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
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

func (dao Dao) FindById(id string) (interface{}, error) {

	result, err := dao.neo4jConnection.Session.Run(FindServiceById, map[string]interface{}{"id": id})
	if err != nil {
		log.Printf("Error :%s", err)
	}
	for result.Next() {
		return result.Record().GetByIndex(0), nil
	}
	if err != nil {
		return nil, err
	}
	dao.neo4jConnection.Close()

	return nil, err
}

func (dao Dao) Insert(service *Service, user *User) error {
	result, err := dao.neo4jConnection.
		Session.
		Run(InsertServiceWithUserRelation,
			map[string]interface{}{

				//Service information
				"id":               service.Id,
				"description":      service.Description,
				"value":            service.Value,
				"initialDateTime":  service.InitialDateTime,
				"finalDateTime":    service.FinalDateTime,
				"minSubscriptions": service.MinSubscriptions,
				"maxSubscriptions": service.MaxSubscriptions,
				"createdAt":        service.CreatedAt,

				//User information
				"name":  user.Name,
				"email": user.Email,
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

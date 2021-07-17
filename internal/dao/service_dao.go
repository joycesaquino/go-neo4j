package dao

import (
	"dera-services-api/internal/model"
	"dera-services-api/internal/persistence"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"log"
)

type Dao struct {
	neo4jConnection *persistence.Neo4Go
}

func (dao Dao) FindById(id string) (*model.Class, error) {

	result, err := dao.neo4jConnection.Session.Run(FindServiceById, map[string]interface{}{"id": id})
	if err != nil {
		log.Printf("Error :%s", err)
	}

	for result.Next() {
		returnedMap := result.Record().GetByIndex(0).(neo4j.Node)
		service := returnedMap.Props()
		return &model.Class{
			Id:               service["id"].(string),
			Description:      service["description"].(string),
			Value:            service["value"].(float64),
			InitialDateTime:  service["initialDateTime"].(neo4j.LocalDateTime),
			FinalDateTime:    service["finalDateTime"].(neo4j.LocalDateTime),
			Subscriptions:    service["subscriptions"].(int64),
			MinSubscriptions: service["minSubscriptions"].(int64),
			MaxSubscriptions: service["maxSubscriptions"].(int64),
			CreatedAt:        service["createdAt"].(neo4j.LocalDateTime),
		}, nil
	}

	_ = dao.neo4jConnection.Close()

	return nil, err
}
func (dao Dao) InsertService(service *model.Service, user *model.User) error {
	result, err := dao.neo4jConnection.
		Session.
		Run(CreateService,
			map[string]interface{}{

				//Class information
				"id":          service.Id,
				"description": service.Description,
				"value":       service.Value,
				"createdAt":   service.CreatedAt,

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
		return err
	}

	_ = dao.neo4jConnection.Close()

	return nil

}

func (dao Dao) InsertClass(class *model.Class, id string) error {
	result, err := dao.neo4jConnection.
		Session.
		Run(InsertClassWithServiceRelation,
			map[string]interface{}{

				//Class information
				"id":               class.Id,
				"description":      class.Description,
				"value":            class.Value,
				"initialDateTime":  class.InitialDateTime,
				"finalDateTime":    class.FinalDateTime,
				"subscriptions":    class.Subscriptions,
				"minSubscriptions": class.MinSubscriptions,
				"maxSubscriptions": class.MaxSubscriptions,
				"createdAt":        class.CreatedAt,

				//Service id for realtions
				"service_id": id,
			})
	if err != nil {
		return err
	}
	for result.Next() {
		log.Printf("Created Item with Id = '%s' and Description = '%s'\n", result.Record().GetByIndex(0).(string), result.Record().GetByIndex(1).(string))
	}

	if err = result.Err(); err != nil {
		return err
	}

	_ = dao.neo4jConnection.Close()

	return nil
}

func NewDao() *Dao {
	return &Dao{neo4jConnection: persistence.NewNeo4Go()}
}

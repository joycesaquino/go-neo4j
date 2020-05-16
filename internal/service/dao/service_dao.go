package dao

import (
	"dera-services-api/internal/persistence"
	"dera-services-api/internal/service"
	"dera-services-api/internal/service/query"
	"log"
)

type Dao struct {
	neo4jConnection *persistence.Neo4Go
}

func (dao Dao) Insert(service service.Service, user service.User) error {
	result, err := dao.neo4jConnection.
		Session.
		Run(query.InsertServiceWithUserRelation,
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

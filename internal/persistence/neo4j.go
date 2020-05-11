package persistence

import (
	"dera-services-api/internal/service"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Config struct {
	User     string `env:"USER_DATABASE"`
	Password string `env:"PASSWORD_DATABASE"`
	Uri      string `env:"URI_DATABASE"`
}

type Neo4Go struct {
	config  Config
	driver  neo4j.Driver
	session neo4j.Session
	result  neo4j.Result
}

func (neo4go Neo4Go) neo4jConnection() (neo4j.Session, error) {
	sess, err := neo4go.driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func (neo4go Neo4Go) insert(s service.Service) error {
	_, err := neo4go.session.Run("CREATE (n:Service { id: $id, description: $description, value: $value, date: $date, minSubscriptions: $minSubscriptions, maxSubscriptions: $maxSubscriptions}) RETURN n.id",
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

	return nil
}

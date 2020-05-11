package persistence

import (
	"dera-services-api/internal/service"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Config struct {
	User     string `env:"USER_DATABASE"`
	Password string `env:"PASSWORD_DATABASE"`
}

type Neo4Go struct {
	config  Config
	driver  neo4j.Driver
	session neo4j.Session
	result  neo4j.Result
}

func (neo4go Neo4Go) insert(service service.Service) {
}

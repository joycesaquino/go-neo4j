package database

import (
	"github.com/caarlos0/env"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
)

type Config struct {
	User         string `env:"USER_DATABASE,required"`
	Password     string `env:"PASSWORD_DATABASE,required"`
	Uri          string `env:"URI_DATABASE,required"`
	DatabaseName string `env:"DATABASE_NAME,required"`
}

type Neo4Go struct {
	config  Config
	Session neo4j.Session
}

func neo4jConnection(config Config) (neo4j.Session, error) {

	authToken := neo4j.BasicAuth(config.User, config.Password, "")
	driver, err := neo4j.NewDriver(config.Uri, authToken)
	if err != nil {
		return nil, err
	}
	sess := driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: config.DatabaseName,
	})

	return sess, nil
}

func (neo4go Neo4Go) Close() error {
	err := neo4go.Session.Close()
	if err != nil {
		return err
	}
	return nil
}

func NewNeo4Go() *Neo4Go {
	var config Config
	err := env.Parse(&config)
	if err != nil {
		log.Fatalf("Error to configure Neo4J cliente. Error : %s", err)
	}

	sess, err := neo4jConnection(config)
	if err != nil {
		log.Fatalf("Error to configure Neo4J cliente. Error : %s", err)
	}

	return &Neo4Go{config: config, Session: sess}
}

package persistence

import (
	"github.com/caarlos0/env"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"log"
)

type Config struct {
	User     string `env:"USER_DATABASE,required"`
	Password string `env:"PASSWORD_DATABASE,required"`
	Uri      string `env:"URI_DATABASE,required"`
}

type Neo4Go struct {
	config  Config
	Session neo4j.Session
}

func neo4jConnection(config Config) (neo4j.Session, error) {
	useConsoleLogger := func(level neo4j.LogLevel) func(config *neo4j.Config) {
		return func(config *neo4j.Config) {
			config.Log = neo4j.ConsoleLogger(level)
		}
	}

	driver, err := neo4j.NewDriver(config.Uri, neo4j.BasicAuth(config.User, config.Password, ""), useConsoleLogger(neo4j.DEBUG))
	if err != nil {
		return nil, err
	}

	sess, err := driver.Session(neo4j.AccessModeWrite)
	if err != nil {
		return nil, err
	}

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

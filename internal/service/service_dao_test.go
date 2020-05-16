package service

import (
	"dera-services-api/internal/persistence"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"os"
	"testing"
	"time"
)

func beforeInsert() {
	_ = os.Setenv("USER_DATABASE", "neo4j")
	_ = os.Setenv("PASSWORD_DATABASE", "dera")
	_ = os.Setenv("URI_DATABASE", "bolt://0.0.0.0:7687")
}

func TestDao_Insert(t *testing.T) {
	beforeInsert()
	type fields struct {
		neo4jConnection *persistence.Neo4Go
	}
	type args struct {
		service *Service
		user    *User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "Insert service object on database", fields: fields{neo4jConnection: persistence.NewNeo4Go()}, args: args{
			service: &Service{
				Id:               "0002",
				Description:      "Meditação Guiada Nível Intermediário",
				Value:            120.90,
				InitialDateTime:  neo4j.LocalDateTimeOf(time.Now()),
				FinalDateTime:    neo4j.LocalDateTimeOf(time.Now()),
				MinSubscriptions: 20,
				MaxSubscriptions: 5,
				CreatedAt:        neo4j.LocalDateTimeOf(time.Now()),
			},
			user: &User{
				Name:  "Joyce Aquino",
				Email: "joycesaquino@gmail.com",
			},
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := NewDao()
			if err := dao.Insert(tt.args.service, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

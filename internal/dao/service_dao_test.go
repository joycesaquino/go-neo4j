package dao

import (
	"dera-services-api/internal/database"
	"dera-services-api/internal/model"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"os"
	"testing"
	"time"
)

func init() {
	_ = os.Setenv("USER_DATABASE", "neo4j")
	_ = os.Setenv("PASSWORD_DATABASE", "dera")
	_ = os.Setenv("URI_DATABASE", "bolt://0.0.0.0:7687")
	_ = os.Setenv("DATABASE_NAME", "neo4j")
}

func TestDao_CreateClass(t *testing.T) {
	type fields struct {
		neo4jConnection *database.Neo4Go
	}
	type args struct {
		service *model.Class
		id      string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "CreateClass service object on database", fields: fields{neo4jConnection: database.NewNeo4Go()}, args: args{
			service: &model.Class{
				Id:               "0001",
				Description:      "Fundamentos de Backend",
				Value:            180.50,
				InitialDateTime:  neo4j.LocalDateTimeOf(time.Time{}),
				FinalDateTime:    neo4j.LocalDateTimeOf(time.Time{}),
				Subscriptions:    0,
				MinSubscriptions: 100,
				MaxSubscriptions: 20,
				CreatedAt:        neo4j.LocalDateTimeOf(time.Time{}),
			},
			id: "IDS001",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := NewDao()
			if err := dao.InsertClass(tt.args.service, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("CreateClass() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDao_FindById(t *testing.T) {
	type fields struct {
		neo4jConnection *database.Neo4Go
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Class
		wantErr bool
	}{
		{name: "Finding service by id", fields: fields{neo4jConnection: database.NewNeo4Go()}, args: args{"0001"}, want: &model.Class{
			Id:               "0001",
			Description:      "Fundamentos de Backend",
			Value:            180.50,
			InitialDateTime:  neo4j.LocalDateTimeOf(time.Time{}),
			FinalDateTime:    neo4j.LocalDateTimeOf(time.Time{}),
			Subscriptions:    0,
			MinSubscriptions: 100,
			MaxSubscriptions: 20,
			CreatedAt:        neo4j.LocalDateTimeOf(time.Time{}),
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := Dao{
				neo4jConnection: tt.fields.neo4jConnection,
			}
			got, err := dao.FindById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				if err != nil {
					t.Errorf("FindById() got = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}

func TestDao_CreateService(t *testing.T) {
	type fields struct {
		neo4jConnection *database.Neo4Go
	}
	type args struct {
		service *model.Service
		user    *model.User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		//want Service
		wantErr bool
	}{
		{name: "User Create Service", fields: fields{neo4jConnection: database.NewNeo4Go()}, args: args{
			service: &model.Service{
				Id:          "IDS001",
				Description: "Aula de culinária com a Palmirinha e Anna Maria",
				Value:       80.00,
				CreatedAt:   neo4j.LocalDateTimeOf(time.Now()),
			},
			user: &model.User{
				Name:  "Palmirinha Maria",
				Email: "palmirinha@gmail.com",
			},
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := Dao{
				neo4jConnection: tt.fields.neo4jConnection,
			}
			if err := dao.InsertService(tt.args.service, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("InsertService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

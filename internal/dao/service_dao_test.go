package dao

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
				Id:               "0001",
				Description:      "Fundamentos de Backend",
				Value:            180.50,
				InitialDateTime:  neo4j.LocalDateTimeOf(time.Time{}),
				FinalDateTime:    neo4j.LocalDateTimeOf(time.Time{}),
				MinSubscriptions: 100,
				MaxSubscriptions: 20,
				CreatedAt:        neo4j.LocalDateTimeOf(time.Time{}),
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

func TestDao_FindById(t *testing.T) {
	beforeInsert()
	type fields struct {
		neo4jConnection *persistence.Neo4Go
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Service
		wantErr bool
	}{
		{name: "Finding service by id", fields: fields{neo4jConnection: persistence.NewNeo4Go()}, args: args{"0001"}, want: &Service{
			Id:               "0001",
			Description:      "Fundamentos de Backend",
			Value:            180.50,
			InitialDateTime:  neo4j.LocalDateTimeOf(time.Time{}),
			FinalDateTime:    neo4j.LocalDateTimeOf(time.Time{}),
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

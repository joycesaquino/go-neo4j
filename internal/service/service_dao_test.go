package service

import (
	"dera-services-api/internal/persistence"
	"os"
	"testing"
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
		s Service
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "Insert service object on database", fields: fields{neo4jConnection: persistence.NewNeo4Go()}, args: args{
			s: Service{
				Id:               "0001",
				Description:      "Aula de JAVA e Orientação a Objetos",
				Value:            120.90,
				Date:             "time.Now()",
				MinSubscriptions: 20,
				MaxSubscriptions: 5,
			},
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dao := NewDao()
			if err := dao.Insert(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

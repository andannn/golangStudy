package database

import (
	"example.com/internal/infra/database/ent"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func NewDatabaseClient() *ent.Client {
	client, err := ent.Open("mysql", "root:poi@tcp(localhost:3306)/db_ent_demo?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalf("error when close connection to mysql%v", err)
		}
	}(client)

	return client
}

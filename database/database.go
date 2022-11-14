package database

import (
	"context"
	"log"
	"telegra/ent"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var EntClient *ent.Client

func init() {
	// host := config.Config("DB_HOST")
	// port := config.Config("DB_PORT")
	// user := config.Config("DB_USER")
	// password := config.Config("DB_PASSWORD")
	// dbname := config.Config("DB_NAME")
	// connectString := fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=disable", host, port, dbname, user, password)

	// add sslmode=disable to resolve error pq: SSL is not enabled on the server
	// client, err := ent.Open("postgres", connectString)
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	EntClient = client
}

//docker run -d -p 5432:5432 --name postgres_DB -e POSTGRES_PASSWORD=mysecretpassword postgres
//docker exec -it postgres_DB bash

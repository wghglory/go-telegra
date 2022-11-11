package database

import (
	"context"
	"log"
	"telegra/ent"

	_ "github.com/lib/pq"
)

var EntClient *ent.Client

func init() {
	// add sslmode=disable to resolve error pq: SSL is not enabled on the server
	client, err := ent.Open("postgres", "host=localhost port=5432 user=postgres dbname=telegra password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	EntClient = client
}

//docker run -d -p 5432:5432 --name postgres_DB -e POSTGRES_PASSWORD=mysecretpassword postgres
//docker exec -it postgres_DB bash

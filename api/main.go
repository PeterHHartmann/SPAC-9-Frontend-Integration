//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema

package main

import (
	"api/db"
	"api/ent"
	"api/grpc"
	"context"
	"log"

	_ "github.com/lib/pq"
)

type MovieQuote struct {
	ID       	int    `json:"id"`
	Quote    	string `json:"quote"`
	Movie    	string `json:"movie"`
	Character 	string `json:"character"`
	Year     	int    `json:"year"`
}

func main() {
	// Get DB connection string
	connStr := db.GetConnectionString()

	client, err := ent.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}
	defer client.Close()

	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("❌ Failed to run migrations: %v", err)
	}

	if err := db.Seed(ctx, client); err != nil {
		log.Fatalf("❌ Failed to seed data: %v", err)
	}

	grpc.StartGRPC(client)
}

package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"hello/pkg/ent"
	"hello/pkg/environment"

	// Import postgres dialect for ent
	_ "github.com/lib/pq"
)

func NewClient() *ent.Client {
	connstring := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_SSL"),
	)
	postgres, err := ent.Open("postgres", connstring)
	if err != nil {
		log.Fatal(err)
	}

	if environment.IsDevelopment() {
		if err := postgres.Schema.Create(context.Background()); err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	}
	return postgres
}

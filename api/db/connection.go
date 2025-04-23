package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

//Gets connectionString for database from .env file or return default
func GetConnectionString() string {
	if os.Getenv("DEV_ENV") == "production" {
		err := godotenv.Load("./../.env.production")
		if err != nil {
			log.Fatal(err)
			log.Fatal("❌ Error loading .env.production file")
		}
	}

	postgresDB := os.Getenv("POSTGRES_DB")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresHost := os.Getenv("POSTGRES_HOST")

	if postgresDB == "" || postgresUser == "" || postgresPassword == "" || postgresPort == "" || postgresHost == "" {
		log.Println("⚠️ One or more environment variables are missing, using default fallback.")
		return "host=localhost port=5432 user=exampleuser dbname=exampledb password=examplepwd sslmode=disable"
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgresHost, postgresPort, postgresUser, postgresPassword, postgresDB,
	)
}
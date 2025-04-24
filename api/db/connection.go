package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

//Gets connectionString for database from .env file or return default
func GetConnectionString() string {
	devenv := getEnv("DEV_ENV", "example")

	env_file := fmt.Sprintf(
		"./../.env.%s",
		devenv,
	)

	err := godotenv.Load(env_file)
	if err != nil {
		log.Fatal(err)
		log.Fatal("❌ Error loading .env file")
	}

	postgresDB := os.Getenv("POSTGRES_DB")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresHost := os.Getenv("POSTGRES_HOST")

	// return fmt.Sprintf(
	// 	"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	postgresHost, postgresPort, postgresUser, postgresPassword, postgresDB,
	// )
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?search_path=public&sslmode=disable",
		postgresUser, postgresPassword, postgresHost, postgresPort, postgresDB,
	)
}
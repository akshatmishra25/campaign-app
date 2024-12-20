package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	MongoDBURI  string
	DatabaseName string
}

var AppConfig Config

func LoadConfig() {

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found. Using system environment variables.")
	}


	AppConfig = Config{
		Port:        getEnv("PORT", "8080"),
		MongoDBURI:  getEnv("MONGODB_URI", ""),
		DatabaseName: getEnv("DATABASE_NAME", "campaignDB"),
	}

	if AppConfig.MongoDBURI == "" {
		log.Fatal("MONGODB_URI is not set in the environment")
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

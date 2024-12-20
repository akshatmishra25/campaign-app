package database

import (
	"context"
	"log"
	"time"

	"campaign-app.local/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() {
	clientOptions := options.Client().ApplyURI(config.AppConfig.MongoDBURI)

	var err error
	DB, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = DB.Ping(ctx, nil); err!=nil {
		log.Fatalf("Ping to MongoDB failed: %v", err)
	}
	
	log.Println("Connected to MongoDB")
}

func GetCollection(collectionName string) *mongo.Collection {
	if DB == nil {
		log.Fatal("Database connection is not initialized")
	}
	return DB.Database(config.AppConfig.DatabaseName).Collection(collectionName)
}

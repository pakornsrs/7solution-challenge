package configs

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongoDBConnection() *mongo.Client {
	mongoDBConnectionString, err := getMongoDBConnectionString()
	if err != nil {
		log.Fatal(fmt.Printf("MongoDB connection error: %s", err.Error()))
	}

	opts := options.Client().ApplyURI(mongoDBConnectionString)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		log.Fatal(fmt.Printf("MongoDB connection error: %s", err.Error()))
	}

	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(fmt.Printf("MongoDB connection error: %s", err.Error()))
	} else {
		log.Println("Successfully connected to MongoDB")
	}

	return client
}

func getMongoDBConnectionString() (string, error) {

	ip := os.Getenv("MONGO_IP")
	port := os.Getenv("MONGO_PORT")
	username := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")

	return buildUpConnectionString(ip, port, username, password)
}

func buildUpConnectionString(ip string, port string, username string, password string) (string, error) {

	if ip == "" || port == "" || username == "" || password == "" {
		return "", errors.New("cannot get connection string for sql db")
	}

	var builder strings.Builder
	builder.WriteString("mongodb://")
	builder.WriteString(username + ":")
	builder.WriteString(password)
	builder.WriteString("@" + ip + ":" + port)

	return builder.String(), nil

}

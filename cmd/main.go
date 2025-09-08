package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"pakornssn/7solution-challenge/configs"
	routes "pakornssn/7solution-challenge/internal/routers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Setup Service Routes
	router := gin.Default()
	router.Use(cors.Default())

	// implement mondoDB
	mongoClient := configs.SetupMongoDBConnection()
	defer mongoClient.Disconnect(context.Background())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := mongoClient.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB connection failed (ping): %v", err)
	}

	// gRPC Client
	routes.SetupServiceHealthCheckRoute(router)

	// Starting Service
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := router.Run(":" + port); err != nil {
		log.Fatal(fmt.Printf("Starting service error: %s", err.Error()))
	}

	log.Println("Wealth-Up server started at port: ", port)

}

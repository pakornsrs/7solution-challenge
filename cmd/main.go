package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"pakornssn/7solution-challenge/configs"
	"pakornssn/7solution-challenge/internal/handlers"
	"pakornssn/7solution-challenge/internal/repositories"
	routes "pakornssn/7solution-challenge/internal/routers"
	"pakornssn/7solution-challenge/internal/services"
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

	// implement service client
	userRepositoryClient := repositories.NewUserRepository(mongoClient)

	userServiceClient := services.NewUserService(userRepositoryClient)
	authServiceClient := services.NewAuthenticationService(userRepositoryClient)

	authUser := handlers.NewAuthenticationHandler(authServiceClient)
	userHandler := handlers.NewUserHandler(userServiceClient)

	// gRPC Client
	routes.SetupServiceHealthCheckRoute(router)
	routes.SetupAuthenticationServiceRoute(router, authUser)
	routes.SetupUserServiceRoute(router, userHandler)

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

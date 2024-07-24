package main

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"CRUD_echo_ent/database"
	"CRUD_echo_ent/handlers"
)

func main() {
	// Connect to the database
	client, err := database.ConnectDatabase()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Create Echo instance
	e := echo.New()

	// Initialize handlers
	h := handlers.NewHandler(client)

	// Routes
	e.POST("/users", h.CreateUser)
	e.GET("/users/:id", h.GetUser)
	e.PUT("/users/:id", h.UpdateUser)
	e.DELETE("/users/:id", h.DeleteUser)
	e.GET("/users", h.ListUsers)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
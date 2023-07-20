package http

import (
	"context"
	"log"
	"net/http"
	"rest-api/database"
	"rest-api/internal/user/application"
	"rest-api/internal/user/infrastructure/mariadb"
	"rest-api/settings"
)

func StartServer() {

	myContext := context.Background()

	myConfig, err := settings.New()
	if err != nil {
		log.Panicf("failed to load settings %s", err)
	}

	myConnection, err := database.New(myContext, myConfig)
	if err != nil {
		log.Panicf("failed to start database %s", err)
	}

	postgresRepo := mariadb.NewPostgresRepository(myConnection)
	userService := application.NewUserService(postgresRepo)

	userHandler := NewUserHandler(userService)
	http.HandleFunc("/users", userHandler.CreateUserHandler)

	log.Println("Server listening on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

package http

import (
	"context"
	"log"
	"net/http"
	"rest-api/internal/database"
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

	userHandler := NewHandler(userService)

	router := NewRouter()
	router.Handle("/users", http.HandlerFunc(userHandler.CreateUser))
	router.Handle("/user/:id", http.HandlerFunc(userHandler.GetUserByID))

	log.Println("Server listening on :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}

}

package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"rest-api/database"
	"rest-api/internal/user/application"
	"rest-api/internal/user/infrastructure/mariadb"
	"rest-api/settings"
)

func StartServer() {
	myContext := context.Background()
	var userService = &application.UserService{}
	var mariaDBRepo = &mariadb.MariaDBRepository{}

	myConfig, err := settings.New()
	if err != nil {
		log.Panicf("failed to load settings %s", err)
	}

	myConnection, err := database.New(myContext, myConfig)
	if err != nil {
		log.Panicf("failed to start database %s", err)
	}

	switch myConfig.DB.Engine {
	case "mariadb":
		mariaDBRepo = mariadb.NewMariaDBRepository(myConnection)
		userService = application.NewUserService(mariaDBRepo)
		fmt.Println("mariadb is connected")

	case "postgres":
		fmt.Println("postgresSQL is connected")

	default:
		panic("not engine case contemplated")
	}

	mux := http.NewServeMux()
	userH := &userHandler{userService}
	mux.Handle("/users", userH)

	log.Println("Server listening on :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}

}

package http

import (
	"context"
	"encoding/json"
	"net/http"
	"rest-api/internal/user/application"
)

type UserHandler struct {
	userService *application.UserService
}

func NewUserHandler(userService *application.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	// Parsear el cuerpo de la solicitud para obtener los datos del usuario
	var userParams struct {
		Name     string
		Email    string
		Password string
	}

	err := json.NewDecoder(r.Body).Decode(&userParams)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Llamar al servicio para registrar el usuario
	err = h.userService.SaveUser(ctx, userParams.Name, userParams.Email, userParams.Password)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

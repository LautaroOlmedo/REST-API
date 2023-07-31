package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api/internal/user/application"
	"strconv"
)

type Handler struct {
	userService *application.UserService
}

func NewHandler(userService *application.UserService) *Handler {
	return &Handler{userService: userService}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}
	fmt.Fprintf(w, "Hello there %s", "visitor")
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

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

	err = h.userService.RegisterUser(ctx, userParams.Name, userParams.Email, userParams.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	idStr := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userService.GetByID(ctx, userID)
	if err != nil {

		http.Error(w, fmt.Sprintf("Failed to get user: %v", err), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

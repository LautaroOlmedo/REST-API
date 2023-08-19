package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"rest-api/internal/user/application"
	"rest-api/internal/user/application/DTOs"
)

var (
	listUserRe   = regexp.MustCompile(`^\/users[\/]*$`)
	getUserRe    = regexp.MustCompile(`^\/users\/(\d+)$`) // ---> /users/123
	createUserRe = regexp.MustCompile(`^\/users[\/]*$`)
	updateUserRe = regexp.MustCompile(`^\/users\/(\d+)$`) // ---> /users/123
)

type userHandler struct {
	userService *application.UserService
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch {
	case r.Method == http.MethodGet && listUserRe.MatchString(r.URL.Path):
		h.List(w, r)
		return
	case r.Method == http.MethodGet && getUserRe.MatchString(r.URL.Path):
		h.Get(w, r)
		return

	case r.Method == http.MethodPost && createUserRe.MatchString(r.URL.Path):
		h.Create(w, r)
		return

	case r.Method == http.MethodPut && updateUserRe.MatchString(r.URL.Path):
		h.Update(w, r)
		return
	default:
		http.NotFound(w, r)
		return
	}

}

func (h *userHandler) List(w http.ResponseWriter, r *http.Request) {
	myContext := context.Background()
	users, err := h.userService.GetAll(myContext)

	if err != nil {
		internalServerError(w, r)
		return
	}
	// Serialize the user map in JSON
	usersJSON, jsonErr := json.Marshal(users)
	if jsonErr != nil {
		internalServerError(w, r)
		return
	}

	// Configures the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send the JSON response to the client
	_, _ = w.Write(usersJSON)
}

func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal("Hello world!!")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

	/*myContext := context.Background()
	matches := getUserRe.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		http.NotFound(w, r)
		return
	}
	params := strings.Split(matches[1], "/")
	idStr := params[len(params[1])]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	user, err := h.userService.GetByID(myContext, id)
	if err != nil {
		if !errors.Is(err, application.InternalServerError) {
			if errors.Is(err, application.UserNotFound) {
				userNotFound(w, r)
				return
			}
		} else {
			internalServerError(w, r)
			return
		}
	}
	fmt.Println(*user)
	w.WriteHeader(http.StatusOK)*/
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	myContext := context.Background()
	var user DTOs.UserDTO
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.userService.RegisterUser(myContext, user.Name, user.Email, user.Password)
	if err != nil {
		fmt.Println(err)
		if !errors.Is(err, application.InternalServerError) {
			if errors.Is(err, application.UserAlreadyExist) {
				http.Error(w, "error: user already exists", http.StatusBadRequest)
				return
			} else {
				fmt.Println(err)
				http.Error(w, "missing required fields", http.StatusBadRequest)
				return
			}
		} else {
			internalServerError(w, r)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal("Hello world from update user!!")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func userNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`error: user not found`))
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(`error: internal server error`))
}

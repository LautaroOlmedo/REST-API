package http

import (
	"encoding/json"
	"net/http"
	"regexp"
	"rest-api/internal/user/application"
)

var (
	listUserRe   = regexp.MustCompile(`^\/users[\/]*$`)
	getUserRe    = regexp.MustCompile(`^\/users\/(\d+)$`) // ---> /users/123
	createUserRe = regexp.MustCompile(`^\/users[\/]*$`)
)

type userHandler struct {
	userService *application.UserService
}

func (h *userHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
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
	default:
		http.NotFound(w, r)
		return
	}

}

func (h *userHandler) List(w http.ResponseWriter, r *http.Request) {
	jsonBytes, err := json.Marshal("Hello world!!")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}

func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {

}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {

}

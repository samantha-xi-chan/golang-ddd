package http

import (
	"encoding/json"
	"golang-ddd/internal/app"
	"golang-ddd/internal/domain"
	"net/http"
)

type UserHandler struct {
	PUserService *app.UserService
}

func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/user", h.handleUser)
}

func (h *UserHandler) handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var user domain.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err := h.PUserService.CreateUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

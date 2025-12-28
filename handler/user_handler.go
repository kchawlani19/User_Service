package handler

import (
	"encoding/json"
	"go-basic-user-service/model"
	"go-basic-user-service/service"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) HandleUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost && r.URL.Path == "/users" {
		var user model.User
		json.NewDecoder(r.Body).Decode(&user)

		if err := h.service.Create(user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/users/") {
		id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/users/"))

		switch r.Method {

		case http.MethodGet:
			user, err := h.service.Get(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(user)

		case http.MethodPut:
			var body struct {
				Name string `json:"name"`
			}
			json.NewDecoder(r.Body).Decode(&body)

			if err := h.service.Update(id, body.Name); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)

		case http.MethodDelete:
			if err := h.service.Delete(id); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		}
		return
	}

	http.NotFound(w, r)
}

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

		err := h.service.Create(user)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	}

	//get /users/{id}
	if r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/users/") {
		idStr := strings.TrimPrefix(r.URL.Path, "/users/")
		id, _ := strconv.Atoi(idStr)

		user, err := h.service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)
		return
	}

	if r.Method == http.MethodPut && strings.HasPrefix(r.URL.Path, "/users/") {
		idStr := strings.TrimPrefix(r.URL.Path, "/users/")
		id, _ := strconv.Atoi(idStr)

		var body struct {
			Name string `json:"name"`
		}
		json.NewDecoder(r.Body).Decode(&body)

		err := h.service.Update(id, body.Name)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	//delete

	if r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/users/") {
		idStr := strings.TrimPrefix(r.URL.Path, "/users/")
		id, _ := strconv.Atoi(idStr)

		err := h.service.Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
		//w.WriteHeader(http.StatusNotFound)
	}
}

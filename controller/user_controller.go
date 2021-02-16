package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gugazimmermann/xteam-bootcamp-go/service"
)

// UserGet return one or all users
func UserGet(r *http.Request) (int, string) {
	email := r.URL.Query().Get("email")
	if email != "" {
		return service.GetByEmail(email)
	}
	return service.GetAll()
}

// UserPost create one user
func UserPost(r *http.Request) (int, string) {
	var user service.User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	if user.Email != "" && user.Name != "" {
		return service.AddOne(user)
	}
	return http.StatusTeapot, `{Message: "email and name must be provided"}`
}

// UserUpdate update one user
func UserUpdate(r *http.Request) (int, string) {
	var user service.User
	json.NewDecoder(r.Body).Decode(&user)
	if user.ID != "" && user.Email != "" && user.Name != "" {
		return service.UpdateByID(user)
	}
	return http.StatusTeapot, `{Message: "ID, email and name must be provided"}`
}

// UserDelete delete one user by id
func UserDelete(r *http.Request) (int, string) {
	id := r.URL.Query().Get("id")
	return service.DeleteByID(id)
}

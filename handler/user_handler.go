package handler

import (
	"fmt"
	"net/http"

	"github.com/gugazimmermann/xteam-bootcamp-go/controller"
)

// UserHandler handle resquest to '/user'
func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		status, res := controller.UserGet(r)
		sendresponse(w, status, res)
	case "POST":
		status, res := controller.UserPost(r)
		sendresponse(w, status, res)
	case "PUT":
		status, res := controller.UserUpdate(r)
		sendresponse(w, status, res)
	case "DELETE":
		status, res := controller.UserDelete(r)
		sendresponse(w, status, res)
	default:
		sendresponse(w, http.StatusMethodNotAllowed, `{Message: "method not allowed"}`)
	}
}

func sendresponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v", message)
}

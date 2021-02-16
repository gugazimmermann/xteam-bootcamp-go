package handler

import (
	"fmt"
	"net/http"
)

// RootHandler handle resquest to '/'
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the website")
}

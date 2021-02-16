package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gugazimmermann/xteam-bootcamp-go/handler"
	"github.com/joho/godotenv"
)

// RunWeb starts running the web portal on address addr
func RunWeb() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	address := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")
	http.HandleFunc("/", handler.RootHandler)
	http.HandleFunc("/user", handler.UserHandler)
	log.Printf("Starting webserver on %s:%s", address, port)
	return http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), nil)
}

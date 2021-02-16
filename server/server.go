package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gugazimmermann/xteam-bootcamp-go/handler"
)

type configuration struct {
	Address string
	Port    int
}

// RunWeb starts running the web portal on address addr
func RunWeb() error {
	file, err := os.Open("./webserver.json")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler.RootHandler)
	http.HandleFunc("/user", handler.UserHandler)
	config := new(configuration)
	json.NewDecoder(file).Decode(config)
	log.Printf("Starting webserver on %s:%d", config.Address, config.Port)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", config.Address, config.Port), nil)
}

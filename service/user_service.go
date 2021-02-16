package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/google/uuid"
)

// User the user model
type User struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getUsersFile() ([]User, error) {
	file, err := os.Open("./users.json")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	users := new([]User)
	json.NewDecoder(file).Decode(users)
	return *users, nil
}

func saveUsersFile(users []User) {
	file, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile("./users.json", file, 0644)
	if err != nil {
		log.Println(err)
	}
}

func remove(users []User, i int) []User {
	copy(users[i:], users[i+1:])
	return users[:len(users)-1]
}

func internalServerError() (int, string) {
	return http.StatusInternalServerError, `{"message": "Internal Server Error"}`
}

func notFound() (int, string) {
	return http.StatusNotFound, `{Message: "user not found"}`
}

// GetByEmail return one user by email
func GetByEmail(email string) (int, string) {
	users, err := getUsersFile()
	if err != nil {
		return internalServerError()
	}
	var user = User{}
	for i := range users {
		if users[i].Email == email {
			user = users[i]
			break
		}
	}
	if reflect.DeepEqual(user, User{}) {
		return notFound()
	}
	userjson, _ := json.Marshal(user)
	return http.StatusOK, string(userjson)
}

// GetAll return all users
func GetAll() (int, string) {
	users, err := getUsersFile()
	if err != nil {
		return internalServerError()
	}
	usersjson, _ := json.Marshal(users)
	if string(usersjson) == "null" {
		log.Println("GetAll :", string(usersjson))
		return http.StatusOK, `[]`
	}
	return http.StatusOK, string(usersjson)
}

// AddOne add one user to the user list
func AddOne(user User) (int, string) {
	user.ID = uuid.NewString()
	users, err := getUsersFile()
	if err != nil {
		return internalServerError()
	}
	users = append(users, user)
	saveUsersFile(users)
	userjson, _ := json.Marshal(user)
	return http.StatusCreated, string(userjson)
}

// UpdateByID update one user by ID
func UpdateByID(user User) (int, string) {
	users, err := getUsersFile()
	if err != nil {
		return internalServerError()
	}
	for i := range users {
		if users[i].ID == user.ID {
			fmt.Println(user)
			users[i] = user
			saveUsersFile(users)
			userjson, _ := json.Marshal(user)
			return http.StatusOK, string(userjson)
		}
	}
	return notFound()
}

// DeleteByID delete one user by ID
func DeleteByID(id string) (int, string) {
	users, err := getUsersFile()
	if err != nil {
		return internalServerError()
	}
	for i := range users {
		if users[i].ID == id {
			users = remove(users, i)
			saveUsersFile(users)
			return http.StatusNoContent, `{}`
		}
	}
	return notFound()
}

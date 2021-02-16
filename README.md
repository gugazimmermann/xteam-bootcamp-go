# X-Team Go Bootcamp

https://github.com/julianosouza/go-crash-course/tree/master/08-http

## HTTP

https://gobyexample.com/http-clients

https://gobyexample.com/http-servers

#### web Server

rename `.env.sample` to `.env`

open http://localhost:8080 in your favorite browser

#### `/user` endpoint

For this project I'm using a json file to store user's data

* **GET ALL** - `curl -X GET http://localhost:8080/user`

* **GET BY EMAIL** - `curl -X GET http://localhost:8080/user?email=EMAIL` - change EMAIL to one from the GET ALL

* **ADD** - `curl -X POST --data "{\"name\":\"NAME\", \"email\":\"EMAIL\"}" --header "Content-Type:application/json" http://localhost:8080/user` - change NAME and EMAIL

* **UPDATE BY ID** - `curl -X PUT --data "{\"id\":\"ID\",\"name\":\"NAME\", \"email\":\"EMAIL\"}" --header "Content-Type:application/json" http://localhost:8080/user` - change ID, NAME and EMAIL

* **REMOVE BY ID** - `curl -X DELETE http://localhost:8080/user?id=ID` - change ID



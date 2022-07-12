package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kartien/go-gorm-restapi/db"
	"github.com/kartien/go-gorm-restapi/models"
	"github.com/kartien/go-gorm-restapi/routes"
)


func main() {
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", routes.HomeHandler)
	
	r.HandleFunc("/users", routes.GetrUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetrUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)

}




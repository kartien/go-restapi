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
	//Users 
	r.HandleFunc("/users", routes.GetrUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetrUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	//Tasks 
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks/{id}", routes.GetTasksHandler).Methods("GET")
	r.HandleFunc("/tasks", routes.PostTaskHandler).Methods("POST")
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}




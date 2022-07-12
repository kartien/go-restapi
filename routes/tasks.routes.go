package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kartien/go-gorm-restapi/db"
	"github.com/kartien/go-gorm-restapi/models"
)


func GetTasksHandler(w http.ResponseWriter, r *http.Request){
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request){
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	json.NewEncoder(w).Encode(&task)
}

func PostTaskHandler(w http.ResponseWriter, r *http.Request){
	var task models.Task

	json.NewDecoder(r.Body).Decode(&task)

	createTask := db.DB.Create(&task)
	err := createTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400 
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&task)

	w.Write([]byte("Post"))
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request){
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])

	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Task not found"))
		return
	}
	db.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusNoContent)
}

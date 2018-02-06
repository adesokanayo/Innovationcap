package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Task is defined.
type Task struct {
	ID          string      `json:"id,omitempty"`
	Title       string      `json:"title,omitempty"`
	Description string      `json:"description,omitempty"`
	Completed   bool        `json:"completed,omitempty"`
	TaskDetail  *TaskDetail `json:"taskdetail,omitempty"`
}

//TaskDetail is exported
type TaskDetail struct {
	Category string `json:"category,omitempty"`
	Owner    string `json:"owner,omitempty"`
}

var work []Task

//GetAllTask returns all the task.
func GetAllTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(work)

}

//GetSingleTask returns a task by ID
func GetSingleTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range work {
		if item.ID == params["id"] {

			json.NewEncoder(w).Encode(item)
			return
		}

	}

	json.NewEncoder(w).Encode(&Task{})

}

//CreateTask is to create a new task
//func CreateTask(w http.ResponseWriter, r *http.request) {

//}

//DeleteTask is used to delete a task
//func DeleteTask(w http.ResponseWriter, r *http.request) {

//}

func main() {

	work = append(work, Task{ID: "1", Title: "Deployment", Description: "Deployment of XYZ", Completed: true, TaskDetail: &TaskDetail{Category: "QA", Owner: "Odunayo"}})
	work = append(work, Task{ID: "2", Title: "Testing", Description: "Jmeter", Completed: false, TaskDetail: &TaskDetail{Category: "QA", Owner: "Collins"}})
	work = append(work, Task{ID: "3", Title: "Copying ", Description: "TIIS Setup", Completed: true, TaskDetail: &TaskDetail{Category: "QA", Owner: "Jesupelumi"}})

	router := mux.NewRouter()
	router.HandleFunc("/tasks", GetAllTask).Methods("GET")
	router.HandleFunc("/task/{id}", GetSingleTask).Methods("GET")
	//router.HandleFunc("/createtask/{id}", CreateTask).Methods("POST")
	//router.HandleFunc("/deletetask/{id}", DeleteTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

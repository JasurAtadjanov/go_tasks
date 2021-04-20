package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllTasks")
	json.NewEncoder(w).Encode(getAllTasks())
}

func createNewTask(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this and add to DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	var task Task
	json.Unmarshal(reqBody, &task)
	fmt.Println("Endpoint Hit: createNewTask: ", task)

	createTask(task)

	json.NewEncoder(w).Encode(getAllTasks())
}

func editTask(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// unmarshal this and add to DB
	reqBody, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 0)

	CheckError(err)

	var task Task
	json.Unmarshal(reqBody, &task)
	fmt.Println("Endpoint Hit: Task: ", task)

	updateTask(id, task)

	json.NewEncoder(w).Encode(getAllTasks())
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 0)

	CheckError(err)

	delTask(id)

	json.NewEncoder(w).Encode(getAllTasks())
}

func returnSingleTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleTask")
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 0)

	CheckError(err)

	json.NewEncoder(w).Encode(getTask(id))
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/tasks", returnAllTasks)
	myRouter.HandleFunc("/task", createNewTask).Methods("POST")
	myRouter.HandleFunc("/task/{id}", editTask).Methods("PUT")
	myRouter.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")
	myRouter.HandleFunc("/tasks/{id}", returnSingleTask)
	log.Fatal(http.ListenAndServe(":81", myRouter))
}

func main() {
	connectDb()

	defer closeDb()

	handleRequest()
}

package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "postgres"
)

var db *sql.DB
var err error

type Task struct {
	Id          int       `json:"Id"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	Duedate     time.Time `json:"Duedate"`
	Status      string    `json:"Status"`
}

func connectDb() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlconn)
	CheckError(err)
}

func closeDb() {
	db.Close()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func createTask(task Task) {
	insertStmt := `insert into task(name, description, duedate, status) values($1, $2, $3, $4)`
	fmt.Println("New Task inserting")
	_, e := db.Exec(insertStmt, task.Name, task.Description, task.Duedate, task.Status)
	CheckError(e)
}

func updateTask(id int64, task Task) {
	insertStmt := `update task set name = $1, description = $2, duedate = $3, status = $4 where id_ = $5`
	fmt.Println("New Task inserting")
	_, e := db.Exec(insertStmt, task.Name, task.Description, task.Duedate, task.Status, id)
	CheckError(e)
}

func delTask(id int64) {
	delStmt := `delete from task where id_ = $1`
	fmt.Println("Deleting Task:", id)
	_, e := db.Exec(delStmt, id)
	CheckError(e)
}

func getTask(id int64) Task {
	var task Task
	rows, err := db.Query(`SELECT id_, name, description, duedate, status FROM task where id_ = $1`, id)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var id_ int
		var name string
		var description string
		var duedate time.Time
		var status string

		err = rows.Scan(&id_, &name, &description, &duedate, &status)
		CheckError(err)

		task = Task{Id: id_, Name: name, Description: description, Duedate: duedate, Status: status}
	}

	CheckError(err)

	return task
}

func getAllTasks() []Task {
	var allTasks = []Task{}
	rows, err := db.Query(`SELECT id_, name, description, duedate, status FROM task`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		var id_ int
		var name string
		var description string
		var duedate time.Time
		var status string

		err = rows.Scan(&id_, &name, &description, &duedate, &status)
		CheckError(err)

		var task Task = Task{Id: id_, Name: name, Description: description, Duedate: duedate, Status: status}

		allTasks = append(allTasks, task)
	}

	CheckError(err)

	return allTasks
}

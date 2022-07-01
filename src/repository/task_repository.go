package repository

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Task struct {
	Id          int
	Description string
	Is_ready    bool
}

func GetTasks() []Task {
	var tasks []Task
	rows, errQuery := db.Query("SELECT id, description, is_ready FROM jira.tasks ORDER BY is_ready ASC, id DESC")
	if errQuery != nil {
		return tasks
	}
	for rows.Next() {
		var task Task
		rows.Scan(&task.Id, &task.Description, &task.Is_ready)
		tasks = append(tasks, task)
	}
	return tasks
}

func AddTask(description string) {
	result, err := db.Exec("INSERT INTO jira.tasks(description, is_ready) VALUES (?, ?)", description, false)
	if err != nil {
		log.Fatal(err.Error())
	}
	id, err := result.LastInsertId()
	_ = id
	if err != nil {
		log.Fatal(err.Error())
	}
	// log.Println(id)
}

func DeleteTask(id int) {
	result, err := db.Exec("DELETE FROM jira.tasks WHERE id = ?", id)
	_ = result
	if err != nil {
		log.Fatal(err.Error())
	}
}

func UpdateTask(id int) {
	result, err := db.Exec("UPDATE jira.tasks SET is_ready = 1 WHERE id = ?", id)
	_ = result
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Connect() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "go-mysql:3306",
		DBName:               "jira",
		AllowNativePasswords: true,
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("error open", err.Error())
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}

package repository

import (
	"log"
)

type Task struct {
	Id          int
	Description string
	Is_ready    bool
	User        User
}

func GetTasks() []Task {
	var tasks []Task
	rows, errQuery := db.Query(`
		SELECT t.id, t.description, t.is_ready, u.id, u.name, u.role
		FROM jira.tasks t
			LEFT JOIN jira.user_tasks ut ON ut.task_id = t.id
			LEFT JOIN jira.users u ON ut.user_id = u.id
		ORDER BY t.is_ready ASC, t.id DESC
	`)
	if errQuery != nil {
		return tasks
	}
	for rows.Next() {
		var task Task
		rows.Scan(&task.Id, &task.Description, &task.Is_ready, &task.User.Id, &task.User.Name, &task.User.Role)
		tasks = append(tasks, task)
	}
	return tasks
}

func AddTask(description string, user_id int) {
	result, err := db.Exec("INSERT INTO jira.tasks(description, is_ready) VALUES (?, ?)", description, false)
	if err != nil {
		log.Fatal(err.Error())
	}
	id, err := result.LastInsertId()
	_ = id
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = db.Exec("INSERT INTO jira.user_tasks(task_id, user_id) VALUES (?, ?)", id, user_id)
	if err != nil {
		log.Fatal(err.Error())
	}
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

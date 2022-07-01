package main

import (
	"html/template"
	"net/http"
	"path"
	taskRepository "popug-jira/src/repository"
	"strconv"
)

type templateData struct {
	Tasks []taskRepository.Task
}

func executeTemplate(w http.ResponseWriter) {
	t, _ := template.ParseFiles("templates/index.html")
	data := templateData{
		Tasks: taskRepository.GetTasks(),
	}
	t.Execute(w, data)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		executeTemplate(w)
	case http.MethodPost:
		taskRepository.AddTask(r.FormValue("description"))
		http.Redirect(w, r, "/tasks", http.StatusFound)
	case http.MethodPut:
		id, _ := strconv.Atoi(path.Base(r.RequestURI))
		taskRepository.UpdateTask(id)
	case http.MethodDelete:
		id, _ := strconv.Atoi(path.Base(r.RequestURI))
		taskRepository.DeleteTask(id)
	}

}

func main() {
	taskRepository.Connect()
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/tasks", httpHandler)
	http.HandleFunc("/tasks/", httpHandler)
	server.ListenAndServe()
}

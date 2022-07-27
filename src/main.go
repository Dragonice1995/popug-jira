package main

import (
	"html/template"
	"net/http"
	"path"
	"popug-jira/src/repository"
	"strconv"
)

type templateData struct {
	Tasks []repository.Task
	Users []repository.User
}

func executeTemplate(w http.ResponseWriter) {
	t, _ := template.ParseFiles("templates/index.html")
	data := templateData{
		Tasks: repository.GetTasks(),
		Users: repository.GetUsersByRole("popug"),
	}
	t.Execute(w, data)
}

func httpHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		executeTemplate(w)
	case http.MethodPost:
		user_id, _ := strconv.Atoi(r.FormValue("user_id"))
		repository.AddTask(r.FormValue("description"), user_id)
		http.Redirect(w, r, "/tasks", http.StatusFound)
	case http.MethodPut:
		id, _ := strconv.Atoi(path.Base(r.RequestURI))
		repository.UpdateTask(id)
	case http.MethodDelete:
		id, _ := strconv.Atoi(path.Base(r.RequestURI))
		repository.DeleteTask(id)
	}

}

func main() {
	repository.Connect()
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/tasks", httpHandler)
	http.HandleFunc("/tasks/", httpHandler)
	server.ListenAndServe()
}

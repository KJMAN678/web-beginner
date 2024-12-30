package main

import (
	"html/template"
	"log"
	"net/http"
)

var todoList []string

func handleTodo(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("ch06/TodoListAddPostRedirectGet/templates/todo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, todoList)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	todo := r.Form.Get("todo")
	todoList = append(todoList, todo)
	// handleTodo(w, r)
	http.Redirect(w, r, "/todo", 303)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("ch06/TodoListAddPostRedirectGet/static"))))
	http.HandleFunc("/todo", handleTodo)
	http.HandleFunc("/add", handleAdd)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start:", err)
	}
}

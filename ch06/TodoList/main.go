package main

import (
	"html/template"
	"log"
	"net/http"
)

var todoList []string

func handleTodo(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("ch06/TodoList/templates/todo.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, todoList)
}

func main() {
	todoList = append(todoList, "顔を洗う", "朝食を食べる", "歯を磨く")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("ch06/TodoList/static"))))
	http.HandleFunc("/todo", handleTodo)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start:", err)
	}
}

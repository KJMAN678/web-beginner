package main

import (
	"html"
	"html/template"
	"net/http"
	"strings"
)

// セッションIDをキー、TODO リストをバリューとして保持するマップ
var todoLists = make(map[string][]string)

// セッションIDに紐付く Todo リストを取得する
func getTodoList(sessionId string) []string {
	todos, ok := todoLists[sessionId]

	if !ok {
		todos = []string{}
		todoLists[sessionId] = todos
	}
	return todos
}

// ToDo リストを返す
func handleTodo(w http.ResponseWriter, r *http.Request) {
	// session.go で設定した、クッキーからセッションIDを取得する関数を呼び出す
	sessionId, err := ensureSession(w, r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	todos := getTodoList(sessionId)

	t, _ := template.ParseFiles("ch06/TodoListSession/templates/todo.html")
	t.Execute(w, todos)
}

// セッション上の ToDo リストにToDo を追加する
func handleAdd(w http.ResponseWriter, r *http.Request) {
	// session.go で設定した、クッキーからセッションIDを取得する関数を呼び出す
	sessionId, err := ensureSession(w, r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	todos := getTodoList(sessionId)

	r.ParseForm()
	todo := strings.TrimSpace(html.EscapeString(r.Form.Get("todo")))
	if todo != "" {
		todoLists[sessionId] = append(todos, todo)
	}
	http.Redirect(w, r, "/todo", 303)
}

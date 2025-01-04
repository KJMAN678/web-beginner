package main

import (
	"html"
	"net/http"
	"strings"
)

type TodoPageData struct {
	UserId   string
	Expires  string
	ToDoList []string
}

func handleTodo(w http.ResponseWriter, r *http.Request) {
	// セッション情報が存在するか、有効性をチェック
	session, err := checkSession(w, r)
	if err != nil {
		return
	}
	// 認証済み/ログイン済みかどうかをチェック
	if !isAuthenticated(w, r, session) {
		return
	}

	// セッションが有効で、かつ、ログイン済みの場合、以下の処理を実行
	// ToDoリストの表示
	pageData := TodoPageData{
		UserId:   session.UserAccount.Id,
		Expires:  session.UserAccount.ExpiresText(),
		ToDoList: session.UserAccount.ToDoList,
	}

	templates.ExecuteTemplate(w, "todo.html", pageData)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	session, err := checkSession(w, r)
	if err != nil {
		return
	}
	if !isAuthenticated(w, r, session) {
		return
	}

	r.ParseForm()
	todo := strings.TrimSpace(html.EscapeString(r.Form.Get("todo")))
	if todo != "" {
		session.UserAccount.ToDoList = append(session.UserAccount.ToDoList, todo)
	}
	http.Redirect(w, r, "/todo", http.StatusSeeOther)
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	session, err := checkSession(w, r)
	if err != nil {
		return
	}

	sessionManager.RevokeSession(w, session.SessionId)
	sessionManager.StartSession(w)

	http.Redirect(w, r, "/todo", http.StatusSeeOther)
}

package main

import (
	"html"
	"log"
	"net/http"
	"strings"
)

type TodoPageData struct {
	UserId   string
	Expires  string
	ToDoList []*ToDoItem
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
		ToDoList: session.UserAccount.ToDoList.Items,
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
		item := session.UserAccount.ToDoList.Append(todo)
		log.Printf("Todo item added. sessionId=%s itemId=%s todo=%s", session.SessionId, item.Id, item.Todo)
	}
	http.Redirect(w, r, "/todo", http.StatusSeeOther)
}

func handleEdit(w http.ResponseWriter, r *http.Request) {
	// POST メソッドによるリクエストであることの確認
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// セッション情報を取得
	session, err := checkSession(w, r)
	logRequest(r, session)
	if err != nil {
		return
	}
	// 認証チェック
	if !isAuthenticated(w, r, session) {
		return
	}

	// POSTパラメーターを解析し、編集対象の ToDo 項目 ID と編集内容を取得する
	r.ParseForm()
	todoId := r.Form.Get("id")
	todo := r.Form.Get("todo")

	// ToDoList 構造体の Update() 関数を使って、Todo 項目を更新
	_, err = session.UserAccount.ToDoList.Update(todoId, todo)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Todo item updated. sessionId=%s itemId=%s todo=%s", session.SessionId, todoId, todo)

	// 成功したことを表す HTTP ステータスコード 200 OK をレスポンスとして返す
	w.WriteHeader(http.StatusOK)
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

package main

import (
	"log"
	"net/http"
)

type LoginPageData struct {
	UserId       string
	ErrorMessage string
}

// ログインに関するリクエスト処理
func handleLogin(w http.ResponseWriter, r *http.Request) {
	// セッションが必ず開始されている状態にする
	session, err := ensureSession(w, r)
	if err != nil {
		return
	}

	switch r.Method {
	// GET リクエスト: ログイン画面の表示
	case http.MethodGet:
		showLogin(w, r, session)
		return

	// POST リクエスト: ログイン処理
	case http.MethodPost:
		login(w, r, session)
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

// ログイン画面を表示する
func showLogin(w http.ResponseWriter, r *http.Request, session *HttpSession) {
	var pageData LoginPageData
	if p, ok := session.PageData.(LoginPageData); ok {
		pageData = p
	} else {
		pageData = LoginPageData{}
	}

	templates.ExecuteTemplate(w, "login.html", pageData)
	session.ClearPageData()
}

// ログイン処理を行う
func login(w http.ResponseWriter, r *http.Request, session *HttpSession) {
	// POST パラメータを取得
	r.ParseForm()
	userId := r.Form.Get("userId")
	password := r.Form.Get("password")

	// 認証処理
	log.Printf("login attempt: %s\n", userId)
	// ユーザーIDとパスワードのチェック
	account, err := accountManager.Authenticate(userId, password)
	if err != nil {
		log.Printf("login failed: %s\n", userId)

		// ログイン画面に表示するエラーメッセージを LoginPageData という構造体に格納してセッションに保存
		// Post-Redirect-Get による画面遷移は、遷移前後でエラーメッセージのような一時的な情報を保持することができない。
		// ブラウザは繊維前後のページに独立して GET リクエストを送信するため、前のページで処理した結果を後続のページに
		// 受け渡すには、セッションを使う
		session.PageData = LoginPageData{
			ErrorMessage: "ユーザーIDまたはパスワードが違います",
		}
		// ログイン画面にリダイレクト
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// ログイン成功
	// 現在のセッションにユーザーの UserAccount インスタンスを紐付ける
	session.UserAccount = account

	log.Printf("login success: %s\n", account.Id)
	http.Redirect(w, r, "/todo", http.StatusSeeOther)

	return
}

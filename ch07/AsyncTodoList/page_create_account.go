package main

import (
	"log"
	"net/http"
)

type CreateUserAccountPageData struct {
	ErrorMessage string
}

// ユーザー作成に関するリクエスト処理
func handleCreateUserAccount(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// GET リクエスト: ユーザー作成画面の表示
	case http.MethodGet:
		session, err := ensureSession(w, r)
		if err != nil {
			return
		}
		showCreateUserAccount(w, session)
		return

	// POST リクエスト: ユーザー作成処理
	case http.MethodPost:
		session, err := checkSession(w, r)
		if err != nil {
			return
		}
		createNewUserAccount(w, r, session)
		return

	default:
		w.WriteHeader((http.StatusMethodNotAllowed))
		return
	}
}

// ユーザー作成画面を表示する
func showCreateUserAccount(w http.ResponseWriter, session *HttpSession) {
	pageData := CreateUserAccountPageData{}

	if p, ok := session.PageData.(CreateUserAccountPageData); ok {
		pageData = p
	}
	templates.ExecuteTemplate(w, "create-user-account.html", pageData)
	session.ClearPageData()
}

// ユーザーを作成する
func createNewUserAccount(w http.ResponseWriter, r *http.Request, session *HttpSession) {
	r.ParseForm()
	userId := r.Form.Get("userId")
	password := MakePassword()

	user, err := accountManager.NewUserAccount(userId, password)
	if err != nil {
		// ユーザーアカウント作成失敗
		pageData := CreateUserAccountPageData{}
		log.Printf("create user failed: userId=%s cause=%v\n", userId, err)
		if err == ErrUserAlreadyExists {
			pageData.ErrorMessage = "すでに使われているユーザーIDです。他のIDを試してください。"
		} else if err == ErrInvalidUserIdFormat {
			pageData.ErrorMessage = "ユーザーIDの形式が違います。"
		} else {
			pageData.ErrorMessage = err.Error()
		}
		session.PageData = pageData
		http.Redirect(w, r, "/create-user-account", http.StatusSeeOther)
		return
	}
	// ユーザーアカウント作成成功
	// リダイレクト先画面で表示するためにユーザー情報をセッションへ格納
	session.PageData = NewUserAccountPageData{
		UserId:   user.Id,
		Password: password,
		Expires:  user.ExpiresText(),
	}
	http.Redirect(w, r, "/new-user-account", http.StatusSeeOther)
	return
}

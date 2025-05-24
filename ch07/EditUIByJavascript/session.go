package main

import (
	"fmt"
	"net/http"
	"time"
)

const CookieNameSessionId = "sessionId"

// ユーザアカウント情報を保持する構造体。
type UserAccount struct {
	// ユーザID
	Id string
	// ハッシュ化されたパスワード
	HashedPassword string
	// アカウントの有効期限
	Expires time.Time
	// ToDoリスト
	ToDoList []string
}

const cookieSessionId = "sessionId"
const sessionIdSecret = 123456789

func ensureSession(w http.ResponseWriter, r *http.Request) (sessionId string, err error) { // <1>
	c, err := r.Cookie(cookieSessionId)
	if err == http.ErrNoCookie { // <2>
		sessionId, err = startSession(w)
		return
	}
	if err == nil { // <3>
		sessionId = c.Value
		if ok, _ := verifySessionId(sessionIdSecret, sessionId); !ok {
			return "", fmt.Errorf("invalid session id")
		}
		return
	}
	return
}

// セッション情報を保持する構造体。
type HttpSession struct {
	// セッションID
	SessionId string
	// セッションの有効期限(時刻)
	Expires time.Time
	// Post-Redirect-Getでの遷移先に表示するデータ
	PageData any
	// ユーザアカウント情報への参照
	UserAccount *UserAccount
	// セッションの有効時間
	validityTime time.Duration
	// Cookieにsecure属性を付与するかどうか
	useSecureCookie bool
}

// 新しいセッション情報を生成する。
func NewHttpSession(sessionId string, validityTime time.Duration, useSecureCookie bool) *HttpSession {
	session := &HttpSession{
		SessionId:       sessionId,
		validityTime:    validityTime,
		PageData:        "",
		useSecureCookie: useSecureCookie,
	}
	session.Extend()
	return session
}

// 有効期限を延長する。
func (s *HttpSession) Extend() {
	s.Expires = time.Now().Add(s.validityTime)
}

// ページデータを削除する。
func (s *HttpSession) ClearPageData() {
	s.PageData = ""
}

// セッションIDをCookieへ書き込む。
func (s HttpSession) SetCookie(w http.ResponseWriter) {
	// 同一リクエスト処理内で既にCookieが書き込まれていた場合は削除する
	w.Header().Del("Set-Cookie")
	cookie := &http.Cookie{
		Name:     CookieNameSessionId,
		Value:    s.SessionId,
		Expires:  s.Expires,
		HttpOnly: true,
		Secure:   s.useSecureCookie,
	}
	http.SetCookie(w, cookie)
}

func startSession(w http.ResponseWriter) (string, error) {
	sessionId := generateSessionId(sessionIdSecret)
	cookie := &http.Cookie{
		Name:     cookieSessionId,
		Value:    sessionId,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
	return sessionId, nil
}

func generateSessionId(secret int) string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func verifySessionId(secret int, sessionId string) (bool, error) {
	// 簡易的な実装
	return true, nil
}

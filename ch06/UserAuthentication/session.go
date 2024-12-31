package main

import (
	"net/http"
	"time"
)

const CookieNameSessionId = "sessionId"

// セッション情報を保持する構造体
type HttpSession struct {
	// セッションID
	SessionId string
	// セッションの有効機縁
	Expires time.Time
	// Post-Redirect-Get での遷移先画面で表示するデータ
	PageData any
	// セッションを利用中のユーザーアカウント
	UserAccount *UserAccount
}

// 新しいセッション情報を生成する
func NewHttpSession(sessionId string, validityTime time.Duration) *HttpSession {
	session := &HttpSession{
		SessionId: sessionId,
		Expires:   time.Now().Add(validityTime),
		PageData:  "",
	}
	return session
}

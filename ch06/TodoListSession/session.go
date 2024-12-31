package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"time"
)

const cookieSessionId = "sessionId"

// セッションが開始されていることを保証する
// セッションが存在しなければ、新しく発行する
// これにより、現在のクライアントに紐づくセッションIDを取得できる
func ensureSession(w http.ResponseWriter, r *http.Request) (string, error) {
	cookie, err := r.Cookie(cookieSessionId)

	// Cookie にセッションIDが入ってない場合は、新しく発行して返す
	if err == http.ErrNoCookie {
		sessionId, err := startSession(w)
		return sessionId, err
	}
	// Cookie にセッションIDが入っている場合は、それを返す
	if err == nil {
		sessionId := cookie.Value
		return sessionId, nil
	}
	return "", nil
}

// (セッションIDが存在しなかったときに) 新たにセッションIDを発行する
func startSession(w http.ResponseWriter) (string, error) {
	// セッションIDを生成
	sessionId, err := makeSessionId()
	if err != nil {
		return "", err
	}

	// セッションIDを Cookie に格納
	cookie := &http.Cookie{
		Name:  cookieSessionId,
		Value: sessionId,
		// Cookie の有効期限 (30分)
		Expires: time.Now().Add(60 * 30 * time.Second),
		// JavaScript からのアクセスを禁止
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	return sessionId, nil
}

// 実際にセッションIDを生成する処理
func makeSessionId() (string, error) {
	// 16バイトのランダム値を生成
	randBytes := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, randBytes); err != nil {
		return "", err
	}
	// Base64 URL エンコードしたものを Session ID として返す
	SessionId := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(randBytes)
	return SessionId, nil
}

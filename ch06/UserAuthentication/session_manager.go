package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	ErrSessionExpired  = errors.New("session expired")
	ErrSessionNotFound = errors.New("session not found")
	ErrInvalidSession  = errors.New("invalid session id")
)

// セッションを管理する構造体
type HttpSessionManager struct {
	// セッションIDをキーとしてセッション情報を保持するマップ
	sessions map[string]*HttpSession
}

func NewHttpSessionManager() *HttpSessionManager {
	mgr := &HttpSessionManager{
		sessions := make(map[string]*HttpSession)
	}
	return mgr
}

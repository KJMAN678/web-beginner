package main

import (
	"golang.org/x/crypto/bcrypt"
	"time"
	_ "time/tzdata"
)

// アカウントを削除する時間 (分)
const UserAccountLimitInMinute = 60

// パスワードの長さ
const PasswordLength = 10

const PasswordChars = "23456789abcdefghijkmnpqrstuvwxyz"

// ユーザーアカウント情報を保持する構造体
type UserAccount struct {
	// ユーザーID
	Id string

	// ハッシュ化されたパスワード
	HashedPassword string

	// アカウントの有効期限
	Expires time.Time

	// Todoリスト
	ToDoList []string
}

// ユーザーアカウント情報を生成する
func NewUserAccount(userId string, plainPassword string, expires time.Time) *UserAccount {
	// bcrypt アルゴリズムでパスワードをハッシュ化する
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	account := &UserAccount{
		Id:             userId,
		HashedPassword: string(hashedPassword),
		Expires:        expires,
		ToDoList:       make([]string, 0, 10),
	}
	return account
}

func (u UserAccount) ExpiresText() string {
	return u.Expires.Format("2006/01/02 15:04:05")
}

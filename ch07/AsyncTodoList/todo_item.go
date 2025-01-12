package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

// ToDo 項目を表す構造体
type ToDoItem struct {
	Id   string `json:"id"`
	Todo string `json:"todo"`
}

// 新しい ToDoItem を作成する
func NewToDoItem(todo string) *ToDoItem {
	id := MakeToDoId(todo)
	return &ToDoItem{
		Id:   id,
		Todo: todo,
	}
}

// ToDo 項目の ID を生成する
// ID は現在時刻と ToDo の文字列を MD5 ハッシュで変換したもの
func MakeToDoId(todo string) string {
	timeBytes := []byte(fmt.Sprintf("%d", time.Now().UnixNano()))
	hasher := md5.New()
	hasher.Write(timeBytes)
	hasher.Write([]byte(todo))
	return hex.EncodeToString(hasher.Sum(nil))
}

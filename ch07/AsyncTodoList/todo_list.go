package main

import (
	"fmt"
	"sync"
)

// ToDoList を保持する構造体
type ToDoList struct {
	lock  sync.Mutex
	Items []*ToDoItem
}

// 新しい ToDoList を生成する
func NewToDoList() *ToDoList {
	list := &ToDoList{
		Items: make([]*ToDoItem, 0, 10),
	}
	return list
}

// ToDo を追加する
func (t *ToDoList) Append(todo string) *ToDoItem {
	t.lock.Lock()
	defer t.lock.Unlock()

	todoItem := NewToDoItem(todo)
	t.Items = append(t.Items, todoItem)
	return todoItem
}

// ToDo 項目を取得する
func (t *ToDoList) Get(id string) (*ToDoItem, error) {
	for _, todo := range t.Items {
		if todo.Id == id {
			return todo, nil
		}
	}
	return nil, fmt.Errorf("todo not found. itemId=%s", id)
}

// ToDo 項目を更新する
func (t *ToDoList) Update(id string, newTodo string) (*ToDoItem, error) {
	t.lock.Lock()
	defer t.lock.Unlock()

	todoItem, err := t.Get(id)
	if err != nil {
		return nil, err
	}

	todoItem.Todo = newTodo
	return todoItem, nil
}

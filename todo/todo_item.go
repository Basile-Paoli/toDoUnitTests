package todo

import (
	"errors"
	"time"
)

type TodoItem struct {
	Name         string
	Content      string
	CreationDate time.Time
}

func NewTodoItem(name string, content string) (TodoItem, error) {
	if len(content) > 1000 {
		return TodoItem{}, errors.New("content must me under 1000 characters")
	}

	item := TodoItem{
		Name:         name,
		Content:      content,
		CreationDate: time.Now(),
	}
	return item, nil
}

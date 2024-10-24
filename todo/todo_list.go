package todo

import (
	"errors"
	"time"
)

type timeProvider interface {
	Now() time.Time
}

type realTimeProvider struct {
}

func (p realTimeProvider) Now() time.Time {
	return time.Now()
}

type ToDoList struct {
	Items        []TodoItem
	LastAddedAt  time.Time
	timeProvider timeProvider
}

func (t *ToDoList) AddItem(name string, content string) error {

	now := t.timeProvider.Now()
	if err := t.assertCanAddItem(name, now); err != nil {
		return err
	}

	item, err := NewTodoItem(name, content)
	if err != nil {
		return err
	}

	t.Items = append(t.Items, item)
	t.LastAddedAt = now

	return nil
}

func (t *ToDoList) assertCanAddItem(name string, now time.Time) error {
	if !t.LastAddedAt.Add(30 * time.Minute).Before(now) {
		return errors.New("cannot add new item yet")
	}

	if t.containsName(name) {
		return errors.New("item with name \"" + name + "\" already exists")
	}

	if len(t.Items) >= 10 {
		return errors.New("todoList cannot contain more than 10 items")
	}

	return nil
}

func (t *ToDoList) containsName(name string) bool {
	for _, item := range t.Items {
		if item.Name == name {
			return true
		}
	}
	return false
}

func newTodoList() *ToDoList {
	return &ToDoList{timeProvider: realTimeProvider{}}
}

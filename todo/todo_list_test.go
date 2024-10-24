package todo

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"strconv"
	"testing"
	"time"
)

type mockTimeProvider struct {
	mock.Mock
}

func (p *mockTimeProvider) Now() time.Time {
	args := p.Called()
	return args.Get(0).(time.Time)
}

func TestAddItemName(t *testing.T) {
	list := newTodoList()
	err := list.AddItem("foo", "test")
	assert.Nil(t, err)

	assert.True(t, list.containsName("foo"))
}

func TestAddItemsAtTheSameTime(t *testing.T) {
	mockTime := &mockTimeProvider{}
	list := &ToDoList{timeProvider: mockTime}

	mockTime.On("Now").Return(time.Now())
	list.AddItem("foo", "Lorem")
	mockTime.AssertNumberOfCalls(t, "Now", 1)

	err := list.AddItem("bar", "test")
	assert.NotNil(t, err)
	mockTime.AssertNumberOfCalls(t, "Now", 2)
}

func TestAddItemsSeparately(t *testing.T) {
	mockTime := &mockTimeProvider{}
	list := &ToDoList{timeProvider: mockTime}

	mockTime.On("Now").Once().Return(time.Now())
	list.AddItem("foo", "Lorem")
	mockTime.AssertNumberOfCalls(t, "Now", 1)

	mockTime.On("Now").Return(time.Now().Add(time.Hour))
	err := list.AddItem("bar", "test")
	assert.Nil(t, err)
}

func TestAddToManyItems(t *testing.T) {

	mockTime := &mockTimeProvider{}
	list := &ToDoList{timeProvider: mockTime}

	for i := 0; i < 10; i++ {
		mockTime.On("Now").Once().Return(time.Now().Add(time.Hour * time.Duration(i)))
		err := list.AddItem("foo"+strconv.Itoa(i), "Lorem ipsum")
		assert.Nil(t, err)
	}

	mockTime.On("Now").Once().Return(time.Now().Add(time.Hour * time.Duration(10)))
	assert.NotNil(t, list.AddItem("11th item", "Lorem ipsum"))
}

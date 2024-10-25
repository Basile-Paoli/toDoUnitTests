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

type mockRepository struct {
	mock.Mock
}

func (r *mockRepository) Save(item TodoItem) error {
	args := r.Called(item)
	return args.Error(0)
}

func newTestTodoList() *ToDoList {
	mockRepo := &mockRepository{}
	mockRepo.On("Save", mock.Anything).Return(nil)
	res := newTodoList()
	res.repository = mockRepo

	return res
}

func TestContainsName(t *testing.T) {
	list := newTestTodoList()
	err := list.AddItem("foo", "test")
	assert.Nil(t, err)

	assert.True(t, list.containsName("foo"))
	assert.False(t, list.containsName("bar"))
}

func TestAddItemsAtTheSameTime(t *testing.T) {
	list := newTestTodoList()

	list.AddItem("foo", "Lorem")

	err := list.AddItem("bar", "test")
	assert.NotNil(t, err)
	assert.Equal(t, "cannot add new item yet", err.Error())
}

func TestAddItemsSeparately(t *testing.T) {
	mockTime := &mockTimeProvider{}
	list := newTestTodoList()
	list.timeProvider = mockTime

	mockTime.On("Now").Once().Return(time.Now())
	list.AddItem("foo", "Lorem")
	mockTime.AssertNumberOfCalls(t, "Now", 1)

	mockTime.On("Now").Return(time.Now().Add(time.Hour))
	err := list.AddItem("bar", "test")
	assert.Nil(t, err)
}

func TestAddTooManyItems(t *testing.T) {

	mockTime := &mockTimeProvider{}
	list := newTestTodoList()
	list.timeProvider = mockTime

	for i := 0; i < 10; i++ {
		mockTime.On("Now").Once().Return(time.Now().Add(time.Hour * time.Duration(i)))
		err := list.AddItem("foo"+strconv.Itoa(i), "Lorem ipsum")
		assert.Nil(t, err)
	}

	mockTime.On("Now").Once().Return(time.Now().Add(time.Hour * 10))
	err := list.AddItem("11th item", "Lorem ipsum")

	assert.NotNil(t, err)
	assert.Equal(t, "todoList cannot contain more than 10 items", err.Error())
}

func TestSave(t *testing.T) {
	list := newTestTodoList()
	mockRepo := &mockRepository{}
	list.repository = mockRepo

	mockRepo.On("Save", mock.Anything).Return(nil)
	list.AddItem("foo", "Lorem ipsum")
	mockRepo.AssertCalled(t, "Save", mock.Anything)
}

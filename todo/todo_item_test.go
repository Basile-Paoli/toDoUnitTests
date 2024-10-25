package todo

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTodoItemContentLength(t *testing.T) {
	asserts := assert.New(t)

	_, err := NewTodoItem("foo", "Lorem Ipsum")
	asserts.Nil(err)

	_, err = NewTodoItem("bar", strings.Repeat("gaga", 500))
	asserts.NotNil(err)
}

package json_markd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Components from https://flaviocopes.com/golang-data-structure-stack/

func initStack() *itemStack {
	var s itemStack
	if s.items == nil {
		s = itemStack{}
		s.new()
	}
	return &s
}

func TestSize(t *testing.T) {
	t.Run("when popping from empty stack", func(t *testing.T) {
		s := initStack()
		t.Run("it should return error", func(t *testing.T) {
			s.push(1)
			s.push(2)
			s.push(3)
			assert.Equal(t, 3, s.size())
		})
	})

}

func TestPush(t *testing.T) {
	t.Run("when popping from empty stack", func(t *testing.T) {
		s := initStack()
		t.Run("it should return error", func(t *testing.T) {
			s.push(1)
			assert.Equal(t, 1, s.size())
			s.push(2)
			assert.Equal(t, 2, s.size())
			s.push(3)
			assert.Equal(t, 3, s.size())
		})
	})
}

func TestTop(t *testing.T) {
	t.Run("when popping from empty stack", func(t *testing.T) {
		s := initStack()
		t.Run("it should return error", func(t *testing.T) {
			s.push(1)
			s.push(2)
			item := (*s.top()).(int)
			assert.Equal(t, 2, item)
		})
	})
}

func TestPop(t *testing.T) {
	t.Run("when popping from empty stack", func(t *testing.T) {
		s := initStack()
		t.Run("it should return error", func(t *testing.T) {
			_, err := s.pop()
			assert.Equal(t, errors.New(".errors.stack_empty"), err)
		})
	})

	t.Run("when popping from a stack", func(t *testing.T) {
		s := initStack()
		t.Run("it should remove one element", func(t *testing.T) {
			s.push(1)
			s.push(2)
			item, _ := s.pop()
			assert.Equal(t, 2, (*item).(int))
		})
	})

}

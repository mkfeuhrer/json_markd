package json_markd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Components from https://flaviocopes.com/golang-data-structure-stack/

func initStack() *ItemStack {
	var s ItemStack
	if s.items == nil {
		s = ItemStack{}
		s.New()
	}
	return &s
}

func TestSize(t *testing.T) {
	t.Run("when popping from empty stack", func(t *testing.T) {
		s := initStack()
		t.Run("it should return error", func(t *testing.T) {
			s.Push(1)
			s.Push(2)
			s.Push(3)
			assert.Equal(t, 3, s.Size())
		})
	})

}

func TestPush(t *testing.T) {
	t.Run("when popping from empty stack", func(t *testing.T) {
		s := initStack()
		t.Run("it should return error", func(t *testing.T) {
			s.Push(1)
			assert.Equal(t, 1, s.Size())
			s.Push(2)
			assert.Equal(t, 2, s.Size())
			s.Push(3)
			assert.Equal(t, 3, s.Size())
		})
	})
}

func TestTop(t *testing.T) {
	t.Run("when popping from empty stack", func(t *testing.T) {
		s := initStack()
		t.Run("it should return error", func(t *testing.T) {
			s.Push(1)
			s.Push(2)
			item := (*s.Top()).(int)
			assert.Equal(t, 2, item)
		})
	})
}

func TestPop(t *testing.T) {
	t.Run("when popping from empty stack", func(t *testing.T) {
		s := initStack()
		t.Run("it should return error", func(t *testing.T) {
			_, err := s.Pop()
			assert.Equal(t, errors.New(".errors.stack_empty"), err)
		})
	})

	t.Run("when popping from a stack", func(t *testing.T) {
		s := initStack()
		t.Run("it should remove one element", func(t *testing.T) {
			s.Push(1)
			s.Push(2)
			item, _ := s.Pop()
			assert.Equal(t, 2, (*item).(int))
		})
	})

}

package stack

import (
	"errors"
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Components from https://flaviocopes.com/golang-data-structure-stack/

// Item the type of the stack
type Item generic.Type

// ItemStack the stack of Items
type ItemStack struct {
	items []Item
	lock  sync.RWMutex
}

// New creates a new ItemStack
func (s *ItemStack) New() *ItemStack {
	s.items = []Item{}
	return s
}

// Push adds an Item to the top of the stack
func (s *ItemStack) Push(t Item) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

func (s *ItemStack) Size() int {
	return len(s.items)
}

func (s *ItemStack) Top() *Item {
	item := s.items[len(s.items)-1]
	return &item
}

// Pop removes an Item from the top of the stack
func (s *ItemStack) Pop() (*Item, error) {
	s.lock.Lock()
	if s.Size() == 0 {
		return nil, errors.New(".errors.stack_empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.lock.Unlock()
	return &item, nil
}

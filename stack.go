package json_markd

import (
	"errors"
	"sync"

	"github.com/cheekybits/genny/generic"
)

// Components from https://flaviocopes.com/golang-data-structure-stack/

// Item the type of the stack
type item generic.Type

// ItemStack the stack of Items
type itemStack struct {
	items []item
	lock  sync.RWMutex
}

// New creates a new ItemStack
func (s *itemStack) new() *itemStack {
	s.items = []item{}
	return s
}

// Push adds an Item to the top of the stack
func (s *itemStack) push(t item) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

func (s *itemStack) size() int {
	return len(s.items)
}

func (s *itemStack) top() *item {
	item := s.items[len(s.items)-1]
	return &item
}

// Pop removes an Item from the top of the stack
func (s *itemStack) pop() (*item, error) {
	s.lock.Lock()
	if s.size() == 0 {
		return nil, errors.New(".errors.stack_empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.lock.Unlock()
	return &item, nil
}

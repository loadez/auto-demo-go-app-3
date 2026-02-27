package todo

import (
	"sync"
	"sync/atomic"
)

type Item struct {
	ID    int
	Title string
	Done  bool
}

type Store struct {
	mu    sync.RWMutex
	items map[int]*Item
	seq   atomic.Int64
}

func NewStore() *Store {
	return &Store{items: make(map[int]*Item)}
}

func (s *Store) Add(title string) *Item {
	id := int(s.seq.Add(1))
	item := &Item{ID: id, Title: title}
	s.mu.Lock()
	s.items[id] = item
	s.mu.Unlock()
	return item
}

func (s *Store) All() []*Item {
	s.mu.RLock()
	defer s.mu.RUnlock()

	items := make([]*Item, 0, len(s.items))
	for _, item := range s.items {
		items = append(items, item)
	}
	return items
}

func (s *Store) Get(id int) (*Item, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	item, ok := s.items[id]
	return item, ok
}

func (s *Store) Toggle(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	item, ok := s.items[id]
	if !ok {
		return false
	}
	item.Done = !item.Done
	return true
}

func (s *Store) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.items[id]; !ok {
		return false
	}
	delete(s.items, id)
	return true
}

func (s *Store) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.items)
}

func (s *Store) Pending() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	count := 0
	for _, item := range s.items {
		if !item.Done {
			count++
		}
	}
	return count
}

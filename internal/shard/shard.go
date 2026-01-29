package shard

import (
	"errors"
	"sync"
	"time"

	"container/list"
)

type items struct {
	key       string
	value     any
	expiresAt time.Time
}

type Shard struct {
	mu       sync.Mutex
	lru      *list.List
	items    map[string]*list.Element
	capacity int
}

func New(capacity int) *Shard {
	return &Shard{
		lru:      list.New(),
		items:    make(map[string]*list.Element),
		capacity: capacity,
	}
}

func (s *Shard) Get(key string) (interface{}, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	el, ok := s.items[key]
	if !ok {
		return nil, errors.New("Not found")
	}
	expire := el.Value.(*items).expiresAt
	if time.Now().After(expire) {
		s.lru.Remove(el)
		delete(s.items, key)
		return nil, errors.New("Time expired")
	}
	if ok {
		s.lru.MoveToFront(el)
		return el.Value.(*items).value, nil
	}

	return nil, errors.New("Not found")
}

func (s *Shard) Set(key string, value any, ttl time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if el, _ := s.items[key]; el != nil {
		s.lru.MoveToFront(el)
		el.Value.(*items).value = value
		el.Value.(*items).expiresAt = time.Now().Add(ttl)
		return
	}

	el := s.lru.PushFront(&items{key: key, value: value, expiresAt: time.Now().Add(ttl)}) //Add it to a list
	s.items[key] = el                                                                     //then we link it to the map

	if s.lru.Len() > s.capacity {
		old := s.lru.Back()
		if old != nil {
			s.lru.Remove(old)
			kv := old.Value.(*items)
			delete(s.items, kv.key)
		}
	}

}

func (s *Shard) Delete(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	el, ok := s.items[key]
	if !ok {
		return errors.New("Not found")
	}

	s.lru.Remove(el)
	//kv := el.Value.(*items)
	delete(s.items, key)

	return nil
}

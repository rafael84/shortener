package persistence

import "sync"

func NewMemory() *Memory {
	return &Memory{items: make(map[string]string)}
}

type Memory struct {
	mu    sync.Mutex
	items map[string]string
}

func (m *Memory) Set(alias, url string) error {
	m.mu.Lock()
	m.items[alias] = url
	m.mu.Unlock()
	return nil
}

func (m *Memory) Get(alias string) (url string, found bool) {
	m.mu.Lock()
	url, found = m.items[alias]
	m.mu.Unlock()
	return
}

func (m *Memory) Count() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.items)
}

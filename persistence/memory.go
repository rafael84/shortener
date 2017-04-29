package persistence

func NewMemory() Memory {
	return make(Memory)
}

type Memory map[string]string

func (m Memory) Set(alias, url string) error {
	m[alias] = url
	return nil
}

func (m Memory) Get(alias string) (url string, found bool) {
	url, found = m[alias]
	return
}

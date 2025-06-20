package di

import "sync"

var container = make(map[string]any)
var mu sync.RWMutex

func Provide[T any](key string, service T) {
	mu.Lock()
	defer mu.Unlock()
	container[key] = service
}

func Get[T any](key string) T {
	mu.RLock()
	defer mu.RUnlock()
	if val, ok := container[key]; ok {
		return val.(T)
	}

	panic("dependency not found: " + key)
}

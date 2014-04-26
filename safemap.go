package safemap

import (
	"sync"
)

type SafeMap struct {
	mu      sync.RWMutex
	objects map[string]interface{} // map of objects
}

// NewSafeMap returns instance of SafeMap
func New() *SafeMap {
	return &SafeMap{objects: make(map[string]interface{})}
}

// GetObject returns the object for the given key is one exists along with the boolean indicating that it was found or not
func (m *SafeMap) GetObject(key string) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.objects[key]
	return val, ok
}

// SetObject sets the given object for the given key
func (m *SafeMap) SetObject(key string, obj interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.objects[key] = obj
	return nil
}

// RemoveObject removes the object for the given key
func (m *SafeMap) RemoveObject(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.objects, key)
	return nil
}

// Values returns all the objects in the underlying map in a slice
func (m *SafeMap) Values() []interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	values := []interface{}{}
	for _, v := range m.objects {
		values = append(values, v)
	}
	return values
}

// Keys returns all keys for the map in a slice
func (m *SafeMap) Keys() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	keys := []string{}
	for k, _ := range m.objects {
		keys = append(keys, k)
	}
	return keys
}

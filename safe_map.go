package map_manager

import (
	"sync"
)

type SafeMap struct {
	mu      sync.RWMutex
	objects map[string]interface{} // map of objects
}

// NewMapManager - instantiates a new MapManager
func NewSafeMap() *SafeMap {
	return &SafeMap{objects: make(map[string]interface{})}
}

/*
Manager run loop
*/

/*
func (m *MapManager) Start() {
	m.running = true
	for {
		select {
		case qSt := <-m.queryChan:
			// operate on the object that matches the query
			parts := strings.Split(qSt, "|")

			switch parts[0] {
			case "get":
				w, ok := m.objects[parts[1]]
				if ok {
					m.recvObjChan <- w
				} else {
					m.recvObjChan <- nil
				}
			case "values":
				values := []interface{}{}
				for _, v := range m.objects {
					values = append(values, v)
				}
				m.recvObjChan <- values
			case "keys":
				keys := []interface{}{}
				for k, _ := range m.objects {
					keys = append(keys, k)
				}
				m.recvObjChan <- keys
			case "remove":
				_, ok := m.objects[parts[1]]
				if ok {
					delete(m.objects, parts[1])
				}
				m.respChan <- SUCCESS
			}
		case msg := <-m.sendObjChan:
			obj := msg.(ObjMessage)
			m.objects[obj.key] = obj.value
			m.respChan <- SUCCESS
		}
	}
}*/

func (m *SafeMap) GetObject(key string) (interface{}, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.objects[key]
	return val, ok
}

func (m *SafeMap) SetObject(key string, obj interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.objects[key] = obj
	return nil
}

func (m *SafeMap) RemoveObject(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.objects, key)
	return nil
}

func (m *SafeMap) Values() []interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	values := []interface{}{}
	for _, v := range m.objects {
		values = append(values, v)
	}
	return values
}

func (m *SafeMap) Keys() []interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	keys := []interface{}{}
	for k, _ := range m.objects {
		keys = append(keys, k)
	}
	return keys
}

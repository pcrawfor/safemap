/*
A channel based access object for shared use of a map of arbitrary objects
*/
package processor

import (
	"errors"
	"strings"
)

const SUCCESS = "success"

/*
Type definition
*/
type MapManager struct {
	objects     map[string]interface{} // map of objects
	sendObjChan chan interface{}
	recvObjChan chan interface{}
	queryChan   chan string
	respChan    chan string
}

type ObjMessage struct {
	value interface{}
	key   string
}

func NewMapManager() *MapManager {
	return &MapManager{sendObjChan: make(chan interface{}), recvObjChan: make(chan interface{}), respChan: make(chan string), queryChan: make(chan string), objects: make(map[string]interface{})}
}

/*
Manager run loop
*/
func (m *MapManager) Start() {
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
}

func (m *MapManager) GetObject(key string) (interface{}, bool) {
	m.queryChan <- "get:" + key
	msg := <-m.recvObjChan
	return msg, (nil != msg)
}

func (m *MapManager) SetObject(key string, obj interface{}) error {
	m.sendObjChan <- ObjMessage{value: obj, key: key}
	s := <-m.respChan
	if s == SUCCESS {
		return nil
	}
	return errors.New("Error adding obj")
}

func (m *MapManager) RemoveObject(key string) error {
	m.queryChan <- "remove:" + key
	s := <-m.respChan
	if s == SUCCESS {
		return nil
	}
	return errors.New("Error removing obj")
}

func (m *MapManager) Values() []interface{} {
	m.queryChan <- "values:_"
	vals := <-m.recvObjChan
	return vals.([]interface{})
}

func (m *MapManager) Keys() []interface{} {
	m.queryChan <- "keys:_"
	keys := <-m.recvObjChan
	return keys.([]interface{})
}

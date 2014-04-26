safemap
===========

A simple wrapper around a map keyed by strings and containing arbitrary objects (interface{}) that is goroutine safe to work with.

    PACKAGE DOCUMENTATION

    package safemap
      import "github.com/pcrawfor/safemap"


    TYPES

    type SafeMap struct {
      // contains filtered or unexported fields
    }


    func New() *SafeMap
      NewSafeMap returns instance of SafeMap


    func (m *SafeMap) GetObject(key string) (interface{}, bool)
      GetObject returns the object for the given key if one exists along with
      the boolean indicating that it was found or not

    func (m *SafeMap) Keys() []string
      Keys returns all keys for the map in a slice

    func (m *SafeMap) RemoveObject(key string) error
      RemoveObject removes the object for the given key

    func (m *SafeMap) SetObject(key string, obj interface{}) error
      SetObject sets the given object for the given key

    func (m *SafeMap) Values() []interface{}
      Values returns all the objects in the underlying map in a slice

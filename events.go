package events

import (
	"errors"
	"sync"
)

var ErrMissingKey = errors.New("key not found")

type (
	Map map[string]interface{}

	Handler func(target interface{}, data Map) error

	Events struct {
		hm  map[string][]Handler
		rwm *sync.RWMutex
	}
)

func NewEventsManager() *Events {
	return &Events{
		hm: make(map[string][]Handler),
		rwm: new(sync.RWMutex),
	}
}

func (e *Events) Subscribe(key string, handler Handler)  {
	e.rwm.Lock()
	defer e.rwm.Unlock()

	if _, ok := e.hm[key]; !ok {
		e.hm[key] = []Handler{}
	}

	e.hm[key] = append(e.hm[key], handler)
}

func (e *Events) Dispatch(key string, target interface{}, data Map) (Map, error) {
	e.rwm.RLock()
	defer e.rwm.RUnlock()

	if _, ok := e.hm[key]; !ok {
		return nil, ErrMissingKey
	}

	for _, handler := range e.hm[key] {
		if err := handler(target, data); err != nil {
			return nil, err
		}
	}

	return data, nil
}

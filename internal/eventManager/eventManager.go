package eventManager

import "sync"

type EventType string

const (
	ConfigUpdated EventType = "configUpdated"
)

type EventHandler func()

type EventManager struct {
	listeners map[EventType][]EventHandler
	lock      sync.Mutex
}

func NewEventManager() *EventManager {
	return &EventManager{
		listeners: make(map[EventType][]EventHandler),
	}
}

func (m *EventManager) Subscribe(eventType EventType, handler EventHandler) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.listeners[eventType] = append(m.listeners[eventType], handler)
}

func (m *EventManager) Emit(eventType EventType) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if handlers, exists := m.listeners[eventType]; exists {
		for _, handler := range handlers {
			go handler()
		}
	}
}

package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	handlers map[string][]IEventHandler
}

func (ed *EventDispatcher) Register(eventName string, eventHandler IEventHandler) error {
	if _, ok := ed.handlers[eventName]; ok {
		for _, handler := range ed.handlers[eventName] {
			if handler == eventHandler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], eventHandler)
	return nil
}

func (ed *EventDispatcher) Clear() {
	ed.handlers = make(map[string][]IEventHandler)
}

func (ed *EventDispatcher) Has(eventName string, handler IEventHandler) bool {
	if _, ok := ed.handlers[eventName]; ok {
		for _, h := range ed.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) Dispatch(event IEvent) error {
	if handlers, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

func (ed *EventDispatcher) Remove(eventName string, handler IEventHandler) error {
	if _, ok := ed.handlers[eventName]; ok {
		for i, h := range ed.handlers[eventName] {
			if h == handler {
				ed.handlers[eventName] = append(ed.handlers[eventName][:i], ed.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}
	return nil
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]IEventHandler),
	}
}

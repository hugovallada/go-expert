package events

import (
	"sync"
	"time"
)

type IEvent interface {
	GetName() string        // nome do evento
	GetDateTime() time.Time // hora do evento
	GetPayload() any        // payload do evento
}

type IEventHandler interface {
	Handle(event IEvent, wg *sync.WaitGroup) // método q executa a operação
}

// Gerenciador de Eventos
type IEventDispatcher interface {
	Register(eventName string, handler IEventHandler) error // Método q registra o evento
	Dispatch(event IEvent) error                            // Método que executa o evento e os handlers sejam executados
	Remove(eventName string, handler IEventHandler) error   // Remove um evento da fila
	Has(eventName string, handler IEventHandler) bool       // Verifica se tem um eventName para o handler
	Clear()                                                 // Limpa o event dispatcher, matando todos os eventos que estão registrados
}

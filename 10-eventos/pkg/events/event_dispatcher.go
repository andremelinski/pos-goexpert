package events

import "errors"

// 1 evento pode ter diversos handlers registrados
type EventDispatcher struct{
	handlers map[string][]EventHandlerInterface
}
/*
{eventName: []}
*/

func NewEventDispatcher() *EventDispatcher{
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	// nome do evento ja existe
	if eventHandlerInterfaceArr, ok := ed.handlers[eventName]; ok {
		// verifica se o handler ja existe
		for _, h := range eventHandlerInterfaceArr {
			if h == handler {
				return errors.New("handlers already registered")
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)

	return nil
}

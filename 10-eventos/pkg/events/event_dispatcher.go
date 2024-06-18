package events

import (
	"errors"
	"fmt"
	"sync"
)

// 1 evento pode ter diversos handlers registrados
type EventDispatcher struct{
	handlers map[string][]EventHandlerInterface
}
/*
{
	handlers: {
		eventName: [&{111111}, &{2}],
	}
}
*/

func NewEventDispatcher() *EventDispatcher{
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ed *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	// nome do evento ja existe
	//handler e h  possuem valores = &{11111111}. Compara se os objetos com a info TestEventHandler sao iguais
	// se nao for, adiciona no arr e fica: [0xc000110d00], 
	// se nao, fica:&{map[eventName:[0xc00009f200 0xc00009f208]]} ou &{map[eventName:[0xc000119200] eventName2:[0xc000119208]]}
	
	if eventHandlerInterfaceArr, ok := ed.handlers[eventName]; ok {
		fmt.Println(eventHandlerInterfaceArr)
		for _, h := range eventHandlerInterfaceArr {
			fmt.Println(h)
			if h == handler {
				return errors.New("handlers already registered")
			}
		}
	}
	ed.handlers[eventName] = append(ed.handlers[eventName], handler)
	// fmt.Println(ed)

	return nil
}

func (ed *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	// nome do evento ja existe
	if eventHandlerInterfaceArr, ok := ed.handlers[eventName]; ok {
		// verifica se o handler ja existe
		for _, h := range eventHandlerInterfaceArr {
			if h == handler {
				return true
			}
		}
	}
	return false
}

func (ed *EventDispatcher) Dispatch(event EventInterface) error {
	if eventHandlerInterfaceArr, ok := ed.handlers[event.GetName()]; ok {
		wg := &sync.WaitGroup{}
		wg.Add(len(eventHandlerInterfaceArr))
		for _, handler := range eventHandlerInterfaceArr {
			go handler.Handle(event, wg) 
		}
		wg.Wait()
	}
	return nil
}

func (ed *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	// nome do evento ja existe
	if eventHandlerInterfaceArr, ok := ed.handlers[eventName]; ok {
		// verifica se o handler ja existe
		for i, h := range eventHandlerInterfaceArr {
			if h == handler {
				eventHandlerInterfaceArr = append(eventHandlerInterfaceArr[:i], eventHandlerInterfaceArr[i+1:]...)
                return nil
			}
		}
	}
	return errors.New("handler not registered")
}

func (ed *EventDispatcher) Clear() {
	// nome do evento ja existe
	ed.handlers = make(map[string][]EventHandlerInterface)
}

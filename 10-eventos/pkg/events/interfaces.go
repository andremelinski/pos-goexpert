package events

import "time"

// Evento -> carrega Dados
type EventInterface interface{
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{}
}

//  operacoes que sao executados quando evento (Handle) eh chamado
// Handle executa a operacao a partir dos eventos, os quais entao no EventInterface 
type EventHandlerInterface interface{
	Handle(EventInterface) 
}

// Gerenciador dos eventos -> 
// registrar os eventos e sua operacoes; 
// Despachar no evento para que suas operacoes sejam executadas;

type EventDispatcherInterface interface{
	// registra um novo evento. Quando esse evento for executado, executa o handler pra ele
	Register(eventName string, handler EventHandlerInterface) error
	// fez com que os handlers sejam executados
	Dispatch(event EventInterface) error
	// remove evento do event Dispatcher
	Remove(eventName string, handler EventHandlerInterface) error
	// verifica se tem ou nao o evento com esse handler
	Has(eventName string, handler EventHandlerInterface) bool
	// limpa o event dispatcher
	Clear() error
	
}

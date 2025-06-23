package event

// EventInterface define o contrato para qualquer evento que ocorre na aplicação.
type EventInterface interface {
	GetName() string
	GetPayload() interface{}
}

// EventHandlerInterface define o contrato para qualquer componente que lida com um evento.
type EventHandlerInterface interface {
	Handle(event EventInterface, ch chan<- error)
}

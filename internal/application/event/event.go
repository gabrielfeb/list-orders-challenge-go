package event

// EventInterface defines the contract for any event
type EventInterface interface {
	GetName() string
	GetPayload() interface{}
}

// EventHandlerInterface defines the contract for any component that handles an event
type EventHandlerInterface interface {
	Handle(event EventInterface, ch chan<- error)
}

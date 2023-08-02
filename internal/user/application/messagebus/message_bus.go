package messagebus

import "rest-api/internal/user/application/events"

type MessageBus struct {
	handlers []UserEventHandler
}

func NewMessageBus() *MessageBus {
	return &MessageBus{}
}

func (mb *MessageBus) Subscribe(handler UserEventHandler) {
	mb.handlers = append(mb.handlers, handler)
}

func (mb *MessageBus) PublishUserCreatedEvent(event *events.UserCreatedEvent) {
	for _, handler := range mb.handlers {
		handler.HandleUserCreatedEvent(event)
	}
}

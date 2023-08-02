package messagebus

import "rest-api/internal/user/application/events"

type UserEventHandler interface {
	HandleUserCreatedEvent(event *events.UserCreatedEvent)
}

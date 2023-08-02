package events

type UserCreatedEvent struct {
	UserID int
	Name   string
	Email  string
}

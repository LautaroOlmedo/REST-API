package customer

import (
	"errors"
	"github.com/google/uuid"
	"rest-api/entity"
)

var (
	ErrInvalidPerson = errors.New("Required name, email and password")
)

type Customer struct {
	person   *entity.Person
	products []*entity.Item
}

func NewCustomer(name, email, password string) (Customer, error) {
	if name == "" || password == "" {
		return Customer{}, ErrInvalidPerson
	} else if email == "" || email == " " {
		return Customer{}, ErrInvalidPerson
	}
	person := &entity.Person{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: password,
	}
	return Customer{
		person:   person,
		products: make([]*entity.Item, 0),
	}, nil
}

func getId(c *Customer) uuid.UUID {
	return c.person.ID
}

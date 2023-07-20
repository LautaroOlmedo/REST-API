package customer

import (
	"github.com/google/uuid"
)

var ()

type Repository interface {
	getAll([]Customer)
	getByID(uuid uuid.UUID) (Customer, error)
}

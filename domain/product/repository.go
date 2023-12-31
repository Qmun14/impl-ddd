package product

import (
	"errors"

	"github.com/Qmun14/ddd-go/aggregate"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound    = errors.New("No such Product")
	ErrProductAlredyExist = errors.New("Product is already exist")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}

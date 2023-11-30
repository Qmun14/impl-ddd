package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound    = errors.New("No such Product")
	ErrProductAlredyExist = errors.New("Product is already exist")
)

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByID(id uuid.UUID) (Product, error)
	Add(product Product) error
	Update(product Product) error
	Delete(id uuid.UUID) error
}

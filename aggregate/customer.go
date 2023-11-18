// Package aggregate menampung semua aggregate yangg menggabungkan banyak entitas menjadi sebuah object yang utuh (full object)
package aggregate

import (
	"errors"

	"github.com/Qmun14/ddd-go/entity"
	"github.com/Qmun14/ddd-go/valueobject"
	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("a customer has to a valid name")
)

type Customer struct {
	// person adalah root entities dari Customer
	// yang berarti person.ID adalah pengidentifikasi utama untuk customer
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

// NewCustomer adalah sebuah factory (design-pattern) untuk membuat sebuah aggregat Customer baru dan akan memvalidasi bahwa field namanya tidak kosong
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

func (c *Customer) GetName() string {
	return c.person.Name
}

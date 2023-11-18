package services

import (
	"testing"

	"github.com/Qmun14/ddd-go/aggregate"
	"github.com/google/uuid"
)

func init_Products(t *testing.T) []aggregate.Product {
	coffe, err := aggregate.NewProduct("Coffe", "Healthy Baverage", 1.99)
	if err != nil {
		t.Fatal(err)
	}
	potato, err := aggregate.NewProduct("Fried Potatoes", "Delicious snack", 1.45)
	if err != nil {
		t.Fatal(err)
	}

	peanuts, err := aggregate.NewProduct("Peanuts", "Healthy snack", 1.15)
	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{
		coffe, potato, peanuts,
	}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_Products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("Ma'mun")

	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)

	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)

	if err != nil {
		t.Error(err)
	}
}

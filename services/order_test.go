package services

import (
	"testing"

	"github.com/Qmun14/kedai/domain/customer"
	"github.com/Qmun14/kedai/domain/product"
	"github.com/google/uuid"
)

func init_Products(t *testing.T) []product.Product {
	coffe, err := product.NewProduct("Coffe", "Healthy Baverage", 1.99)
	if err != nil {
		t.Fatal(err)
	}
	potato, err := product.NewProduct("Fried Potatoes", "Delicious snack", 1.45)
	if err != nil {
		t.Fatal(err)
	}

	peanuts, err := product.NewProduct("Peanuts", "Healthy snack", 1.15)
	if err != nil {
		t.Fatal(err)
	}

	return []product.Product{
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

	cust, err := customer.NewCustomer("Ma'mun")

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

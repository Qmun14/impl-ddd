package services

import (
	"context"
	"testing"

	"github.com/Qmun14/kedai/domain/customer"
	"github.com/google/uuid"
)

func Test_Kedai(t *testing.T) {
	products := init_Products(t)

	connStr := "root:root@tcp(localhost:3306)/ddd"

	os, err := NewOrderService(
		// WithMemoryCustomerRepository(),
		WithSqlCustomerRepository(context.Background(), connStr),
		WithMemoryProductRepository(products),
	)

	kedai, err := NewKedai(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	cust, err := customer.NewCustomer("Ma'mun")
	if err != nil {
		t.Fatal(err)
	}

	if err = os.customers.Add(cust); err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = kedai.Order(cust.GetID(), order)
	if err != nil {
		t.Fatal(err)
	}

}

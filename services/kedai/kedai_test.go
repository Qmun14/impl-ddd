package kedai

import (
	"context"
	"testing"

	"github.com/Qmun14/kedai/domain/product"
	"github.com/Qmun14/kedai/services/order"
	"github.com/google/uuid"
)

func Test_Kedai(t *testing.T) {
	products := init_Products(t)

	connStr := "root:root@tcp(localhost:3306)/ddd"

	os, err := order.NewOrderService(
		// WithMemoryCustomerRepository(),
		order.WithSqlCustomerRepository(context.Background(), connStr),
		order.WithMemoryProductRepository(products),
	)

	kedai, err := NewKedai(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	uid, err := os.AddCustomer("Ma'mun")
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = kedai.Order(uid, order)
	if err != nil {
		t.Fatal(err)
	}

}

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

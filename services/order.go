package services

import (
	"context"
	"log"

	"github.com/Qmun14/kedai/domain/customer"
	"github.com/Qmun14/kedai/domain/customer/memory"
	"github.com/Qmun14/kedai/domain/customer/mysql"
	"github.com/Qmun14/kedai/domain/product"
	prodmem "github.com/Qmun14/kedai/domain/product/memory"
	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository

	// billing billing.Service
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	// Todo: Loop ke semua config dan menerapkannya
	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

// * WithCustomerRepository akan menerapkan sebuah customer repository ke OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// Todo: akan me-return fungsi yg cocok dengan orderconfiguration alias
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithSqlCustomerRepository(ctx context.Context, connStr string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mysql.New(ctx, connStr)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}

}

func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()

		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	// Todo: Fetch the customer
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}
	// Todo: Get each Product,
	var products []product.Product
	var total float64

	for _, id := range productsIDs {
		p, err := o.products.GetByID(id)

		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))
	return total, nil
}

package kedai

import (
	"log"

	"github.com/Qmun14/kedai/services/order"
	"github.com/google/uuid"
)

type KedaiConfiguration func(os *Kedai) error

type Kedai struct {
	// orderservice untuk menerima pesanan
	OrderService *order.OrderService

	// BillingService
	BillingService interface{}
}

func NewKedai(cfgs ...KedaiConfiguration) (*Kedai, error) {
	k := &Kedai{}

	for _, cfg := range cfgs {
		if err := cfg(k); err != nil {
			return nil, err
		}
	}

	return k, nil
}

func WithOrderService(os *order.OrderService) KedaiConfiguration {
	return func(k *Kedai) error {
		k.OrderService = os
		return nil
	}
}

func (k *Kedai) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := k.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}

	// Todo: coba buat implementasi mysql repository customer
	log.Printf("\nBill the customer: %0.0f\n", price)
	// Todo: next coba buat Implementasi Billing Service
	return nil
}

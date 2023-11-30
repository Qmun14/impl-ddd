package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/Qmun14/kedai/domain/customer"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type MysqlRepository struct {
	db *sql.DB
}

// mysqlCustomer adalah  internal type yang digunakan untuk menyimpan suatu Customercustomer di dalam Repository ini
// sehingga implementasi ini tidak boleh memiliki coupling apapun ke aggregat
type mysqlCustomer struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func NewFromCustomer(c customer.Customer) mysqlCustomer {
	return mysqlCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

func (m mysqlCustomer) Tocustomer() customer.Customer {
	c := customer.Customer{}

	c.SetID(m.ID)
	c.SetName(m.Name)
	return c
}

func New(ctx context.Context, conncetionString string) (*MysqlRepository, error) {
	client, err := sql.Open("mysql", conncetionString)
	if err != nil {
		return nil, err
	}

	// db := client.Ping()

	// return client, nil
	return &MysqlRepository{
		db: client,
	}, nil
}

func (mr *MysqlRepository) Get(id uuid.UUID) (customer.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	SQL := "select id, name from customers where id = ?"
	result, err := mr.db.QueryContext(ctx, SQL, id)
	if err != nil {
		return customer.Customer{}, err
	}
	var c mysqlCustomer

	if result.Next() {
		err := result.Scan(&c.ID, &c.Name)
		if err != nil {
			return c.Tocustomer(), err
		}
	}
	return customer.Customer{}, nil
}

func (mr *MysqlRepository) Add(c customer.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	internal := NewFromCustomer(c)
	SQL := "insert into customers(id, name) values(?,?)"

	_, err := mr.db.ExecContext(ctx, SQL, internal.ID, internal.Name)
	if err != nil {
		return err
	}
	return nil
}

func (mr *MysqlRepository) Update(c customer.Customer) error {
	panic("In Implmented")
}

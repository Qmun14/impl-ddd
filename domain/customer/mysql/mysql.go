package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/Qmun14/ddd-go/aggregate"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type MysqlRepository struct {
	db *sql.DB
}

// mysqlCustomer adalah  internal type yang digunakan untuk menyimpan suatu CustomerAggregate di dalam Repository ini
// sehingga implementasi ini tidak boleh memiliki coupling apapun ke aggregat
type mysqlCustomer struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func NewFromCustomer(c aggregate.Customer) mysqlCustomer {
	return mysqlCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

func (m mysqlCustomer) ToAggregate() aggregate.Customer {
	c := aggregate.Customer{}

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

func (mr *MysqlRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	SQL := "select id, name from customers where id = ?"
	result, err := mr.db.QueryContext(ctx, SQL, id)
	if err != nil {
		return aggregate.Customer{}, err
	}
	var c mysqlCustomer

	if result.Next() {
		err := result.Scan(&c.ID, &c.Name)
		if err != nil {
			return c.ToAggregate(), err
		}
	}
	return aggregate.Customer{}, nil
}

func (mr *MysqlRepository) Add(c aggregate.Customer) error {
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

func (mr *MysqlRepository) Update(c aggregate.Customer) error {
	panic("In Implmented")
}

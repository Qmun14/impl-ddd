package memory

import (
	"errors"
	"testing"

	"github.com/Qmun14/ddd-go/aggregate"
	"github.com/Qmun14/ddd-go/domain/customer"
	"github.com/google/uuid"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	c, err := aggregate.NewCustomer("Ma'mun")
	if err != nil {
		t.Fatal(err)
	}

	id := c.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{
			id: c,
		},
	}

	testCases := []testCase{
		{
			name:        "no customer by id",
			id:          uuid.MustParse("83064af3-bb81-4514-a6d4-afba340825cd"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			_, err := repo.Get(tC.id)
			if !errors.Is(err, tC.expectedErr) {
				t.Errorf("expected error %v, got %v", tC.expectedErr, err)
			}
		})
	}

}

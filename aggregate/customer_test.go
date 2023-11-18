package aggregate_test

import (
	"errors"
	"testing"

	"github.com/Qmun14/ddd-go/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}

	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: aggregate.ErrInvalidPerson,
		}, {
			test:        "Valid name",
			name:        "Ma'mun Ramdhan",
			expectedErr: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tC.name)

			if !errors.Is(err, tC.expectedErr) {
				t.Errorf("expected error %v, got %v", tC.expectedErr, err)
			}
		})
	}
}

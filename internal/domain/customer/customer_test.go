package customer

import (
	"errors"
	"testing"
)

func Test_NewCustomer(t *testing.T) {
	type testCase struct {
		test          string
		name          string
		email         string
		password      string
		expectedError error
	}

	testCases := []testCase{
		{
			test:          "Valid test",
			name:          "Lautaro",
			email:         "lautaroolmedo77@gmail.com",
			password:      "1234",
			expectedError: nil,
		},
		{
			test:          "Invalid name",
			name:          "",
			email:         "lautaroolmedo77@gmail.com",
			password:      "1234",
			expectedError: ErrInvalidPerson,
		},
		{
			test:          "Invalid email",
			name:          "",
			email:         "lautaroolmedo77@gmail.com",
			password:      "1234",
			expectedError: ErrInvalidPerson,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.test, func(t *testing.T) {
			t.Parallel()
			_, err := NewCustomer(tc.name, tc.email, tc.password)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected %v, got %v", tc.expectedError, err)
			}

		})
	}

}

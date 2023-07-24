package application

import (
	"context"
	"errors"
	"rest-api/internal/user/domain"
	"testing"
)

func TestUserService_GetAllUsers(t *testing.T) {

}

func TestUserService_SaveUser(t *testing.T) {

	repo := &RepositoryMocked{}
	var s domain.Repository = NewUserService(repo)

	myContext := context.Background()

	type testCase struct {
		test          string
		name          string
		email         string
		password      string
		expectedError error
	}

	testCases := []testCase{
		{
			test:          "Valid user",
			name:          "Lautaro",
			email:         "lautaroolmedo77@gmail.com",
			password:      "1234",
			expectedError: nil,
		},
		{
			test:          "Error. Missing name parameter",
			name:          "",
			email:         "lautaroolmedo77@gmail.com",
			password:      "1234",
			expectedError: InvalidParameter,
		},

		{
			test:          "Error. Invalid email parameter",
			name:          "Lautaro",
			email:         "lautaroolmedo77gmail.com",
			password:      "1234",
			expectedError: InvalidEmail,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.test, func(t *testing.T) {
			t.Parallel()
			err := s.SaveUser(myContext, tc.name, tc.email, tc.password)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected %v, got %v", tc.expectedError, err)
			}

		})
	}
}

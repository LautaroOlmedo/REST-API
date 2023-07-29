package application

import (
	"context"
	"errors"
	"os"
	"rest-api/internal/user/domain"
	"testing"
)

var s domain.Repository

func TestMain(m *testing.M) {
	repo := &domain.MockRepository{}
	s = NewUserService(repo)

	code := m.Run()
	os.Exit(code)
}

func TestUserService_GetAllUsers(t *testing.T) {

}

func TestUserService_CreateUser(t *testing.T) {

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
			err := s.CreateUser(myContext, tc.name, tc.email, tc.password)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected %v, got %v", tc.expectedError, err)
			}

		})
	}
}

func TestUserService_GetUserByID(t *testing.T) {
	myContext := context.Background()

	type testCase struct {
		test          string
		id            int
		expectedError error
	}

	testCases := []testCase{
		{
			test:          "valid id",
			id:            1,
			expectedError: nil,
		},
		{
			test:          "invalid id",
			id:            -4,
			expectedError: InvalidID,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.test, func(t *testing.T) {
			t.Parallel()
			_, err := s.GetUserByID(myContext, tc.id)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}

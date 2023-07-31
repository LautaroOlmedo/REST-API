package application

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"os"
	"rest-api/internal/user/domain"
	"testing"
)

var userRepo *domain.MockRepository

func TestMain(m *testing.M) {
	userRepo = &domain.MockRepository{}
	userRepo.On("GetUserByEmail", mock.Anything, "lautaroolmedo77@gmail.com").Return(nil, nil)
	userRepo.On("GetUserByEmail", mock.Anything, "javierMiner@gmail.com").Return(&domain.User{Email: "javierMiner@gmail.com"}, nil)
	userRepo.On("GetUserByEmail", mock.Anything, "joelquinteros99gmail.com").Return(nil, InvalidEmail)
	userRepo.On("CreateUser", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	userRepo.On("GetUserByID", mock.Anything, 1).Return(&domain.User{ID: 1}, nil)
	userRepo.On("GetUserByID", mock.Anything, 3).Return(nil, UserNotFound)

	code := m.Run()
	os.Exit(code)
}

func TestUserService_RegisterUser(t *testing.T) {
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
			test:          "PASS. valid user",
			name:          "Lautaro",
			email:         "lautaroolmedo77@gmail.com",
			password:      "1234",
			expectedError: nil,
		},
		{
			test:          "ERROR. invalid name",
			name:          "",
			email:         "lautaroolmedo77@gmail.com",
			password:      "1234",
			expectedError: InvalidName,
		},
		{
			test:          "ERROR. empty email",
			name:          "Lautaro",
			email:         "",
			password:      "1234",
			expectedError: InvalidEmail,
		},
		{
			test:          "ERROR. invalid email",
			name:          "Joel",
			email:         "joelquinteros99gmail.com",
			password:      "1234",
			expectedError: InvalidEmail,
		},
		{
			test:          "ERROR. user already exist",
			name:          "Javier",
			email:         "javierMiner@gmail.com",
			password:      "1234",
			expectedError: UserAlreadyExist,
		},
		{
			test:          "ERROR. invalid password",
			name:          "Lautaro",
			email:         "lautaroolmedo77@gmail.com",
			password:      "",
			expectedError: InvalidPassword,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.test, func(t *testing.T) {
			t.Parallel()
			userRepo.Mock.Test(t)

			serv := NewUserService(userRepo)

			err := serv.RegisterUser(myContext, tc.name, tc.email, tc.password)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected %v, got %v", tc.expectedError, err)
			}

		})
	}
}

func TestUserService_GetByID(t *testing.T) {
	myContext := context.Background()

	type testCase struct {
		test          string
		id            int
		expectedError error
	}

	testCases := []testCase{
		{
			test:          "PASS. valid id",
			id:            1,
			expectedError: nil,
		},
		{
			test:          "ERROR. user not found",
			id:            3,
			expectedError: UserNotFound,
		},
		{
			test:          "ERROR. invalid id",
			id:            -1,
			expectedError: InvalidID,
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.test, func(t *testing.T) {
			t.Parallel()
			userRepo.Mock.Test(t)
			serv := NewUserService(userRepo)

			_, err := serv.GetByID(myContext, tc.id)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected %v, got %v", tc.expectedError, err)
			}

		})
	}
}

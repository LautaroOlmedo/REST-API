package application

import (
	"context"
	"fmt"
	"github.com/asaskevich/govalidator"
	"rest-api/internal/user/domain"
	"rest-api/internal/user/domain/model"
)

type UserService struct {
	userRepository domain.Repository
	//messageBus     *messagebus.MessageBus
}

func NewUserService(userRepository domain.Repository) *UserService {
	return &UserService{userRepository: userRepository}
}

//func (userService *UserService) GetAll(ctx context.Context) (map[int]*model.User, error) {
//	usersMap := make(map[int]*model.User)
//	users, err := userService.userRepository.GetAllUsers(ctx)
//
//	if err != nil {
//		return nil, err
//	}
//	return nil, nil
//
//	for i := 0; i < len(users); i++ {
//		usersMap[users[i].ID] = &model.User{ID: users[i].ID, Name: users[i].Name, Email: users[i].Email}
//	}
//
//	return usersMap, nil
//}

func (userService *UserService) GetByID(ctx context.Context, id int) (*model.User, error) {
	if id < 1 {
		return nil, InvalidID
	}
	user, err := userService.userRepository.GetUserByID(ctx, id)

	if err != nil {
		fmt.Println("ERROR:", err)
		return nil, err
	}

	return &model.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (userService *UserService) LoginUser(ctx context.Context, email, password string) (*model.User, error) {
	user, err := userService.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// ---> DECRYPT PASSWORD

	if user.Password != password {
		return nil, InvalidPassword
	}

	return &model.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (userService *UserService) RegisterUser(ctx context.Context, name, email, password string) error {

	if name == "" {
		return InvalidName
	} else if email == "" || !govalidator.IsEmail(email) {
		return InvalidEmail
	} else if password == "" {
		return InvalidPassword
	}
	user, err := userService.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	if user != nil {
		return UserAlreadyExist
	}

	// ---> HASH PASSWORD

	return userService.userRepository.CreateUser(ctx, name, email, password)
}

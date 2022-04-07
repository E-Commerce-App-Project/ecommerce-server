package service

import (
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/database"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	GetAll() ([]payload.UserModel, error)
	GetUserById(id int) (payload.UserModel, error)
	CreateUser(user payload.RegisterPayload) (payload.UserModel, error)
	DeleteUser(id int) error
	UpdateUser(id int, userplyd payload.RegisterPayload) (payload.UserModel, error)
}

type userService struct {
	opt Option
}

func NewUserService(opt Option) IUserService {
	return &userService{
		opt: opt,
	}

}

func (us *userService) GetAll() ([]payload.UserModel, error) {
	users, err := us.opt.Repository.User.GetAll()
	usersModel := []payload.UserModel{}
	for i := 0; i < len(users); i++ {
		usersModel = append(usersModel, payload.UserModel{
			UserID:      users[i].ID,
			Email:       users[i].Email,
			PhoneNumber: users[i].PhoneNumber,
			Address:     users[i].Address,
		})
	}
	return usersModel, err
}

func (us *userService) GetUserById(id int) (payload.UserModel, error) {
	user, err := us.opt.Repository.User.GetUserById(id)
	userModel := payload.UserModel{
		UserID:      user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
	}
	return userModel, err
}

func (us *userService) CreateUser(userplyd payload.RegisterPayload) (payload.UserModel, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(userplyd.Password), bcrypt.DefaultCost)
	userEntity := database.UserEntity{
		Name:        userplyd.Name,
		Email:       userplyd.Email,
		Password:    string(password),
		PhoneNumber: userplyd.Phone,
		Address:     userplyd.Address,
	}
	user, err := us.opt.Repository.User.CreateUser(userEntity)

	userModel := payload.UserModel{
		UserID:      user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
	}
	return userModel, err
}

func (us *userService) DeleteUser(id int) error {
	err := us.opt.Repository.User.DeleteUser(id)
	return err
}

func (us *userService) UpdateUser(id int, userplyd payload.RegisterPayload) (payload.UserModel, error) {

	password, _ := bcrypt.GenerateFromPassword([]byte(userplyd.Password), bcrypt.DefaultCost)
	userEntity := database.UserEntity{
		Name:        userplyd.Name,
		Email:       userplyd.Email,
		Password:    string(password),
		PhoneNumber: userplyd.Phone,
		Address:     userplyd.Address,
	}
	user, err := us.opt.Repository.User.UpdateUser(id, userEntity)

	userModel := payload.UserModel{
		UserID:      user.ID,
		Email:       user.Email,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
	}
	return userModel, err
}

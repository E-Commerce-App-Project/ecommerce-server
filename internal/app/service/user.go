package service

import (
	"fmt"
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

func (uuc *userService) GetAll() ([]payload.UserModel, error) {
	users, err := uuc.opt.Repository.User.GetAll()
	usersmodel := []payload.UserModel{}
	for i := 0; i < len(users); i++ {
		usersmodel = append(usersmodel, payload.UserModel{
			UserID:      users[i].ID,
			Email:       users[i].Email,
			PhoneNumber: users[i].PhoneNumber,
			Address:     users[i].Address,
		})
	}
	return usersmodel, err
}

func (uuc *userService) GetUserById(id int) (payload.UserModel, error) {
	user, err := uuc.opt.Repository.User.GetUserById(id)
	userModel := payload.UserModel{
		UserID:      user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
	}
	return userModel, err
}

func (uuc *userService) CreateUser(userplyd payload.RegisterPayload) (payload.UserModel, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(userplyd.Password), 14)
	userEntity := database.UserEntity{
		Name:        userplyd.Name,
		Email:       userplyd.Email,
		Password:    string(password),
		PhoneNumber: userplyd.Phone,
		Address:     userplyd.Address,
	}
	user, err := uuc.opt.Repository.User.CreateUser(userEntity)

	userModel := payload.UserModel{
		UserID:      user.ID,
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
	}
	return userModel, err
}

func (uuc *userService) DeleteUser(id int) error {
	err := uuc.opt.Repository.User.DeleteUser(id)
	return err
}

func (uuc *userService) UpdateUser(id int, userplyd payload.RegisterPayload) (payload.UserModel, error) {
	userEntity := database.UserEntity{
		Name:        userplyd.Name,
		Email:       userplyd.Email,
		Password:    userplyd.Password,
		PhoneNumber: userplyd.Phone,
		Address:     userplyd.Address,
	}
	user, err := uuc.opt.Repository.User.UpdateUser(id, userEntity)
	fmt.Println(err)
	userModel := payload.UserModel{
		UserID:      user.ID,
		Email:       user.Email,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
	}
	return userModel, err
}

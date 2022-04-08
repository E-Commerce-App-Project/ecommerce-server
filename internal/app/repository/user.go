package repository

import (
	"fmt"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/database"
)

type IUserRepository interface {
	GetAll() ([]database.UserEntity, error)
	GetUserById() (database.UserEntity, error)
	CreateUser(user database.UserEntity) (database.UserEntity, error)
	DeleteUser(id int) error
	UpdateUser(id int, user database.UserEntity) (database.UserEntity, error)
}

type userRepository struct {
	opt Option
}

func NewUserRepository(opt Option) *userRepository {
	return &userRepository{
		opt: opt,
	}
}

func (ur *userRepository) GetAll() ([]database.UserEntity, error) {
	var users []database.UserEntity
	tx := ur.opt.DbMysql.Find(&users)
	if tx.Error != nil {
		return nil, commons.ErrGetAll
	}
	return users, nil

}

func (ur *userRepository) GetUserById() (database.UserEntity, error) {
	var users database.UserEntity
	tx := ur.opt.DbMysql.First(&users)
	if tx.Error != nil {
		return users, commons.ErrGetUserByID
	}
	return users, nil

}

func (ur *userRepository) CreateUser(user database.UserEntity) (database.UserEntity, error) {

	tx := ur.opt.DbMysql.Create(&user)
	fmt.Println(tx.Error)
	if tx.Error != nil {
		return user, commons.ErrCreateUser
	}
	return user, nil

}

func (ur *userRepository) DeleteUser(id int) error {
	var users database.UserEntity
	tx := ur.opt.DbMysql.Where("id = ?", id).Delete(&users)

	if tx.RowsAffected == 0 {
		return commons.ErrDeleteUser
	}
	return nil
}

func (ur *userRepository) UpdateUser(id int, user database.UserEntity) (database.UserEntity, error) {

	tx := ur.opt.DbMysql.Where("id = ?", id).Updates(&user)
	if tx.Error != nil {
		return user, commons.ErrUpdate
	}
	if tx.RowsAffected == 0 {

		return user, commons.ErrUpdate
	}
	return user, nil

}

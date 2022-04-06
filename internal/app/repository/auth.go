package repository

import (
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/database"
	"github.com/go-sql-driver/mysql"
)

type IAuthRepository interface {
	GetUserByEmail(email string) (database.UserEntity, error)
	RegisterUser(user database.UserEntity) (database.UserEntity, error)
}

type authRepository struct {
	opt Option
}

func NewAuthRepository(opt Option) IAuthRepository {
	return &authRepository{
		opt: opt,
	}
}

func (r *authRepository) GetUserByEmail(email string) (database.UserEntity, error) {
	var user database.UserEntity
	err := r.opt.DbMysql.First(&user, "email = ?", email).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *authRepository) RegisterUser(user database.UserEntity) (database.UserEntity, error) {
	userData := database.UserEntity{
		Email:       user.Email,
		Password:    user.Password,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
	}
	result := r.opt.DbMysql.Create(&userData)
	if result.Error != nil {
		mysqlErrorNumber := result.Error.(*mysql.MySQLError).Number
		switch mysqlErrorNumber {
		case 1062:
			return userData, commons.ErrEmailExists
		}
		return userData, result.Error
	}

	return userData, nil
}
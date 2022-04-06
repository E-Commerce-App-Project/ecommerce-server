package repository

import (
	"strings"

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
		Address:     user.Address,
	}
	result := r.opt.DbMysql.Create(&userData)
	if result.Error != nil {
		mysqlErrorNumber := result.Error.(*mysql.MySQLError).Number
		switch mysqlErrorNumber {
		case 1062:
			phoneError := strings.Split(result.Error.Error(), "phone_number")
			if len(phoneError) > 1 {
				return userData, commons.ErrPhoneExistError
			}
			emailError := strings.Split(result.Error.Error(), "email")
			if len(emailError) > 1 {
				return userData, commons.ErrEmailExists
			}
		case 1406:
			return userData, commons.ErrInvalidData
		}

		return userData, result.Error
	}

	return userData, nil
}

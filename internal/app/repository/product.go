package repository

import (
	"fmt"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/database"
)

type IProductRepository interface {
	GetAllProduct() ([]database.ProductEntity, error)
	GetProductById(id int) (database.ProductEntity, error)
	GetProductByIdUser(id int) ([]database.ProductEntity, error)
	CreateProduct(product database.ProductEntity) (database.ProductEntity, error)
	DeleteProduct(id, userID int) error
	UpdateProduct(id, userID int, product database.ProductEntity) (database.ProductEntity, error)
}

type productRepository struct {
	opt Option
}

func NewProductRepository(opt Option) *productRepository {
	return &productRepository{
		opt: opt,
	}
}

func (ur *productRepository) GetAllProduct() ([]database.ProductEntity, error) {
	var products []database.ProductEntity
	tx := ur.opt.DbMysql.Find(&products)
	if tx.Error != nil {
		return nil, commons.ErrGetAll
	}
	return products, nil

}

func (ur *productRepository) GetProductById(id int) (database.ProductEntity, error) {
	var products database.ProductEntity
	tx := ur.opt.DbMysql.First(&products)
	if tx.Error != nil {
		return products, commons.ErrGetUserByID
	}
	return products, nil

}

func (ur *productRepository) GetProductByIdUser(id int) ([]database.ProductEntity, error) {
	var products []database.ProductEntity
	tx := ur.opt.DbMysql.Where("user_id = ?", id).Find(&products)
	if tx.Error != nil {
		return products, commons.ErrGetUserByID
	}
	return products, nil

}

func (ur *productRepository) CreateProduct(product database.ProductEntity) (database.ProductEntity, error) {

	tx := ur.opt.DbMysql.Create(&product)
	fmt.Println(tx.Error)
	if tx.Error != nil {
		return product, commons.ErrCreateUser
	}
	return product, nil

}

func (ur *productRepository) DeleteProduct(id, userID int) error {
	var products database.ProductEntity
	tx := ur.opt.DbMysql.Where("id =? and user_id = ?", id, userID).Delete(&products)

	if tx.RowsAffected == 0 {
		return commons.ErrDeleteUser
	}
	return nil
}

func (ur *productRepository) UpdateProduct(id, userID int, product database.ProductEntity) (database.ProductEntity, error) {

	tx := ur.opt.DbMysql.Where("id =? and user_id = ?", id, userID).Updates(&product)
	fmt.Println(tx.Error)
	fmt.Println(id)
	if tx.Error != nil {
		return product, commons.ErrUpdate
	}
	if tx.RowsAffected == 0 {

		return product, commons.ErrUpdate
	}
	return product, nil

}

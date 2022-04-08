package service

import (
	"fmt"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/database"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
)

type IProductService interface {
	GetAllProduct() ([]payload.ProductModel, error)
	GetProductById(id int) (payload.ProductModel, error)
	GetProductByIdUser() ([]payload.ProductModel, error)
	CreateProduct(product payload.CreateProductPayload) (payload.ProductModel, error)
	DeleteProduct(id, userID int) error
	UpdateProduct(id, userID int, productplyd payload.CreateProductPayload) (payload.ProductModel, error)
}

type productService struct {
	opt Option
}

func NewProductService(opt Option) IProductService {
	return &productService{
		opt: opt,
	}

}

func (puc *productService) GetAllProduct() ([]payload.ProductModel, error) {
	product, err := puc.opt.Repository.Product.GetAllProduct()
	productmodel := []payload.ProductModel{}
	for i := 0; i < len(product); i++ {
		productmodel = append(productmodel, payload.ProductModel{
			Name:        product[i].Name,
			Price:       product[i].Price,
			Description: product[i].Description,
			Image:       product[i].Image,
			Stock:       product[i].Stock,
			UserID:      product[i].UserID,
		})
	}
	return productmodel, err
}

func (puc *productService) GetProductById(id int) (payload.ProductModel, error) {
	product, err := puc.opt.Repository.Product.GetProductById(id)
	productModel := payload.ProductModel{
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Image:       product.Image,
		Stock:       product.Stock,
		UserID:      product.UserID,
	}
	return productModel, err
}

func (puc *productService) GetProductByIdUser() ([]payload.ProductModel, error) {
	product, err := puc.opt.Repository.Product.GetProductByIdUser()
	productmodel := []payload.ProductModel{}
	for i := 0; i < len(product); i++ {
		productmodel = append(productmodel, payload.ProductModel{
			Name:        product[i].Name,
			Price:       product[i].Price,
			Description: product[i].Description,
			Image:       product[i].Image,
			Stock:       product[i].Stock,
			UserID:      product[i].UserID,
		})
	}
	return productmodel, err
}

func (puc *productService) CreateProduct(productplyd payload.CreateProductPayload) (payload.ProductModel, error) {

	productEntity := database.ProductEntity{
		UserID:      productplyd.UserID,
		Name:        productplyd.Name,
		Price:       productplyd.Price,
		Description: productplyd.Description,
		Image:       productplyd.Image,
		Stock:       productplyd.Stock,
	}
	product, err := puc.opt.Repository.Product.CreateProduct(productEntity)

	productModel := payload.ProductModel{
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Image:       product.Image,
		Stock:       product.Stock,
		UserID:      product.UserID,
	}

	return productModel, err
}

func (puc *productService) DeleteProduct(id, userID int) error {
	err := puc.opt.Repository.Product.DeleteProduct(id, userID)
	return err
}

func (puc *productService) UpdateProduct(id, userID int, productplyd payload.CreateProductPayload) (payload.ProductModel, error) {
	productEntity := database.ProductEntity{
		Name:        productplyd.Name,
		Price:       productplyd.Price,
		Description: productplyd.Description,
		Image:       productplyd.Image,
		Stock:       productplyd.Stock,
		UserID:      productplyd.UserID,
	}
	product, err := puc.opt.Repository.Product.UpdateProduct(id, userID, productEntity)
	fmt.Println(err)
	productModel := payload.ProductModel{
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Image:       product.Image,
		Stock:       product.Stock,
		UserID:      product.UserID,
	}
	return productModel, err
}

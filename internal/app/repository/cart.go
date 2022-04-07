package repository

import (
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
)

type ICartRepository interface {
	GetCart(getCartPyld payload.GetUserCartPayload) ([]payload.CartModel, error)
	// AddToCart(cartPyld payload.AddProductToCartPayload) (payload.CartModel, error)
	// UpdateCart(cartPyld payload.AddProductToCartPayload) (payload.CartModel, error)
	// DeleteCart(cartPyld payload.AddProductToCartPayload) (payload.CartModel, error)
	// ValidateCart(getCartPyld payload.GetUserCartPayload) (payload.CartModel, error)
}

type cartRepository struct {
	opt Option
}

func NewCartRepository(opt Option) ICartRepository {
	return &cartRepository{
		opt: opt,
	}
}

func (cr *cartRepository) GetCart(getCartPyld payload.GetUserCartPayload) ([]payload.CartModel, error) {
	var carts []payload.CartModel
	productJoin := "JOIN product_entities ON product_entities.id = cart_item_entities.product_id"
	cartSelect := "cart_item_entities.id as id, cart_item_entities.user_id as user_id, product_entities.id as product_id, product_entities.name as name, cart_item_entities.total_price as price, product_entities.image as image, cart_item_entities.total_product as quantity, product_entities.deleted_at as deleted_at, product_entities.stock as stock"
	query := cr.opt.DbMysql.Table("cart_item_entities").Select(cartSelect).Joins(productJoin)
	query = query.Where("cart_item_entities.user_id = ? AND deleted_at IS NULL AND stock >= cart_item_entities.total_product", getCartPyld.UserID)
	err := query.Find(&carts).Error
	if err != nil {
		return carts, err
	}

	return carts, nil
}

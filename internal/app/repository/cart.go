package repository

import (
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/database"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
)

type ICartRepository interface {
	GetCart(getCartPyld payload.GetUserCartPayload) ([]payload.CartModel, error)
	AddToCart(cartPyld payload.AddProductToCartPayload) ([]payload.CartModel, error)
	UpdateCart(updatePyld payload.UpdateCartPayload) ([]payload.CartModel, error)
	DeleteCart(delPyld payload.DeleteUserCartPayload) bool
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

func (cr *cartRepository) AddToCart(cartPyld payload.AddProductToCartPayload) ([]payload.CartModel, error) {
	cartItem, err := cr.FindProductInCart(cartPyld.ProductID, cartPyld.UserID)
	if err == nil {
		// get cart
		return cr.UpdateCart(payload.UpdateCartPayload{
			CartID:   int(cartItem.ID),
			UserID:   cartPyld.UserID,
			Quantity: cartPyld.Quantity + cartItem.TotalProduct,
		})
	}
	var product database.ProductEntity
	err = cr.opt.DbMysql.Model(&database.ProductEntity{}).Where("id = ?", cartPyld.ProductID).First(&product).Error
	if err != nil {
		return []payload.CartModel{}, err
	}
	if product.Stock < cartPyld.Quantity {
		return []payload.CartModel{}, commons.ErrOutOfStock
	}

	cartItem.UserID = uint(cartPyld.UserID)
	cartItem.ProductID = uint(cartPyld.ProductID)
	cartItem.TotalProduct = cartPyld.Quantity
	cartItem.TotalPrice = product.Price

	err = cr.opt.DbMysql.Create(&cartItem).Error
	if err != nil {
		return []payload.CartModel{}, commons.ErrAddCart
	}

	return cr.GetCart(payload.GetUserCartPayload{
		UserID: cartPyld.UserID,
	})
}

func (cr *cartRepository) FindProductInCart(productID int, userID int) (database.CartItemEntity, error) {
	var cart database.CartItemEntity
	query := cr.opt.DbMysql.Table("cart_item_entities").Where("product_id = ? AND user_id = ?", productID, userID)
	tx := query.Find(&cart)
	if tx.Error != nil {
		return cart, tx.Error
	}

	if tx.RowsAffected == 0 {
		return cart, commons.ErrNotFound
	}

	return cart, nil
}

func (cr *cartRepository) UpdateCart(cartPyld payload.UpdateCartPayload) ([]payload.CartModel, error) {

	var cartItem database.CartItemEntity
	err := cr.opt.DbMysql.Preload("Product").Model(&database.CartItemEntity{}).Where("id = ? AND user_id = ?", cartPyld.CartID, cartPyld.UserID).First(&cartItem).Error

	if cartItem.Product.Stock < cartPyld.Quantity {
		return []payload.CartModel{}, commons.ErrOutOfStock
	}

	if err != nil {
		return []payload.CartModel{}, commons.ErrUpdateCart
	}
	cartItem.TotalProduct = cartPyld.Quantity

	err = cr.opt.DbMysql.Save(&cartItem).Error
	if err != nil {
		return []payload.CartModel{}, commons.ErrUpdateCart
	}

	return cr.GetCart(payload.GetUserCartPayload{
		UserID: cartPyld.UserID,
	})
}

func (cr *cartRepository) DeleteCart(delPyld payload.DeleteUserCartPayload) bool {
	var cartItem database.CartItemEntity
	err := cr.opt.DbMysql.Model(&database.CartItemEntity{}).Where("id = ?", delPyld.CartItemID).First(&cartItem).Error
	if err != nil {
		return false
	}

	err = cr.opt.DbMysql.Delete(&cartItem).Error

	return err != nil
}

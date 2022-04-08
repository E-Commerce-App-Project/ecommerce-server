package service

import (
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/database"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
)

type ICartService interface {
	GetCart(getCartPyld payload.GetUserCartPayload) ([]payload.CartModel, error)
	Checkout(checkoutPyld payload.CheckoutPayload) (database.TransactionEntity, error)
	AddToCart(addToCartPayload payload.AddProductToCartPayload) ([]payload.CartModel, error)
	UpdateCart(updateCartPayload payload.UpdateCartPayload) ([]payload.CartModel, error)
	DeleteCart(deleteCartPayload payload.DeleteUserCartPayload) bool
}

type cartService struct {
	opt Option
}

func NewCartService(opt Option) ICartService {
	return &cartService{
		opt: opt,
	}
}

func (cs *cartService) GetCart(getCartPyld payload.GetUserCartPayload) ([]payload.CartModel, error) {
	carts, err := cs.opt.Repository.Cart.GetCart(getCartPyld)
	if err != nil {
		return carts, commons.ErrQueryDB
	}

	return carts, nil
}

func (cs *cartService) AddToCart(addToCartPayload payload.AddProductToCartPayload) ([]payload.CartModel, error) {
	cart, err := cs.opt.Repository.Cart.AddToCart(addToCartPayload)
	if err != nil {
		return cart, commons.ErrQueryDB
	}

	return cart, nil
}

func (cs *cartService) UpdateCart(updateCartPayload payload.UpdateCartPayload) ([]payload.CartModel, error) {
	cart, err := cs.opt.Repository.Cart.UpdateCart(updateCartPayload)
	if err != nil {
		return cart, commons.ErrQueryDB
	}

	return cart, nil
}

func (cs *cartService) DeleteCart(deleteCartPayload payload.DeleteUserCartPayload) bool {
	return cs.opt.Repository.Cart.DeleteCart(deleteCartPayload)
}

func (cs *cartService) Checkout(checkoutPyld payload.CheckoutPayload) (database.TransactionEntity, error) {
	carts, err := cs.opt.Repository.Cart.GetCart(payload.GetUserCartPayload{
		UserID: checkoutPyld.UserID,
	})
	if err != nil {
		return database.TransactionEntity{}, commons.ErrQueryDB
	}

	if len(carts) == 0 {
		return database.TransactionEntity{}, commons.ErrEmptyCart
	}

	pickedItem := []payload.CartModel{}
	for i, cartItemID := range checkoutPyld.CartItemIDs {
		for _, cart := range carts {
			if cart.ID == uint(cartItemID) {
				pickedItem = append(pickedItem, carts[i])
			}
		}
	}

	if len(pickedItem) == 0 {
		return database.TransactionEntity{}, commons.ErrInvalidCartItem
	}

	if len(pickedItem) != len(checkoutPyld.CartItemIDs) {
		return database.TransactionEntity{}, commons.ErrInvalidCartItem
	}
	cartItems := []payload.TransactionItemPayload{}
	totalPrice := int(0)
	for _, cartItem := range pickedItem {
		cartItems = append(cartItems, payload.TransactionItemPayload{
			ProductID: int(cartItem.ProductID),
			Quantity:  int(cartItem.Quantity),
			Price:     cartItem.Price,
		})
		totalPrice += cartItem.Price * int(cartItem.Quantity)
	}
	// add transaction
	transaction, err := cs.opt.Repository.Transaction.AddTransaction(payload.TransactionPayload{
		UserID:     checkoutPyld.UserID,
		TotalPrice: totalPrice,
		CartItems:  cartItems,
	})
	if err != nil {
		return database.TransactionEntity{}, commons.ErrAddTransaction
	}

	// delete cart
	for _, cartItem := range pickedItem {
		deleted := cs.opt.Repository.Cart.DeleteCart(payload.DeleteUserCartPayload{
			UserID:     checkoutPyld.UserID,
			CartItemID: int(cartItem.ID),
		})
		if !deleted {
			return database.TransactionEntity{}, commons.ErrDeleteCart
		}
	}

	return transaction, nil
}

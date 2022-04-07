package service

import (
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/commons"
	"github.com/E-Commerce-App-Project/ecommerce-server/internal/app/payload"
)

type ICartService interface {
	GetCart(getCartPyld payload.GetUserCartPayload) ([]payload.CartModel, error)
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

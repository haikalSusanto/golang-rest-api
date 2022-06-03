package cart

import (
	"github.com/pkg/errors"
)

type service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return &service{
		repo: repo,
	}
}

type Service interface {
	GetOngoingCart(username string) (*OngoingCart, error)
	AddItemToCart(cart_id int, product_id int, quantity int) (*CartItem, error)
}

func (s *service) GetOngoingCart(username string) (*OngoingCart, error) {
	var cart *OngoingCart

	user_id, err := s.repo.GetUserWithUsername(username)
	if err != nil {
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}

	cart, err = s.repo.GetUserOngoingCart(user_id)
	if err != nil {
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}

	if cart == nil {
		cart, err = s.repo.CreateNewCart(user_id)
		if err != nil {
			return nil, errors.Wrap(ErrInternalServer, err.Error())
		}
	}

	return cart, nil
}

func (s *service) AddItemToCart(cart_id int, product_id int, quantity int) (*CartItem, error) {
	currentCartItem, err := s.repo.CheckCartItem(cart_id, product_id)
	if err != nil {
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}

	var cartItem *CartItem
	if currentCartItem != nil {
		newQuantity := currentCartItem.Quantity + quantity
		cartItem, err = s.repo.UpdateItemQuantity(cart_id, product_id, newQuantity)
		if err != nil {
			return nil, errors.Wrap(ErrInternalServer, err.Error())
		}
	} else {
		cartItem, err = s.repo.AddItemToCart(cart_id, product_id, quantity)
		if err != nil {
			return nil, errors.Wrap(ErrInternalServer, err.Error())
		}
	}

	return cartItem, nil
}

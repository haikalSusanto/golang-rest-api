package cart

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

type Repo interface {
	GetUserWithUsername(username string) (int, error)
	GetUserOngoingCart(user_id int) (*OngoingCart, error)
	CreateNewCart(user_id int) (*OngoingCart, error)
	AddItemToCart(cart_id int, product_id int, quantity int) (*CartItem, error)
	UpdateItemQuantity(cart_id int, product_id int, quantity int) (*CartItem, error)
	CheckCartItem(cart_id int, product_id int) (*CartItem, error)
	GetAllCartItems(cart_id int) (*ListCartItem, error)
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{
		db: db,
	}
}

func (r *repo) GetUserWithUsername(username string) (int, error) {
	var user User
	err := r.db.Table("users").Find(&user, "username = ?", username).Error
	if err != nil {
		return 0, errors.Wrap(ErrInternalServer, err.Error())
	}

	return user.ID, nil
}

func (r *repo) GetUserOngoingCart(user_id int) (*OngoingCart, error) {
	var cart OngoingCart
	err := r.db.Table("carts").First(&cart, "user_id = ? and status = ?", user_id, 1).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}

	return &cart, nil
}

func (r *repo) CreateNewCart(user_id int) (*OngoingCart, error) {
	newCart := Cart{
		User_id:     user_id,
		Status:      1,
		Total_price: 0,
	}
	err := r.db.Table("carts").Create(&newCart).Error
	if err != nil {
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}

	var cart OngoingCart
	err = r.db.Table("carts").First(&cart, "user_id = ? and status = ?", user_id, 1).Error
	if err != nil {
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}

	return &cart, nil
}

func (r *repo) AddItemToCart(cart_id int, product_id int, quantity int) (*CartItem, error) {
	newCartItem := CartItem{
		Cart_id:    cart_id,
		Product_id: product_id,
		Quantity:   quantity,
	}
	err := r.db.Table("cart_items").Create(&newCartItem).Error
	if err != nil {
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}
	return &newCartItem, nil
}

func (r *repo) UpdateItemQuantity(cart_id int, product_id int, quantity int) (*CartItem, error) {
	var cartItem CartItem
	err := r.db.Table("cart_items").First(&cartItem, "cart_id = ? and product_id = ?", cart_id, product_id).Update("quantity", quantity).Error
	if err != nil {
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}

	err = r.db.Table("cart_items").First(&cartItem, "cart_id = ? and product_id = ?", cart_id, product_id).Error
	if err != nil {
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}

	return &cartItem, nil
}

func (r *repo) CheckCartItem(cart_id int, product_id int) (*CartItem, error) {
	var cartItem CartItem
	err := r.db.Table("cart_items").First(&cartItem, "cart_id = ? and product_id = ?", cart_id, product_id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}
	return &cartItem, nil
}

func (r *repo) GetAllCartItems(cart_id int) (*ListCartItem, error) {
	var listCartItem ListCartItem
	err := r.db.Table("cart_items").Find(&listCartItem.CartItems, "cart_id = ?", cart_id).Error
	if err != nil {
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}

	return &listCartItem, nil
}

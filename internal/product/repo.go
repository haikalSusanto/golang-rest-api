package product

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

type Repo interface {
	GetAllProducts() (*ListProduct, error)
	GetProductsByCategory(category string) (*ListProduct, error)
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{
		db: db,
	}
}

func (r *repo) GetAllProducts() (*ListProduct, error) {
	var listProduct ListProduct
	err := r.db.Table("products").Find(&listProduct.Products).Error
	if err != nil {
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}

	return &listProduct, nil
}

func (r *repo) GetProductsByCategory(category string) (*ListProduct, error) {
	var listProduct ListProduct
	err := r.db.Table("products").Find(&listProduct.Products, "category = ?", category).Error
	if err != nil {
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}

	return &listProduct, nil
}

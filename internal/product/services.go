package product

import "github.com/pkg/errors"

type service struct {
	repo Repo
}

func NewService(repo Repo) Service {
	return &service{
		repo: repo,
	}
}

type Service interface {
	GetAllProducts() (*ListProduct, error)
	GetProductsByCategory(category string) (*ListProduct, error)
}

func (s *service) GetAllProducts() (*ListProduct, error) {
	listProduct, err := s.repo.GetAllProducts()
	if err != nil {
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}

	return listProduct, nil
}

func (s *service) GetProductsByCategory(category string) (*ListProduct, error) {
	listProduct, err := s.repo.GetProductsByCategory(category)
	if err != nil {
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}

	return listProduct, nil
}

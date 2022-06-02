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
}

func (s *service) GetAllProducts() (*ListProduct, error) {
	listProduct, err := s.repo.GetAllProducts()
	if err != nil {
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}

	return listProduct, nil
}

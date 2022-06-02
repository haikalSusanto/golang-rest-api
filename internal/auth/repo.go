package auth

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

type Repo interface {
	GetCustomerByUsername(username string) (*User, error)
	CreateNewUser(username string, hashedPassword string, name string) (*User, error)
}

func NewRepo(db *gorm.DB) Repo {
	return &repo{
		db: db,
	}
}

func (r repo) GetCustomerByUsername(username string) (*User, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrap(ErrNotFound, err.Error())
		}
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}

	return &user, nil
}

func (r repo) CreateNewUser(username string, hashedPassword string, name string) (*User, error) {
	newUser := User{
		Username: username,
		Password: hashedPassword,
		Name:     name,
	}
	err := r.db.Table("users").Create(&newUser).Error
	if err != nil {
		return nil, errors.Wrap(ErrInternalServer, err.Error())
	}

	return &newUser, nil
}

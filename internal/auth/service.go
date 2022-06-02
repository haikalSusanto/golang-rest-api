package auth

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo Repo
}

// NewService for initialize service
func NewService(repo Repo) Service {
	return &service{
		repo: repo,
	}
}

// Service will contain all the function that can be used by service
type Service interface {
	GetCustomerByUsername(username string, password string) (string, error)
	CreateNewUser(request RegisterRequest) (string, error)
}

func (s service) GetCustomerByUsername(username string, password string) (string, error) {
	user, err := s.repo.GetCustomerByUsername(username)
	if err != nil {
		if errors.Cause(err) == ErrNotFound {
			return "", errors.Wrap(ErrNotFound, err.Error())
		}
		return "", errors.Wrap(ErrInternalServer, err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.Wrap(ErrInputValidation, err.Error())
	}

	stringToken, err := getToken(user.Username)
	if err != nil {
		return "", err
	}

	return stringToken, nil
}

func (s service) CreateNewUser(request RegisterRequest) (string, error) {
	var errorList []string

	if len(errorList) > 0 {
		return "", errors.Wrap(ErrInputValidation, strings.Join(errorList, ","))
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	hashedPassword := string(bytes)

	user, err := s.repo.CreateNewUser(request.Username, hashedPassword, request.Name)
	if err != nil {
		return "", errors.Wrap(ErrInternalServer, err.Error())
	}

	stringToken, err := getToken(user.Username)
	if err != nil {
		return "", err
	}

	return stringToken, nil
}

func getToken(username string) (string, error) {
	claim := &JwtClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	stringToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return stringToken, nil
}

package middleware

import (
	"fmt"
	"net/http"
	"rest-api/internal/auth"
	"rest-api/util"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	authorizationTypeBearer = "bearer"
)

type AuthMiddleware struct {
	userRepo auth.Repo
}

func NewAuthMiddleware(userRepo auth.Repo) AuthMiddleware {
	return AuthMiddleware{userRepo: userRepo}
}

func (a AuthMiddleware) AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			for key, values := range ctx.Request().Header {
				if key == "Authorization" {
					if len(values) < 1 || values[0] == "" {
						err := errors.New("invalid authorization header format")
						return ctx.JSON(http.StatusUnauthorized, util.APIResponse{
							Status:  http.StatusUnauthorized,
							Message: "unauthorized",
							Errors:  []string{err.Error()},
						})
					}

					authHeader := strings.Split(values[0], " ")

					authorizationType := strings.ToLower(authHeader[0])
					if authorizationType != authorizationTypeBearer {
						err := fmt.Errorf("unsupported authorization type %s", authorizationType)
						return ctx.JSON(http.StatusUnauthorized, util.APIResponse{
							Status:  http.StatusUnauthorized,
							Message: "unauthorized",
							Errors:  []string{err.Error()},
						})
					}

					tokenString := authHeader[1]
					claims := auth.JwtClaims{}
					_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
						return []byte("secret"), nil
					})
					if err != nil {
						return ctx.JSON(http.StatusUnauthorized, util.APIResponse{
							Status:  http.StatusUnauthorized,
							Message: "unauthorized",
							Errors:  []string{err.Error()},
						})
					}
					username := claims.Username

					user, err := a.userRepo.GetCustomerByUsername(username)
					if err != nil && errors.Cause(err) != auth.ErrNotFound {
						logrus.Error("[failed to get data from user repo] ", err.Error())
						return ctx.JSON(http.StatusInternalServerError, util.APIResponse{
							Status:  http.StatusInternalServerError,
							Message: "internal server error",
						})
					}

					ctx.Set("userDatabase", user)
					return next(ctx)
				}
			}

			err := errors.New("authorization header is not provided")
			return ctx.JSON(http.StatusUnauthorized, util.APIResponse{
				Status:  http.StatusUnauthorized,
				Message: "unauthorized",
				Errors:  []string{err.Error()},
			})
		}
	}
}

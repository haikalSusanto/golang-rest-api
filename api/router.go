package api

import (
	"net/http"
	"rest-api/internal/auth"
	"rest-api/util"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	Router      *echo.Echo
	authHandler *auth.Handler
}

func NewRoutes(router *echo.Echo, authHandler *auth.Handler) *Routes {
	return &Routes{
		Router:      router,
		authHandler: authHandler,
	}
}

func (r *Routes) Init() {
	r.Router.GET("/", test_handler)

	v1 := r.Router.Group("/api/v1")
	{
		v1.POST("/login", r.authHandler.Login)
		v1.POST("/register", r.authHandler.Register)
	}
}

func test_handler(c echo.Context) error {
	return c.JSON(http.StatusOK, util.APIResponse{
		Status:  200,
		Message: "Berhasil",
	})
}

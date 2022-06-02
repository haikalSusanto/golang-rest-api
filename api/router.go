package api

import (
	"net/http"
	"rest-api/internal/auth"
	"rest-api/internal/product"
	"rest-api/util"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	Router         *echo.Echo
	authHandler    *auth.Handler
	productHandler *product.Handler
}

func NewRoutes(router *echo.Echo, authHandler *auth.Handler, productHandler *product.Handler) *Routes {
	return &Routes{
		Router:         router,
		authHandler:    authHandler,
		productHandler: productHandler,
	}
}

func (r *Routes) Init() {
	r.Router.GET("/", test_handler)

	v1 := r.Router.Group("/api/v1")
	{
		v1.POST("/login", r.authHandler.Login)
		v1.POST("/register", r.authHandler.Register)

		productRoutes := v1.Group("/products")
		{
			productRoutes.GET("/", r.productHandler.GetAllProducts)
		}
	}
}

func test_handler(c echo.Context) error {
	return c.JSON(http.StatusOK, util.APIResponse{
		Status:  200,
		Message: "Berhasil",
	})
}

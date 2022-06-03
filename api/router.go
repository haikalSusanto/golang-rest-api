package api

import (
	"net/http"
	"rest-api/internal/auth"
	"rest-api/internal/cart"
	"rest-api/internal/product"
	"rest-api/middleware"
	"rest-api/util"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	Router         *echo.Echo
	authHandler    *auth.Handler
	authMiddleware middleware.AuthMiddleware
	productHandler *product.Handler
	cartHandler    *cart.Handler
}

func NewRoutes(router *echo.Echo, authHandler *auth.Handler, authMiddleware middleware.AuthMiddleware, productHandler *product.Handler, cartHandler *cart.Handler) *Routes {
	return &Routes{
		Router:         router,
		authHandler:    authHandler,
		authMiddleware: authMiddleware,
		productHandler: productHandler,
		cartHandler:    cartHandler,
	}
}

func (r *Routes) Init() {
	r.Router.GET("/", test_handler)

	v1 := r.Router.Group("/api/v1")
	{
		v1.POST("/login", r.authHandler.Login)
		v1.POST("/register", r.authHandler.Register)

		productRoutes := v1.Group("/products", r.authMiddleware.AuthMiddleware())
		{
			productRoutes.GET("/", r.productHandler.GetAllProducts)
		}

		cartRoutes := v1.Group("/cart", r.authMiddleware.AuthMiddleware())
		{
			cartRoutes.GET("/", r.cartHandler.GetAllCartItems)
			cartRoutes.POST("/add", r.cartHandler.AddItemToCart)
			cartRoutes.DELETE("/delete", r.cartHandler.DeleteCartItem)
		}

	}
}

func test_handler(c echo.Context) error {
	return c.JSON(http.StatusOK, util.APIResponse{
		Status:  200,
		Message: "Berhasil",
	})
}

package api

import (
	"net/http"
	"rest-api/database/postgres"
	"rest-api/internal/auth"
	"rest-api/internal/cart"
	"rest-api/internal/product"
	"rest-api/middleware"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Server struct {
	Router *echo.Echo
}

func NewServer(router *echo.Echo) *Server {
	return &Server{
		Router: router,
	}
}

var (
	authRepo       auth.Repo
	authService    auth.Service
	authHandler    *auth.Handler
	authMiddleware middleware.AuthMiddleware

	productRepo    product.Repo
	productService product.Service
	productHandler *product.Handler

	cartRepo    cart.Repo
	cartService cart.Service
	cartHandler *cart.Handler
)

func (s *Server) Init() {
	db := postgres.Init()

	// Auth Modul
	authRepo = auth.NewRepo(db)
	authService = auth.NewService(authRepo)
	authHandler = auth.NewHandler(authService)

	authMiddleware = middleware.NewAuthMiddleware(authRepo)

	productRepo = product.NewRepo(db)
	productService = product.NewService(productRepo)
	productHandler = product.NewHandler(productService)

	cartRepo = cart.NewRepo(db)
	cartService = cart.NewService(cartRepo)
	cartHandler = cart.NewHandler(cartService)

	r := NewRoutes(s.Router, authHandler, authMiddleware, productHandler, cartHandler)
	r.Init()
}

func (s Server) RunServer(port string) {
	if err := s.Router.Start(":" + port); err != http.ErrServerClosed {
		logrus.Error(err)
	}
}

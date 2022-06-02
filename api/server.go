package api

import (
	"net/http"
	"rest-api/database/mysql"
	"rest-api/internal/auth"
	"rest-api/internal/product"

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
	authRepo    auth.Repo
	authService auth.Service
	authHandler *auth.Handler

	productRepo    product.Repo
	productService product.Service
	productHandler *product.Handler
)

func (s *Server) Init() {
	db := mysql.Init()

	// Auth Modul
	authRepo = auth.NewRepo(db)
	authService = auth.NewService(authRepo)
	authHandler = auth.NewHandler(authService)

	productRepo = product.NewRepo(db)
	productService = product.NewService(productRepo)
	productHandler = product.NewHandler(productService)

	r := NewRoutes(s.Router, authHandler, productHandler)
	r.Init()
}

func (s Server) RunServer(port string) {
	if err := s.Router.Start(":" + port); err != http.ErrServerClosed {
		logrus.Error(err)
	}
}

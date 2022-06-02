package api

import (
	"net/http"
	"rest-api/database/mysql"

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

func (s *Server) Init() {
	_ = mysql.Init()

	r := NewRoutes(s.Router)
	r.Init()
}

func (s Server) RunServer(port string) {
	if err := s.Router.Start(":" + port); err != http.ErrServerClosed {
		logrus.Error(err)
	}
}

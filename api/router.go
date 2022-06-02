package api

import (
	"net/http"
	"rest-api/util"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	Router *echo.Echo
}

func NewRoutes(router *echo.Echo) *Routes {
	return &Routes{
		Router: router,
	}
}

func (r *Routes) Init() {
	r.Router.GET("/", test_handler)

	v1 := r.Router.Group("/api/v1")
	{
		v1.GET("/", test_handler2)
	}
}

func test_handler(c echo.Context) error {
	return c.JSON(http.StatusOK, util.APIResponse{
		Status:  200,
		Message: "Berhasil",
	})
}

func test_handler2(c echo.Context) error {
	return c.JSON(http.StatusOK, util.APIResponse{
		Status:  200,
		Message: "In API",
	})
}

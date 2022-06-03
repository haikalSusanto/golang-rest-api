package main

import (
	"fmt"
	"os"
	"rest-api/api"

	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()

	s := api.NewServer(router)
	s.Init()

	s.RunServer(fmt.Sprintf("%s", os.Getenv("PORT")))
}

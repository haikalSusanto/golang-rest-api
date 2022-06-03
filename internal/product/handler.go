package product

import (
	"net/http"
	"rest-api/util"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetAllProducts(c echo.Context) error {
	var listProducts *ListProduct
	var err error

	category := c.QueryParam("category")

	if category != "" {
		listProducts, err = h.service.GetProductsByCategory(category)
	} else {
		listProducts, err = h.service.GetAllProducts()
	}

	if err != nil {
		return util.ErrorWrapWithContext(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, util.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    listProducts.Products,
	})
}

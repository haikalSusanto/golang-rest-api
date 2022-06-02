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
	listProducts, err := h.service.GetAllProducts()
	if err != nil {
		return util.ErrorWrapWithContext(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, util.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    listProducts,
	})
}

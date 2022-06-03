package cart

import (
	"net/http"
	"rest-api/middleware"
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

func (h *Handler) AddItemToCart(c echo.Context) error {
	username := middleware.ParseUserData(c)

	var req AddItemRequest
	if err := c.Bind(&req); err != nil {
		return util.ErrorWrapWithContext(c, http.StatusInternalServerError, ErrInternalServer, err.Error())
	}

	onGoingCart, err := h.service.GetOngoingCart(username)
	if err != nil {
		return util.ErrorWrapWithContext(c, http.StatusInternalServerError, err)
	}

	cartItem, err := h.service.AddItemToCart(onGoingCart.ID, req.ProductID, req.Quantity)
	if err != nil {
		return util.ErrorWrapWithContext(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, util.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    cartItem,
	})
}

func (h *Handler) GetAllCartItems(c echo.Context) error {
	username := middleware.ParseUserData(c)

	onGoingCart, err := h.service.GetOngoingCart(username)
	if err != nil {
		return util.ErrorWrapWithContext(c, http.StatusInternalServerError, err)
	}

	listCartItem, err := h.service.GetAllCartItems(onGoingCart.ID)
	if err != nil {
		return util.ErrorWrapWithContext(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, util.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    listCartItem.CartItems,
	})
}

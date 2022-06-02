package auth

import (
	"net/http"
	"rest-api/util"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Login(c echo.Context) error {

	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return util.ErrorWrapWithContext(c, http.StatusInternalServerError, ErrInternalServer, err.Error())
	}

	stringToken, err := h.service.GetCustomerByUsername(req.Username, req.Password)
	if err != nil {
		return util.ErrorWrapWithContext(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, util.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: TokenResponse{
			Token: stringToken,
		},
	},
	)
}

func (h *Handler) Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return util.ErrorWrapWithContext(c, http.StatusInternalServerError, ErrInternalServer, err.Error())
	}

	_, err := h.service.GetCustomerByUsername(req.Username, req.Password)
	if err != nil && errors.Cause(err) != ErrNotFound {
		return util.ErrorWrapWithContext(c, http.StatusInternalServerError, err)
	}

	stringToken, err := h.service.CreateNewUser(req)
	if err != nil {
		return util.ErrorWrapWithContext(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, util.APIResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data: TokenResponse{
			Token: stringToken,
		},
	})
}

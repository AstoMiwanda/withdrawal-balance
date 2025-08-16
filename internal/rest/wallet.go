package rest

import (
	"context"
	"net/http"
	"withdrawal-balance/internal/domain/wallet"
	"withdrawal-balance/internal/rest/middleware"

	"github.com/labstack/echo/v4"
)

type WalletService interface {
	InquiryBalance(ctx context.Context, req *wallet.InquiryBalanceRequest) (result wallet.InquiryBalanceResponse, err error)
}

type WalletHandler struct {
	Service WalletService
}

func NewWalletHandler(
	e *echo.Echo,
	svc WalletService,
) {
	handler := &WalletHandler{
		Service: svc,
	}
	e.POST("/v1/wallet/balance", handler.InquiryBalance)
}

func (s *WalletHandler) InquiryBalance(c echo.Context) (err error) {
	var request wallet.InquiryBalanceRequest
	err = c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, middleware.ResponseError{Message: err.Error()})
	}

	var ok bool
	if ok, err = middleware.IsRequestValid(&request); !ok {
		return c.JSON(http.StatusBadRequest, middleware.ResponseError{Message: err.Error()})
	}

	ctx := c.Request().Context()
	resp, err := s.Service.InquiryBalance(ctx, &request)
	if err != nil {
		return c.JSON(middleware.GetStatusCode(err), middleware.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

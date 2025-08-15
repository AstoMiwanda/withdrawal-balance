package rest

import (
	"context"
	"net/http"
	"withdrawal-balance/internal/domain/withdrawalhistory"
	"withdrawal-balance/internal/rest/middleware"

	"github.com/labstack/echo/v4"
)

type WithdrawalHistoryService interface {
	RequestWithdrawal(ctx context.Context, req *withdrawalhistory.CreateWithdrawalBalanceRequest) (result withdrawalhistory.CreateWithdrawalBalanceResponse, err error)
}

type UserService interface {
}

type WalletService interface {
}

type WithdrawalHistoryHandler struct {
	Service       WithdrawalHistoryService
	UserService   UserService
	WalletService WalletService
}

func NewWithdrawalHistoryHandler(
	e *echo.Echo,
	svc WithdrawalHistoryService,
	userSvc UserService,
	walletSvc WalletService,
) {
	handler := &WithdrawalHistoryHandler{
		Service:       svc,
		UserService:   userSvc,
		WalletService: walletSvc,
	}
	e.POST("/v1/wallet/withdrawal-balance", handler.RequestWithdrawal)
}

func (s *WithdrawalHistoryHandler) RequestWithdrawal(c echo.Context) (err error) {
	var request withdrawalhistory.CreateWithdrawalBalanceRequest
	err = c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, middleware.ResponseError{Message: err.Error()})
	}

	var ok bool
	if ok, err = middleware.IsRequestValid(&request); !ok {
		return c.JSON(http.StatusBadRequest, middleware.ResponseError{Message: err.Error()})
	}

	ctx := c.Request().Context()
	resp, err := s.Service.RequestWithdrawal(ctx, &request)
	if err != nil {
		return c.JSON(middleware.GetStatusCode(err), middleware.ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

package rest

import (
	"net/http"
	"withdrawal-balance/internal/rest/middleware"
	"withdrawal-balance/model"

	"github.com/labstack/echo/v4"
)

type WithdrawalHistoryService interface {
	RequestWithdrawal(c echo.Context) (err error)
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
	e.POST("/calculate-installments", handler.RequestWithdrawal)
}

func (s *WithdrawalHistoryHandler) RequestWithdrawal(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, middleware.ResponseError{Message: model.ErrInternalServerError.Error()})
}

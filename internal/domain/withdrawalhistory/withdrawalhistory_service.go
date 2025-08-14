package withdrawalhistory

import (
	"context"
	"net/http"
	"withdrawal-balance/model"

	"github.com/labstack/echo/v4"
)

type RepositoryInterface interface {
	Fetch(ctx context.Context) (res []model.WithdrawalHistory, err error)
	GetByID(ctx context.Context, id int64) (model.WithdrawalHistory, error)
	Update(ctx context.Context, req *model.WithdrawalHistory) error
	Store(ctx context.Context, req *model.WithdrawalHistory) error
	Delete(ctx context.Context, id int64) error
}

type Service struct {
	withdrawalHistoryRepo RepositoryInterface
}

func NewService(a RepositoryInterface) *Service {
	return &Service{
		withdrawalHistoryRepo: a,
	}
}

func (s *Service) Fetch(ctx context.Context) (res []model.WithdrawalHistory, err error) {
	res, err = s.withdrawalHistoryRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) GetByID(ctx context.Context, id int64) (res model.WithdrawalHistory, err error) {
	res, err = s.withdrawalHistoryRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (s *Service) Update(ctx context.Context, req *model.WithdrawalHistory) (err error) {
	return s.withdrawalHistoryRepo.Update(ctx, req)
}

func (s *Service) Store(ctx context.Context, req *model.WithdrawalHistory) (err error) {
	err = s.withdrawalHistoryRepo.Store(ctx, req)
	return
}

func (s *Service) Delete(ctx context.Context, id int64) (err error) {
	existedWithdrawalHistory, err := s.withdrawalHistoryRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedWithdrawalHistory == (model.WithdrawalHistory{}) {
		return model.ErrNotFound
	}
	return s.withdrawalHistoryRepo.Delete(ctx, id)
}

func (s *Service) RequestWithdrawal(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, model.ErrInternalServerError)
}

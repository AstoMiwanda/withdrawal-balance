package wallet

import (
	"context"
	"errors"
	"time"
	"withdrawal-balance/internal/constant"
	"withdrawal-balance/pkg"

	"withdrawal-balance/model"

	"github.com/sirupsen/logrus"
)

type RepositoryInterface interface {
	Fetch(ctx context.Context) (res []model.Wallet, err error)
	GetByID(ctx context.Context, id int64) (model.Wallet, error)
	Update(ctx context.Context, ar *model.Wallet) error
	Store(ctx context.Context, a *model.Wallet) error
	Delete(ctx context.Context, id int64) error
}

type Service struct {
	walletRepo RepositoryInterface
}

func NewService(a RepositoryInterface) *Service {
	return &Service{
		walletRepo: a,
	}
}

func (s *Service) Fetch(ctx context.Context) (res []model.Wallet, err error) {
	res, err = s.walletRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) GetByID(ctx context.Context, id int64) (res model.Wallet, err error) {
	res, err = s.walletRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (s *Service) Update(ctx context.Context, req *model.Wallet) (err error) {
	return s.walletRepo.Update(ctx, req)
}

func (s *Service) Store(ctx context.Context, req *model.Wallet) (err error) {
	err = s.walletRepo.Store(ctx, req)
	return
}

func (s *Service) Delete(ctx context.Context, id int64) (err error) {
	existedWallet, err := s.walletRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedWallet == (model.Wallet{}) {
		return model.ErrNotFound
	}
	return s.walletRepo.Delete(ctx, id)
}

func (s *Service) UpdateBalance(ctx context.Context, req *UpdateBalanceRequest) (err error) {
	if req == nil {
		return model.ErrBadParamInput
	}
	wallet, err := s.walletRepo.GetByID(ctx, req.WalletID)
	if err != nil {
		logrus.Error(err)
		return errors.New("wallet not found")
	}
	if wallet.UserID != req.UserID {
		return errors.New("invalid wallet")
	}

	if req.Type == constant.WithdrawTransactionTypeWallet {
		req.Amount = pkg.NegativeOf(req.Amount)
	}
	wallet.Balance += req.Amount

	payloadUpdate := &model.Wallet{
		ID:        req.WalletID,
		UserID:    req.UserID,
		Balance:   wallet.Balance,
		UpdatedAt: pkg.PointerTime(time.Now()),
	}
	return s.walletRepo.Update(ctx, payloadUpdate)
}

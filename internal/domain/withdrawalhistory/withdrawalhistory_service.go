package withdrawalhistory

import (
	"context"
	"errors"
	"withdrawal-balance/internal/api/xendit"
	"withdrawal-balance/internal/constant"
	"withdrawal-balance/internal/domain/wallet"
	"withdrawal-balance/model"

	"github.com/sirupsen/logrus"
)

type RepositoryInterface interface {
	Fetch(ctx context.Context) (res []model.WithdrawalHistory, err error)
	GetByID(ctx context.Context, id int64) (model.WithdrawalHistory, error)
	Update(ctx context.Context, req *model.WithdrawalHistory) error
	Store(ctx context.Context, req *model.WithdrawalHistory) error
	Delete(ctx context.Context, id int64) error
}

type UserService interface {
	GetByID(ctx context.Context, id int64) (model.User, error)
}

type WalletService interface {
	GetByID(ctx context.Context, id int64) (res model.Wallet, err error)
	UpdateBalance(ctx context.Context, req *wallet.UpdateBalanceRequest) (err error)
}

type XenditService interface {
	PayoutRequest(ctx context.Context, request *xendit.CreatePayoutRequest) (result xendit.CreatePayoutResponse, err error)
}

type Service struct {
	repo          RepositoryInterface
	userService   UserService
	walletService WalletService
	xenditService XenditService
}

func NewService(
	repo RepositoryInterface,
	userService UserService,
	walletService WalletService,
	xenditService XenditService,
) *Service {
	return &Service{
		repo:          repo,
		userService:   userService,
		walletService: walletService,
		xenditService: xenditService,
	}
}

func (s *Service) Fetch(ctx context.Context) (res []model.WithdrawalHistory, err error) {
	res, err = s.repo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (s *Service) GetByID(ctx context.Context, id int64) (res model.WithdrawalHistory, err error) {
	res, err = s.repo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (s *Service) Update(ctx context.Context, req *model.WithdrawalHistory) (err error) {
	return s.repo.Update(ctx, req)
}

func (s *Service) Store(ctx context.Context, req *model.WithdrawalHistory) (err error) {
	err = s.repo.Store(ctx, req)
	return
}

func (s *Service) Delete(ctx context.Context, id int64) (err error) {
	existedWithdrawalHistory, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedWithdrawalHistory == (model.WithdrawalHistory{}) {
		return model.ErrNotFound
	}
	return s.repo.Delete(ctx, id)
}

func (s *Service) RequestWithdrawal(ctx context.Context, req *CreateWithdrawalBalanceRequest) (result CreateWithdrawalBalanceResponse, err error) {
	if req == nil {
		return result, model.ErrBadParamInput
	}

	dataUser, err := s.userService.GetByID(ctx, req.UserID)
	if err != nil {
		logrus.Error(err)
		return result, errors.New("user not found")
	}

	dataWallet, err := s.walletService.GetByID(ctx, req.WalletID)
	if err != nil {
		logrus.Error(err)
		return result, errors.New("wallet not found")
	}

	if dataWallet.UserID != dataUser.ID {
		return result, errors.New("invalid wallet")
	}

	payloadPayoutRequest := &xendit.CreatePayoutRequest{
		BankAccountNumber: req.BankAccountNumber,
		BankAccountName:   req.BankAccountName,
		Amount:            req.Amount,
	}
	payout, err := s.xenditService.PayoutRequest(ctx, payloadPayoutRequest)
	if err != nil {
		logrus.Error(err)
		return result, model.ErrInternalServerError
	}

	payloadWithdrawalHistory := &model.WithdrawalHistory{
		UserID:               dataUser.ID,
		WalletID:             dataWallet.ID,
		Amount:               req.Amount,
		BankAccountNumber:    req.BankAccountName,
		BankCode:             req.BankCode,
		Status:               payout.Status,
		TransactionReference: payout.ReferenceId,
		PayoutId:             payout.Id,
	}
	err = s.repo.Store(ctx, payloadWithdrawalHistory)
	if err != nil {
		logrus.Error(err)
		return result, model.ErrInternalServerError
	}

	payloadUpdateBalance := &wallet.UpdateBalanceRequest{
		WalletID: dataWallet.ID,
		UserID:   dataWallet.UserID,
		Amount:   req.Amount,
		Type:     constant.WithdrawTransactionTypeWallet,
	}
	err = s.walletService.UpdateBalance(ctx, payloadUpdateBalance)
	if err != nil {
		logrus.Error(err)
		return result, model.ErrInternalServerError
	}

	result = CreateWithdrawalBalanceResponse{
		ID:              payout.Id,
		UserID:          dataWallet.UserID,
		WalletID:        dataWallet.ID,
		Amount:          req.Amount,
		Status:          payout.Status,
		ReferenceNumber: payout.ReferenceId,
	}

	return result, nil
}

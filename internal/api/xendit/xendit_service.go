package xendit

import (
	"context"
	"withdrawal-balance/internal/constant"
	"withdrawal-balance/pkg"

	"withdrawal-balance/model"

	"github.com/sirupsen/logrus"
)

type RepositoryInterface interface {
	Payout(ctx context.Context, request PayoutPayload) (*PayoutResponse, error)
}

type Service struct {
	repo RepositoryInterface
}

func NewService(repo RepositoryInterface) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) PayoutRequest(ctx context.Context, request *CreatePayoutRequest) (result CreatePayoutResponse, err error) {
	payoutPayload := PayoutPayload{
		ReferenceId: pkg.GeneratePrefixString("WD"),
		ChannelCode: constant.XenditChannelCode,
		ChannelProperties: ChannelProperties{
			AccountNumber:     request.BankAccountNumber,
			AccountHolderName: request.BankAccountName,
		},
		Amount:              request.Amount,
		Description:         constant.XenditWithdrawalDesc,
		Currency:            constant.XenditCurrency,
		ReceiptNotification: ReceiptNotification{},
		Metadata: Metadata{
			OutletNo: constant.XenditOutletNo,
		},
	}

	payout, err := s.repo.Payout(ctx, payoutPayload)
	if err != nil {
		logrus.Error(err)
		return result, model.ErrInternalServerError
	}

	result = CreatePayoutResponse{
		Id:          payout.Id,
		ReferenceId: payout.ReferenceId,
		Status:      payout.Status,
	}
	return
}

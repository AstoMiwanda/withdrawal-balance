package user

import (
	"context"

	"withdrawal-balance/model"
)

type RepositoryInterface interface {
	Fetch(ctx context.Context) (res []model.User, err error)
	GetByID(ctx context.Context, id int64) (model.User, error)
	Update(ctx context.Context, ar *model.User) error
	Store(ctx context.Context, a *model.User) error
	Delete(ctx context.Context, id int64) error
}

type Service struct {
	userRepo RepositoryInterface
}

func NewService(a RepositoryInterface) *Service {
	return &Service{
		userRepo: a,
	}
}

func (a *Service) Fetch(ctx context.Context) (res []model.User, err error) {
	res, err = a.userRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (a *Service) GetByID(ctx context.Context, id int64) (res model.User, err error) {
	res, err = a.userRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (a *Service) Update(ctx context.Context, ar *model.User) (err error) {
	return a.userRepo.Update(ctx, ar)
}

func (a *Service) Store(ctx context.Context, m *model.User) (err error) {
	err = a.userRepo.Store(ctx, m)
	return
}

func (a *Service) Delete(ctx context.Context, id int64) (err error) {
	existedUser, err := a.userRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedUser == (model.User{}) {
		return model.ErrNotFound
	}
	return a.userRepo.Delete(ctx, id)
}

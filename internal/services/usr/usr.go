// Package usr or services/usr provides a interface for interacting with a user repo and http handlers.
package usr

import (
	"context"

	"github.com/Z3DRP/zedsync/internal/domain"
	"github.com/Z3DRP/zedsync/internal/repos/usr"
	"github.com/google/uuid"
)

type UserService struct {
	repo usr.UserRepo
}

func New(r usr.UserRepo) UserService {
	return UserService{
		repo: r,
	}
}

func (us UserService) Create(ctx context.Context, usr *domain.User) (*domain.User, error) {
	var err error
	usr.UID, err = uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	u, err := us.repo.Save(ctx, usr)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us UserService) Update(ctx context.Context, usr domain.User) (*domain.User, error) {
	u, err := us.repo.Update(ctx, usr)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us UserService) Get(ctx context.Context, id string) (*domain.User, error) {
	u, err := us.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us UserService) Fetch(ctx context.Context) ([]*domain.User, error) {
	usrs, err := us.repo.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	return usrs, nil
}

func (us UserService) Delete(ctx context.Context, id string) error {
	if err := us.repo.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

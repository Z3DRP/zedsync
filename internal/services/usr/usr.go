// Package usr or services/usr provides a interface for interacting with a user repo and http handlers.
package usr

import (
	"context"
	"errors"

	"github.com/Z3DRP/zedsync/internal/auth"
	"github.com/Z3DRP/zedsync/internal/domain"
	"github.com/Z3DRP/zedsync/internal/dto"
	"github.com/Z3DRP/zedsync/internal/repos/cfg"
	"github.com/Z3DRP/zedsync/internal/repos/usr"
	v "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserService struct {
	repo      usr.UserRepo
	cfgRepo   cfg.ConfigRepo
	validator *v.Validate
}

func New(r usr.UserRepo, cfg cfg.ConfigRepo, v *v.Validate) UserService {
	return UserService{
		repo:      r,
		cfgRepo:   cfg,
		validator: v,
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

func (us UserService) Authenticate(ctx context.Context, usrname, pwd string) (bool, domain.User, error) {
	usr, err := us.repo.FetchByUsername(ctx, usrname)
	if err != nil {
		return false, domain.User{}, err
	}

	isMatch, err := auth.VerifyHash(usr.Password, pwd)
	if err != nil {
		return false, domain.User{}, err
	}

	return isMatch, usr, nil
}

func (us UserService) ValidateClaims(ctx context.Context, token string) (domain.User, error) {
	claims := auth.ParseAuthToken(token)
	usr, err := us.Get(ctx, claims.ID)

	if err != nil {
		return domain.User{}, err
	}

	if usr.Username != claims.Username {
		return domain.User{}, errors.New("claims do not match username")
	}

	if usr.Role.Name != claims.Role {
		return domain.User{}, errors.New("claims do not match role")
	}

	return *usr, nil
}

func (us UserService) SignupAdapter(signupDto dto.SignupDto) *domain.User {
	return &domain.User{
		Username:  signupDto.Username,
		Password:  signupDto.Password,
		Email:     signupDto.Email,
		FirstName: signupDto.FirstName,
		LastName:  signupDto.LastName,
	}
}

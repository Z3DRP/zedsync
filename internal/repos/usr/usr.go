// Package usr uses basic repo to make a user repository
package usr

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Z3DRP/zedsync/internal/database/store"
	"github.com/Z3DRP/zedsync/internal/domain"
	"github.com/Z3DRP/zedsync/internal/repos"
	"github.com/uptrace/bun"
)

// NOTE: i should be able to overrite a method like get if i specify it here
// maybe i should so i can use uuid to do lookups

type UserRepo struct {
	repo *repos.BasicRepo[string, domain.User]
}

func New(p store.Persister) UserRepo {
	return UserRepo{
		repo: repos.New[string, domain.User](p),
	}
}

func (u UserRepo) Get(ctx context.Context, uid string) (*domain.User, error) {
	var usr domain.User
	err := u.repo.BnDB().NewSelect().Model(&usr).Where("? = ?", bun.Ident("uid"), uid).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repos.ErrNoRecords
		}
		return nil, errors.Join(repos.ErrDBRead, err)
	}

	return &usr, nil
}

func (u UserRepo) Fetch(ctx context.Context) ([]*domain.User, error) {
	usrs, err := u.repo.List(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, repos.ErrNoRecords
		}
		return nil, errors.Join(repos.ErrDBRead, err)
	}

	return usrs, nil
}

func (u UserRepo) Save(ctx context.Context, user *domain.User) (*domain.User, error) {
	var usr = user
	tx, err := u.repo.BnDB().BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, errors.Join(repos.ErrFailedTransaction, err)
	}

	err = tx.NewInsert().Model(&usr).Returning("*").Scan(ctx, &usr)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return nil, errors.Join(repos.ErrFailedTransaction, err)
		}
		return nil, err
	}

	if err = tx.Commit(); err != nil {
		return nil, err
	}

	return usr, nil
}

func (u UserRepo) Update(ctx context.Context, usr domain.User) (*domain.User, error) {
	var user domain.User
	tx, err := u.repo.BnDB().BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, errors.Join(repos.ErrFailedTransaction, err)
	}

	err = tx.NewUpdate().Model(&user).Where("? = ?", bun.Ident("uid"), usr.ID).Scan(ctx, &user)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return nil, errors.Join(repos.ErrFailedRollback, err)
		}
		return nil, errors.Join(repos.ErrDBWrite, err)
	}

	return &user, nil
}

func (u UserRepo) Delete(ctx context.Context, id string) error {
	var usr domain.User
	tx, err := u.repo.BnDB().BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return errors.Join(repos.ErrFailedTransaction, err)
	}

	_, err = tx.NewDelete().Model(&usr).Where("? = ?", bun.Ident("uid"), id).Exec(ctx)
	if err != nil {
		return errors.Join(repos.ErrDBDelete, err)
	}

	return nil
}

func (u UserRepo) FetchByUsername(ctx context.Context, usrname string) (domain.User, error) {
	var usr domain.User
	err := u.repo.BnDB().NewSelect().Model(&domain.User{}).Where("? = ?", bun.Ident("username"), usrname).Scan(ctx, &usr)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, fmt.Errorf("could not find user %w", err)
		}
		return domain.User{}, err
	}

	return domain.User{}, nil
}

package cfg

import (
	"context"

	"github.com/Z3DRP/zedsync/internal/database/store"
	"github.com/Z3DRP/zedsync/internal/domain"
	"github.com/Z3DRP/zedsync/internal/repos"
	"github.com/uptrace/bun"
)

type ConfigRepo struct {
	repo *repos.BasicRepo[int, domain.Role]
}

func New(p store.Persister) ConfigRepo {
	return ConfigRepo{
		repo: repos.New[int, domain.Role](p),
	}
}

func (c ConfigRepo) NewRole(ctx context.Context, name string) (domain.Role, error) {
	r := domain.Role{
		Name: name,
	}

	role, err := c.repo.Save(ctx, &r)
	if err != nil {
		return domain.Role{}, err
	}

	return *role, nil
}

func (c ConfigRepo) GetRole(ctx context.Context, name string) (domain.Role, error) {
	var role domain.Role
	err := c.repo.BnDB().NewSelect().Model(&role).Where("? = ?", bun.Ident("name"), name).Scan(ctx, &role)

	if err != nil {
		return domain.Role{}, err
	}

	return role, nil
}

func (c ConfigRepo) GetRoleByID(ctx context.Context, id int) (domain.Role, error) {
	role, err := c.repo.Get(ctx, id)
	if err != nil {
		return domain.Role{}, err
	}

	return *role, nil
}

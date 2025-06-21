// Package usr uses basic repo to make a user repository
package usr

import (
	"github.com/Z3DRP/zedsync/internal/database/store"
	"github.com/Z3DRP/zedsync/internal/domain"
	"github.com/Z3DRP/zedsync/internal/repos"
)

// NOTE: i should be able to overrite a method like get if i specify it here
// maybe i should so i can use uuid to do lookups

type UserRepo struct {
	*repos.BasicRepo[int, domain.Users]
}

func New(p store.Persister) *UserRepo {
	return &UserRepo{
		BasicRepo: repos.New[int, domain.Users](p),
	}
}

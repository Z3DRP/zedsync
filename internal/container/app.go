// Package container provices a app container for repos and services
package container

import (
	"fmt"

	"github.com/Z3DRP/zedsync/internal/crane"
	"github.com/Z3DRP/zedsync/internal/database/store"
	urepo "github.com/Z3DRP/zedsync/internal/repos/usr"
	"github.com/Z3DRP/zedsync/internal/services"
	"github.com/Z3DRP/zedsync/internal/services/usr"
)

type Container struct {
	store     store.Persister
	logger    *crane.Zlogrus
	endpoints []services.API
}

func (c Container) Endpoints() []services.API {
	return c.endpoints
}

func New(s store.Persister, l *crane.Zlogrus) *Container {
	return &Container{
		store:  s,
		logger: l,
	}
}

func (c Container) RegisterServices(names []string) error {
	for _, name := range names {
		service, err := c.createService(name)
		if err != nil {
			return err
		}
		c.endpoints = append(c.endpoints, service)
	}

	return nil
}

func (c Container) createService(name string) (services.API, error) {
	switch name {
	case "user":
		userRepo := urepo.New(c.store)
		usrService := usr.New(userRepo)
		return usr.Initialize(usrService, c.logger), nil
	default:
		return nil, fmt.Errorf("unknown service %v", name)
	}
}

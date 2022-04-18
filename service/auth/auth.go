package service

import (
	"context"

	klog "github.com/go-kit/kit/log"
	"github.com/hambyhacks/CrimsonIMS/app/models"
)

type AuthService interface {
	AddUser(ctx context.Context, user models.User) (string, error)
	GetByEmail(ctx context.Context, email string) (interface{}, error)
	UpdateUser(ctx context.Context, user models.User) (string, error)
}

type AuthSrv struct {
	repo   AuthRepository
	logger klog.Logger
}

func NewAuthSrv(repo AuthRepository, logger klog.Logger) AuthService {
	return &AuthSrv{
		repo:   repo,
		logger: logger,
	}
}

// AddUser implements AuthService
func (*AuthSrv) AddUser(ctx context.Context, user models.User) (string, error) {
	panic("unimplemented")
}

// GetByEmail implements AuthService
func (*AuthSrv) GetByEmail(ctx context.Context, email string) (interface{}, error) {
	panic("unimplemented")
}

// UpdateUser implements AuthService
func (*AuthSrv) UpdateUser(ctx context.Context, user models.User) (string, error) {
	panic("unimplemented")
}

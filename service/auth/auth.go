package service

import (
	"context"
	"log"
	"time"

	klog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/hambyhacks/CrimsonIMS/app/models"
)

type AuthService interface {
	AddUser(ctx context.Context, user models.User) (string, error)
	GetByEmail(ctx context.Context, email string) (interface{}, error)
	UpdateUser(ctx context.Context, user models.User) (string, error)
}

type AuthServ struct {
	repo   AuthRepository
	logger klog.Logger
}

func NewAuthSrv(repo AuthRepository, logger klog.Logger) AuthService {
	return &AuthServ{
		repo:   repo,
		logger: logger,
	}
}

// AddUser implements AuthService
func (a *AuthServ) AddUser(ctx context.Context, user models.User) (string, error) {
	log.Println("[i] Endpoint: /v1/admin/users/add")
	logger := klog.With(a.logger, "method", "add user")
	msg := "successfully added user"
	userDetails := models.User{
		ID:        user.ID,
		CreatedAt: time.Now().UTC(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Activated: user.Activated,
	}

	err := a.repo.AddUser(ctx, userDetails)
	if err != nil {
		level.Error(logger).Log("repository-error", err)
		return "unable to process request", err
	}
	return msg, nil
}

// GetByEmail implements AuthService
func (a *AuthServ) GetByEmail(ctx context.Context, email string) (interface{}, error) {
	log.Println("[i] Endpoint: /v1/admin/users/email")
	logger := klog.With(a.logger, "method", "get user by email")

	user, err := a.repo.GetByEmail(ctx, email)
	if err != nil {
		level.Error(logger).Log("repository-error", err)
		return nil, err
	}
	return user, nil
}

// UpdateUser implements AuthService
func (a *AuthServ) UpdateUser(ctx context.Context, user models.User) (string, error) {
	log.Println("[i] Endpoint: /v1/admin/user/update/")
	logger := klog.With(a.logger, "method", "update user by ")
	msg := "successfully updated user"
	userDetails := models.User{
		ID:        user.ID,
		CreatedAt: time.Now().UTC(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Activated: user.Activated,
	}

	msg, err := a.repo.UpdateUser(ctx, userDetails)
	if err != nil {
		level.Error(logger).Log("repository-error", err)
		return "unable to process request", err
	}
	return msg, nil
}

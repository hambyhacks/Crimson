package service

import (
	"context"
	"database/sql"
	"errors"

	klog "github.com/go-kit/log"
	"github.com/hambyhacks/CrimsonIMS/app/models"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
	ErrRepo           = errors.New("unable to process database query")
)

type AuthRepository interface {
	AddUser(ctx context.Context, user models.User) error
	GetByEmail(ctx context.Context, email string) (interface{}, error)
	UpdateUser(ctx context.Context, user models.User) (string, error)
}

type authRepo struct {
	db     *sql.DB
	logger klog.Logger
}

func NewAuthRepo(db *sql.DB, logger klog.Logger) (AuthRepository, error) {
	return &authRepo{db: db, logger: klog.With(logger, "repo", "auth_svc")}, nil
}

func (a *authRepo) AddUser(ctx context.Context, user models.User) error {
	panic("unimplemented")
}

// GetByEmail implements AuthRepository
func (*authRepo) GetByEmail(ctx context.Context, email string) (interface{}, error) {
	panic("unimplemented")
}

// UpdateUser implements AuthRepository
func (*authRepo) UpdateUser(ctx context.Context, user models.User) (string, error) {
	panic("unimplemented")
}

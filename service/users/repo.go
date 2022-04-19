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
	ErrNotFound       = errors.New("user not found")
)

type UserRepository interface {
	AddUser(ctx context.Context, user models.User) error
	GetByEmail(ctx context.Context, email string) (interface{}, error)
	UpdateUser(ctx context.Context, user models.User) (string, error)
}

type userRepo struct {
	db     *sql.DB
	logger klog.Logger
}

func NewUserRepo(db *sql.DB, logger klog.Logger) (UserRepository, error) {
	return &userRepo{db: db, logger: klog.With(logger, "repo", "auth_svc")}, nil
}

func (r *userRepo) AddUser(ctx context.Context, user models.User) error {
	q := `INSERT INTO users (username, email, password_hash, activated) 
		  VALUES ($1,$2,$3,$4)
		  RETURNING id, created_at`
	args := []interface{}{user.Name, user.Email, user.Password.Hash, user.Activated}

	err := r.db.QueryRowContext(ctx, q, args...).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return ErrRepo
		}
	}
	return nil
}

// GetByEmail implements AuthRepository
func (r *userRepo) GetByEmail(ctx context.Context, email string) (interface{}, error) {
	q := `SELECT id, created_at, username, email, password_hash, activated 
		  FROM users
		  WHERE email = $1`
	var user models.User
	err := r.db.QueryRowContext(ctx, q, email).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.Name,
		&user.Email,
		&user.Password.Hash,
		&user.Activated,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, ErrRepo
		}
	}
	return &user, nil
}

// UpdateUser implements AuthRepository
func (r *userRepo) UpdateUser(ctx context.Context, user models.User) (string, error) {
	q := `UPDATE users
		  SET username = $1, email = $2
		      password_hash = $3, activated = $4,
		  WHERE id = $5 RETURNING id`

	args := []interface{}{user.Name, user.Email, user.Password.Hash, user.Activated, user.ID}
	err := r.db.QueryRowContext(ctx, q, args...).Scan(&user.ID)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "users_email_key"`:
			return "unable to process request due to duplicate email", ErrDuplicateEmail
		default:
			return "unable to process request", ErrRepo
		}
	}
	return "", nil
}

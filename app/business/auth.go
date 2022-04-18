package app

import (
	"errors"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"

	"github.com/hambyhacks/CrimsonIMS/app/models"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Password models.Password
	logger   log.Logger
}

func (p *Password) SetPassword(plaintext string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintext), 12)
	if err != nil {
		level.Error(p.logger).Log("auth-err", err)
		return err
	}
	p.Password.Plaintext = &plaintext
	p.Password.Hash = hash
	return nil
}

func (p *Password) CheckHash(plaintext string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.Password.Hash, []byte(plaintext))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

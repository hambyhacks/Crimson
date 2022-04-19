package auth

import "github.com/hambyhacks/CrimsonIMS/app/models"

type (
	AddUserRequest struct {
		User models.User
	}

	GetUserByEmailRequest struct {
		Email string `json:"email"`
	}

	UpdateUserRequest struct {
		User models.User
	}
)

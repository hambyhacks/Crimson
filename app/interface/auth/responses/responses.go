package responses

type (
	AddUserResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error,omitempty"`
	}

	GetUserByEmailResponse struct {
		User interface{} `json:"user,omitempty"`
		Err  error       `json:"error,omitempty"`
	}

	UpdateUserResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error,omitempty"`
	}
)

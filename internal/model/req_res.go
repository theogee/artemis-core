package model

// error messages to be used by the client
const (
	UsernameCantBeEmpty  = "username can't be empty"
	PasswordCantBeEmpty  = "password can't be empty"
	UsernameAlreadyExist = "username already exist"
	EmailCantBeEmpty     = "email can't be empty"
	IncorrectCredential  = "incorrect username or password"
)

const (
	UnauthorizedAccess = "unauthorized access"
)

const (
	AdminCreatedSuccessfully = "new admin with username: %v has been created"
)

type (
	DefaultResponse struct {
		Success   bool        `json:"success"`
		ServError []string    `json:"servError,omitempty"`
		Data      interface{} `json:"data,omitempty"`
	}

	LoginAsAdminRequest struct {
		Username string
		Password string
	}

	LoginAsAdminResponse struct {
		ErrMessage    []string `json:"errMessage,omitempty"`
		Message       []string `json:"message,omitempty"`
		UsernameError string   `json:"usernameError,omitempty"`
		PasswordError string   `json:"passwordError,omitempty"`
	}

	RegisterAsAdminRequest struct {
		Username string
		Password string
		Email    string
	}

	RegisterAsAdminResponse struct {
		ErrMessage []string `json:"errMessage,omitempty"`
		Message    []string `json:"message,omitempty"`
	}

	AuthenticateResponse struct {
		ErrMessage []string `json:"errMessage,omitempty"`
		Message    []string `json:"message,omitempty"`
	}

	LogoutResponse struct {
		ErrMessage []string `json:"errMessage,omitempty"`
		Message    []string `json:"message,omitempty"`
	}
)

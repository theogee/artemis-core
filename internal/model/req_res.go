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
	ForbiddenAccess    = "forbidden access"
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

	LoginRequest struct {
		Username string
		Password string
	}

	LoginAsAdminRequest struct {
		Username string
		Password string
	}

	LoginResponse struct {
		ErrMessage    []string `json:"errMessage,omitempty"`
		Message       []string `json:"message,omitempty"`
		UsernameError string   `json:"usernameError,omitempty"`
		PasswordError string   `json:"passwordError,omitempty"`
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

	AuthorizeResponse struct {
		ErrMessage []string `json:"errMessage,omitempty"`
		Message    []string `json:"message,omitempty"`
	}

	LogoutResponse struct {
		ErrMessage []string `json:"errMessage,omitempty"`
		Message    []string `json:"message,omitempty"`
	}

	GetMetaResponse struct {
		UserType string `json:"userType,omitempty"`
	}

	GetStudentsRequest struct {
		Limit        int64
		Page         int64
		SGUMajorID   int64
		ExchangeYear int64
		StudentID    int64
		Name         string
	}

	GetStudentsResponse struct {
		ErrMessage   []string         `json:"errMessage,omitempty"`
		Students     []*StudentSimple `json:"students"`
		TotalStudent int              `json:"totalStudent,omitempty"`
	}

	StudentSimple struct {
		StudentID uint32 `json:"studentID"`
		Name      string `json:"name"`
		SGUMajor  string `json:"sguMajor"`
		SGUEmail  string `json:"sguEmail"`
		// set to +62 if +49 is not available
		MobilePhone  string `json:"mobilePhone"`
		ExchangeYear int16  `json:"exchangeYear"`
	}

	GetSGUMajorsResponse struct {
		ErrMessage []string    `json:"errMessage,omitempty"`
		Message    []string    `json:"message,omitempty"`
		Majors     []*SGUMajor `json:"majors"`
	}

	GetExchangeYearResponse struct {
		ErrMessage   []string `json:"errMessage,omitempty"`
		Message      []string `json:"message,omitempty"`
		ExchangeYear []int    `json:"exchangeYear"`
	}

	RegisterStudentByCSV struct {
		ErrMessage []string `json:"errMessage,omitempty"`
		Message    []string `json:"message,omitempty"`
	}
)

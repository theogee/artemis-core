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
	AdminCreatedSuccessfully         = "new admin with username: %v has been created"
	StudentIDDoesNotExist            = "studentID: %v doesn't exist"
	StudentWithIDUpdatedSuccessfully = "studentID: %v has been updated"
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

	GetStudentByIDResponse struct {
		ErrMessage []string       `json:"errMessage,omitempty"`
		Student    *StudentDetail `json:"studentData,omitempty"`
	}

	StudentDetail struct {
		GivenName                 string `json:"givenName"`
		Surname                   string `json:"surname"`
		Gender                    string `json:"gender"`
		SGUMajor                  string `json:"sguMajor"`
		SGUMajorInitial           string `json:"sguMajorInitial"`
		FHDepartment              string `json:"fhDepartment"`
		StudentID                 uint32 `json:"studentID"`
		DateOfBirth               string `json:"dateOfBirth"`
		CityOfBirth               string `json:"cityOfBirth"`
		PassportNumber            string `json:"passportNumber"`
		DateOfIssue               string `json:"dateOfIssue"`
		DateOfExpiry              string `json:"dateOfExpiry"`
		IssuingOffice             string `json:"issuingOffice"`
		PrivateEmail              string `json:"privateEmail"`
		SGUEmail                  string `json:"sguEmail"`
		Username                  string `json:"username"`
		FHEmail                   string `json:"fhEmail"`
		IBAN                      string `json:"iban"`
		MobilePhone               string `json:"mobilePhone"`
		MobilePhoneDE             string `json:"mobilePhoneDE"`
		CurrentAddress            string `json:"currentAddress"`
		CurrentPostcode           string `json:"currentPostcode"`
		CurrentCity               string `json:"currentCity"`
		CoName                    string `json:"coName"`
		InternshipCompany         string `json:"internshipCompany"`
		InternshipStartDate       string `json:"internshipStartDate"`
		InternshipEndDate         string `json:"internshipEndDate"`
		InternshipCompanyAddress  string `json:"internshipCompanyAddress"`
		InternshipCompanyPostcode string `json:"internshipCompanyPostcode"`
		InternshipCompanyCity     string `json:"internshipCompanyCity"`
		InternshipSupervisorName  string `json:"internshipSupervisorName"`
		InternshipSupervisorEmail string `json:"internshipSupervisorEmail"`
		InternshipSupervisorPhone string `json:"internshipSupervisorPhone"`
		ExchangeYear              int16  `json:"exchangeYear"`
	}

	UpdateStudentByIDRequest struct {
		StudentID                 int64
		MobilePhone               string
		MobilePhoneDE             string
		PrivateEmail              string
		CurrentAddress            string
		CurrentPostcode           interface{}
		CurrentCity               string
		CoName                    string
		InternshipCompany         string
		InternshipStartDate       interface{}
		InternshipEndDate         interface{}
		InternshipCompanyAddress  string
		InternshipCompanyPostcode interface{}
		InternshipCompanyCity     string
		InternshipSupervisorName  string
		InternshipSupervisorEmail string
		InternshipSupervisorPhone string
	}

	UpdateStudentByIDResponse struct {
		ErrMessage []string `json:"errMessage,omitempty"`
		Message    []string `json:"message,omitempty"`
	}
)

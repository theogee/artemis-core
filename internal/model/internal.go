package model

import "database/sql"

var (
	SGUMajors = map[string]int{
		"Mechatronics":                    1,
		"Business Management":             2,
		"Global Strategic Communications": 3,
		"Information Technology":          4,
		"Biomedical Engineering":          5,
		"Pharmaceutical Engineering":      6,
		"Food Technology":                 7,
		"Accounting":                      8,
		"Industrial Engineering":          9,
	}

	FHDepartments = map[string]int{
		"Wirtschaftsingenieurwesen":                1,
		"Business Administration with Informatics": 2,
	}
)

const (
	UserTypeStudent = "student"
	UserTypeAdmin   = "admin"
)

type (
	Student struct {
		GivenName                 string         `csv:"given_name" db:"given_name"`
		Surname                   sql.NullString `csv:"surname" db:"surname"`
		Gender                    sql.NullString `csv:"gender" db:"gender"`
		SGUMajorID                uint32         `db:"sgu_major_id"`
		SGUMajor                  string         `csv:"sgu_major" db:"sgu_major"`
		FHDepartmentID            uint32         `db:"fh_department_id"`
		FHDepartment              string         `csv:"fh_department"`
		StudentID                 uint32         `csv:"student_id" db:"student_id"`
		DateOfBirth               sql.NullString `csv:"date_of_birth" db:"date_of_birth"`
		CityOfBirth               sql.NullString `csv:"city_of_birth" db:"city_of_birth"`
		PassportNumber            sql.NullString `csv:"passport_number" db:"passport_number"`
		DateOfIssue               sql.NullString `csv:"date_of_issue" db:"date_of_issue"`
		DateOfExpiry              sql.NullString `csv:"date_of_expiry" db:"date_of_expiry"`
		IssuingOffice             sql.NullString `csv:"issuing_office" db:"issuing_office"`
		PrivateEmail              sql.NullString `csv:"private_email" db:"private_email"`
		SGUEmail                  sql.NullString `csv:"sgu_email" db:"sgu_email"`
		Username                  string         `csv:"username" db:"username"`
		Password                  string         `csv:"password" db:"password"`
		FHEmail                   sql.NullString `csv:"fh_email" db:"fh_email"`
		IBAN                      sql.NullString `csv:"iban" db:"iban"`
		MobilePhone               sql.NullString `csv:"mobile_phone" db:"mobile_phone"`
		MobilePhoneDE             sql.NullString `csv:"mobile_phone_de" db:"mobile_phone_de"`
		CurrentAddress            sql.NullString `db:"current_address"`
		CurrentPostcode           sql.NullString `db:"current_postcode"`
		CurrentCity               sql.NullString `db:"current_city"`
		CoName                    sql.NullString `db:"co_name"`
		InternshipCompany         sql.NullString `db:"internship_company"`
		InternshipStartDate       sql.NullString `db:"internship_start_date"`
		InternshipEndDate         sql.NullString `db:"internship_end_date"`
		InternshipCompanyAddress  sql.NullString `db:"internship_company_address"`
		InternshipCompanyPostcode sql.NullString `db:"internship_company_postcode"`
		InternshipCompanyCity     sql.NullString `db:"internship_company_city"`
		InternshipSupervisorName  sql.NullString `db:"internship_supervisor_name"`
		InternshipSupervisorEmail sql.NullString `db:"internship_supervisor_email"`
		InternshipSupervisorPhone sql.NullString `db:"internship_supervisor_phone"`
		DateOfDeparture           sql.NullString `db:"date_of_departure"`
		DepartureAirline          sql.NullString `db:"departure_airline"`
		DepartureFlightNumber     sql.NullString `db:"departure_flight_number"`
		DateOfArrival             sql.NullString `db:"date_of_arrival"`
		ArrivalFlightNumber       sql.NullString `db:"arrival_flight_number"`
		ArrivalAirport            sql.NullString `db:"arrival_airport"`
		SGUWPickup                sql.NullBool   `db:"sguw_pickup"`
		ExchangeYear              sql.NullInt16  `db:"exchange_year"`
	}

	StudentCSV struct {
		GivenName      string `csv:"given_name" db:"given_name"`
		Surname        string `csv:"surname" db:"surname"`
		Gender         string `csv:"gender" db:"gender"`
		SGUMajorID     uint32 `db:"sgu_major_id"`
		SGUMajor       string `csv:"sgu_major" db:"sgu_major"`
		FHDepartmentID uint32 `db:"fh_department_id"`
		FHDepartment   string `csv:"fh_department"`
		StudentID      uint32 `csv:"student_id" db:"student_id"`
		DateOfBirth    string `csv:"date_of_birth" db:"date_of_birth"`
		CityOfBirth    string `csv:"city_of_birth" db:"city_of_birth"`
		PassportNumber string `csv:"passport_number" db:"passport_number"`
		DateOfIssue    string `csv:"date_of_issue" db:"date_of_issue"`
		DateOfExpiry   string `csv:"date_of_expiry" db:"date_of_expiry"`
		IssuingOffice  string `csv:"issuing_office" db:"issuing_office"`
		PrivateEmail   string `csv:"private_email" db:"private_email"`
		SGUEmail       string `csv:"sgu_email" db:"sgu_email"`
		Username       string `csv:"username" db:"username"`
		Password       string `csv:"password" db:"password"`
		FHEmail        string `csv:"fh_email" db:"fh_email"`
		IBAN           string `csv:"iban" db:"iban"`
		MobilePhone    string `csv:"mobile_phone" db:"mobile_phone"`
		MobilePhoneDE  string `csv:"mobile_phone_de" db:"mobile_phone_de"`
	}

	Admin struct {
		AdminID     uint32         `db:"admin_id" json:"admin_id"`
		Username    string         `db:"username" json:"username"`
		Password    string         `db:"password" json:"-"`
		Email       string         `db:"email" json:"email"`
		MobilePhone sql.NullString `db:"mobile_phone" json:"mobile_phone"`
		Address     sql.NullString `db:"address" json:"address"`
		City        sql.NullString `db:"city" json:"city"`
		Postcode    sql.NullInt16  `db:"postcode" json:"postcode"`
	}

	UserCache struct {
		UID      string `json:"uid"`
		UserType string `josn:"userType"`
	}

	SGUMajor struct {
		MajorID   int    `db:"major_id" json:"majorID"`
		MajorName string `db:"major_name" json:"majorName"`
		MajorCode string `db:"major_code" json:"majorCode"`
	}
)

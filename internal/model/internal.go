package model

type (
	Student struct {
		GivenName      string `csv:"given_name"`
		Surname        string `csv:"surname"`
		Gender         string `csv:"gender"`
		SGUMajor       string `csv:"sgu_major"`
		FHDepartment   string `csv:"fh_department"`
		StudentID      uint32 `csv:"student_id"`
		DateOfBirth    string `csv:"date_of_birth"`
		CityOfBirth    string `csv:"city_of_birth"`
		PassportNumber string `csv:"passport_number"`
		DateOfIssue    string `csv:"date_of_issue"`
		DateOfExpiry   string `csv:"date_of_expiry"`
		IssuingOffice  string `csv:"issuing_office"`
		PrivateEmail   string `csv:"private_email"`
		SGUEmail       string `csv:"sgu_email"`
		Username       string `csv:"username"`
		Password       string `csv:"password"`
		FHEmail        string `csv:"fh_email"`
		IBAN           string `csv:"iban"`
		MobilePhone    string `csv:"mobile_phone"`
		MobilePhoneDE  string `csv:"mobile_phone_de"`
	}
)

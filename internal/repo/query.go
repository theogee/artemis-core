package artemis

const (
	// to be completed by the process
	InsertStudentsQuery = `
		INSERT INTO students (given_name,surname,gender,sgu_major_id,fh_department_id,student_id,date_of_birth,city_of_birth,passport_number,date_of_issue,date_of_expiry,issueing_office,private_email,sgu_email,username,password,fh_email,iban,mobile_phone,mobile_phone_de)
		VALUES
	`

	GetAdminByUsernameQuery = `
		SELECT admin_id, username, password, email, mobile_phone, address, city, postcode
		FROM admins
		WHERE username = $1
	`

	InsertAdminQuery = `
		INSERT INTO admins (admin_id, username, password, email)
		VALUES (DEFAULT, $1, $2, $3)
	`
)

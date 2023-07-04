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

	GetStudentByUsernameQuery = `
		SELECT *
		FROM students
		WHERE username = $1
	`

	GetStudentsQuery = `
		SELECT s.student_id, s.given_name, s.surname, m.major_name as sgu_major, s.sgu_email, s.mobile_phone_de, s.mobile_phone, s.exchange_year
		FROM students s, sgu_majors m
		WHERE s.sgu_major_id = m.major_id
	`

	GetStudentsCountQuery = `
		SELECT COUNT(student_id)
		FROM students s
		WHERE 1 = 1
	`

	GetSGUMajorsQuery = `
		SELECT major_id, major_name, major_code
		FROM sgu_majors
	`
)

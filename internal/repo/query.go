package artemis

const (
	// to be completed by the process
	InsertStudentsQuery = `
		INSERT INTO students (given_name,surname,gender,sgu_major_id,fh_department_id,student_id,date_of_birth,city_of_birth,passport_number,date_of_issue,date_of_expiry,issuing_office,private_email,sgu_email,username,password,fh_email,iban,mobile_phone,mobile_phone_de, exchange_year)
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

	GetExchangeYearQuery = `
		SELECT exchange_year
		FROM students
		GROUP BY exchange_year
	`

	GetStudentByIDQuery = `
		SELECT 
		s.student_id,
		s.username,
		s.gender,
		s.mobile_phone,
		s.mobile_phone_de,
		s.given_name,
		s.surname,
		sgu.major_name as sgu_major,
		sgu.major_code as sgu_major_code,
		fh.department_name as fh_department,
		s.private_email,
		s.sgu_email,
		s.fh_email,
		s.iban,
		s.exchange_year,
		s.current_address,
		s.current_postcode,
		s.current_city,
		s.co_name,
		s.date_of_birth,
		s.city_of_birth,
		s.passport_number,
		s.date_of_issue,
		s.date_of_expiry,
		s.issuing_office,
		s.internship_company,
		s.internship_start_date,
		s.internship_end_date,
		s.internship_company_address,
		s.internship_company_postcode,
		s.internship_company_city,
		s.internship_supervisor_name,
		s.internship_supervisor_email,
		s.internship_supervisor_phone
		FROM students s, sgu_majors sgu, fh_departments fh 
		WHERE s.student_id = $1 AND s.sgu_major_id = sgu.major_id AND fh.department_id = s.fh_department_id
	`

	UpdateStudentByIDQuery = `
		UPDATE students
		SET
		mobile_phone = $1,
		mobile_phone_de = $2,
		private_email = $3,
		current_address = $4,
		current_postcode = $5,
		current_city = $6,
		co_name = $7,
		internship_company = $8,
		internship_start_date = $9,
		internship_end_date = $10,
		internship_company_address = $11,
		internship_company_postcode = $12,
		internship_company_city = $13,
		internship_supervisor_name = $14,
		internship_supervisor_email = $15,
		internship_supervisor_phone = $16
		WHERE student_id = $17
	`
)

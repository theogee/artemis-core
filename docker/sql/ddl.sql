CREATE TABLE "students" (
	student_id INT NOT NULL PRIMARY KEY,
	username VARCHAR UNIQUE NOT NULL,
	password VARCHAR NOT NULL,
	private_email VARCHAR,
	sgu_email VARCHAR,
	fh_email VARCHAR,
	gender VARCHAR,
	given_name VARCHAR NOT NULL,
	surname VARCHAR,
	sgu_major_id INT NOT NULL,
	fh_department_id INT NOT NULL,
	iban VARCHAR,
	mobile_phone VARCHAR,
	mobile_phone_de VARCHAR,
	current_address TEXT,
	current_postcode INT,
	current_city VARCHAR,
	co_name VARCHAR,
	-- support for passport data
	date_of_birth DATE,
	city_of_birth VARCHAR,
	passport_number VARCHAR,
	date_of_issue DATE,
	date_of_expiry DATE,
	issueing_office VARCHAR,
	-- support for internship data
	internship_company VARCHAR,
	internship_start_date DATE,
	internship_end_date DATE,
	internship_company_address VARCHAR,
	internship_company_postcode INT,
	internship_company_city VARCHAR,
	internship_supervisor_name VARCHAR,
	internship_supervisor_email VARCHAR,
	internship_supervisor_phone VARCHAR,
	-- support for departure data
	date_of_departure TIMESTAMP,
	departure_airline VARCHAR,
	departure_flight_number VARCHAR,
	date_of_arrival TIMESTAMP,
	arrival_flight_number VARCHAR,
	arrival_airport VARCHAR,
	sguw_pickup BOOLEAN,
	exchange_year INT DEFAULT date_part('year', CURRENT_DATE)
);

CREATE TABLE "sgu_majors" (
	major_id SERIAL PRIMARY KEY,
	major_name VARCHAR NOT NULL,
	major_code VARCHAR
);

CREATE TABLE "fh_departments" (
	department_id SERIAL PRIMARY KEY,
	department_name VARCHAR NOT NULL
);

CREATE TABLE "admins" (
	admin_id SERIAL PRIMARY KEY,
	username VARCHAR NOT NULL UNIQUE,
	password VARCHAR NOT NULL,
	email VARCHAR NOT NULL,
	mobile_phone VARCHAR,
	address VARCHAR,
	city VARCHAR,
	postcode INT
);
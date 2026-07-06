--Profile table on database
CREATE TABLE IF NOT EXISTS profile (
	profile_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	user_ID INT REFERENCES users(user_id),
	f_name VARCHAR(100),
	l_name VARCHAR(100),
	user_name VARCHAR(100) UNIQUE,
	date_of_birth DATE,
	is_active BOOLEAN DEFAULT TRUE,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    modified_at TIMESTAMP WITH TIME ZONE
)
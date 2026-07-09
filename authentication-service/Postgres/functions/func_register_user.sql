CREATE OR REPLACE FUNCTION func_register_user(in_f_name varchar(100), in_l_name varchar(100), in_email varchar(100), in_hashed_password varchar(255), in_username varchar(100), in_dob date)
RETURNS VOID 
AS 
$$
	DECLARE
			new_user_id INT;
	BEGIN
		INSERT INTO users(email,hashed_password)
		VALUES(in_email,in_hashed_password)
		RETURNING user_id INTO new_user_id;

		INSERT INTO profiles(user_id,f_name,l_name,user_name,date_of_birth)
		VALUES(new_user_id,in_f_name,in_l_name,in_user_name,in_dob);
	END;
$$ LANGUAGE plpgsql;

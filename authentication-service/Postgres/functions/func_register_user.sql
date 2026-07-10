-- FUNCTION: public.func_register_user(character varying, character varying, character varying, character varying, character varying, date)

-- DROP FUNCTION IF EXISTS public.func_register_user(character varying, character varying, character varying, character varying, character varying, date);

CREATE OR REPLACE FUNCTION public.func_register_user(
	in_f_name character varying,
	in_l_name character varying,
	in_email character varying,
	in_hashed_password character varying,
	in_username character varying,
	in_dob date)
    RETURNS void
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
	DECLARE
			new_user_id INT;
	BEGIN
		INSERT INTO users(email,username,hashed_password)
		VALUES(in_email,in_username,in_hashed_password)
		RETURNING user_id INTO new_user_id;

		INSERT INTO profiles(user_id,f_name,l_name,date_of_birth)
		VALUES(new_user_id,in_f_name,in_l_name,in_dob);
	END;
$BODY$;

ALTER FUNCTION public.func_register_user(character varying, character varying, character varying, character varying, character varying, date)
    OWNER TO postgres;

-- FUNCTION: public.func_fetch_profile(integer, text)

-- DROP FUNCTION IF EXISTS public.func_fetch_profile(integer, text);

CREATE OR REPLACE FUNCTION public.func_fetch_profile(
	in_user_id integer,
	in_user_uid text)
    RETURNS TABLE(user_email character varying, full_name text, dob text) 
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
    ROWS 1000

AS $BODY$
BEGIN
	RETURN QUERY
	SELECT
		u.email AS user_email,
		CONCAT_WS(' ',p.f_name, p.l_name) AS full_name,
		date_of_birth::text AS dob
	FROM users u
	JOIN profiles p ON u.user_id = p.user_id
	WHERE u.user_uid = in_user_uid::uuid AND p.user_id = in_user_id;
END;
$BODY$;

ALTER FUNCTION public.func_fetch_profile(integer, text)
    OWNER TO postgres;

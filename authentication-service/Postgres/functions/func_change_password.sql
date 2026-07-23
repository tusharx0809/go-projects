-- FUNCTION: public.func_change_password(text, integer, text)

-- DROP FUNCTION IF EXISTS public.func_change_password(text, integer, text);

CREATE OR REPLACE FUNCTION public.func_change_password(
	new_password_hash text,
	in_user_id integer,
	in_user_uid text)
    RETURNS void
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
BEGIN
	UPDATE users
	SET hashed_password = new_password_hash,
	modified_at = NOW()
	WHERE user_id = in_user_id AND user_uid = in_user_uid::uuid;
END;
$BODY$;

ALTER FUNCTION public.func_change_password(text, integer, text)
    OWNER TO postgres;

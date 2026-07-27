-- FUNCTION: public.func_add_refresh_token(integer, text)

-- DROP FUNCTION IF EXISTS public.func_add_refresh_token(integer, text);

CREATE OR REPLACE FUNCTION public.func_add_refresh_token(
	in_user_id integer,
	in_token text)
    RETURNS void
    LANGUAGE 'plpgsql'
    COST 100
    VOLATILE PARALLEL UNSAFE
AS $BODY$
DECLARE
	prev_refresh_token_id INT;
BEGIN
	SELECT refresh_token_id INTO prev_refresh_token_id FROM refresh_tokens WHERE user_id = in_user_id;
	IF prev_refresh_token_id IS NOT NULL THEN
		UPDATE refresh_tokens
		SET refresh_token = in_token,
			expires_at = NOW() + INTERVAL '7 days',
			revoked = false,
			created_at = NOW()
		WHERE refresh_token_id = prev_refresh_token_id;
	ELSE
		INSERT INTO refresh_tokens(user_id, refresh_token, expires_at, revoked, created_at)
		VALUES(in_user_id, in_token, NOW()+INTERVAL '7 days', false, NOW());

	END IF;
END;
$BODY$;

ALTER FUNCTION public.func_add_refresh_token(integer, text)
    OWNER TO postgres;

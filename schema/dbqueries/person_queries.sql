-- name: GetPublicDataByEmail :one
SELECT * FROM public_data WHERE email = $1;

-- name: UpsertPublicDataByUUID :exec
INSERT INTO public.public_data
(uuid, email, name, additional_info, birthdate, acc_balance, acc_balance_null, user_credentials_id, comments, comments_null)
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT(email) DO UPDATE SET
    email = EXCLUDED.email,
    name = EXCLUDED.name,
    additional_info = EXCLUDED.additional_info,
    birthdate = EXCLUDED.birthdate,
    acc_balance = EXCLUDED.acc_balance,
    acc_balance_null = EXCLUDED.acc_balance_null,
    user_credentials_id = EXCLUDED.user_credentials_id,
    comments = EXCLUDED.comments,
    comments_null = EXCLUDED.comments_null
;

-- name: DeletePublicDataByUUID :exec
DELETE FROM public.public_data
WHERE uuid = $1;

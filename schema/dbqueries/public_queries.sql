-- name: GetPublicDataByEmail :one
SELECT * FROM public_data WHERE email = $1;

-- name: UpsertPublicDataByUUID :exec
INSERT INTO public.public_data
(uuid, email, name, birthdate, acc_balance, acc_balance_null,
    hobbies, user_credentials_id, important_notes, additional_info
    )
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
ON CONFLICT(email) DO UPDATE SET
    name = EXCLUDED.name,
    email = EXCLUDED.email,
    birthdate = EXCLUDED.birthdate,
    acc_balance = EXCLUDED.acc_balance,
    acc_balance_null = EXCLUDED.acc_balance_null,
    hobbies = EXCLUDED.hobbies,
    user_credentials_id = EXCLUDED.user_credentials_id,
    important_notes = EXCLUDED.important_notes,
    additional_info = EXCLUDED.additional_info
;

-- name: DeletePublicDataByUUID :exec
DELETE FROM public.public_data
WHERE uuid = $1;

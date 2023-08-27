// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: public_queries.sql

package generated

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

const deletePublicDataByUUID = `-- name: DeletePublicDataByUUID :exec
DELETE FROM public.public_data
WHERE uuid = $1
`

func (q *Queries) DeletePublicDataByUUID(ctx context.Context, uuid pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deletePublicDataByUUID, uuid)
	return err
}

const getPublicDataByEmail = `-- name: GetPublicDataByEmail :one
SELECT id, uuid, name, email, birthdate, acc_balance, acc_balance_null, user_credentials_id, hobbies, important_notes, additional_info FROM public_data WHERE email = $1
`

func (q *Queries) GetPublicDataByEmail(ctx context.Context, email string) (PublicDatum, error) {
	row := q.db.QueryRow(ctx, getPublicDataByEmail, email)
	var i PublicDatum
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Name,
		&i.Email,
		&i.Birthdate,
		&i.AccBalance,
		&i.AccBalanceNull,
		&i.UserCredentialsID,
		&i.Hobbies,
		&i.ImportantNotes,
		&i.AdditionalInfo,
	)
	return i, err
}

const upsertPublicDataByUUID = `-- name: UpsertPublicDataByUUID :exec
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
`

type UpsertPublicDataByUUIDParams struct {
	Uuid              pgtype.UUID         `db:"uuid"`
	Email             string              `db:"email"`
	Name              pgtype.Text         `db:"name"`
	Birthdate         pgtype.Timestamptz  `db:"birthdate"`
	AccBalance        decimal.Decimal     `db:"acc_balance"`
	AccBalanceNull    decimal.NullDecimal `db:"acc_balance_null"`
	Hobbies           []string            `db:"hobbies"`
	UserCredentialsID int64               `db:"user_credentials_id"`
	ImportantNotes    []byte              `db:"important_notes"`
	AdditionalInfo    []byte              `db:"additional_info"`
}

func (q *Queries) UpsertPublicDataByUUID(ctx context.Context, arg UpsertPublicDataByUUIDParams) error {
	_, err := q.db.Exec(ctx, upsertPublicDataByUUID,
		arg.Uuid,
		arg.Email,
		arg.Name,
		arg.Birthdate,
		arg.AccBalance,
		arg.AccBalanceNull,
		arg.Hobbies,
		arg.UserCredentialsID,
		arg.ImportantNotes,
		arg.AdditionalInfo,
	)
	return err
}
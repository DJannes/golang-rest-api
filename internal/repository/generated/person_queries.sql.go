// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: person_queries.sql

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
SELECT id, uuid, email, name, additional_info, birthdate, acc_balance_null, acc_balance, user_credentials_id, comments_null, comments FROM public_data WHERE email = $1
`

func (q *Queries) GetPublicDataByEmail(ctx context.Context, email string) (PublicDatum, error) {
	row := q.db.QueryRow(ctx, getPublicDataByEmail, email)
	var i PublicDatum
	err := row.Scan(
		&i.ID,
		&i.Uuid,
		&i.Email,
		&i.Name,
		&i.AdditionalInfo,
		&i.Birthdate,
		&i.AccBalanceNull,
		&i.AccBalance,
		&i.UserCredentialsID,
		&i.CommentsNull,
		&i.Comments,
	)
	return i, err
}

const upsertPublicDataByUUID = `-- name: UpsertPublicDataByUUID :exec
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
`

type UpsertPublicDataByUUIDParams struct {
	Uuid              pgtype.UUID         `db:"uuid"`
	Email             string              `db:"email"`
	Name              pgtype.Text         `db:"name"`
	AdditionalInfo    []string            `db:"additional_info"`
	Birthdate         pgtype.Timestamptz  `db:"birthdate"`
	AccBalance        decimal.Decimal     `db:"acc_balance"`
	AccBalanceNull    decimal.NullDecimal `db:"acc_balance_null"`
	UserCredentialsID int64               `db:"user_credentials_id"`
	Comments          []byte              `db:"comments"`
	CommentsNull      []byte              `db:"comments_null"`
}

func (q *Queries) UpsertPublicDataByUUID(ctx context.Context, arg UpsertPublicDataByUUIDParams) error {
	_, err := q.db.Exec(ctx, upsertPublicDataByUUID,
		arg.Uuid,
		arg.Email,
		arg.Name,
		arg.AdditionalInfo,
		arg.Birthdate,
		arg.AccBalance,
		arg.AccBalanceNull,
		arg.UserCredentialsID,
		arg.Comments,
		arg.CommentsNull,
	)
	return err
}

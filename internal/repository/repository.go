package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	public "gitlab.com/janneseffendi/rest-api/internal/repository/generated"
)

type Repo struct {
	PublicRepo *public.Queries
}

func NewRepo(con *pgxpool.Pool) *Repo {
	return &Repo{
		PublicRepo: public.New(con),
	}
}

package repository

import (
	"gitlab.com/janneseffendi/rest-api/depedency"
	public "gitlab.com/janneseffendi/rest-api/internal/repository/generated"
)

type Repository struct {
	PublicRepo *public.Queries
}

func NewRepository() *Repository {
	con := depedency.GetDbConnection()
	return &Repository{
		PublicRepo: public.New(con),
	}
}

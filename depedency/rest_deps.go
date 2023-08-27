package depedency

import (
	"github.com/go-playground/validator/v10"
	"gitlab.com/janneseffendi/rest-api/internal/repository"
)

type RestDeps struct {
	Repo      *repository.Repo
	Validator *validator.Validate
}

func GetRestDeps() *RestDeps {
	// Translator
	InitTranslator()

	restDep := new(RestDeps)
	restDep.Repo = repository.NewRepo(GetConnectionPool())
	restDep.Validator = GetValidator()

	return restDep
}

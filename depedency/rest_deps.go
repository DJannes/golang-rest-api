package depedency

import (
	"github.com/go-playground/validator/v10"
	"gitlab.com/janneseffendi/rest-api/internal/repository"
	"gitlab.com/janneseffendi/rest-api/internal/security"
)

type RestDeps struct {
	Repo      *repository.Repo
	Validator *validator.Validate
	Paseto    *security.PasetoGen
}

func GetRestDeps() *RestDeps {
	// Translator
	InitTranslator()

	restDep := new(RestDeps)
	restDep.Repo = repository.NewRepo(GetConnectionPool())
	restDep.Validator = GetValidator()
	restDep.Paseto = security.NewMockPasetoGen()

	return restDep
}

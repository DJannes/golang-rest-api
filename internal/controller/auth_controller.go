package controller

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gitlab.com/janneseffendi/rest-api/internal/internal_utils"
	"gitlab.com/janneseffendi/rest-api/internal/service"
)

type AuthController struct {
	publicService *service.PublicService

	tokenGen *internal_utils.PasetoGen
}

func NewAuthController() *AuthController {
	return &AuthController{
		publicService: service.NewPublicService(),
		tokenGen:      internal_utils.NewMockPasetoGen(),
	}
}

func AddAuthRouter(r chi.Router) chi.Router {
	c := NewAuthController()
	r.Route("/auth", func(router chi.Router) {
		router.Get("/", c.GetToken)
		router.With(internal_utils.TokenAuth).Get("/test-token", c.TestToken)
	})

	return r
}

func (c *AuthController) GetToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="basic auth"`))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if password != "mockpassword" {
		w.Header().Add("WWW-Authenticate", fmt.Sprintf(`Basic realm="basic auth"`))
		w.Write([]byte("Password not matched"))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	token, err := c.tokenGen.CreateToken(username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("error creating token"))
		return
	}

	w.Write([]byte(token))
}

func (c *AuthController) TestToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

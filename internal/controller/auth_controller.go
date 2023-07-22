package controller

import (
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

// GetToken godoc
// @Security BasicAuth
//
//	@Summary		Get Token for authentication
//	@Description	Get Token for authentication
//	@Tags			0. Auth
//	@Accept			json
//	@Produce		plain
//
//	@Success		200	{string}	string "Success Response"
//	@Router			/auth [get]
func (c *AuthController) GetToken(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Add("WWW-Authenticate", `Basic realm="basic auth"`)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if password != "mockpassword" {
		w.Header().Add("WWW-Authenticate", `Basic realm="basic auth"`)
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

// TestToken godoc
// @Security TokenAuth
//
//	@Summary		Test Token for authentication
//	@Description	Test Token for auth
//	@Tags			0. Auth
//	@Accept			json
//	@Produce		plain
//
//	@Success		200	{string}	string "Success Response"
//	@Router			/auth/test-token [get]
func (c *AuthController) TestToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

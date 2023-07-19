package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"gitlab.com/janneseffendi/rest-api/depedency"
	"gitlab.com/janneseffendi/rest-api/internal/dto"
	"gitlab.com/janneseffendi/rest-api/internal/service"
)

type PublicController struct {
	publicService *service.PublicService

	validator *validator.Validate
}

func NewPublicController() *PublicController {
	return &PublicController{
		publicService: service.NewPublicService(),
		validator:     depedency.GetValidator(),
	}
}

func AddPublicRouter(r chi.Router) chi.Router {
	c := NewPublicController()
	r.Route("/public", func(router chi.Router) {
		router.Get("/", c.GetPublicData)
		router.Post("/", c.SavePublicData)
	})

	return r
}

func (c *PublicController) GetPublicData(w http.ResponseWriter, r *http.Request) {
	reqTime := time.Now()
	w.Header().Add("Content-Type", "application/json")
	publicData, err := c.publicService.GetPublicData(r.Context())
	if err != nil {
		render.Render(w, r, dto.ResponseFailBuilder(err, reqTime))
		return
	}

	render.Render(w, r, dto.ResponseOK(reqTime, publicData))
}

func (c *PublicController) SavePublicData(w http.ResponseWriter, r *http.Request) {
	reqTime := time.Now()
	reqJson, err := io.ReadAll(r.Body)
	if err != nil {
		res := dto.ResponseFailBuilder(err, reqTime)
		render.Render(w, r, res)
		return
	}
	defer r.Body.Close()

	var req dto.SavePublicData
	if err := json.Unmarshal(reqJson, &req); err != nil {
		res := dto.ResponseFailBuilder(err, reqTime)
		render.Render(w, r, res)
		return
	}

	if err := c.validator.Struct(req); err != nil {
		res := dto.ResponseFailBuilder(err, reqTime)
		render.Render(w, r, res)
		return
	}

	res := c.publicService.SavePublicData(req)
	httpRes := dto.ResponseOK(reqTime, res)
	render.Render(w, r, httpRes)
	w.Header().Add("Content-Type", "application/json")
}

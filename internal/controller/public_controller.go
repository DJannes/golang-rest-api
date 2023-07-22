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
	"gitlab.com/janneseffendi/rest-api/internal/internal_utils"
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

		router.Group(func(r chi.Router) {
			r.Use(internal_utils.TokenAuth)
			r.Post("/", c.SavePublicData)
			r.Delete("/", c.DeletePublicData)
		})
	})

	return r
}

const (
	mockUUID = "15ea0672-24c7-4429-813a-056c58f09ffb"
)

// GetPublicData godoc
//
//	@Summary		Get public data
//	@Description	Getting Public Data
//	@Tags			PublicData
//	@Accept			json
//	@Produce		json
//
// @Param uuid header string true "uuid in header"
//
//	@Success		200	{object}	dto.RestResponse{data=dto.PublicData,errors=dto.Error} "Success Response On public data"
//	@Failure		400	{object}	dto.ErrResponseForSwaggerDocsOnly
//	@Router			/public [get]
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

// SavePublicData godoc
// @Security TokenAuth
//
//	@Summary		Save public data
//	@Description	Save Public Data
//	@Tags			PublicData
//	@Accept			json
//	@Produce		json
//
// @Param reqBody body dto.SavePublicData true "Save Public Data Request Body"
//
//	@Success		200	{object}	dto.RestResponse{data=dto.PublicData} "Success Response"
//	@Router			/public [post]
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

	if err := c.publicService.SavePublicData(r.Context(), mockUUID, req); err != nil {
		res := dto.ResponseFailBuilder(err, reqTime)
		render.Render(w, r, res)
		return
	}

	httpRes := dto.ResponseOK(reqTime, nil)
	render.Render(w, r, httpRes)
	w.Header().Add("Content-Type", "application/json")
}

// DeletePublicData godoc
// @Security TokenAuth
//
//	@Summary		Save public data
//	@Description	Save Public Data
//	@Tags			PublicData
//	@Accept			json
//	@Produce		json
//
// @Param uuid header string true "uuid data to be deleted"
//
//	@Success		200	{object}	dto.RestResponse{data=dto.PublicData} "Success Response"
//	@Router			/public [delete]
func (c *PublicController) DeletePublicData(w http.ResponseWriter, r *http.Request) {
	reqTime := time.Now()
	if err := c.publicService.DeletePublicData(r.Context(), mockUUID); err != nil {
		res := dto.ResponseFailBuilder(err, reqTime)
		render.Render(w, r, res)
		return
	}

	httpRes := dto.ResponseOK(reqTime, nil)
	render.Render(w, r, httpRes)
	w.Header().Add("Content-Type", "application/json")
}

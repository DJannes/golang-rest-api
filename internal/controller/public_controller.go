package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	w.Header().Add("Content-Type", "application/json")
	publicData := c.publicService.GetPublicData()
	jsonData, err := json.Marshal(publicData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Json Marshal Failed: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (c *PublicController) SavePublicData(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	reqJson, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error reading request body"))
		return
	}
	defer r.Body.Close()

	var req dto.SavePublicData
	if err := json.Unmarshal(reqJson, &req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error unmarshaling request"))
		return
	}

	if err := c.validator.Struct(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error bad request" + err.Error()))
		return
	}

	res := c.publicService.SavePublicData(req)
	jsonData, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Json Marshal Failed: " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

package dto

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"gitlab.com/janneseffendi/rest-api/depedency"
)

const (
	responseTimeFormat = "2006-01-02T15:04:05"
)

type RestResponse struct {
	Err            error `json:"-"`
	HttpStatusCode int   `json:"-"`

	Errors       []Error `json:"errors,omitempty" swaggerignore:"true"`
	Code         string  `json:"code" example:"200"`
	Status       string  `json:"status" example:"OK"`
	Message      string  `json:"message" example:"Success"`
	Data         any     `json:"data"`
	RequestTime  string  `json:"requestTime" example:"2010-01-25T13:00:00"`
	ResponseTime string  `json:"responseTime" example:"2010-01-25T13:00:00"`
}

type ErrResponseForSwaggerDocsOnly struct {
	Errors       []Error `json:"errors,omitempty"`
	Code         string  `json:"code" example:"400"`
	Status       string  `json:"status" example:"Bad Request"`
	Message      string  `json:"message" example:"Request Body Is not Valid"`
	Data         any     `json:"data"`
	RequestTime  string  `json:"requestTime" example:"2010-01-25T13:00:00"`
	ResponseTime string  `json:"responseTime" example:"2010-01-25T13:00:00"`
}

type Error struct {
	FieldName  string `json:"fieldName" example:"publicName"`
	FieldError string `json:"fieldError" example:"publicName is required"`
}

func (e *RestResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HttpStatusCode)
	return nil
}

func ResponseOK(reqTime time.Time, data any) *RestResponse {
	return &RestResponse{
		Err:            nil,
		HttpStatusCode: http.StatusOK,

		Data:         data,
		Code:         fmt.Sprint(http.StatusOK),
		Status:       http.StatusText(http.StatusOK),
		Message:      http.StatusText(http.StatusOK),
		RequestTime:  reqTime.Format(responseTimeFormat),
		ResponseTime: time.Now().Format(responseTimeFormat),
	}
}

func ResponseBadRequest(reqTime time.Time, badErrs []Error) *RestResponse {
	return &RestResponse{
		Err:            nil,
		HttpStatusCode: http.StatusBadRequest,

		Errors:       badErrs,
		Data:         nil,
		Code:         fmt.Sprint(http.StatusBadRequest),
		Status:       http.StatusText(http.StatusBadRequest),
		Message:      http.StatusText(http.StatusBadRequest),
		RequestTime:  reqTime.Format(responseTimeFormat),
		ResponseTime: time.Now().Format(responseTimeFormat),
	}
}

func ResponseInternalError(err error, reqTime time.Time) *RestResponse {
	return &RestResponse{
		Err:            err,
		HttpStatusCode: http.StatusInternalServerError,

		Data:         nil,
		Code:         fmt.Sprint(http.StatusInternalServerError),
		Status:       http.StatusText(http.StatusInternalServerError),
		Message:      err.Error(),
		RequestTime:  reqTime.Format(responseTimeFormat),
		ResponseTime: time.Now().Format(responseTimeFormat),
	}
}

func ResponseFailBuilder(err error, reqTime time.Time) *RestResponse {
	if err == nil {
		return ResponseOK(reqTime, nil)
	}

	var validationErr validator.ValidationErrors
	if errors.As(err, &validationErr) {
		englishTranslator := depedency.GetTranslator()
		badReqErrors := make([]Error, 0)
		for _, e := range validationErr {
			fieldErr := e.Translate(englishTranslator)
			badReqErrors = append(badReqErrors, Error{
				FieldName:  e.Field(),
				FieldError: fieldErr,
			})
		}
		return ResponseBadRequest(reqTime, badReqErrors)
	}

	return ResponseInternalError(err, reqTime)
}

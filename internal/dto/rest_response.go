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
	SUCCESS_CODE       = "000"
	SUCCESS_MESSAGE    = "SUCCESS"
)

type Error struct {
	FieldName string `json:"fieldName"`
	Message   string `json:"message"`
}

type RestResponse struct {
	Err            error `json:"-"`
	HttpStatusCode int   `json:"-"`

	Errors       []Error `json:"errors,omitempty"`
	Code         string  `json:"code"`
	Status       string  `json:"status"`
	Data         any     `json:"data"`
	RequestTime  string  `json:"requestTime"`
	ResponseTime string  `json:"responseTime"`
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
		Code:         SUCCESS_CODE,
		Status:       SUCCESS_MESSAGE,
		RequestTime:  reqTime.Format(responseTimeFormat),
		ResponseTime: time.Now().Format(responseTimeFormat),
	}
}

func ResponseBadRequest(reqTime time.Time, badErrs []Error) *RestResponse {
	return &RestResponse{
		Err:            nil,
		HttpStatusCode: http.StatusBadRequest,

		Errors:       badErrs,
		Code:         fmt.Sprint(http.StatusBadRequest),
		Status:       http.StatusText(http.StatusBadRequest),
		RequestTime:  reqTime.Format(responseTimeFormat),
		ResponseTime: time.Now().Format(responseTimeFormat),
	}
}

func ResponseInternalError(err error, reqTime time.Time) *RestResponse {
	return &RestResponse{
		Err:            err,
		HttpStatusCode: http.StatusInternalServerError,

		Code:   fmt.Sprint(http.StatusInternalServerError),
		Status: http.StatusText(http.StatusInternalServerError),
		Errors: []Error{
			Error{Message: http.StatusText(http.StatusInternalServerError)},
		},
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
		englishTranslator := depedency.GetEnTranslator()
		badReqErrors := make([]Error, 0)

		for _, e := range validationErr {
			fieldErr := e.Translate(englishTranslator)
			badReqErrors = append(badReqErrors, Error{
				FieldName: e.Field(),
				Message:   fieldErr,
			})
		}
		return ResponseBadRequest(reqTime, badReqErrors)
	}

	return ResponseInternalError(err, reqTime)
}

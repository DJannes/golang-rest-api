package dto

type ErrResponseForSwaggerDocs struct {
	Errors       []Error `json:"errors,omitempty"`
	Code         string  `json:"code" example:"400"`
	Status       string  `json:"status" example:"Bad Request"`
	Message      string  `json:"message" example:"Request Body Is not Valid"`
	Data         any     `json:"data"`
	RequestTime  string  `json:"requestTime" example:"2010-01-25T13:00:00"`
	ResponseTime string  `json:"responseTime" example:"2010-01-25T13:00:00"`
}

type SuccessResponseForSwaggerDocs struct {
	Errors       []Error `json:"errors,omitempty" swaggerignore:"true"` // Hide Errors Field
	Code         string  `json:"code" example:"000"`
	Status       string  `json:"status" example:"SUCCESS"`
	Message      string  `json:"message" example:"SUCCESS"`
	Data         any     `json:"data"`
	RequestTime  string  `json:"requestTime" example:"2010-01-25T13:00:00"`
	ResponseTime string  `json:"responseTime" example:"2010-01-25T13:00:00"`
}

type ErrorSwaggerDocs struct {
	FieldName  string `json:"fieldName" example:"publicName"`
	FieldError string `json:"fieldError" example:"publicName is required"`
}

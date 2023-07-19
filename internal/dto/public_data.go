package dto

import "time"

type PublicData struct {
	PublicName         string    `json:"publicName"`
	PublicDate         time.Time `json:"publicDate"`
	PublicDescriptions []string  `json:"publicDescriptions"`
}

type SavePublicData struct {
	PublicName         string   `json:"publicName" validate:"required"`
	PublicDate         string   `json:"publicDate" validate:"required"`
	PublicDescriptions []string `json:"publicDescriptions" validate:"required"`
}

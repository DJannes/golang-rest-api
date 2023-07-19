package dto

import "time"

type PublicData struct {
	PublicName         string    `json:"publicName"`
	PublicDate         time.Time `json:"publicDate"`
	PublicDescriptions []string  `json:"publicDescriptions"`
}

type SavePublicData struct {
	PublicName         string   `json:"publicName"`
	PublicDate         string   `json:"publicDate"`
	PublicDescriptions []string `json:"publicDescriptions"`
}

package dto

type PublicData struct {
	Email             string   `json:"email"`
	Name              string   `json:"name"`
	AdditionalInfo    []string `json:"additional_info"`
	Birthdate         string   `json:"birthdate"`
	AccBalanceNull    string   `json:"accBalanceNull"`
	AccBalance        string   `json:"accBalance"`
	UserCredentialsID int64    `json:"userCredentialsId"`
	CommentsNull      []byte   `json:"commentsNull"`
	Comments          []byte   `json:"comments"`
}

type SavePublicData struct {
	PublicName         string   `json:"publicName" validate:"required"`
	PublicDate         string   `json:"publicDate" validate:"required"`
	PublicDescriptions []string `json:"publicDescriptions" validate:"required"`
}

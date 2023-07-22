package dto

type PublicData struct {
	Email             string   `json:"email" example:"Email Public" format:"email"`
	Name              string   `json:"name" example:"Nama Public"`
	AdditionalInfo    []string `json:"additional_info" example:"Info1, Info2, Info3"`
	Birthdate         string   `json:"birthdate" example:"2010-03-10"`
	AccBalanceNull    string   `json:"accBalanceNull" example:"0.20"`
	AccBalance        string   `json:"accBalance" example:"10000.20"`
	UserCredentialsID int64    `json:"userCredentialsId" example:"1"`
	CommentsNull      []byte   `json:"commentsNull" swaggertype:"string" example:"{}"`
	Comments          []byte   `json:"comments" swaggertype:"string" example:"{a: abc}"`
}

type SavePublicData struct {
	PublicName         string   `json:"publicName" validate:"required" minLength:"4" maxLength:"16" example:"Nama Publik"`
	PublicDate         string   `json:"publicDate" validate:"required" minLength:"4" maxLength:"16" example:"2010-10-12"`
	PublicDescriptions []string `json:"publicDescriptions" validate:"required" minLength:"4" maxLength:"16" example:"data,data2,data3"`
}

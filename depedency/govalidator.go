package depedency

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en_US"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	supportedTranslator *ut.UniversalTranslator
)

func InitTranslator() {
	english := en_US.New()
	supportedTranslator = ut.New(english, english)
}

func GetValidator() *validator.Validate {
	v := validator.New()

	// Custom Field Name
	v.RegisterTagNameFunc(jsonFieldName)

	// Setup Translator
	en_translations.RegisterDefaultTranslations(v, GetEnTranslator())
	return v
}

func GetEnTranslator() ut.Translator {
	trans, _ := supportedTranslator.GetTranslator("en_US")
	return trans
}

func jsonFieldName(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return fld.Name
	}
	return name
}

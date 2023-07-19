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
	v             *validator.Validate
	supportedLang *ut.UniversalTranslator
	englishTrans  ut.Translator
)

func InitializedValidator() {
	v = validator.New()

	// Custom Field Name
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})
}

func InitializedTranslator() {
	english := en_US.New()
	supportedLang = ut.New(english, english)

	englishTrans, _ = supportedLang.GetTranslator("en_US")
	en_translations.RegisterDefaultTranslations(v, englishTrans)
}

func GetValidator() *validator.Validate {
	return v
}

func GetTranslator() ut.Translator {
	return englishTrans
}

package depedency

import (
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

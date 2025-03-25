package validation

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	es_translations "github.com/go-playground/validator/v10/translations/es"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	TransES  ut.Translator
	TransEN  ut.Translator
)

func NewValidator() {
	enLocale := en.New()
	esLocale := es.New()

	uni = ut.New(enLocale, enLocale, esLocale)

	var ok bool
	TransEN, ok = uni.GetTranslator("en")
	if !ok {
		panic("Cannot get translator for English")
	}

	TransES, ok = uni.GetTranslator("es")
	if !ok {
		panic("Cannot get translator for Spanish")
	}

	validate = validator.New()

	if err := en_translations.RegisterDefaultTranslations(validate, TransEN); err != nil {
		panic(err)
	}

	if err := es_translations.RegisterDefaultTranslations(validate, TransES); err != nil {
		panic(err)
	}
}

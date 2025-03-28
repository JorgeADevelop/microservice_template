package translater

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var Bundle *i18n.Bundle

func NewTranslater() {
	bundle := i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	bundle.MustLoadMessageFile("pkg/translater/locales/en.json")
	bundle.MustLoadMessageFile("pkg/translater/locales/es.json")

	Bundle = bundle
}

func TranslateMessage(lang string, messageID string) string {
	localizer := i18n.NewLocalizer(Bundle, lang)
	message, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: messageID,
	})

	if err != nil {
		return messageID
	}
	return message
}

package i18n_test

import (
	"context"
	"testing"

	"github.com/monstrum/stick/i18n"
)

func TestTranslator(t *testing.T) {
	dummyTranslation := map[string]string{
		"Hello":     "World",
		"MyArgsKey": "Hello %text%",
	}
	dummyDeTranslation := map[string]string{
		"Hello": "Welt",
	}
	translator := &i18n.Translator{
		Locale:       "en",
		Domain:       "messages",
		Translations: dummyTranslation,
		LocaleTranslations: map[i18n.Locale]i18n.Translations{
			"en": dummyTranslation,
			"de": dummyDeTranslation,
		},
		DomainTranslations: map[i18n.Domain]i18n.LocaleTranslations{
			"messages": map[i18n.Locale]i18n.Translations{
				"en": dummyTranslation,
				"de": dummyDeTranslation,
			},
			"custom": map[i18n.Locale]i18n.Translations{
				"en": map[string]string{
					"Hello": "Lala World",
				},
			},
		},
	}

	locale := translator.GetLocale("")
	if locale != "en" {
		t.Errorf("Expected 'en', got '%s'", locale)
	}

	domain := translator.GetDomain("")
	if domain != "messages" {
		t.Errorf("Expected 'messages', got '%s'", domain)
	}

	translationTests := map[string]struct {
		key      string
		locale   i18n.Locale
		domain   i18n.Domain
		args     map[string]string
		expected string
	}{
		"basic key": {
			key:      "Hello",
			locale:   "",
			domain:   "",
			expected: "World",
		},
		"with locale": {
			key:      "Hello",
			locale:   "de",
			domain:   "",
			expected: "Welt",
		},
		"with domain": {
			key:      "Hello",
			locale:   "",
			domain:   "messages",
			expected: "World",
		},
		"with domain and locale": {
			key:      "Hello",
			locale:   "de",
			domain:   "messages",
			expected: "Welt",
		},
		"with other domain and locale": {
			key:      "Hello",
			locale:   "en",
			domain:   "custom",
			expected: "Lala World",
		},
		"with other domain and de locale with fallback": {
			key:      "Hello",
			locale:   "de",
			domain:   "custom",
			expected: "Welt",
		},
		"with domain with args": {
			key:      "MyArgsKey",
			locale:   "en",
			domain:   "messages",
			args:     map[string]string{},
			expected: "Hello %text%",
		},
	}

	for key, test := range translationTests {
		t.Run(key, func(t *testing.T) {
			translated := translator.Translate(context.Background(), test.key, test.args, test.locale, test.domain)
			if translated != test.expected {
				t.Errorf("Expected '%s', got '%s'", test.expected, translated)
			}
		})
	}
}

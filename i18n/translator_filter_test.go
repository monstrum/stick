package i18n_test

import (
	"testing"

	"github.com/monstrum/stick"
	"github.com/monstrum/stick/i18n"
)

func TestTranslatorFilter(t *testing.T) {
	filter := &i18n.TranslationFilter{
		Translator: i18n.Translator{},
	}

	filters := map[string]stick.Filter{}
	filter.Register(filters)

	if _, ok := filters["trans"]; !ok {
		t.Error("Filter not registered")
	}
}

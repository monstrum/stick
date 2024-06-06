package i18n_test

import (
	"github.com/monstrum/stick/i18n"
	"testing"
)

func TestLoadTranslationFiles(t *testing.T) {
	translations := i18n.DomainTranslations{}
	i18n.LoadTranslationFiles("./test/translations", translations, nil)

	if len(translations) == 0 {
		t.Errorf("Expected translations to be loaded")
	}

	if translations["messages"] == nil {
		t.Errorf("Expected messages domain to be loaded")
	}

	if len(translations["messages"]) != 2 {
		t.Errorf("Expected 2 locale with default translations to be loaded")
	}

	if result, ok := translations["messages"]["en"]["hello"]; !ok || result != "World" {
		t.Errorf("Expected translations to be loaded")
	}

	if result, ok := translations["messages"]["de"]["hello"]; !ok || result != "Welt" {
		t.Errorf("Expected translations to be loaded")
	}
}

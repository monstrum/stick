package i18n

import (
	"github.com/monstrum/stick/twig/factory"
)

type OptionFn func(*Option)

// TranslatorFilter
// {{ message|trans(arguments = [], domain = null, locale = null) }}
func TranslatorFilter(options ...OptionFn) factory.AppendFilterFn {
	o := new(Option)
	for _, option := range options {
		option(o)
	}
	translationFilter := TranslationFilter{*NewTranslator(options...)}
	return translationFilter.Register
}

func WithTranslator(translator *Translator) factory.AppendFilterFn {
	translationFilter := TranslationFilter{*translator}
	return translationFilter.Register
}

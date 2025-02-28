package i18n

import (
	"github.com/monstrum/stick"
)

type TranslationFilter struct {
	Translator
}

func (f *TranslationFilter) parseTranslationFilterArguments(
	t Translator,
	args ...stick.Value,
) (arguments map[string]string, locale Locale, domain Domain) {
	if len(args) >= 3 {
		locale = t.GetLocale(Locale(args[2].(string)))
	}

	if len(args) >= 2 {
		domain = t.GetDomain(Domain(args[1].(string)))
	}

	if len(args) >= 1 {
		if values, ok := args[0].(map[string]stick.Value); ok && len(values) > 0 {
			arguments = make(map[string]string)
			for key, value := range values {
				arguments[key] = stick.CoerceString(value)
			}
		}
	}
	return
}

func (f *TranslationFilter) Register(filters map[string]stick.Filter) {
	filters["trans"] = f.Trans
}

func (f *TranslationFilter) Trans(ctx stick.Context, value stick.Value, args ...stick.Value) stick.Value {
	message, ok := value.(string)
	if !ok {
		return value
	}

	f.SetLocale("")
	if lang, ok := ctx.Context().Value("locale").(string); ok && lang != "" {
		f.SetLocale(Locale(lang))
	}

	arguments, locale, domain := f.parseTranslationFilterArguments(f.Translator, args...)
	return stick.Value(f.Translator.Translate(ctx.Context(), message, arguments, locale, domain))
}

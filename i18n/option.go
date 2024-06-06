package i18n

import (
	"encoding/json"
)

const defaultDomain = "messages"

type (
	Domain string
	Locale string

	Translations       map[string]string
	LocaleTranslations map[Locale]Translations
	DomainTranslations map[Domain]LocaleTranslations

	Option struct {
		DefaultLocale      Locale
		DefaultDomain      Domain
		DomainTranslations DomainTranslations
		Translation        LocaleTranslations
	}
)

func (o *Option) Value(key string, locale Locale, domain Domain) (string, bool) {
	if locale == "" {
		locale = o.DefaultLocale
	}
	if domain == "" {
		return o.defaultValue(key, locale)
	}
	return o.domainValue(key, locale, domain)
}

func (o *Option) defaultValue(key string, locale Locale) (string, bool) {
	return o.localeValue(key, locale, o.Translation)
}

func (o *Option) domainValue(key string, locale Locale, domain Domain) (string, bool) {
	translations, ok := o.DomainTranslations[domain]
	if !ok {
		return key, false
	}

	return o.localeValue(key, locale, translations)
}

func (o *Option) localeValue(key string, locale Locale, translations LocaleTranslations) (string, bool) {
	translation, ok := translations[locale]
	if !ok {
		return key, false
	}

	return o.value(key, translation)
}

func (o *Option) value(key string, translations Translations) (string, bool) {
	translation, ok := translations[key]
	if !ok {
		return key, false
	}

	return translation, true
}

func WithLocale(locale string) OptionFn {
	return func(o *Option) {
		o.DefaultLocale = Locale(locale)
	}
}

func WithDomain(domain string) OptionFn {
	return func(o *Option) {
		o.DefaultDomain = Domain(domain)
	}
}

func WithDefaultDomain() OptionFn {
	return func(o *Option) {
		o.DefaultDomain = defaultDomain
	}
}

func WithTranslations(translations DomainTranslations) OptionFn {
	return func(o *Option) {
		o.DomainTranslations = translations
	}
}

func WithDefaultTranslations(translations LocaleTranslations) OptionFn {
	return func(o *Option) {
		o.Translation = translations
		if o.DomainTranslations == nil {
			o.DomainTranslations = DomainTranslations{
				defaultDomain: translations,
			}
		}
	}
}

func WithLocalizationPath(path string) OptionFn {
	return func(o *Option) {
		if o.DomainTranslations == nil {
			o.DomainTranslations = make(DomainTranslations)
		}

		LoadTranslationFiles(path, o.DomainTranslations, nil)
		if o.Translation == nil && o.DomainTranslations[defaultDomain] != nil {
			o.Translation = o.DomainTranslations[defaultDomain]
			o.DefaultDomain = defaultDomain
		}
	}
}

func AddDomainTranslations(domain, locale string, loader Loader) OptionFn {
	d := Domain(domain)
	l := Locale(locale)
	return func(o *Option) {
		if o.DefaultDomain == "" {
			o.DefaultDomain = d
		}
		if o.DefaultLocale == "" {
			o.DefaultLocale = l
		}

		if o.DomainTranslations == nil {
			o.DomainTranslations = make(DomainTranslations)
		}

		localeTranslations := o.DomainTranslations[d]
		if localeTranslations == nil {
			o.DomainTranslations[d] = make(LocaleTranslations)
		}

		translations := o.DomainTranslations[d][l]
		if translations == nil {
			o.DomainTranslations[d][l] = make(Translations)
		}

		err := json.Unmarshal(loader(), &translations)
		if err != nil {
			panic(err)
			return
		}
		o.DomainTranslations[d][l] = translations
		if o.Translation == nil {
			o.Translation = o.DomainTranslations[d]
		}
	}
}

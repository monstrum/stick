package i18n

import (
	"context"
	"strings"
)

func NewTranslator(opts ...OptionFn) *Translator {
	o := &Option{}
	for _, opt := range opts {
		opt(o)
	}

	return &Translator{
		Domain:             o.DefaultDomain,
		Locale:             o.DefaultLocale,
		DomainTranslations: o.DomainTranslations,
		LocaleTranslations: o.Translation,
	}
}

type Translator struct {
	Domain Domain
	Locale Locale

	Translations       map[string]string
	LocaleTranslations map[Locale]Translations
	DomainTranslations map[Domain]LocaleTranslations
}

func (t *Translator) domainValue(key string, locale Locale, domain Domain) (string, bool) {
	locale = t.GetLocale(locale)
	domain = t.GetDomain(domain)

	_, ok := t.DomainTranslations[domain][locale]
	if ok {
		return t.localeValue(key, locale, t.DomainTranslations[domain])
	}

	_, ok = t.DomainTranslations[defaultDomain][locale]
	if ok {
		return t.localeValue(key, locale, t.DomainTranslations[defaultDomain])
	}
	return key, false
}

func (t *Translator) localeValue(key string, locale Locale, translations LocaleTranslations) (string, bool) {
	translation, ok := translations[locale]
	if !ok {
		return key, false
	}

	return t.value(key, translation)
}

func (t *Translator) value(key string, translations Translations) (string, bool) {
	translation, ok := translations[key]
	if !ok {
		return key, false
	}

	return translation, true
}

func (t *Translator) GetDomain(domain Domain) Domain {
	if domain == "" {
		return defaultDomain
	}
	return domain
}

func (t *Translator) GetLocale(locale Locale) Locale {
	if locale == "" {
		return t.Locale
	}
	return locale
}

func (t *Translator) Translate(
	ctx context.Context,
	key string,
	args map[string]string,
	locale Locale,
	domain Domain,
) string {
	translated, ok := t.domainValue(key, locale, domain)
	if !ok {
		return key
	}

	if args == nil {
		return translated
	}

	for k, v := range args {
		translated = strings.Replace(translated, "%"+k+"%", v, -1)
	}

	return translated
}

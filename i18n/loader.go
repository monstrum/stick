package i18n

import (
	"encoding/json"
	"os"
	"strings"
)

type Loader = func() []byte

func FileLoader(path string) Loader {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return func() []byte {
		return data
	}
}

func LoadTranslationFiles(path string, translations DomainTranslations, domain *Domain) {
	paths, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	d := Domain(defaultDomain)
	if domain != nil {
		d = *domain
	}

	for _, p := range paths {
		if p.IsDir() {
			newDomain := Domain(p.Name())
			LoadTranslationFiles(path+"/"+p.Name(), translations, &newDomain)
			return
		}

		if !strings.Contains(p.Name(), ".json") {
			continue
		}

		// get file name as locale example de.json as de locale
		locale := strings.Split(p.Name(), ".")[0]
		loader := FileLoader(path + "/" + p.Name())
		loadTranslations(translations, d, Locale(locale), loader)
	}
}

func loadTranslations(domainTranslations DomainTranslations, domain Domain, locale Locale, loader Loader) {
	localeTranslations := domainTranslations[domain]
	if localeTranslations == nil {
		domainTranslations[domain] = make(LocaleTranslations)
	}

	translations := domainTranslations[domain][locale]
	if translations == nil {
		domainTranslations[domain][locale] = make(Translations)
	}

	err := json.Unmarshal(loader(), &translations)
	if err != nil {
		panic(err)
		return
	}

	domainTranslations[domain][locale] = translations
}

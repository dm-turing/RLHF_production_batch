package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	languagesDirectory = "translations"
	sourceLanguage     = "en"
)

type Translation struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Plural string `json:"plural,omitempty"`
}

func validateTranslation(lang, id, text string) error {
	// Simple validation to ensure the text is not empty
	if text == "" {
		return fmt.Errorf("empty translation for ID '%s' in language '%s'", id, lang)
	}
	return nil
}

func loadTranslations(dir string) (map[string][]Translation, error) {
	translations := make(map[string][]Translation)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".json" {
			continue
		}

		lang := strings.TrimSuffix(file.Name(), ".json")
		data, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
		if err != nil {
			return nil, err
		}

		var translationsForLang []Translation
		if err := json.Unmarshal(data, &translationsForLang); err != nil {
			return nil, err
		}

		translations[lang] = translationsForLang
	}

	return translations, nil
}

func validateTranslations(translations map[string][]Translation) error {
	for lang, trans := range translations {
		for _, translation := range trans {
			if err := validateTranslation(lang, translation.ID, translation.Text); err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	translationsDir, err := filepath.Abs(languagesDirectory)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting absolute path of translations directory:", err)
		os.Exit(1)
	}

	translations, err := loadTranslations(translationsDir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error loading translations:", err)
		os.Exit(1)
	}

	if err := validateTranslations(translations); err != nil {
		fmt.Fprintln(os.Stderr, "Error validating translations:", err)
		os.Exit(1)
	}

	fmt.Println("Translations are valid.")
}

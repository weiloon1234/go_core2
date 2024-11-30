package localization

import (
	"context"
	"net/http"
	"strings"
)

type Config struct {
	DefaultLanguage    string
	SupportedLanguages []string
}

var LocaleContextKey string = "locale"
var defaultLanguage string
var supportedLanguages []string

func Init(config Config) {
	defaultLanguage = config.DefaultLanguage
	supportedLanguages = config.SupportedLanguages
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Accept-Language")
		if !isSupported(lang) {
			lang = defaultLanguage
		}
		r = r.WithContext(context.WithValue(r.Context(), LocaleContextKey, lang))
		next.ServeHTTP(w, r)
	})
}

func isSupported(lang string) bool {
	for _, l := range supportedLanguages {
		if strings.HasPrefix(lang, l) {
			return true
		}
	}
	return false
}

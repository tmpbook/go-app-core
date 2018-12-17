package decorator

import (
	"log"
	"net/http"
)

// ErrorCatcher Function adapter. Wrap type http.HandlerFunc
func ErrorCatcher(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf(`%d %q ERROR %q: %v`, http.StatusInternalServerError, r.Method, r.RequestURI, err)
		} else {
			log.Printf(`%d %q %q %q`, http.StatusOK, r.Method, http.StatusText(http.StatusOK), r.RequestURI)
		}
	}
}

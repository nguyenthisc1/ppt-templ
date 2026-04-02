package webapp

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
)

func renderComponent(w http.ResponseWriter, r *http.Request, status int, component templ.Component) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("render error: %v", err)
	}
}

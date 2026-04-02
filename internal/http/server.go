package webapp

import (
	"net/http"

	"github.com/a-h/templ"

	"ppt-templ/internal/views/pages"
	"ppt-templ/internal/views/partials"
)

type Config struct {
	Addr      string
	SiteName  string
	AssetsDir string
}

func NewServer(cfg Config) *http.ServeMux {
	mux := http.NewServeMux()
	staticFS := http.FileServer(http.Dir(cfg.AssetsDir))

	mux.Handle("/assets/", http.StripPrefix("/", staticFS))
	mux.HandleFunc("/", pageHandler(func() templ.Component {
		return pages.Home(cfg.SiteName)
	}))
	mux.HandleFunc("/about", pageHandler(func() templ.Component {
		return pages.About(cfg.SiteName)
	}))
	mux.HandleFunc("/contact", pageHandler(func() templ.Component {
		return pages.Contact(cfg.SiteName)
	}))
	mux.HandleFunc("/fragments/highlight", func(w http.ResponseWriter, r *http.Request) {
		audience := r.URL.Query().Get("audience")
		if audience == "" {
			audience = "founders"
		}

		if r.Header.Get("HX-Request") == "true" {
			renderComponent(w, r, http.StatusOK, partials.HighlightPanel(audience))
			return
		}

		renderComponent(w, r, http.StatusOK, pages.HighlightPreview(cfg.SiteName, audience))
	})

	return mux
}

func pageHandler(view func() templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderComponent(w, r, http.StatusOK, view())
	}
}

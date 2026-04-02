package webapp

import (
	"net/http"
	"strings"

	"github.com/a-h/templ"

	"ppt-templ/internal/content"
	"ppt-templ/internal/views/pages"
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
	mux.HandleFunc("/listings", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		category := r.URL.Query().Get("category")

		renderComponent(w, r, http.StatusOK, pages.Listings(cfg.SiteName, query, category, content.FilterListings(query, category)))
	})
	mux.HandleFunc("/listings/", func(w http.ResponseWriter, r *http.Request) {
		slug := strings.TrimPrefix(r.URL.Path, "/listings/")
		if slug == "" || strings.Contains(slug, "/") {
			http.NotFound(w, r)
			return
		}

		item, ok := content.FindListingBySlug(slug)
		if !ok {
			http.NotFound(w, r)
			return
		}

		renderComponent(w, r, http.StatusOK, pages.ListingDetail(cfg.SiteName, item))
	})
	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		status := r.URL.Query().Get("status")

		renderComponent(w, r, http.StatusOK, pages.Projects(cfg.SiteName, query, status, content.FilterProjects(query, status)))
	})
	mux.HandleFunc("/news", pageHandler(func() templ.Component {
		return pages.News(cfg.SiteName, content.News())
	}))

	return mux
}

func pageHandler(view func() templ.Component) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderComponent(w, r, http.StatusOK, view())
	}
}

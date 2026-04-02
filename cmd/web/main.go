package main

import (
	"log"
	"net/http"
	"os"

	webapp "ppt-templ/internal/http"
)

func main() {
	addr := os.Getenv("APP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	mux := webapp.NewServer(webapp.Config{
		Addr:      addr,
		SiteName:  "PPT Templ",
		AssetsDir: "public",
	})

	log.Printf("listening on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

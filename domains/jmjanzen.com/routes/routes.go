package routes

import (
	"log"
	"net/http"

	"jm/domains/jmjanzen.com/handlers"
)

// Handlers receive and log an HTTP req, then serve our pages (using _render)
func Route(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v%v", r.RemoteAddr, r.Method, r.Host, r.RequestURI)
	handlers.Default(w, r)
}

func Launch() {
	mux := http.NewServeMux()

	// Handle homepage, others
	mux.HandleFunc("/", Route)

	// Handle static resources (js, css)
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handle favicon, home img (anon func)
	//mux.HandleFunc("/favicon.ico", handlers.HandleFile())
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/img/fav/favicon.ico")
	})

	// If any issue starting, log err, and exit(1)
	listenPort := ":6060"
	log.Printf("Listening on port %v", listenPort)
	log.Fatal(http.ListenAndServe(listenPort, mux))
}


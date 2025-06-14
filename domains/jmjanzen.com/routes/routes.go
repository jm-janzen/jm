package routes

import (
	"log"
	"net/http"

	"jm/domains/jmjanzen.com/handlers"
)

func Route(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v %v%v", r.RemoteAddr, r.Method, r.Host, r.RequestURI)
	handlers.Default(w, r)
}

func Launch() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Route)

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/favicon.ico", handlers.HandleFile)

	listenPort := ":6060"
	log.Printf("Listening on port %v", listenPort)
	log.Fatal(http.ListenAndServe(listenPort, mux))
}


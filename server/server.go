package server

import (
	"fmt"
	"github.com/gobuffalo/packr"
	. "github.com/rsteube/tinyelk/config"
	"github.com/rsteube/tinyelk/db"
	"log"
	"net/http"
	"time"
)

type Server struct {
	config *Config
	cache  *db.Cache
}

func Serve(config *Config, cache *db.Cache) {
	server := Server{
		config: config,
		cache:  cache,
	}

	http.Handle("/", http.FileServer(packr.NewBox("./resources")))
	http.HandleFunc("/jq", server.jq)

	log.Fatal(http.ListenAndServe(":7318", nil)) // port T(iny)ELK
}

func (server *Server) jq(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // parse arguments, you have to call this by yourself

	w.Header().Set("Content-Type", "application/json")
	if query, ok := r.URL.Query()["q"]; !ok || len(query) < 1 {
		fmt.Fprintf(w, "fail!") // send data to client side

	} else {
		start := time.Now()
		server.cache.QueryPrefix(w, "2018-02-16T05", query[0])
		log.Printf("path: %s form: %s elapsed-time: %s", r.URL.Path, r.Form, time.Since(start))

		server.cache.SomeTest()
	}
}

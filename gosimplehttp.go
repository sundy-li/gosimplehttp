package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

var (
	port  string
	dir   string
	debug bool
)

func init() {
	flag.StringVar(&port, "p", "8000", "port to bind")
	flag.StringVar(&dir, "d", ".", "directory of the files")
	flag.BoolVar(&debug, "debug", false, "show debug log or not")

	flag.Parse()
	log.SetFlags(log.Ldate | log.Lshortfile)
}
func main() {
	log.Printf("Serving HTTP on 0.0.0.0 port %s ...\n", port)
	h := http.FileServer(http.Dir(dir))
	log.Fatal(http.ListenAndServe(":"+port, &Handler{handler: h}))
}

type Handler struct {
	handler http.Handler
}

func (f *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	st := time.Now().UnixNano()
	f.handler.ServeHTTP(w, r)
	if debug {
		ed := time.Now().UnixNano()
		log.Printf("server %s in %.3f millseconds\n", r.RequestURI, float32(ed-st)*1.0/1e6)
	}
}

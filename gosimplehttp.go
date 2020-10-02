package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/sundy-li/gosimplehttp/util"
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
	log.Printf("Serving HTTP on 0.0.0.0 port %s : %s...\n", port, dir)
	h := http.FileServer(http.Dir(dir))
	log.Fatal(http.ListenAndServe(":"+port, &Handler{handler: h}))
}

type Handler struct {
	handler http.Handler
}

func (f *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	st := time.Now().UnixNano()

	var foundSuffix = false
	for _, s := range suffixes {
		if strings.HasSuffix(r.URL.Path, s) {
			foundSuffix = true
		}
	}

	if foundSuffix {
		err := util.ConvertToPdf(filepath.Join(dir, r.URL.Path), w)
		if err != nil {
			fmt.Fprintf(w, "Error in got %v", err.Error())
			return
		}
	} else {
		f.handler.ServeHTTP(w, r)
	}

	if debug {
		ed := time.Now().UnixNano()
		log.Printf("server %s in %.3f millseconds\n", r.RequestURI, float32(ed-st)*1.0/1e6)
	}
}

var (
	suffixes = []string{".pptx", ".ppt", ".pdf", ".doc", ".docx"}
)

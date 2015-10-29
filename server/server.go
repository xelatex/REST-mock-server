// Copyright 2015-2018 Arthur Chunqi Li. All rights reserved.

package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"../config"
	"../controller"
)

// Server is our main struct.
type Server struct {
	opts config.Options
	c    *controller.Controller
}

// PrintServerAndExit will print our version and exit.
func PrintServerAndExit() {
	fmt.Printf("mock-server version %s\n", config.VERSION)
	os.Exit(0)
}

// New will setup a new server struct after parsing the options.
func New(_opts config.Options) *Server {
	s := &Server{
		opts: _opts,
	}

	s.c = controller.New(_opts)
	return s
}

func (s *Server) Start() {
	var addr string
	addr = fmt.Sprintf("%s:%d", s.opts.Host, s.opts.Port)
	log.Printf("addr: %s\n", addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handler)
	http.ListenAndServe(addr, mux)
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("test", "abcd")
	//	w.Header().Set("test2", "abcdefgh")
	//	w.WriteHeader(http.StatusBadRequest)
	//	fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path)
	//	fmt.Fprintf(w, "Request method: %s\n", r.Method)
	//	fmt.Fprintf(w, "Header: %s\n", r.Header)
	//	body, err := ioutil.ReadAll(r.Body)
	//	if err != nil {
	//		log.Fatalf("Error: %s\n", err)
	//		w.WriteHeader(http.StatusInternalServerError)
	//		return
	//	}
	//	fmt.Fprintf(w, "Body: %s\n", body)

	cr := controller.Request{
		HttpRequest: r,
	}
	s.c.Get(&cr, &w)
}

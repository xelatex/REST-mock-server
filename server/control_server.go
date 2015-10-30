// Copyright 2015-2018 Arthur Chunqi Li. All rights reserved.

package server

import (
	"fmt"
	"net/http"

	"../config"
	"../controller"
	"../log"
)

// Server is our main struct.
type ControlServer struct {
	opts config.Options
	c    *controller.Controller
}

// New will setup a new server struct after parsing the options.
func NewControlServer(_opts config.Options) *ControlServer {
	s := &ControlServer{
		opts: _opts,
	}

	s.c = controller.New(_opts)
	return s
}

func (s *ControlServer) Start() {
	var addr string
	addr = fmt.Sprintf("%s:%d", s.opts.Host, s.opts.ControlPort)
	log.Info.Printf("Control server addr: %s\n", addr)
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.handler)
	http.ListenAndServe(addr, mux)
}

func (s *ControlServer) handler(w http.ResponseWriter, r *http.Request) {
	cr := controller.Request{
		HttpRequest: r,
	}
	if r.Method == "GET" {
		s.c.GetControlMessage(&cr, &w)
	} else if r.Method == "POST" {
		s.c.PostControlMessage(&cr, &w)
	} else if r.Method == "DELETE" {
		s.c.DeleteControlMessage(&cr, &w)
	} else {
		log.Info.Println("Method in control server not supported: " + r.Method)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Method in control server not supported: "+r.Method)
	}
}

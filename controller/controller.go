// Copyright 2015-2018 Arthur Chunqi Li. All rights reserved.

package controller

import (
	"../config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Controller struct {
	configPath string
}

// New will setup a new controller struct after parsing the options.
func New(_opts config.Options) *Controller {
	s := &Controller{
		configPath: _opts.ConfigPath,
	}

	return s
}

func (c *Controller) Get(r *Request, w *http.ResponseWriter) {
	filename := fmt.Sprintf("%s/%s/%s", c.configPath, r.HttpRequest.URL.Path, c.getFilename(r))
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("%s", err)
		(*w).WriteHeader(http.StatusNotFound)
		return
	}

	//	var response2 Response
	//	response2.Content = "{'name':'john'}"
	//	response2.Status = 200
	//	header := make(http.Header)
	//	header.Add("a", "a1")
	//	header.Add("a", "a2")
	//	header.Add("b", "b")
	//	response2.Header = header
	//	response2.ContentType = "text/json; charset=utf-8"
	//	byte, _ := json.Marshal(response2)
	//	fmt.Printf("Contents of response: %s\n", string(byte))

	var response Response
	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Printf("%s", err)
		(*w).WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Println("Read content:", response)
	for key, values := range response.Header {
		for i := range values {
			(*w).Header().Add(key, values[i])
		}
	}
	if response.ContentType != "" {
		(*w).Header().Set("Content-Type", response.ContentType)
	}
	if response.Status == 0 {
		response.Status = http.StatusOK
	}
	(*w).WriteHeader(response.Status)
	fmt.Fprintf(*w, "%s", response.Content)
}

func (c *Controller) getFilename(r *Request) string {
	return fmt.Sprintf("?!@#$%%^%s&*().json", r.HttpRequest.Method)
}

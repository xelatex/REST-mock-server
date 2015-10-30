// Copyright 2015-2018 Arthur Chunqi Li. All rights reserved.

package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"../api"
	"../config"
	"../log"
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

func (c *Controller) GetMessage(r *Request, w *http.ResponseWriter) {
	filename := fmt.Sprintf("%s/%s/%s", c.configPath, r.HttpRequest.URL.Path, c.getFilename(r.HttpRequest.Method))
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error.Printf("%s", err)
		fmt.Fprintf(*w, "%s", err)
		(*w).WriteHeader(http.StatusNotFound)
		return
	}

	//	var response2 Message
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

	var msg api.Message
	err = json.Unmarshal(data, &msg)
	if err != nil {
		log.Error.Printf("GetMessage: Json unmarshal failed, %s", err)
		fmt.Fprintf(*w, "Json unmarshal failed, %s", err)
		(*w).WriteHeader(http.StatusNotFound)
		return
	}
	log.Debug.Println("GetMessage: Get message:", msg)
	for key, values := range msg.Header {
		for i := range values {
			(*w).Header().Add(key, values[i])
		}
	}
	if msg.ContentType != "" {
		(*w).Header().Set("Content-Type", msg.ContentType)
	}
	if msg.Status == 0 {
		msg.Status = http.StatusOK
	}
	(*w).WriteHeader(msg.Status)
	fmt.Fprintf(*w, "%s", msg.Content)
}

func (c *Controller) getFilename(method string) string {
	return fmt.Sprintf("?!@#$%%^%s&*().json", method)
}

func (c *Controller) GetControlMessage(r *Request, w *http.ResponseWriter) {
	method := r.HttpRequest.Header.Get("method")
	if method == "" {
		method = "GET"
	}
	filename := fmt.Sprintf("%s/%s/%s", c.configPath, r.HttpRequest.URL.Path, c.getFilename(method))
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error.Printf("GetControlMessage: Read file failed, %s", err)
		fmt.Fprintf(*w, "Read file failed, %s", err)
		(*w).WriteHeader(http.StatusNotFound)
		return
	}

	log.Debug.Println("GetControlMessage: Get control message from", r.HttpRequest.URL.Path, "with method", method, ":", strings.TrimSuffix(string(data), "\n"))
	(*w).WriteHeader(http.StatusOK)
	fmt.Fprintf(*w, "%s", string(data))
}

func (c *Controller) PostControlMessage(r *Request, w *http.ResponseWriter) {
	method := r.HttpRequest.Header.Get("method")
	if method == "" {
		method = "GET"
	}
	filename := fmt.Sprintf("%s/%s/%s", c.configPath, r.HttpRequest.URL.Path, c.getFilename(method))
	data, err := ioutil.ReadAll(r.HttpRequest.Body)
	if err != nil {
		log.Error.Printf("PostControlMessage: Read http request body failed, %s", err)
		fmt.Fprintf(*w, "Read http request body failed, %s", err)
		(*w).WriteHeader(http.StatusBadRequest)
		return
	}
	var msg api.Message
	err = json.Unmarshal(data, &msg)
	if err != nil {
		log.Error.Printf("PostControlMessage: Convert request body to Message failed, %s", err)
		fmt.Fprintf(*w, "Convert request body to Message failed, %s", err)
		(*w).WriteHeader(http.StatusBadRequest)
		return
	}
	msgData, err := json.Marshal(msg)
	if err != nil {
		log.Error.Printf("PostControlMessage: Json marshal failed, %s", err)
		fmt.Fprintf(*w, "Json marshal failed, %s", err)
		(*w).WriteHeader(http.StatusInternalServerError)
		return
	}
	err = ioutil.WriteFile(filename, msgData, 0644)
	if err != nil {
		log.Error.Printf("PostControlMessage: Write file failed, %s", err)
		fmt.Fprintf(*w, "Write file failed, %s", err)
		(*w).WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Debug.Println("Write control message:", strings.TrimSuffix(string(msgData), "\n"))
	(*w).WriteHeader(http.StatusOK)
}

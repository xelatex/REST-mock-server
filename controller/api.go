// Copyright 2015-2018 Arthur Chunqi Li. All rights reserved.

package controller

import (
	"net/http"
)

type Request struct {
	HttpRequest *http.Request
}

type Response struct {
	Status      int
	Content     string
	Header      http.Header
	ContentType string
}

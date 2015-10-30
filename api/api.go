// Copyright 2015-2018 Arthur Chunqi Li. All rights reserved.

package api

import (
	"net/http"
)

type Message struct {
	Status      int
	Content     string
	Header      http.Header
	ContentType string
}

/***************************************************************
 *
 * Copyright (c) 2015, Menglong TAN <tanmenglong@gmail.com>
 *
 * This program is free software; you can redistribute it
 * and/or modify it under the terms of the GPL licence
 *
 **************************************************************/

/**
 * Web UI for hpipe
 *
 * @file webui.go
 * @author Menglong TAN <tanmenglong@gmail.com>
 * @date Sat Jan  3 12:45:39 2015
 *
 **/

package webui

import (
	"../../log"
	"fmt"
	"net/http"
)

//===================================================================
// Public APIs
//===================================================================

type WebUI struct {
	host   string
	port   string
	server *http.Server
}

func NewWebUI(host string, port string) *WebUI {
	ret := new(WebUI)
	ret.host = host
	ret.port = port
	return ret
}

func (this *WebUI) Start() {
	http.HandleFunc("/", this.handleIndex())
	log.Fatal(http.ListenAdnServe(this.host+":"+this.port, nil))
}

//===================================================================
// Private
//===================================================================

func (this *WebUI) handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// Copyright 2017 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"fmt"
	"github.com/pingcap/pd/server"
	"github.com/unrolled/render"
	"net/http"
)

type namespaceHandler struct {
	svr *server.Server
	rd  *render.Render
}

func newNamespaceHandler(svr *server.Server, rd *render.Render) *namespaceHandler {
	return &namespaceHandler{
		svr: svr,
		rd:  rd,
	}
}

// Get list namespace mapping
func (h *namespaceHandler) Get(w http.ResponseWriter, r *http.Request) {
	//TODO
	h.rd.JSON(w, http.StatusOK, "here should display namespace list")
}

// Post create a namespace
func (h *namespaceHandler) Post(w http.ResponseWriter, r *http.Request) {
	cluster := h.svr.GetRaftCluster()
	if cluster == nil {
		h.rd.JSON(w, http.StatusInternalServerError, server.ErrNotBootstrapped.Error())
	}

	var input map[string]string
	if err := readJSON(r.Body, &input); err != nil {
		h.rd.JSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	ns := input["namespace"]

	//TODO create namespace
	fmt.Println(ns)

	h.rd.JSON(w, http.StatusOK, fmt.Sprintf("create namespace %s ok", ns))
}
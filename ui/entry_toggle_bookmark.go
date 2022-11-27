// Copyright 2018 Frédéric Guillot. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package ui // import "miniflux.app/ui"

import (
	"miniflux.app/http/response/html"
	"miniflux.app/integration"
	"miniflux.app/model"
	"net/http"

	"miniflux.app/http/request"
	"miniflux.app/http/response/json"
)

func (h *handler) toggleBookmark(w http.ResponseWriter, r *http.Request) {
	entryID := request.RouteInt64Param(r, "entryID")
	if err := h.store.ToggleBookmark(request.UserID(r), entryID); err != nil {
		json.ServerError(w, r, err)
		return
	}
	builder := h.store.NewEntryQueryBuilder(request.UserID(r))
	builder.WithEntryID(entryID)
	builder.WithoutStatus(model.EntryStatusRemoved)

	entry, err := builder.GetEntry()
	if err != nil {
		html.ServerError(w, r, err)
		return
	}

	settings, err := h.store.Integration(request.UserID(r))
	if err != nil {
		json.ServerError(w, r, err)
		return
	}
	if entry.Starred {
		go func() {
			integration.SendEntry(entry, settings)
		}()
	}

	json.OK(w, r, "OK")
}

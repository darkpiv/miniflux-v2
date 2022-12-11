// Copyright 2018 Frédéric Guillot. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package ui // import "miniflux.app/ui"

import (
	"miniflux.app/http/response/json"
	"net/http"

	"miniflux.app/http/request"
	"miniflux.app/http/response/html"
)

func (h *handler) feedCleanup(w http.ResponseWriter, r *http.Request) {
	err := h.store.CleanupFeed(request.UserID(r))
	if err != nil {
		html.ServerError(w, r, err)
		return
	}
	json.NoContent(w, r)

}

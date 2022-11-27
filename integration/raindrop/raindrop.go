// Copyright 2021 Frédéric Guillot. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package raindrop // import "miniflux.app/integration/raindrop"

import (
	"encoding/json"
	"fmt"
	"io"
	"miniflux.app/logger"
	"miniflux.app/model"
	"net/http"
	"strings"
)

type Raindrop struct {
	Created int64  `json:"created"`
	Tags    string `json:"tags"`
	Title   string `json:"title"`
	Link    string `json:"link"`
}

// PushEntry pushes entry to raindrop chat using integration settings provided
func PushEntry(entry *model.Entry, accessToken string) error {
	url := "https://api.raindrop.io/rest/v1/raindrop"
	data := &Raindrop{
		Created: entry.CreatedAt.UnixMilli(),
		Tags:    "hey-you",
		Title:   entry.Title,
		Link:    entry.URL,
	}
	mar, _ := json.Marshal(data)

	payload := strings.NewReader(string(mar))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(res.Body)

	body, _ := io.ReadAll(res.Body)
	logger.Debug("send raindrop successfully", string(body))

	return nil
}

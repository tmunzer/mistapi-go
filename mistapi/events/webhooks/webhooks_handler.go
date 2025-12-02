// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package webhooks

import (
	"encoding/json"
	"github.com/apimatic/go-core-runtime/https"
	"net/http"
)

// Manages WebhooksHandler event processing and event payload parsing.
type WebhooksHandler struct{}

// NewWebhooksHandler creates a new instance of WebhooksHandler.
func NewWebhooksHandler() *WebhooksHandler {
	return &WebhooksHandler{}
}

// ParseEvent parses the event payload.
func (w *WebhooksHandler) ParseEvent(req *http.Request) WebhooksParsingResult {
	bodyBytes, err := https.ReadRequestBody(req)
	if err != nil {
		return WebhooksParsingResultContainer.MarkUnknown()
	}

	var data WebhooksParsingResult
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return WebhooksParsingResultContainer.MarkUnknown()
	}

	return data
}

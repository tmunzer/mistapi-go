// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package webhooks_integrations

import (
	"bytes"
	"framework_integrations"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newWebhooksHandlerRequest(bodyBytes []byte) *http.Request {
	req := httptest.NewRequest("POST", "/events", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func TestOnAlarms_WebhooksHandler_ReturnsAlarmsEvent(t *testing.T) {
	tests := []struct {
		name string
		app  framework_integrations.AppHandler
	}{
		{"Fiber", framework_integrations.NewFiberApp(WebhooksFiberEventHandler)},
		{"Chi", framework_integrations.NewChiApp(WebhooksChiEventHandler)},
		{"Echo", framework_integrations.NewEchoApp(WebhooksEchoEventHandler)},
		{"Gin", framework_integrations.NewGinApp(WebhooksGinEventHandler)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Arrange
			bodyBytes := []byte(`{"events":[{"aps":["string"],"bssids":["string"],"count":0,"event_id":"a7a26ff2-e851-45b6-9634-d595f45458b7","for_site":true,"id":"489f6eca-6276-4993-bfeb-c3cbbbba1f08","last_seen":0.0,"org_id":"a40f5d1f-d889-42e9-94ea-b9b33585fc6b","site_id":"72771e6a-6f5e-4de4-a5b9-1266c4197811","ssids":["string"],"timestamp":0,"type":"string","update":true}],"topic":"alarms"}`)
			request := newWebhooksHandlerRequest(bodyBytes)
			rec := httptest.NewRecorder()

			// Act
			tt.app.ServeHTTP(rec, request)

			// Assert
			if rec.Code != 200 {
				t.Errorf("%s: expected status 200, got %d", tt.name, rec.Code)
			}
			if len(rec.Body.Bytes()) == 0 {
				t.Errorf("%s: expected non-empty response", tt.name)
			}
		})
	}
}

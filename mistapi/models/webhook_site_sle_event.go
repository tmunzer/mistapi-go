package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// WebhookSiteSleEvent represents a WebhookSiteSleEvent struct.
type WebhookSiteSleEvent struct {
	OrgId                *uuid.UUID              `json:"org_id,omitempty"`
	SiteId               *uuid.UUID              `json:"site_id,omitempty"`
	Sle                  *WebhookSiteSleEventSle `json:"sle,omitempty"`
	Timestamp            *int                    `json:"timestamp,omitempty"`
	AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookSiteSleEvent.
// It customizes the JSON marshaling process for WebhookSiteSleEvent objects.
func (w WebhookSiteSleEvent) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookSiteSleEvent object to a map representation for JSON marshaling.
func (w WebhookSiteSleEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, w.AdditionalProperties)
	if w.OrgId != nil {
		structMap["org_id"] = w.OrgId
	}
	if w.SiteId != nil {
		structMap["site_id"] = w.SiteId
	}
	if w.Sle != nil {
		structMap["sle"] = w.Sle.toMap()
	}
	if w.Timestamp != nil {
		structMap["timestamp"] = w.Timestamp
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookSiteSleEvent.
// It customizes the JSON unmarshaling process for WebhookSiteSleEvent objects.
func (w *WebhookSiteSleEvent) UnmarshalJSON(input []byte) error {
	var temp tempWebhookSiteSleEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "org_id", "site_id", "sle", "timestamp")
	if err != nil {
		return err
	}

	w.AdditionalProperties = additionalProperties
	w.OrgId = temp.OrgId
	w.SiteId = temp.SiteId
	w.Sle = temp.Sle
	w.Timestamp = temp.Timestamp
	return nil
}

// tempWebhookSiteSleEvent is a temporary struct used for validating the fields of WebhookSiteSleEvent.
type tempWebhookSiteSleEvent struct {
	OrgId     *uuid.UUID              `json:"org_id,omitempty"`
	SiteId    *uuid.UUID              `json:"site_id,omitempty"`
	Sle       *WebhookSiteSleEventSle `json:"sle,omitempty"`
	Timestamp *int                    `json:"timestamp,omitempty"`
}

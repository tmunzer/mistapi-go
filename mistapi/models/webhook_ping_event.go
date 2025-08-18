// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// WebhookPingEvent represents a WebhookPingEvent struct.
type WebhookPingEvent struct {
	// Unique ID of the object instance in the Mist Organization
	Id     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	SiteId uuid.UUID `json:"site_id"`
	// Epoch (seconds)
	Timestamp            float64                `json:"timestamp"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookPingEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookPingEvent) String() string {
	return fmt.Sprintf(
		"WebhookPingEvent[Id=%v, Name=%v, SiteId=%v, Timestamp=%v, AdditionalProperties=%v]",
		w.Id, w.Name, w.SiteId, w.Timestamp, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookPingEvent.
// It customizes the JSON marshaling process for WebhookPingEvent objects.
func (w WebhookPingEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"id", "name", "site_id", "timestamp"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookPingEvent object to a map representation for JSON marshaling.
func (w WebhookPingEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	structMap["id"] = w.Id
	structMap["name"] = w.Name
	structMap["site_id"] = w.SiteId
	structMap["timestamp"] = w.Timestamp
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookPingEvent.
// It customizes the JSON unmarshaling process for WebhookPingEvent objects.
func (w *WebhookPingEvent) UnmarshalJSON(input []byte) error {
	var temp tempWebhookPingEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "name", "site_id", "timestamp")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.Id = *temp.Id
	w.Name = *temp.Name
	w.SiteId = *temp.SiteId
	w.Timestamp = *temp.Timestamp
	return nil
}

// tempWebhookPingEvent is a temporary struct used for validating the fields of WebhookPingEvent.
type tempWebhookPingEvent struct {
	Id        *uuid.UUID `json:"id"`
	Name      *string    `json:"name"`
	SiteId    *uuid.UUID `json:"site_id"`
	Timestamp *float64   `json:"timestamp"`
}

func (w *tempWebhookPingEvent) validate() error {
	var errs []string
	if w.Id == nil {
		errs = append(errs, "required field `id` is missing for type `webhook_ping_event`")
	}
	if w.Name == nil {
		errs = append(errs, "required field `name` is missing for type `webhook_ping_event`")
	}
	if w.SiteId == nil {
		errs = append(errs, "required field `site_id` is missing for type `webhook_ping_event`")
	}
	if w.Timestamp == nil {
		errs = append(errs, "required field `timestamp` is missing for type `webhook_ping_event`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

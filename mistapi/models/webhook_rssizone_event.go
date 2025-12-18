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

// WebhookRssizoneEvent represents a WebhookRssizoneEvent struct.
type WebhookRssizoneEvent struct {
	// Client MAC address
	Mac   string    `json:"mac"`
	MapId uuid.UUID `json:"map_id"`
	// RSSI zone name
	RssizoneId uuid.UUID `json:"rssizone_id"`
	SiteId     uuid.UUID `json:"site_id"`
	// Epoch (seconds)
	Timestamp float64 `json:"timestamp"`
	// enum: `enter`, `exit`
	Trigger WebhookZoneEventTriggerEnum `json:"trigger"`
	// Type of client. enum: `asset` (BLE Tag), `sdk`, `wifi`
	Type                 WebhookZoneEventTypeEnum `json:"type"`
	AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookRssizoneEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookRssizoneEvent) String() string {
	return fmt.Sprintf(
		"WebhookRssizoneEvent[Mac=%v, MapId=%v, RssizoneId=%v, SiteId=%v, Timestamp=%v, Trigger=%v, Type=%v, AdditionalProperties=%v]",
		w.Mac, w.MapId, w.RssizoneId, w.SiteId, w.Timestamp, w.Trigger, w.Type, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookRssizoneEvent.
// It customizes the JSON marshaling process for WebhookRssizoneEvent objects.
func (w WebhookRssizoneEvent) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"mac", "map_id", "rssizone_id", "site_id", "timestamp", "trigger", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookRssizoneEvent object to a map representation for JSON marshaling.
func (w WebhookRssizoneEvent) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	structMap["mac"] = w.Mac
	structMap["map_id"] = w.MapId
	structMap["rssizone_id"] = w.RssizoneId
	structMap["site_id"] = w.SiteId
	structMap["timestamp"] = w.Timestamp
	structMap["trigger"] = w.Trigger
	structMap["type"] = w.Type
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookRssizoneEvent.
// It customizes the JSON unmarshaling process for WebhookRssizoneEvent objects.
func (w *WebhookRssizoneEvent) UnmarshalJSON(input []byte) error {
	var temp tempWebhookRssizoneEvent
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "map_id", "rssizone_id", "site_id", "timestamp", "trigger", "type")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.Mac = *temp.Mac
	w.MapId = *temp.MapId
	w.RssizoneId = *temp.RssizoneId
	w.SiteId = *temp.SiteId
	w.Timestamp = *temp.Timestamp
	w.Trigger = *temp.Trigger
	w.Type = *temp.Type
	return nil
}

// tempWebhookRssizoneEvent is a temporary struct used for validating the fields of WebhookRssizoneEvent.
type tempWebhookRssizoneEvent struct {
	Mac        *string                      `json:"mac"`
	MapId      *uuid.UUID                   `json:"map_id"`
	RssizoneId *uuid.UUID                   `json:"rssizone_id"`
	SiteId     *uuid.UUID                   `json:"site_id"`
	Timestamp  *float64                     `json:"timestamp"`
	Trigger    *WebhookZoneEventTriggerEnum `json:"trigger"`
	Type       *WebhookZoneEventTypeEnum    `json:"type"`
}

func (w *tempWebhookRssizoneEvent) validate() error {
	var errs []string
	if w.Mac == nil {
		errs = append(errs, "required field `mac` is missing for type `webhook_rssizone_event`")
	}
	if w.MapId == nil {
		errs = append(errs, "required field `map_id` is missing for type `webhook_rssizone_event`")
	}
	if w.RssizoneId == nil {
		errs = append(errs, "required field `rssizone_id` is missing for type `webhook_rssizone_event`")
	}
	if w.SiteId == nil {
		errs = append(errs, "required field `site_id` is missing for type `webhook_rssizone_event`")
	}
	if w.Timestamp == nil {
		errs = append(errs, "required field `timestamp` is missing for type `webhook_rssizone_event`")
	}
	if w.Trigger == nil {
		errs = append(errs, "required field `trigger` is missing for type `webhook_rssizone_event`")
	}
	if w.Type == nil {
		errs = append(errs, "required field `type` is missing for type `webhook_rssizone_event`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

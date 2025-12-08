// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// WebhookAssetRawRssi represents a WebhookAssetRawRssi struct.
// Sample of the `asset-raw-rssi` webhook payload.
// This webhook topic provides raw data from packets emitted by named and filtered assets.
// Raw data webhooks are a special subset of webhooks that provide insight into raw data packets emitted by a client,
// identified by their advertising MAC address (assets, discovered ble, connected wifi, unconnected wifi).
// The data that client raw data webhooks encompasses are reporting AP information, RSSI Data, and any special packets/telemetry
// packets that the client may emit.
// Note that client raw webhooks are the raw data coming from the client and do not contain the X,Y location data of the client.
// In order to get the location data for a client please see our location webhooks.
// Clients can be identified uniquely across these client raw data topics and location webhook topic using MAC address as the Unique identifier (client identifier).
type WebhookAssetRawRssi struct {
	Events []WebhookAssetRawRssiEvent `json:"events"`
	// enum: `asset-raw-rssi`
	Topic                string                 `json:"topic"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookAssetRawRssi,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookAssetRawRssi) String() string {
	return fmt.Sprintf(
		"WebhookAssetRawRssi[Events=%v, Topic=%v, AdditionalProperties=%v]",
		w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookAssetRawRssi.
// It customizes the JSON marshaling process for WebhookAssetRawRssi objects.
func (w WebhookAssetRawRssi) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"events", "topic"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookAssetRawRssi object to a map representation for JSON marshaling.
func (w WebhookAssetRawRssi) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	structMap["events"] = w.Events
	structMap["topic"] = w.Topic
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookAssetRawRssi.
// It customizes the JSON unmarshaling process for WebhookAssetRawRssi objects.
func (w *WebhookAssetRawRssi) UnmarshalJSON(input []byte) error {
	var temp tempWebhookAssetRawRssi
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "events", "topic")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.Events = *temp.Events
	w.Topic = *temp.Topic
	return nil
}

// tempWebhookAssetRawRssi is a temporary struct used for validating the fields of WebhookAssetRawRssi.
type tempWebhookAssetRawRssi struct {
	Events *[]WebhookAssetRawRssiEvent `json:"events"`
	Topic  *string                     `json:"topic"`
}

func (w *tempWebhookAssetRawRssi) validate() error {
	var errs []string
	if w.Events == nil {
		errs = append(errs, "required field `events` is missing for type `webhook_asset_raw_rssi`")
	}
	if w.Topic == nil {
		errs = append(errs, "required field `topic` is missing for type `webhook_asset_raw_rssi`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

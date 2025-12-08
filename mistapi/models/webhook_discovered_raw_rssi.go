// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// WebhookDiscoveredRawRssi represents a WebhookDiscoveredRawRssi struct.
// Sample of the `discovered-raw-rssi` webhook payload.
type WebhookDiscoveredRawRssi struct {
	Events               []WebhookDiscoveredRawRssiEvent `json:"events,omitempty"`
	Topic                string                          `json:"topic"`
	AdditionalProperties map[string]interface{}          `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookDiscoveredRawRssi,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookDiscoveredRawRssi) String() string {
	return fmt.Sprintf(
		"WebhookDiscoveredRawRssi[Events=%v, Topic=%v, AdditionalProperties=%v]",
		w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookDiscoveredRawRssi.
// It customizes the JSON marshaling process for WebhookDiscoveredRawRssi objects.
func (w WebhookDiscoveredRawRssi) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"events", "topic"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookDiscoveredRawRssi object to a map representation for JSON marshaling.
func (w WebhookDiscoveredRawRssi) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.Events != nil {
		structMap["events"] = w.Events
	}
	structMap["topic"] = w.Topic
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookDiscoveredRawRssi.
// It customizes the JSON unmarshaling process for WebhookDiscoveredRawRssi objects.
func (w *WebhookDiscoveredRawRssi) UnmarshalJSON(input []byte) error {
	var temp tempWebhookDiscoveredRawRssi
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

	w.Events = temp.Events
	w.Topic = *temp.Topic
	return nil
}

// tempWebhookDiscoveredRawRssi is a temporary struct used for validating the fields of WebhookDiscoveredRawRssi.
type tempWebhookDiscoveredRawRssi struct {
	Events []WebhookDiscoveredRawRssiEvent `json:"events,omitempty"`
	Topic  *string                         `json:"topic"`
}

func (w *tempWebhookDiscoveredRawRssi) validate() error {
	var errs []string
	if w.Topic == nil {
		errs = append(errs, "required field `topic` is missing for type `webhook_discovered_raw_rssi`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

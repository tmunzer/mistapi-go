// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// WebhookDeviceEvents represents a WebhookDeviceEvents struct.
// Device event webhook sample
type WebhookDeviceEvents struct {
	// List of events
	Events []WebhookDeviceEventsEvent `json:"events"`
	// Topic subscribed to
	Topic                string                 `json:"topic"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookDeviceEvents,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookDeviceEvents) String() string {
	return fmt.Sprintf(
		"WebhookDeviceEvents[Events=%v, Topic=%v, AdditionalProperties=%v]",
		w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookDeviceEvents.
// It customizes the JSON marshaling process for WebhookDeviceEvents objects.
func (w WebhookDeviceEvents) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"events", "topic"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookDeviceEvents object to a map representation for JSON marshaling.
func (w WebhookDeviceEvents) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	structMap["events"] = w.Events
	structMap["topic"] = w.Topic
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookDeviceEvents.
// It customizes the JSON unmarshaling process for WebhookDeviceEvents objects.
func (w *WebhookDeviceEvents) UnmarshalJSON(input []byte) error {
	var temp tempWebhookDeviceEvents
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

// tempWebhookDeviceEvents is a temporary struct used for validating the fields of WebhookDeviceEvents.
type tempWebhookDeviceEvents struct {
	Events *[]WebhookDeviceEventsEvent `json:"events"`
	Topic  *string                     `json:"topic"`
}

func (w *tempWebhookDeviceEvents) validate() error {
	var errs []string
	if w.Events == nil {
		errs = append(errs, "required field `events` is missing for type `webhook_device_events`")
	}
	if w.Topic == nil {
		errs = append(errs, "required field `topic` is missing for type `webhook_device_events`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

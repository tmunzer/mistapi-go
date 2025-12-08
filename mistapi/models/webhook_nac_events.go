// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// WebhookNacEvents represents a WebhookNacEvents struct.
// Sample of the `nac-events` webhook payload.
type WebhookNacEvents struct {
	Events               []NacClientEvent           `json:"events,omitempty"`
	Topic                *WebhookNacEventsTopicEnum `json:"topic,omitempty"`
	AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookNacEvents,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookNacEvents) String() string {
	return fmt.Sprintf(
		"WebhookNacEvents[Events=%v, Topic=%v, AdditionalProperties=%v]",
		w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookNacEvents.
// It customizes the JSON marshaling process for WebhookNacEvents objects.
func (w WebhookNacEvents) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"events", "topic"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookNacEvents object to a map representation for JSON marshaling.
func (w WebhookNacEvents) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.Events != nil {
		structMap["events"] = w.Events
	}
	if w.Topic != nil {
		structMap["topic"] = w.Topic
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookNacEvents.
// It customizes the JSON unmarshaling process for WebhookNacEvents objects.
func (w *WebhookNacEvents) UnmarshalJSON(input []byte) error {
	var temp tempWebhookNacEvents
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "events", "topic")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.Events = temp.Events
	w.Topic = temp.Topic
	return nil
}

// tempWebhookNacEvents is a temporary struct used for validating the fields of WebhookNacEvents.
type tempWebhookNacEvents struct {
	Events []NacClientEvent           `json:"events,omitempty"`
	Topic  *WebhookNacEventsTopicEnum `json:"topic,omitempty"`
}

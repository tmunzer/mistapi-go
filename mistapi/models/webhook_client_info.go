// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// WebhookClientInfo represents a WebhookClientInfo struct.
// Sample of the `client-info` webhook payload.
type WebhookClientInfo struct {
	Events []WebhookClientInfoEvent `json:"events,omitempty"`
	// enum: `client-info`
	Topic                *WebhookClientInfoTopicEnum `json:"topic,omitempty"`
	AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookClientInfo,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookClientInfo) String() string {
	return fmt.Sprintf(
		"WebhookClientInfo[Events=%v, Topic=%v, AdditionalProperties=%v]",
		w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookClientInfo.
// It customizes the JSON marshaling process for WebhookClientInfo objects.
func (w WebhookClientInfo) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"events", "topic"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookClientInfo object to a map representation for JSON marshaling.
func (w WebhookClientInfo) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookClientInfo.
// It customizes the JSON unmarshaling process for WebhookClientInfo objects.
func (w *WebhookClientInfo) UnmarshalJSON(input []byte) error {
	var temp tempWebhookClientInfo
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

// tempWebhookClientInfo is a temporary struct used for validating the fields of WebhookClientInfo.
type tempWebhookClientInfo struct {
	Events []WebhookClientInfoEvent    `json:"events,omitempty"`
	Topic  *WebhookClientInfoTopicEnum `json:"topic,omitempty"`
}

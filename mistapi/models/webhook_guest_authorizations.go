// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// WebhookGuestAuthorizations represents a WebhookGuestAuthorizations struct.
// Sample of the `guest-authorizations` webhook payload.
type WebhookGuestAuthorizations struct {
	// List of events
	Events []WebhookGuestAuthorizationsEvent `json:"events,omitempty"`
	// enum: `guest-authorizations`
	Topic                *WebhookGuestAuthorizationsTopicEnum `json:"topic,omitempty"`
	AdditionalProperties map[string]interface{}               `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookGuestAuthorizations,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookGuestAuthorizations) String() string {
	return fmt.Sprintf(
		"WebhookGuestAuthorizations[Events=%v, Topic=%v, AdditionalProperties=%v]",
		w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookGuestAuthorizations.
// It customizes the JSON marshaling process for WebhookGuestAuthorizations objects.
func (w WebhookGuestAuthorizations) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"events", "topic"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookGuestAuthorizations object to a map representation for JSON marshaling.
func (w WebhookGuestAuthorizations) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookGuestAuthorizations.
// It customizes the JSON unmarshaling process for WebhookGuestAuthorizations objects.
func (w *WebhookGuestAuthorizations) UnmarshalJSON(input []byte) error {
	var temp tempWebhookGuestAuthorizations
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

// tempWebhookGuestAuthorizations is a temporary struct used for validating the fields of WebhookGuestAuthorizations.
type tempWebhookGuestAuthorizations struct {
	Events []WebhookGuestAuthorizationsEvent    `json:"events,omitempty"`
	Topic  *WebhookGuestAuthorizationsTopicEnum `json:"topic,omitempty"`
}

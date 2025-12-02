// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// WebhookClientSessions represents a WebhookClientSessions struct.
// Sample of the `client-sessions` webhook payload.
type WebhookClientSessions struct {
	Events []WebhookClientSessionsEvent `json:"events"`
	// enum: `client-sessions`
	Topic                string                 `json:"topic"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookClientSessions,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookClientSessions) String() string {
	return fmt.Sprintf(
		"WebhookClientSessions[Events=%v, Topic=%v, AdditionalProperties=%v]",
		w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookClientSessions.
// It customizes the JSON marshaling process for WebhookClientSessions objects.
func (w WebhookClientSessions) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"events", "topic"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookClientSessions object to a map representation for JSON marshaling.
func (w WebhookClientSessions) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	structMap["events"] = w.Events
	structMap["topic"] = w.Topic
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookClientSessions.
// It customizes the JSON unmarshaling process for WebhookClientSessions objects.
func (w *WebhookClientSessions) UnmarshalJSON(input []byte) error {
	var temp tempWebhookClientSessions
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

// tempWebhookClientSessions is a temporary struct used for validating the fields of WebhookClientSessions.
type tempWebhookClientSessions struct {
	Events *[]WebhookClientSessionsEvent `json:"events"`
	Topic  *string                       `json:"topic"`
}

func (w *tempWebhookClientSessions) validate() error {
	var errs []string
	if w.Events == nil {
		errs = append(errs, "required field `events` is missing for type `webhook_client_sessions`")
	}
	if w.Topic == nil {
		errs = append(errs, "required field `topic` is missing for type `webhook_client_sessions`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

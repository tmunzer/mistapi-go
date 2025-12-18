// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// WebhookMxedgeEvents represents a WebhookMxedgeEvents struct.
// Sample of the `mxedge-events` webhook payload.
type WebhookMxedgeEvents struct {
	Events []MxedgeEvent `json:"events"`
	// enum: `mxedge-events`
	Topic                string                 `json:"topic"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookMxedgeEvents,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookMxedgeEvents) String() string {
	return fmt.Sprintf(
		"WebhookMxedgeEvents[Events=%v, Topic=%v, AdditionalProperties=%v]",
		w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookMxedgeEvents.
// It customizes the JSON marshaling process for WebhookMxedgeEvents objects.
func (w WebhookMxedgeEvents) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"events", "topic"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookMxedgeEvents object to a map representation for JSON marshaling.
func (w WebhookMxedgeEvents) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	structMap["events"] = w.Events
	structMap["topic"] = w.Topic
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookMxedgeEvents.
// It customizes the JSON unmarshaling process for WebhookMxedgeEvents objects.
func (w *WebhookMxedgeEvents) UnmarshalJSON(input []byte) error {
	var temp tempWebhookMxedgeEvents
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

// tempWebhookMxedgeEvents is a temporary struct used for validating the fields of WebhookMxedgeEvents.
type tempWebhookMxedgeEvents struct {
	Events *[]MxedgeEvent `json:"events"`
	Topic  *string        `json:"topic"`
}

func (w *tempWebhookMxedgeEvents) validate() error {
	var errs []string
	if w.Events == nil {
		errs = append(errs, "required field `events` is missing for type `webhook_mxedge_events`")
	}
	if w.Topic == nil {
		errs = append(errs, "required field `topic` is missing for type `webhook_mxedge_events`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

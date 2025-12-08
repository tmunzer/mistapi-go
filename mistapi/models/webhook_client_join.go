// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// WebhookClientJoin represents a WebhookClientJoin struct.
// Sample of the `client-join` webhook payload.
type WebhookClientJoin struct {
	Events []WebhookClientJoinEvent `json:"events"`
	// enum: `client-join`
	Topic                string                 `json:"topic"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookClientJoin,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookClientJoin) String() string {
	return fmt.Sprintf(
		"WebhookClientJoin[Events=%v, Topic=%v, AdditionalProperties=%v]",
		w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookClientJoin.
// It customizes the JSON marshaling process for WebhookClientJoin objects.
func (w WebhookClientJoin) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"events", "topic"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookClientJoin object to a map representation for JSON marshaling.
func (w WebhookClientJoin) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	structMap["events"] = w.Events
	structMap["topic"] = w.Topic
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookClientJoin.
// It customizes the JSON unmarshaling process for WebhookClientJoin objects.
func (w *WebhookClientJoin) UnmarshalJSON(input []byte) error {
	var temp tempWebhookClientJoin
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

// tempWebhookClientJoin is a temporary struct used for validating the fields of WebhookClientJoin.
type tempWebhookClientJoin struct {
	Events *[]WebhookClientJoinEvent `json:"events"`
	Topic  *string                   `json:"topic"`
}

func (w *tempWebhookClientJoin) validate() error {
	var errs []string
	if w.Events == nil {
		errs = append(errs, "required field `events` is missing for type `webhook_client_join`")
	}
	if w.Topic == nil {
		errs = append(errs, "required field `topic` is missing for type `webhook_client_join`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

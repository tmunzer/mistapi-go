// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// WebhookLocation represents a WebhookLocation struct.
// Sample of the `location` webhook payload.
type WebhookLocation struct {
	// List of events
	Events               []WebhookLocationEvent `json:"events"`
	Topic                string                 `json:"topic"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookLocation,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookLocation) String() string {
	return fmt.Sprintf(
		"WebhookLocation[Events=%v, Topic=%v, AdditionalProperties=%v]",
		w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookLocation.
// It customizes the JSON marshaling process for WebhookLocation objects.
func (w WebhookLocation) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"events", "topic"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookLocation object to a map representation for JSON marshaling.
func (w WebhookLocation) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	structMap["events"] = w.Events
	structMap["topic"] = w.Topic
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookLocation.
// It customizes the JSON unmarshaling process for WebhookLocation objects.
func (w *WebhookLocation) UnmarshalJSON(input []byte) error {
	var temp tempWebhookLocation
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

// tempWebhookLocation is a temporary struct used for validating the fields of WebhookLocation.
type tempWebhookLocation struct {
	Events *[]WebhookLocationEvent `json:"events"`
	Topic  *string                 `json:"topic"`
}

func (w *tempWebhookLocation) validate() error {
	var errs []string
	if w.Events == nil {
		errs = append(errs, "required field `events` is missing for type `webhook_location`")
	}
	if w.Topic == nil {
		errs = append(errs, "required field `topic` is missing for type `webhook_location`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

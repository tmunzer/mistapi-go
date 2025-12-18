// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// WebhookRssizone represents a WebhookRssizone struct.
// Sample of the `rssizone` webhook payload.
type WebhookRssizone struct {
	Events []WebhookRssizoneEvent `json:"events"`
	// enum: `rssizone`
	Topic                string                 `json:"topic"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookRssizone,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookRssizone) String() string {
	return fmt.Sprintf(
		"WebhookRssizone[Events=%v, Topic=%v, AdditionalProperties=%v]",
		w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookRssizone.
// It customizes the JSON marshaling process for WebhookRssizone objects.
func (w WebhookRssizone) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"events", "topic"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookRssizone object to a map representation for JSON marshaling.
func (w WebhookRssizone) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	structMap["events"] = w.Events
	structMap["topic"] = w.Topic
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookRssizone.
// It customizes the JSON unmarshaling process for WebhookRssizone objects.
func (w *WebhookRssizone) UnmarshalJSON(input []byte) error {
	var temp tempWebhookRssizone
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

// tempWebhookRssizone is a temporary struct used for validating the fields of WebhookRssizone.
type tempWebhookRssizone struct {
	Events *[]WebhookRssizoneEvent `json:"events"`
	Topic  *string                 `json:"topic"`
}

func (w *tempWebhookRssizone) validate() error {
	var errs []string
	if w.Events == nil {
		errs = append(errs, "required field `events` is missing for type `webhook_rssizone`")
	}
	if w.Topic == nil {
		errs = append(errs, "required field `topic` is missing for type `webhook_rssizone`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

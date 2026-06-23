// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// WebhookMinisApplication represents a WebhookMinisApplication struct.
// Sample of the `minis-application` webhook payload.
type WebhookMinisApplication struct {
	// Marvis Minis application test events included in a webhook delivery
	Events []WebhookMinisApplicationEvent `json:"events"`
	// Webhook topic name for Minis application test deliveries. enum: `minis-application`
	Topic                string                 `json:"topic"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookMinisApplication,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookMinisApplication) String() string {
	return fmt.Sprintf(
		"WebhookMinisApplication[Events=%v, Topic=%v, AdditionalProperties=%v]",
		w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookMinisApplication.
// It customizes the JSON marshaling process for WebhookMinisApplication objects.
func (w WebhookMinisApplication) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"events", "topic"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookMinisApplication object to a map representation for JSON marshaling.
func (w WebhookMinisApplication) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	structMap["events"] = w.Events
	structMap["topic"] = w.Topic
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookMinisApplication.
// It customizes the JSON unmarshaling process for WebhookMinisApplication objects.
func (w *WebhookMinisApplication) UnmarshalJSON(input []byte) error {
	var temp tempWebhookMinisApplication
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

// tempWebhookMinisApplication is a temporary struct used for validating the fields of WebhookMinisApplication.
type tempWebhookMinisApplication struct {
	Events *[]WebhookMinisApplicationEvent `json:"events"`
	Topic  *string                         `json:"topic"`
}

func (w *tempWebhookMinisApplication) validate() error {
	var errs []string
	if w.Events == nil {
		errs = append(errs, "required field `events` is missing for type `webhook_minis_application`")
	}
	if w.Topic == nil {
		errs = append(errs, "required field `topic` is missing for type `webhook_minis_application`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// WebhookMinisReachability represents a WebhookMinisReachability struct.
// Sample of the `minis-reachability` webhook payload.
type WebhookMinisReachability struct {
	// Marvis Minis reachability test events included in a webhook delivery
	Events []WebhookMinisReachabilityEvent `json:"events"`
	// Webhook topic name for Minis reachability test deliveries. enum: `minis-reachability`
	Topic                string                 `json:"topic"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookMinisReachability,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookMinisReachability) String() string {
	return fmt.Sprintf(
		"WebhookMinisReachability[Events=%v, Topic=%v, AdditionalProperties=%v]",
		w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookMinisReachability.
// It customizes the JSON marshaling process for WebhookMinisReachability objects.
func (w WebhookMinisReachability) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"events", "topic"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookMinisReachability object to a map representation for JSON marshaling.
func (w WebhookMinisReachability) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	structMap["events"] = w.Events
	structMap["topic"] = w.Topic
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookMinisReachability.
// It customizes the JSON unmarshaling process for WebhookMinisReachability objects.
func (w *WebhookMinisReachability) UnmarshalJSON(input []byte) error {
	var temp tempWebhookMinisReachability
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

// tempWebhookMinisReachability is a temporary struct used for validating the fields of WebhookMinisReachability.
type tempWebhookMinisReachability struct {
	Events *[]WebhookMinisReachabilityEvent `json:"events"`
	Topic  *string                          `json:"topic"`
}

func (w *tempWebhookMinisReachability) validate() error {
	var errs []string
	if w.Events == nil {
		errs = append(errs, "required field `events` is missing for type `webhook_minis_reachability`")
	}
	if w.Topic == nil {
		errs = append(errs, "required field `topic` is missing for type `webhook_minis_reachability`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

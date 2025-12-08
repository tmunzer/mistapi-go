// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// WebhookPing represents a WebhookPing struct.
// Sample of the `ping` webhook payload.\n\nThe `ping` webhook can be manually sent with the following API calls:\n- for a Site level webhook with the [Ping Site Webhook]($e/Orgs%20Webhooks/pingOrgWebhook) endpoint\n- for an Org level webhook with the [Ping Org Webhook]($e/Orgs%20Webhooks/pingOrgWebhook) endpoint
type WebhookPing struct {
	Events               []WebhookPingEvent     `json:"events"`
	Topic                string                 `json:"topic"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookPing,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookPing) String() string {
	return fmt.Sprintf(
		"WebhookPing[Events=%v, Topic=%v, AdditionalProperties=%v]",
		w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookPing.
// It customizes the JSON marshaling process for WebhookPing objects.
func (w WebhookPing) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"events", "topic"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookPing object to a map representation for JSON marshaling.
func (w WebhookPing) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	structMap["events"] = w.Events
	structMap["topic"] = w.Topic
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookPing.
// It customizes the JSON unmarshaling process for WebhookPing objects.
func (w *WebhookPing) UnmarshalJSON(input []byte) error {
	var temp tempWebhookPing
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

// tempWebhookPing is a temporary struct used for validating the fields of WebhookPing.
type tempWebhookPing struct {
	Events *[]WebhookPingEvent `json:"events"`
	Topic  *string             `json:"topic"`
}

func (w *tempWebhookPing) validate() error {
	var errs []string
	if w.Events == nil {
		errs = append(errs, "required field `events` is missing for type `webhook_ping`")
	}
	if w.Topic == nil {
		errs = append(errs, "required field `topic` is missing for type `webhook_ping`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

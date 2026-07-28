// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// WebhookRule represents a WebhookRule struct.
// Filtering rule that permits or blocks webhook events for a topic
type WebhookRule struct {
	// Webhook filtering action. enum: `permit`, `block`
	Action *WebhookActionEnum `json:"action,omitempty"`
	// Event payload matching criteria. Property key is the event field name and the value is the list of accepted values (e.g. matching field `type` to [`AP_DISCONNECTED`])
	Matching map[string][]string `json:"matching,omitempty"`
	// Webhook topic this rule applies to
	Topic                string                 `json:"topic"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookRule,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookRule) String() string {
	return fmt.Sprintf(
		"WebhookRule[Action=%v, Matching=%v, Topic=%v, AdditionalProperties=%v]",
		w.Action, w.Matching, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookRule.
// It customizes the JSON marshaling process for WebhookRule objects.
func (w WebhookRule) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"action", "matching", "topic"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookRule object to a map representation for JSON marshaling.
func (w WebhookRule) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.Action != nil {
		structMap["action"] = w.Action
	}
	if w.Matching != nil {
		structMap["matching"] = w.Matching
	}
	structMap["topic"] = w.Topic
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookRule.
// It customizes the JSON unmarshaling process for WebhookRule objects.
func (w *WebhookRule) UnmarshalJSON(input []byte) error {
	var temp tempWebhookRule
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "action", "matching", "topic")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.Action = temp.Action
	w.Matching = temp.Matching
	w.Topic = *temp.Topic
	return nil
}

// tempWebhookRule is a temporary struct used for validating the fields of WebhookRule.
type tempWebhookRule struct {
	Action   *WebhookActionEnum  `json:"action,omitempty"`
	Matching map[string][]string `json:"matching,omitempty"`
	Topic    *string             `json:"topic"`
}

func (w *tempWebhookRule) validate() error {
	var errs []string
	if w.Topic == nil {
		errs = append(errs, "required field `topic` is missing for type `webhook_rule`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

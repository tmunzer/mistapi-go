// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// WebhookNacAccounting represents a WebhookNacAccounting struct.
// nac-accounting webhook sample
type WebhookNacAccounting struct {
    Events               []WebhookNacAccountingEvent `json:"events,omitempty"`
    Topic                *string                     `json:"topic,omitempty"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookNacAccounting,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookNacAccounting) String() string {
    return fmt.Sprintf(
    	"WebhookNacAccounting[Events=%v, Topic=%v, AdditionalProperties=%v]",
    	w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookNacAccounting.
// It customizes the JSON marshaling process for WebhookNacAccounting objects.
func (w WebhookNacAccounting) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "events", "topic"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookNacAccounting object to a map representation for JSON marshaling.
func (w WebhookNacAccounting) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookNacAccounting.
// It customizes the JSON unmarshaling process for WebhookNacAccounting objects.
func (w *WebhookNacAccounting) UnmarshalJSON(input []byte) error {
    var temp tempWebhookNacAccounting
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

// tempWebhookNacAccounting is a temporary struct used for validating the fields of WebhookNacAccounting.
type tempWebhookNacAccounting  struct {
    Events []WebhookNacAccountingEvent `json:"events,omitempty"`
    Topic  *string                     `json:"topic,omitempty"`
}

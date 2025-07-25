// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// WebhookSiteSle represents a WebhookSiteSle struct.
type WebhookSiteSle struct {
    Events               []WebhookSiteSleEvent  `json:"events,omitempty"`
    Topic                *string                `json:"topic,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookSiteSle,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookSiteSle) String() string {
    return fmt.Sprintf(
    	"WebhookSiteSle[Events=%v, Topic=%v, AdditionalProperties=%v]",
    	w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookSiteSle.
// It customizes the JSON marshaling process for WebhookSiteSle objects.
func (w WebhookSiteSle) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "events", "topic"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookSiteSle object to a map representation for JSON marshaling.
func (w WebhookSiteSle) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookSiteSle.
// It customizes the JSON unmarshaling process for WebhookSiteSle objects.
func (w *WebhookSiteSle) UnmarshalJSON(input []byte) error {
    var temp tempWebhookSiteSle
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

// tempWebhookSiteSle is a temporary struct used for validating the fields of WebhookSiteSle.
type tempWebhookSiteSle  struct {
    Events []WebhookSiteSleEvent `json:"events,omitempty"`
    Topic  *string               `json:"topic,omitempty"`
}

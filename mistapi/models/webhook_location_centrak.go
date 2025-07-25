// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WebhookLocationCentrak represents a WebhookLocationCentrak struct.
// location webhook sample
type WebhookLocationCentrak struct {
    // List of events
    Events               []WebhookLocationCentrakEvent `json:"events"`
    // Topic subscribed to
    Topic                string                        `json:"topic"`
    AdditionalProperties map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookLocationCentrak,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookLocationCentrak) String() string {
    return fmt.Sprintf(
    	"WebhookLocationCentrak[Events=%v, Topic=%v, AdditionalProperties=%v]",
    	w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookLocationCentrak.
// It customizes the JSON marshaling process for WebhookLocationCentrak objects.
func (w WebhookLocationCentrak) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "events", "topic"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookLocationCentrak object to a map representation for JSON marshaling.
func (w WebhookLocationCentrak) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookLocationCentrak.
// It customizes the JSON unmarshaling process for WebhookLocationCentrak objects.
func (w *WebhookLocationCentrak) UnmarshalJSON(input []byte) error {
    var temp tempWebhookLocationCentrak
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

// tempWebhookLocationCentrak is a temporary struct used for validating the fields of WebhookLocationCentrak.
type tempWebhookLocationCentrak  struct {
    Events *[]WebhookLocationCentrakEvent `json:"events"`
    Topic  *string                        `json:"topic"`
}

func (w *tempWebhookLocationCentrak) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `webhook_location_centrak`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `webhook_location_centrak`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

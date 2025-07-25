// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WebhookLocationAsset represents a WebhookLocationAsset struct.
// Location Asset sample
type WebhookLocationAsset struct {
    // List of events
    Events               []WebhookLocationAssetEvent `json:"events"`
    // Topic subscribed to
    Topic                string                      `json:"topic"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookLocationAsset,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookLocationAsset) String() string {
    return fmt.Sprintf(
    	"WebhookLocationAsset[Events=%v, Topic=%v, AdditionalProperties=%v]",
    	w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookLocationAsset.
// It customizes the JSON marshaling process for WebhookLocationAsset objects.
func (w WebhookLocationAsset) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "events", "topic"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookLocationAsset object to a map representation for JSON marshaling.
func (w WebhookLocationAsset) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookLocationAsset.
// It customizes the JSON unmarshaling process for WebhookLocationAsset objects.
func (w *WebhookLocationAsset) UnmarshalJSON(input []byte) error {
    var temp tempWebhookLocationAsset
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

// tempWebhookLocationAsset is a temporary struct used for validating the fields of WebhookLocationAsset.
type tempWebhookLocationAsset  struct {
    Events *[]WebhookLocationAssetEvent `json:"events"`
    Topic  *string                      `json:"topic"`
}

func (w *tempWebhookLocationAsset) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `webhook_location_asset`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `webhook_location_asset`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WebhookOccupancyAlerts represents a WebhookOccupancyAlerts struct.
// Occupancy alert webhook sample
type WebhookOccupancyAlerts struct {
    Events               []WebhookOccupancyAlertsEvent `json:"events"`
    Topic                string                        `json:"topic"`
    AdditionalProperties map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookOccupancyAlerts,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookOccupancyAlerts) String() string {
    return fmt.Sprintf(
    	"WebhookOccupancyAlerts[Events=%v, Topic=%v, AdditionalProperties=%v]",
    	w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookOccupancyAlerts.
// It customizes the JSON marshaling process for WebhookOccupancyAlerts objects.
func (w WebhookOccupancyAlerts) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "events", "topic"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookOccupancyAlerts object to a map representation for JSON marshaling.
func (w WebhookOccupancyAlerts) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookOccupancyAlerts.
// It customizes the JSON unmarshaling process for WebhookOccupancyAlerts objects.
func (w *WebhookOccupancyAlerts) UnmarshalJSON(input []byte) error {
    var temp tempWebhookOccupancyAlerts
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

// tempWebhookOccupancyAlerts is a temporary struct used for validating the fields of WebhookOccupancyAlerts.
type tempWebhookOccupancyAlerts  struct {
    Events *[]WebhookOccupancyAlertsEvent `json:"events"`
    Topic  *string                        `json:"topic"`
}

func (w *tempWebhookOccupancyAlerts) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `webhook_occupancy_alerts`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `webhook_occupancy_alerts`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

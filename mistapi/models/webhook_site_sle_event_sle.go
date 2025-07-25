// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// WebhookSiteSleEventSle represents a WebhookSiteSleEventSle struct.
type WebhookSiteSleEventSle struct {
    ApAvailability       *float64               `json:"ap-availability,omitempty"`
    SuccessfulConnect    *float64               `json:"successful-connect,omitempty"`
    TimeToConnect        *float64               `json:"time-to-connect,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookSiteSleEventSle,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookSiteSleEventSle) String() string {
    return fmt.Sprintf(
    	"WebhookSiteSleEventSle[ApAvailability=%v, SuccessfulConnect=%v, TimeToConnect=%v, AdditionalProperties=%v]",
    	w.ApAvailability, w.SuccessfulConnect, w.TimeToConnect, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookSiteSleEventSle.
// It customizes the JSON marshaling process for WebhookSiteSleEventSle objects.
func (w WebhookSiteSleEventSle) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "ap-availability", "successful-connect", "time-to-connect"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookSiteSleEventSle object to a map representation for JSON marshaling.
func (w WebhookSiteSleEventSle) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.ApAvailability != nil {
        structMap["ap-availability"] = w.ApAvailability
    }
    if w.SuccessfulConnect != nil {
        structMap["successful-connect"] = w.SuccessfulConnect
    }
    if w.TimeToConnect != nil {
        structMap["time-to-connect"] = w.TimeToConnect
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookSiteSleEventSle.
// It customizes the JSON unmarshaling process for WebhookSiteSleEventSle objects.
func (w *WebhookSiteSleEventSle) UnmarshalJSON(input []byte) error {
    var temp tempWebhookSiteSleEventSle
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap-availability", "successful-connect", "time-to-connect")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.ApAvailability = temp.ApAvailability
    w.SuccessfulConnect = temp.SuccessfulConnect
    w.TimeToConnect = temp.TimeToConnect
    return nil
}

// tempWebhookSiteSleEventSle is a temporary struct used for validating the fields of WebhookSiteSleEventSle.
type tempWebhookSiteSleEventSle  struct {
    ApAvailability    *float64 `json:"ap-availability,omitempty"`
    SuccessfulConnect *float64 `json:"successful-connect,omitempty"`
    TimeToConnect     *float64 `json:"time-to-connect,omitempty"`
}

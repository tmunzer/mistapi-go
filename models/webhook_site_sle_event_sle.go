package models

import (
    "encoding/json"
)

// WebhookSiteSleEventSle represents a WebhookSiteSleEventSle struct.
type WebhookSiteSleEventSle struct {
    ApAvailability       *float64       `json:"ap-availability,omitempty"`
    SuccessfulConnect    *float64       `json:"successful-connect,omitempty"`
    TimeToConnect        *float64       `json:"time-to-connect,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookSiteSleEventSle.
// It customizes the JSON marshaling process for WebhookSiteSleEventSle objects.
func (w WebhookSiteSleEventSle) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookSiteSleEventSle object to a map representation for JSON marshaling.
func (w WebhookSiteSleEventSle) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
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
    var temp webhookSiteSleEventSle
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap-availability", "successful-connect", "time-to-connect")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.ApAvailability = temp.ApAvailability
    w.SuccessfulConnect = temp.SuccessfulConnect
    w.TimeToConnect = temp.TimeToConnect
    return nil
}

// webhookSiteSleEventSle is a temporary struct used for validating the fields of WebhookSiteSleEventSle.
type webhookSiteSleEventSle  struct {
    ApAvailability    *float64 `json:"ap-availability,omitempty"`
    SuccessfulConnect *float64 `json:"successful-connect,omitempty"`
    TimeToConnect     *float64 `json:"time-to-connect,omitempty"`
}
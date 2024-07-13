package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebhookOccupancyAlerts represents a WebhookOccupancyAlerts struct.
// occupancy alert webhook sample
type WebhookOccupancyAlerts struct {
    Events               []WebhookOccupancyAlertsEvent `json:"events"`
    Topic                string                        `json:"topic"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookOccupancyAlerts.
// It customizes the JSON marshaling process for WebhookOccupancyAlerts objects.
func (w WebhookOccupancyAlerts) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookOccupancyAlerts object to a map representation for JSON marshaling.
func (w WebhookOccupancyAlerts) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookOccupancyAlerts.
// It customizes the JSON unmarshaling process for WebhookOccupancyAlerts objects.
func (w *WebhookOccupancyAlerts) UnmarshalJSON(input []byte) error {
    var temp webhookOccupancyAlerts
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "events", "topic")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Events = *temp.Events
    w.Topic = *temp.Topic
    return nil
}

// webhookOccupancyAlerts is a temporary struct used for validating the fields of WebhookOccupancyAlerts.
type webhookOccupancyAlerts  struct {
    Events *[]WebhookOccupancyAlertsEvent `json:"events"`
    Topic  *string                        `json:"topic"`
}

func (w *webhookOccupancyAlerts) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `Webhook_Occupancy_Alerts`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `Webhook_Occupancy_Alerts`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
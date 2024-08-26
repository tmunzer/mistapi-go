package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebhookZone represents a WebhookZone struct.
// zone webhook sample
type WebhookZone struct {
    // list of events
    Events               []WebhookZoneEvent `json:"events"`
    // topic subscribed to
    Topic                string             `json:"topic"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookZone.
// It customizes the JSON marshaling process for WebhookZone objects.
func (w WebhookZone) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookZone object to a map representation for JSON marshaling.
func (w WebhookZone) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookZone.
// It customizes the JSON unmarshaling process for WebhookZone objects.
func (w *WebhookZone) UnmarshalJSON(input []byte) error {
    var temp tempWebhookZone
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

// tempWebhookZone is a temporary struct used for validating the fields of WebhookZone.
type tempWebhookZone  struct {
    Events *[]WebhookZoneEvent `json:"events"`
    Topic  *string             `json:"topic"`
}

func (w *tempWebhookZone) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `webhook_zone`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `webhook_zone`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

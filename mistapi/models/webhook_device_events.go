package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebhookDeviceEvents represents a WebhookDeviceEvents struct.
// device event webhook sample
type WebhookDeviceEvents struct {
    // list of events
    Events               []WebhookDeviceEventsEvent `json:"events"`
    // topic subscribed to
    Topic                string                     `json:"topic"`
    AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookDeviceEvents.
// It customizes the JSON marshaling process for WebhookDeviceEvents objects.
func (w WebhookDeviceEvents) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookDeviceEvents object to a map representation for JSON marshaling.
func (w WebhookDeviceEvents) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookDeviceEvents.
// It customizes the JSON unmarshaling process for WebhookDeviceEvents objects.
func (w *WebhookDeviceEvents) UnmarshalJSON(input []byte) error {
    var temp webhookDeviceEvents
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

// webhookDeviceEvents is a temporary struct used for validating the fields of WebhookDeviceEvents.
type webhookDeviceEvents  struct {
    Events *[]WebhookDeviceEventsEvent `json:"events"`
    Topic  *string                     `json:"topic"`
}

func (w *webhookDeviceEvents) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `Webhook_Device_Events`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `Webhook_Device_Events`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

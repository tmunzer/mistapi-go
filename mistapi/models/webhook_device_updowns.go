package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebhookDeviceUpdowns represents a WebhookDeviceUpdowns struct.
// device up/down webhook sample
type WebhookDeviceUpdowns struct {
    Events               []WebhookDeviceUpdownsEvent `json:"events"`
    Topic                string                      `json:"topic"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookDeviceUpdowns.
// It customizes the JSON marshaling process for WebhookDeviceUpdowns objects.
func (w WebhookDeviceUpdowns) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookDeviceUpdowns object to a map representation for JSON marshaling.
func (w WebhookDeviceUpdowns) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookDeviceUpdowns.
// It customizes the JSON unmarshaling process for WebhookDeviceUpdowns objects.
func (w *WebhookDeviceUpdowns) UnmarshalJSON(input []byte) error {
    var temp tempWebhookDeviceUpdowns
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

// tempWebhookDeviceUpdowns is a temporary struct used for validating the fields of WebhookDeviceUpdowns.
type tempWebhookDeviceUpdowns  struct {
    Events *[]WebhookDeviceUpdownsEvent `json:"events"`
    Topic  *string                      `json:"topic"`
}

func (w *tempWebhookDeviceUpdowns) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `webhook_device_updowns`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `webhook_device_updowns`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebhookLocationUnclient represents a WebhookLocationUnclient struct.
// Location Unconnected Client sample
type WebhookLocationUnclient struct {
    // list of events
    Events               []WebhookLocationUnclientEvent `json:"events"`
    // topic subscribed to
    Topic                string                         `json:"topic"`
    AdditionalProperties map[string]any                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookLocationUnclient.
// It customizes the JSON marshaling process for WebhookLocationUnclient objects.
func (w WebhookLocationUnclient) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookLocationUnclient object to a map representation for JSON marshaling.
func (w WebhookLocationUnclient) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookLocationUnclient.
// It customizes the JSON unmarshaling process for WebhookLocationUnclient objects.
func (w *WebhookLocationUnclient) UnmarshalJSON(input []byte) error {
    var temp webhookLocationUnclient
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

// webhookLocationUnclient is a temporary struct used for validating the fields of WebhookLocationUnclient.
type webhookLocationUnclient  struct {
    Events *[]WebhookLocationUnclientEvent `json:"events"`
    Topic  *string                         `json:"topic"`
}

func (w *webhookLocationUnclient) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `Webhook_Location_Unclient`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `Webhook_Location_Unclient`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

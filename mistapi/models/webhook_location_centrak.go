package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebhookLocationCentrak represents a WebhookLocationCentrak struct.
// location webhook sample
type WebhookLocationCentrak struct {
    // list of events
    Events               []WebhookLocationCentrakEvent `json:"events"`
    // topic subscribed to
    Topic                string                        `json:"topic"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookLocationCentrak.
// It customizes the JSON marshaling process for WebhookLocationCentrak objects.
func (w WebhookLocationCentrak) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookLocationCentrak object to a map representation for JSON marshaling.
func (w WebhookLocationCentrak) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "events", "topic")
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
    return errors.New(strings.Join(errs, "\n"))
}

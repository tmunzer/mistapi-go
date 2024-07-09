package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebhookClientSessions represents a WebhookClientSessions struct.
// client session webhook sample
type WebhookClientSessions struct {
    Events               []WebhookClientSessionsEvent `json:"events"`
    Topic                string                       `json:"topic"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookClientSessions.
// It customizes the JSON marshaling process for WebhookClientSessions objects.
func (w WebhookClientSessions) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookClientSessions object to a map representation for JSON marshaling.
func (w WebhookClientSessions) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookClientSessions.
// It customizes the JSON unmarshaling process for WebhookClientSessions objects.
func (w *WebhookClientSessions) UnmarshalJSON(input []byte) error {
    var temp webhookClientSessions
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

// webhookClientSessions is a temporary struct used for validating the fields of WebhookClientSessions.
type webhookClientSessions  struct {
    Events *[]WebhookClientSessionsEvent `json:"events"`
    Topic  *string                       `json:"topic"`
}

func (w *webhookClientSessions) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `Webhook_Client_Sessions`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `Webhook_Client_Sessions`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebhookClientJoin represents a WebhookClientJoin struct.
// client join webhook sample
type WebhookClientJoin struct {
    Events               []WebhookClientJoinEvent `json:"events"`
    Topic                string                   `json:"topic"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookClientJoin.
// It customizes the JSON marshaling process for WebhookClientJoin objects.
func (w WebhookClientJoin) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookClientJoin object to a map representation for JSON marshaling.
func (w WebhookClientJoin) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookClientJoin.
// It customizes the JSON unmarshaling process for WebhookClientJoin objects.
func (w *WebhookClientJoin) UnmarshalJSON(input []byte) error {
    var temp webhookClientJoin
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

// webhookClientJoin is a temporary struct used for validating the fields of WebhookClientJoin.
type webhookClientJoin  struct {
    Events *[]WebhookClientJoinEvent `json:"events"`
    Topic  *string                   `json:"topic"`
}

func (w *webhookClientJoin) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `Webhook_Client_Join`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `Webhook_Client_Join`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

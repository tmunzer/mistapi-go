package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebhookLocationClient represents a WebhookLocationClient struct.
// Location Client sample
type WebhookLocationClient struct {
    // list of events
    Events               []WebhookLocationClientEvent `json:"events"`
    // topic subscribed to
    Topic                string                       `json:"topic"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookLocationClient.
// It customizes the JSON marshaling process for WebhookLocationClient objects.
func (w WebhookLocationClient) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookLocationClient object to a map representation for JSON marshaling.
func (w WebhookLocationClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookLocationClient.
// It customizes the JSON unmarshaling process for WebhookLocationClient objects.
func (w *WebhookLocationClient) UnmarshalJSON(input []byte) error {
    var temp webhookLocationClient
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

// webhookLocationClient is a temporary struct used for validating the fields of WebhookLocationClient.
type webhookLocationClient  struct {
    Events *[]WebhookLocationClientEvent `json:"events"`
    Topic  *string                       `json:"topic"`
}

func (w *webhookLocationClient) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `Webhook_Location_Client`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `Webhook_Location_Client`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
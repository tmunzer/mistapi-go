package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebhookLocationSdk represents a WebhookLocationSdk struct.
// Location SDK sample
type WebhookLocationSdk struct {
    // list of events
    Events               []WebhookLocationSdkEvent `json:"events"`
    // topic subscribed to
    Topic                string                    `json:"topic"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookLocationSdk.
// It customizes the JSON marshaling process for WebhookLocationSdk objects.
func (w WebhookLocationSdk) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookLocationSdk object to a map representation for JSON marshaling.
func (w WebhookLocationSdk) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookLocationSdk.
// It customizes the JSON unmarshaling process for WebhookLocationSdk objects.
func (w *WebhookLocationSdk) UnmarshalJSON(input []byte) error {
    var temp tempWebhookLocationSdk
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

// tempWebhookLocationSdk is a temporary struct used for validating the fields of WebhookLocationSdk.
type tempWebhookLocationSdk  struct {
    Events *[]WebhookLocationSdkEvent `json:"events"`
    Topic  *string                    `json:"topic"`
}

func (w *tempWebhookLocationSdk) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `webhook_location_sdk`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `webhook_location_sdk`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

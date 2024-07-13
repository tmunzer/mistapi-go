package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebhookLocationAsset represents a WebhookLocationAsset struct.
// Location Asset sample
type WebhookLocationAsset struct {
    // list of events
    Events               []WebhookLocationAssetEvent `json:"events"`
    // topic subscribed to
    Topic                string                      `json:"topic"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookLocationAsset.
// It customizes the JSON marshaling process for WebhookLocationAsset objects.
func (w WebhookLocationAsset) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookLocationAsset object to a map representation for JSON marshaling.
func (w WebhookLocationAsset) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookLocationAsset.
// It customizes the JSON unmarshaling process for WebhookLocationAsset objects.
func (w *WebhookLocationAsset) UnmarshalJSON(input []byte) error {
    var temp webhookLocationAsset
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

// webhookLocationAsset is a temporary struct used for validating the fields of WebhookLocationAsset.
type webhookLocationAsset  struct {
    Events *[]WebhookLocationAssetEvent `json:"events"`
    Topic  *string                      `json:"topic"`
}

func (w *webhookLocationAsset) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `Webhook_Location_Asset`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `Webhook_Location_Asset`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

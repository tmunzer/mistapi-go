package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebhookAssetRaw represents a WebhookAssetRaw struct.
// asset raw webhook sample
type WebhookAssetRaw struct {
    // list of events
    Events               []WebhookAssetRawEvent `json:"events"`
    // topic subscribed to
    Topic                string                 `json:"topic"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookAssetRaw.
// It customizes the JSON marshaling process for WebhookAssetRaw objects.
func (w WebhookAssetRaw) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookAssetRaw object to a map representation for JSON marshaling.
func (w WebhookAssetRaw) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookAssetRaw.
// It customizes the JSON unmarshaling process for WebhookAssetRaw objects.
func (w *WebhookAssetRaw) UnmarshalJSON(input []byte) error {
    var temp tempWebhookAssetRaw
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

// tempWebhookAssetRaw is a temporary struct used for validating the fields of WebhookAssetRaw.
type tempWebhookAssetRaw  struct {
    Events *[]WebhookAssetRawEvent `json:"events"`
    Topic  *string                 `json:"topic"`
}

func (w *tempWebhookAssetRaw) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `webhook_asset_raw`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `webhook_asset_raw`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

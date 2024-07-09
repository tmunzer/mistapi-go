package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebhookDiscoveredRawRssi represents a WebhookDiscoveredRawRssi struct.
type WebhookDiscoveredRawRssi struct {
    Events               []WebhookDiscoveredRawRssiEvent `json:"events,omitempty"`
    Topic                string                          `json:"topic"`
    AdditionalProperties map[string]any                  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookDiscoveredRawRssi.
// It customizes the JSON marshaling process for WebhookDiscoveredRawRssi objects.
func (w WebhookDiscoveredRawRssi) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookDiscoveredRawRssi object to a map representation for JSON marshaling.
func (w WebhookDiscoveredRawRssi) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Events != nil {
        structMap["events"] = w.Events
    }
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookDiscoveredRawRssi.
// It customizes the JSON unmarshaling process for WebhookDiscoveredRawRssi objects.
func (w *WebhookDiscoveredRawRssi) UnmarshalJSON(input []byte) error {
    var temp webhookDiscoveredRawRssi
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
    w.Events = temp.Events
    w.Topic = *temp.Topic
    return nil
}

// webhookDiscoveredRawRssi is a temporary struct used for validating the fields of WebhookDiscoveredRawRssi.
type webhookDiscoveredRawRssi  struct {
    Events []WebhookDiscoveredRawRssiEvent `json:"events,omitempty"`
    Topic  *string                         `json:"topic"`
}

func (w *webhookDiscoveredRawRssi) validate() error {
    var errs []string
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `Webhook_Discovered_Raw_Rssi`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

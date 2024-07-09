package models

import (
    "encoding/json"
)

// WebhookSiteSle represents a WebhookSiteSle struct.
type WebhookSiteSle struct {
    Events               []WebhookSiteSleEvent `json:"events,omitempty"`
    Topic                *string               `json:"topic,omitempty"`
    AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookSiteSle.
// It customizes the JSON marshaling process for WebhookSiteSle objects.
func (w WebhookSiteSle) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookSiteSle object to a map representation for JSON marshaling.
func (w WebhookSiteSle) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Events != nil {
        structMap["events"] = w.Events
    }
    if w.Topic != nil {
        structMap["topic"] = w.Topic
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookSiteSle.
// It customizes the JSON unmarshaling process for WebhookSiteSle objects.
func (w *WebhookSiteSle) UnmarshalJSON(input []byte) error {
    var temp webhookSiteSle
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "events", "topic")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Events = temp.Events
    w.Topic = temp.Topic
    return nil
}

// webhookSiteSle is a temporary struct used for validating the fields of WebhookSiteSle.
type webhookSiteSle  struct {
    Events []WebhookSiteSleEvent `json:"events,omitempty"`
    Topic  *string               `json:"topic,omitempty"`
}

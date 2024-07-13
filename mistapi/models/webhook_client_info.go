package models

import (
    "encoding/json"
)

// WebhookClientInfo represents a WebhookClientInfo struct.
type WebhookClientInfo struct {
    Events               []WebhookClientInfoEvent    `json:"events,omitempty"`
    Topic                *WebhookClientInfoTopicEnum `json:"topic,omitempty"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookClientInfo.
// It customizes the JSON marshaling process for WebhookClientInfo objects.
func (w WebhookClientInfo) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookClientInfo object to a map representation for JSON marshaling.
func (w WebhookClientInfo) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookClientInfo.
// It customizes the JSON unmarshaling process for WebhookClientInfo objects.
func (w *WebhookClientInfo) UnmarshalJSON(input []byte) error {
    var temp webhookClientInfo
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

// webhookClientInfo is a temporary struct used for validating the fields of WebhookClientInfo.
type webhookClientInfo  struct {
    Events []WebhookClientInfoEvent    `json:"events,omitempty"`
    Topic  *WebhookClientInfoTopicEnum `json:"topic,omitempty"`
}

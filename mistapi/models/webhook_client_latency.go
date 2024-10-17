package models

import (
	"encoding/json"
)

// WebhookClientLatency represents a WebhookClientLatency struct.
type WebhookClientLatency struct {
	Events               []WebhookClientLatencyEvent `json:"events,omitempty"`
	Topic                *string                     `json:"topic,omitempty"`
	AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookClientLatency.
// It customizes the JSON marshaling process for WebhookClientLatency objects.
func (w WebhookClientLatency) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookClientLatency object to a map representation for JSON marshaling.
func (w WebhookClientLatency) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookClientLatency.
// It customizes the JSON unmarshaling process for WebhookClientLatency objects.
func (w *WebhookClientLatency) UnmarshalJSON(input []byte) error {
	var temp tempWebhookClientLatency
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

// tempWebhookClientLatency is a temporary struct used for validating the fields of WebhookClientLatency.
type tempWebhookClientLatency struct {
	Events []WebhookClientLatencyEvent `json:"events,omitempty"`
	Topic  *string                     `json:"topic,omitempty"`
}

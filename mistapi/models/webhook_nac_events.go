package models

import (
	"encoding/json"
)

// WebhookNacEvents represents a WebhookNacEvents struct.
type WebhookNacEvents struct {
	Events               []WebhookNacEventsEvent `json:"events,omitempty"`
	Topic                *string                 `json:"topic,omitempty"`
	AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookNacEvents.
// It customizes the JSON marshaling process for WebhookNacEvents objects.
func (w WebhookNacEvents) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookNacEvents object to a map representation for JSON marshaling.
func (w WebhookNacEvents) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookNacEvents.
// It customizes the JSON unmarshaling process for WebhookNacEvents objects.
func (w *WebhookNacEvents) UnmarshalJSON(input []byte) error {
	var temp tempWebhookNacEvents
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

// tempWebhookNacEvents is a temporary struct used for validating the fields of WebhookNacEvents.
type tempWebhookNacEvents struct {
	Events []WebhookNacEventsEvent `json:"events,omitempty"`
	Topic  *string                 `json:"topic,omitempty"`
}

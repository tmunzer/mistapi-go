package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// WebhookPing represents a WebhookPing struct.
// ping webhook
type WebhookPing struct {
	Events               []WebhookPingEvent `json:"events"`
	Topic                string             `json:"topic"`
	AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookPing.
// It customizes the JSON marshaling process for WebhookPing objects.
func (w WebhookPing) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookPing object to a map representation for JSON marshaling.
func (w WebhookPing) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, w.AdditionalProperties)
	structMap["events"] = w.Events
	structMap["topic"] = w.Topic
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookPing.
// It customizes the JSON unmarshaling process for WebhookPing objects.
func (w *WebhookPing) UnmarshalJSON(input []byte) error {
	var temp tempWebhookPing
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

// tempWebhookPing is a temporary struct used for validating the fields of WebhookPing.
type tempWebhookPing struct {
	Events *[]WebhookPingEvent `json:"events"`
	Topic  *string             `json:"topic"`
}

func (w *tempWebhookPing) validate() error {
	var errs []string
	if w.Events == nil {
		errs = append(errs, "required field `events` is missing for type `webhook_ping`")
	}
	if w.Topic == nil {
		errs = append(errs, "required field `topic` is missing for type `webhook_ping`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

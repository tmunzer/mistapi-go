package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebhookAudits represents a WebhookAudits struct.
// audit webhook sample
type WebhookAudits struct {
    Events               []WebhookAuditEvent `json:"events"`
    Topic                string              `json:"topic"`
    AdditionalProperties map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookAudits.
// It customizes the JSON marshaling process for WebhookAudits objects.
func (w WebhookAudits) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookAudits object to a map representation for JSON marshaling.
func (w WebhookAudits) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookAudits.
// It customizes the JSON unmarshaling process for WebhookAudits objects.
func (w *WebhookAudits) UnmarshalJSON(input []byte) error {
    var temp webhookAudits
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

// webhookAudits is a temporary struct used for validating the fields of WebhookAudits.
type webhookAudits  struct {
    Events *[]WebhookAuditEvent `json:"events"`
    Topic  *string              `json:"topic"`
}

func (w *webhookAudits) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `Webhook_Audits`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `Webhook_Audits`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
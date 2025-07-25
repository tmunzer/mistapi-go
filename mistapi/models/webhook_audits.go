// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WebhookAudits represents a WebhookAudits struct.
// audit webhook sample
type WebhookAudits struct {
    Events               []WebhookAuditEvent    `json:"events"`
    Topic                string                 `json:"topic"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookAudits,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookAudits) String() string {
    return fmt.Sprintf(
    	"WebhookAudits[Events=%v, Topic=%v, AdditionalProperties=%v]",
    	w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookAudits.
// It customizes the JSON marshaling process for WebhookAudits objects.
func (w WebhookAudits) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "events", "topic"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookAudits object to a map representation for JSON marshaling.
func (w WebhookAudits) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookAudits.
// It customizes the JSON unmarshaling process for WebhookAudits objects.
func (w *WebhookAudits) UnmarshalJSON(input []byte) error {
    var temp tempWebhookAudits
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "events", "topic")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Events = *temp.Events
    w.Topic = *temp.Topic
    return nil
}

// tempWebhookAudits is a temporary struct used for validating the fields of WebhookAudits.
type tempWebhookAudits  struct {
    Events *[]WebhookAuditEvent `json:"events"`
    Topic  *string              `json:"topic"`
}

func (w *tempWebhookAudits) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `webhook_audits`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `webhook_audits`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

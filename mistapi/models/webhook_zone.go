package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// WebhookZone represents a WebhookZone struct.
// Zone webhook sample
type WebhookZone struct {
    // List of events
    Events               []WebhookZoneEvent     `json:"events"`
    // Topic subscribed to
    Topic                string                 `json:"topic"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookZone,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookZone) String() string {
    return fmt.Sprintf(
    	"WebhookZone[Events=%v, Topic=%v, AdditionalProperties=%v]",
    	w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookZone.
// It customizes the JSON marshaling process for WebhookZone objects.
func (w WebhookZone) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "events", "topic"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookZone object to a map representation for JSON marshaling.
func (w WebhookZone) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookZone.
// It customizes the JSON unmarshaling process for WebhookZone objects.
func (w *WebhookZone) UnmarshalJSON(input []byte) error {
    var temp tempWebhookZone
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

// tempWebhookZone is a temporary struct used for validating the fields of WebhookZone.
type tempWebhookZone  struct {
    Events *[]WebhookZoneEvent `json:"events"`
    Topic  *string             `json:"topic"`
}

func (w *tempWebhookZone) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `webhook_zone`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `webhook_zone`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

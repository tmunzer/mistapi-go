package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebhookAlarms represents a WebhookAlarms struct.
// **N.B.**: Fields like `aps`, `bssids`, `ssids` are event specific. They are relevant to this event type ( rogue-ap-detected). For a different event type, different fields may be sent. These don’t contain all affected entities and are representative samples of entities (capped at 10). For marvis action related events, we expose `details` to include more event specific details.
// Events specific fields for other alarm event type can be found with API https://api.mist.com/api/v1/const/alarm_defs, under “fields” array of /alarm_defs response object.
type WebhookAlarms struct {
    // list of events
    Events               []WebhookAlarmEvent `json:"events"`
    // topic subscribed to
    Topic                string              `json:"topic"`
    AdditionalProperties map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookAlarms.
// It customizes the JSON marshaling process for WebhookAlarms objects.
func (w WebhookAlarms) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookAlarms object to a map representation for JSON marshaling.
func (w WebhookAlarms) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["events"] = w.Events
    structMap["topic"] = w.Topic
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookAlarms.
// It customizes the JSON unmarshaling process for WebhookAlarms objects.
func (w *WebhookAlarms) UnmarshalJSON(input []byte) error {
    var temp tempWebhookAlarms
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

// tempWebhookAlarms is a temporary struct used for validating the fields of WebhookAlarms.
type tempWebhookAlarms  struct {
    Events *[]WebhookAlarmEvent `json:"events"`
    Topic  *string              `json:"topic"`
}

func (w *tempWebhookAlarms) validate() error {
    var errs []string
    if w.Events == nil {
        errs = append(errs, "required field `events` is missing for type `webhook_alarms`")
    }
    if w.Topic == nil {
        errs = append(errs, "required field `topic` is missing for type `webhook_alarms`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

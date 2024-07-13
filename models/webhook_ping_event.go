package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// WebhookPingEvent represents a WebhookPingEvent struct.
type WebhookPingEvent struct {
    Id                   uuid.UUID      `json:"id"`
    Name                 string         `json:"name"`
    SiteId               uuid.UUID      `json:"site_id"`
    Timestamp            float64        `json:"timestamp"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookPingEvent.
// It customizes the JSON marshaling process for WebhookPingEvent objects.
func (w WebhookPingEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookPingEvent object to a map representation for JSON marshaling.
func (w WebhookPingEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["id"] = w.Id
    structMap["name"] = w.Name
    structMap["site_id"] = w.SiteId
    structMap["timestamp"] = w.Timestamp
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookPingEvent.
// It customizes the JSON unmarshaling process for WebhookPingEvent objects.
func (w *WebhookPingEvent) UnmarshalJSON(input []byte) error {
    var temp webhookPingEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "name", "site_id", "timestamp")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Id = *temp.Id
    w.Name = *temp.Name
    w.SiteId = *temp.SiteId
    w.Timestamp = *temp.Timestamp
    return nil
}

// webhookPingEvent is a temporary struct used for validating the fields of WebhookPingEvent.
type webhookPingEvent  struct {
    Id        *uuid.UUID `json:"id"`
    Name      *string    `json:"name"`
    SiteId    *uuid.UUID `json:"site_id"`
    Timestamp *float64   `json:"timestamp"`
}

func (w *webhookPingEvent) validate() error {
    var errs []string
    if w.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Webhook_Ping_Event`")
    }
    if w.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Webhook_Ping_Event`")
    }
    if w.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `Webhook_Ping_Event`")
    }
    if w.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Webhook_Ping_Event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
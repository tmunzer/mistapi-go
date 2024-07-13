package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// WebhookZoneEvent represents a WebhookZoneEvent struct.
type WebhookZoneEvent struct {
    // uuid of named asset
    AssetId              *uuid.UUID                  `json:"asset_id,omitempty"`
    // uuid of SDK-client
    Id                   uuid.UUID                   `json:"id"`
    // mac address of wifi client or asset
    Mac                  *string                     `json:"mac,omitempty"`
    // map id
    MapId                uuid.UUID                   `json:"map_id"`
    // name of the client, may be empty
    Name                 *string                     `json:"name,omitempty"`
    SiteId               uuid.UUID                   `json:"site_id"`
    // timestamp of the event, epoch
    Timestamp            int                         `json:"timestamp"`
    // enter / exit
    Trigger              WebhookZoneEventTriggerEnum `json:"trigger"`
    Type                 string                      `json:"type"`
    // zone id
    ZoneId               uuid.UUID                   `json:"zone_id"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookZoneEvent.
// It customizes the JSON marshaling process for WebhookZoneEvent objects.
func (w WebhookZoneEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookZoneEvent object to a map representation for JSON marshaling.
func (w WebhookZoneEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.AssetId != nil {
        structMap["asset_id"] = w.AssetId
    }
    structMap["id"] = w.Id
    if w.Mac != nil {
        structMap["mac"] = w.Mac
    }
    structMap["map_id"] = w.MapId
    if w.Name != nil {
        structMap["name"] = w.Name
    }
    structMap["site_id"] = w.SiteId
    structMap["timestamp"] = w.Timestamp
    structMap["trigger"] = w.Trigger
    structMap["type"] = w.Type
    structMap["zone_id"] = w.ZoneId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookZoneEvent.
// It customizes the JSON unmarshaling process for WebhookZoneEvent objects.
func (w *WebhookZoneEvent) UnmarshalJSON(input []byte) error {
    var temp webhookZoneEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "asset_id", "id", "mac", "map_id", "name", "site_id", "timestamp", "trigger", "type", "zone_id")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.AssetId = temp.AssetId
    w.Id = *temp.Id
    w.Mac = temp.Mac
    w.MapId = *temp.MapId
    w.Name = temp.Name
    w.SiteId = *temp.SiteId
    w.Timestamp = *temp.Timestamp
    w.Trigger = *temp.Trigger
    w.Type = *temp.Type
    w.ZoneId = *temp.ZoneId
    return nil
}

// webhookZoneEvent is a temporary struct used for validating the fields of WebhookZoneEvent.
type webhookZoneEvent  struct {
    AssetId   *uuid.UUID                   `json:"asset_id,omitempty"`
    Id        *uuid.UUID                   `json:"id"`
    Mac       *string                      `json:"mac,omitempty"`
    MapId     *uuid.UUID                   `json:"map_id"`
    Name      *string                      `json:"name,omitempty"`
    SiteId    *uuid.UUID                   `json:"site_id"`
    Timestamp *int                         `json:"timestamp"`
    Trigger   *WebhookZoneEventTriggerEnum `json:"trigger"`
    Type      *string                      `json:"type"`
    ZoneId    *uuid.UUID                   `json:"zone_id"`
}

func (w *webhookZoneEvent) validate() error {
    var errs []string
    if w.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Webhook_Zone_Event`")
    }
    if w.MapId == nil {
        errs = append(errs, "required field `map_id` is missing for type `Webhook_Zone_Event`")
    }
    if w.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `Webhook_Zone_Event`")
    }
    if w.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Webhook_Zone_Event`")
    }
    if w.Trigger == nil {
        errs = append(errs, "required field `trigger` is missing for type `Webhook_Zone_Event`")
    }
    if w.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Webhook_Zone_Event`")
    }
    if w.ZoneId == nil {
        errs = append(errs, "required field `zone_id` is missing for type `Webhook_Zone_Event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
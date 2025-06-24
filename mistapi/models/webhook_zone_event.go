package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// WebhookZoneEvent represents a WebhookZoneEvent struct.
type WebhookZoneEvent struct {
    // Only if `type`==`asset`. UUID of named asset
    AssetId              *uuid.UUID                  `json:"asset_id,omitempty"`
    // Only if `type`==`sdk`. UUID of the SDK Client
    Id                   *uuid.UUID                  `json:"id,omitempty"`
    // MAC address of Wi-Fi client, SDK Client or Asset
    Mac                  *string                     `json:"mac,omitempty"`
    // Map id
    MapId                uuid.UUID                   `json:"map_id"`
    // Name of the client, may be empty
    Name                 *string                     `json:"name,omitempty"`
    SiteId               uuid.UUID                   `json:"site_id"`
    // Epoch (seconds)
    Timestamp            float64                     `json:"timestamp"`
    // enum: `enter`, `exit`
    Trigger              WebhookZoneEventTriggerEnum `json:"trigger"`
    // Type of client. enum: `asset` (BLE Tag), `sdk`, `wifi`
    Type                 WebhookZoneEventTypeEnum    `json:"type"`
    // Zone id
    ZoneId               uuid.UUID                   `json:"zone_id"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookZoneEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookZoneEvent) String() string {
    return fmt.Sprintf(
    	"WebhookZoneEvent[AssetId=%v, Id=%v, Mac=%v, MapId=%v, Name=%v, SiteId=%v, Timestamp=%v, Trigger=%v, Type=%v, ZoneId=%v, AdditionalProperties=%v]",
    	w.AssetId, w.Id, w.Mac, w.MapId, w.Name, w.SiteId, w.Timestamp, w.Trigger, w.Type, w.ZoneId, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookZoneEvent.
// It customizes the JSON marshaling process for WebhookZoneEvent objects.
func (w WebhookZoneEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "asset_id", "id", "mac", "map_id", "name", "site_id", "timestamp", "trigger", "type", "zone_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookZoneEvent object to a map representation for JSON marshaling.
func (w WebhookZoneEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.AssetId != nil {
        structMap["asset_id"] = w.AssetId
    }
    if w.Id != nil {
        structMap["id"] = w.Id
    }
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
    var temp tempWebhookZoneEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "asset_id", "id", "mac", "map_id", "name", "site_id", "timestamp", "trigger", "type", "zone_id")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.AssetId = temp.AssetId
    w.Id = temp.Id
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

// tempWebhookZoneEvent is a temporary struct used for validating the fields of WebhookZoneEvent.
type tempWebhookZoneEvent  struct {
    AssetId   *uuid.UUID                   `json:"asset_id,omitempty"`
    Id        *uuid.UUID                   `json:"id,omitempty"`
    Mac       *string                      `json:"mac,omitempty"`
    MapId     *uuid.UUID                   `json:"map_id"`
    Name      *string                      `json:"name,omitempty"`
    SiteId    *uuid.UUID                   `json:"site_id"`
    Timestamp *float64                     `json:"timestamp"`
    Trigger   *WebhookZoneEventTriggerEnum `json:"trigger"`
    Type      *WebhookZoneEventTypeEnum    `json:"type"`
    ZoneId    *uuid.UUID                   `json:"zone_id"`
}

func (w *tempWebhookZoneEvent) validate() error {
    var errs []string
    if w.MapId == nil {
        errs = append(errs, "required field `map_id` is missing for type `webhook_zone_event`")
    }
    if w.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `webhook_zone_event`")
    }
    if w.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `webhook_zone_event`")
    }
    if w.Trigger == nil {
        errs = append(errs, "required field `trigger` is missing for type `webhook_zone_event`")
    }
    if w.Type == nil {
        errs = append(errs, "required field `type` is missing for type `webhook_zone_event`")
    }
    if w.ZoneId == nil {
        errs = append(errs, "required field `zone_id` is missing for type `webhook_zone_event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

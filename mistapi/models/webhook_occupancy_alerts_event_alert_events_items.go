package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// WebhookOccupancyAlertsEventAlertEventsItems represents a WebhookOccupancyAlertsEventAlertEventsItems struct.
type WebhookOccupancyAlertsEventAlertEventsItems struct {
    CurrentOccupancy     int                           `json:"current_occupancy"`
    MapId                uuid.UUID                     `json:"map_id"`
    OccupancyLimit       int                           `json:"occupancy_limit"`
    OrgId                uuid.UUID                     `json:"org_id"`
    Timestamp            float64                       `json:"timestamp"`
    // enum: `COMPLIANCE-OK`, `COMPLIANCE-VIOLATION`
    Type                 WebhookOccupancyAlertTypeEnum `json:"type"`
    ZoneId               uuid.UUID                     `json:"zone_id"`
    ZoneName             string                        `json:"zone_name"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookOccupancyAlertsEventAlertEventsItems.
// It customizes the JSON marshaling process for WebhookOccupancyAlertsEventAlertEventsItems objects.
func (w WebhookOccupancyAlertsEventAlertEventsItems) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookOccupancyAlertsEventAlertEventsItems object to a map representation for JSON marshaling.
func (w WebhookOccupancyAlertsEventAlertEventsItems) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["current_occupancy"] = w.CurrentOccupancy
    structMap["map_id"] = w.MapId
    structMap["occupancy_limit"] = w.OccupancyLimit
    structMap["org_id"] = w.OrgId
    structMap["timestamp"] = w.Timestamp
    structMap["type"] = w.Type
    structMap["zone_id"] = w.ZoneId
    structMap["zone_name"] = w.ZoneName
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookOccupancyAlertsEventAlertEventsItems.
// It customizes the JSON unmarshaling process for WebhookOccupancyAlertsEventAlertEventsItems objects.
func (w *WebhookOccupancyAlertsEventAlertEventsItems) UnmarshalJSON(input []byte) error {
    var temp webhookOccupancyAlertsEventAlertEventsItems
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "current_occupancy", "map_id", "occupancy_limit", "org_id", "timestamp", "type", "zone_id", "zone_name")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.CurrentOccupancy = *temp.CurrentOccupancy
    w.MapId = *temp.MapId
    w.OccupancyLimit = *temp.OccupancyLimit
    w.OrgId = *temp.OrgId
    w.Timestamp = *temp.Timestamp
    w.Type = *temp.Type
    w.ZoneId = *temp.ZoneId
    w.ZoneName = *temp.ZoneName
    return nil
}

// webhookOccupancyAlertsEventAlertEventsItems is a temporary struct used for validating the fields of WebhookOccupancyAlertsEventAlertEventsItems.
type webhookOccupancyAlertsEventAlertEventsItems  struct {
    CurrentOccupancy *int                           `json:"current_occupancy"`
    MapId            *uuid.UUID                     `json:"map_id"`
    OccupancyLimit   *int                           `json:"occupancy_limit"`
    OrgId            *uuid.UUID                     `json:"org_id"`
    Timestamp        *float64                       `json:"timestamp"`
    Type             *WebhookOccupancyAlertTypeEnum `json:"type"`
    ZoneId           *uuid.UUID                     `json:"zone_id"`
    ZoneName         *string                        `json:"zone_name"`
}

func (w *webhookOccupancyAlertsEventAlertEventsItems) validate() error {
    var errs []string
    if w.CurrentOccupancy == nil {
        errs = append(errs, "required field `current_occupancy` is missing for type `Webhook_Occupancy_Alerts_Event_Alert_Events_Items`")
    }
    if w.MapId == nil {
        errs = append(errs, "required field `map_id` is missing for type `Webhook_Occupancy_Alerts_Event_Alert_Events_Items`")
    }
    if w.OccupancyLimit == nil {
        errs = append(errs, "required field `occupancy_limit` is missing for type `Webhook_Occupancy_Alerts_Event_Alert_Events_Items`")
    }
    if w.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `Webhook_Occupancy_Alerts_Event_Alert_Events_Items`")
    }
    if w.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Webhook_Occupancy_Alerts_Event_Alert_Events_Items`")
    }
    if w.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Webhook_Occupancy_Alerts_Event_Alert_Events_Items`")
    }
    if w.ZoneId == nil {
        errs = append(errs, "required field `zone_id` is missing for type `Webhook_Occupancy_Alerts_Event_Alert_Events_Items`")
    }
    if w.ZoneName == nil {
        errs = append(errs, "required field `zone_name` is missing for type `Webhook_Occupancy_Alerts_Event_Alert_Events_Items`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

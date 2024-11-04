package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// WebhookLocationSdkEvent represents a WebhookLocationSdkEvent struct.
type WebhookLocationSdkEvent struct {
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID     `json:"id,omitempty"`
    MapId                *uuid.UUID     `json:"map_id,omitempty"`
    Name                 *string        `json:"name,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    Timestamp            *int           `json:"timestamp,omitempty"`
    Type                 *string        `json:"type,omitempty"`
    // x, in meter
    X                    *float64       `json:"x,omitempty"`
    // y, in meter
    Y                    *float64       `json:"y,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookLocationSdkEvent.
// It customizes the JSON marshaling process for WebhookLocationSdkEvent objects.
func (w WebhookLocationSdkEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookLocationSdkEvent object to a map representation for JSON marshaling.
func (w WebhookLocationSdkEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Id != nil {
        structMap["id"] = w.Id
    }
    if w.MapId != nil {
        structMap["map_id"] = w.MapId
    }
    if w.Name != nil {
        structMap["name"] = w.Name
    }
    if w.SiteId != nil {
        structMap["site_id"] = w.SiteId
    }
    if w.Timestamp != nil {
        structMap["timestamp"] = w.Timestamp
    }
    if w.Type != nil {
        structMap["type"] = w.Type
    }
    if w.X != nil {
        structMap["x"] = w.X
    }
    if w.Y != nil {
        structMap["y"] = w.Y
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookLocationSdkEvent.
// It customizes the JSON unmarshaling process for WebhookLocationSdkEvent objects.
func (w *WebhookLocationSdkEvent) UnmarshalJSON(input []byte) error {
    var temp tempWebhookLocationSdkEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "map_id", "name", "site_id", "timestamp", "type", "x", "y")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Id = temp.Id
    w.MapId = temp.MapId
    w.Name = temp.Name
    w.SiteId = temp.SiteId
    w.Timestamp = temp.Timestamp
    w.Type = temp.Type
    w.X = temp.X
    w.Y = temp.Y
    return nil
}

// tempWebhookLocationSdkEvent is a temporary struct used for validating the fields of WebhookLocationSdkEvent.
type tempWebhookLocationSdkEvent  struct {
    Id        *uuid.UUID `json:"id,omitempty"`
    MapId     *uuid.UUID `json:"map_id,omitempty"`
    Name      *string    `json:"name,omitempty"`
    SiteId    *uuid.UUID `json:"site_id,omitempty"`
    Timestamp *int       `json:"timestamp,omitempty"`
    Type      *string    `json:"type,omitempty"`
    X         *float64   `json:"x,omitempty"`
    Y         *float64   `json:"y,omitempty"`
}

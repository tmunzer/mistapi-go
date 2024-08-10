package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// WebhookAlarmEvent represents a WebhookAlarmEvent struct.
type WebhookAlarmEvent struct {
    Aps                  []string           `json:"aps,omitempty"`
    Bssids               []string           `json:"bssids,omitempty"`
    // If present, represents number of events of given type occurred in current interval, default=1
    Count                *int               `json:"count,omitempty"`
    // event id
    EventId              *uuid.UUID         `json:"event_id,omitempty"`
    ForSite              *bool              `json:"for_site,omitempty"`
    Id                   uuid.UUID          `json:"id"`
    LastSeen             float64            `json:"last_seen"`
    // only for HA. enum: `node0`, `node1`
    Node                 *HaClusterNodeEnum `json:"node,omitempty"`
    OrgId                uuid.UUID          `json:"org_id"`
    SiteId               uuid.UUID          `json:"site_id"`
    Ssids                []string           `json:"ssids,omitempty"`
    Timestamp            int                `json:"timestamp"`
    // event type
    Type                 string             `json:"type"`
    // If presents, represents that this is an update to event with given id sent earlier. default=false
    Update               *bool              `json:"update,omitempty"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookAlarmEvent.
// It customizes the JSON marshaling process for WebhookAlarmEvent objects.
func (w WebhookAlarmEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookAlarmEvent object to a map representation for JSON marshaling.
func (w WebhookAlarmEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Aps != nil {
        structMap["aps"] = w.Aps
    }
    if w.Bssids != nil {
        structMap["bssids"] = w.Bssids
    }
    if w.Count != nil {
        structMap["count"] = w.Count
    }
    if w.EventId != nil {
        structMap["event_id"] = w.EventId
    }
    if w.ForSite != nil {
        structMap["for_site"] = w.ForSite
    }
    structMap["id"] = w.Id
    structMap["last_seen"] = w.LastSeen
    if w.Node != nil {
        structMap["node"] = w.Node
    }
    structMap["org_id"] = w.OrgId
    structMap["site_id"] = w.SiteId
    if w.Ssids != nil {
        structMap["ssids"] = w.Ssids
    }
    structMap["timestamp"] = w.Timestamp
    structMap["type"] = w.Type
    if w.Update != nil {
        structMap["update"] = w.Update
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookAlarmEvent.
// It customizes the JSON unmarshaling process for WebhookAlarmEvent objects.
func (w *WebhookAlarmEvent) UnmarshalJSON(input []byte) error {
    var temp tempWebhookAlarmEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "aps", "bssids", "count", "event_id", "for_site", "id", "last_seen", "node", "org_id", "site_id", "ssids", "timestamp", "type", "update")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Aps = temp.Aps
    w.Bssids = temp.Bssids
    w.Count = temp.Count
    w.EventId = temp.EventId
    w.ForSite = temp.ForSite
    w.Id = *temp.Id
    w.LastSeen = *temp.LastSeen
    w.Node = temp.Node
    w.OrgId = *temp.OrgId
    w.SiteId = *temp.SiteId
    w.Ssids = temp.Ssids
    w.Timestamp = *temp.Timestamp
    w.Type = *temp.Type
    w.Update = temp.Update
    return nil
}

// tempWebhookAlarmEvent is a temporary struct used for validating the fields of WebhookAlarmEvent.
type tempWebhookAlarmEvent  struct {
    Aps       []string           `json:"aps,omitempty"`
    Bssids    []string           `json:"bssids,omitempty"`
    Count     *int               `json:"count,omitempty"`
    EventId   *uuid.UUID         `json:"event_id,omitempty"`
    ForSite   *bool              `json:"for_site,omitempty"`
    Id        *uuid.UUID         `json:"id"`
    LastSeen  *float64           `json:"last_seen"`
    Node      *HaClusterNodeEnum `json:"node,omitempty"`
    OrgId     *uuid.UUID         `json:"org_id"`
    SiteId    *uuid.UUID         `json:"site_id"`
    Ssids     []string           `json:"ssids,omitempty"`
    Timestamp *int               `json:"timestamp"`
    Type      *string            `json:"type"`
    Update    *bool              `json:"update,omitempty"`
}

func (w *tempWebhookAlarmEvent) validate() error {
    var errs []string
    if w.Id == nil {
        errs = append(errs, "required field `id` is missing for type `webhook_alarm_event`")
    }
    if w.LastSeen == nil {
        errs = append(errs, "required field `last_seen` is missing for type `webhook_alarm_event`")
    }
    if w.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `webhook_alarm_event`")
    }
    if w.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `webhook_alarm_event`")
    }
    if w.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `webhook_alarm_event`")
    }
    if w.Type == nil {
        errs = append(errs, "required field `type` is missing for type `webhook_alarm_event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

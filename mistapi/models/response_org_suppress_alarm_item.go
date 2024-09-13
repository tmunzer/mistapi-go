package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ResponseOrgSuppressAlarmItem represents a ResponseOrgSuppressAlarmItem struct.
type ResponseOrgSuppressAlarmItem struct {
    // duration, in seconds. Maximum duration is 86400 * 14 (14 days). 0 is to un-suppress alarms.
    Duration             *int                      `json:"duration,omitempty"`
    ExpireTime           *int                      `json:"expire_time,omitempty"`
    ScheduledTime        *int                      `json:"scheduled_time,omitempty"`
    // level of scope. enum: `org`, `site`
    Scope                *SuppressedAlarmScopeEnum `json:"scope,omitempty"`
    SiteId               *uuid.UUID                `json:"site_id,omitempty"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseOrgSuppressAlarmItem.
// It customizes the JSON marshaling process for ResponseOrgSuppressAlarmItem objects.
func (r ResponseOrgSuppressAlarmItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseOrgSuppressAlarmItem object to a map representation for JSON marshaling.
func (r ResponseOrgSuppressAlarmItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Duration != nil {
        structMap["duration"] = r.Duration
    }
    if r.ExpireTime != nil {
        structMap["expire_time"] = r.ExpireTime
    }
    if r.ScheduledTime != nil {
        structMap["scheduled_time"] = r.ScheduledTime
    }
    if r.Scope != nil {
        structMap["scope"] = r.Scope
    }
    if r.SiteId != nil {
        structMap["site_id"] = r.SiteId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOrgSuppressAlarmItem.
// It customizes the JSON unmarshaling process for ResponseOrgSuppressAlarmItem objects.
func (r *ResponseOrgSuppressAlarmItem) UnmarshalJSON(input []byte) error {
    var temp tempResponseOrgSuppressAlarmItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "duration", "expire_time", "scheduled_time", "scope", "site_id")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Duration = temp.Duration
    r.ExpireTime = temp.ExpireTime
    r.ScheduledTime = temp.ScheduledTime
    r.Scope = temp.Scope
    r.SiteId = temp.SiteId
    return nil
}

// tempResponseOrgSuppressAlarmItem is a temporary struct used for validating the fields of ResponseOrgSuppressAlarmItem.
type tempResponseOrgSuppressAlarmItem  struct {
    Duration      *int                      `json:"duration,omitempty"`
    ExpireTime    *int                      `json:"expire_time,omitempty"`
    ScheduledTime *int                      `json:"scheduled_time,omitempty"`
    Scope         *SuppressedAlarmScopeEnum `json:"scope,omitempty"`
    SiteId        *uuid.UUID                `json:"site_id,omitempty"`
}

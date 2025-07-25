// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// ResponseOrgSuppressAlarmItem represents a ResponseOrgSuppressAlarmItem struct.
type ResponseOrgSuppressAlarmItem struct {
    // Duration, in seconds. Maximum duration is 86400 * 14 (14 days). 0 is to un-suppress alarms.
    Duration             *int                      `json:"duration,omitempty"`
    ExpireTime           *int                      `json:"expire_time,omitempty"`
    ScheduledTime        *int                      `json:"scheduled_time,omitempty"`
    // level of scope. enum: `org`, `site`
    Scope                *SuppressedAlarmScopeEnum `json:"scope,omitempty"`
    SiteId               *uuid.UUID                `json:"site_id,omitempty"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseOrgSuppressAlarmItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseOrgSuppressAlarmItem) String() string {
    return fmt.Sprintf(
    	"ResponseOrgSuppressAlarmItem[Duration=%v, ExpireTime=%v, ScheduledTime=%v, Scope=%v, SiteId=%v, AdditionalProperties=%v]",
    	r.Duration, r.ExpireTime, r.ScheduledTime, r.Scope, r.SiteId, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseOrgSuppressAlarmItem.
// It customizes the JSON marshaling process for ResponseOrgSuppressAlarmItem objects.
func (r ResponseOrgSuppressAlarmItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "duration", "expire_time", "scheduled_time", "scope", "site_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseOrgSuppressAlarmItem object to a map representation for JSON marshaling.
func (r ResponseOrgSuppressAlarmItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "duration", "expire_time", "scheduled_time", "scope", "site_id")
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

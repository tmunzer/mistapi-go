// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// WebhookDeviceEventsEvent represents a WebhookDeviceEventsEvent struct.
type WebhookDeviceEventsEvent struct {
    // (will be deprecated soon; please use mac instead) ap mac
    Ap                   *string                            `json:"ap,omitempty"`
    // (will be deprecated soon; please use device_name instead) ap name
    ApName               *string                            `json:"ap_name,omitempty"`
    // (optional) audit id
    AuditId              *uuid.UUID                         `json:"audit_id,omitempty"`
    // Device name
    DeviceName           string                             `json:"device_name"`
    // enum: `ap`, `gateway`, `switch`
    DeviceType           DeviceTypeEnum                     `json:"device_type"`
    // (optional) event advisory. enum: `notice`, `warn`
    EvType               WebhookDeviceEventsEventEvTypeEnum `json:"ev_type"`
    // Device mac
    Mac                  string                             `json:"mac"`
    OrgId                uuid.UUID                          `json:"org_id"`
    // (optional) event reason
    Reason               *string                            `json:"reason,omitempty"`
    SiteId               *uuid.UUID                         `json:"site_id,omitempty"`
    // Site name
    SiteName             *string                            `json:"site_name,omitempty"`
    // (optional) event description
    Text                 *string                            `json:"text,omitempty"`
    // Epoch (seconds)
    Timestamp            float64                            `json:"timestamp"`
    // Event type
    Type                 string                             `json:"type"`
    AdditionalProperties map[string]interface{}             `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookDeviceEventsEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookDeviceEventsEvent) String() string {
    return fmt.Sprintf(
    	"WebhookDeviceEventsEvent[Ap=%v, ApName=%v, AuditId=%v, DeviceName=%v, DeviceType=%v, EvType=%v, Mac=%v, OrgId=%v, Reason=%v, SiteId=%v, SiteName=%v, Text=%v, Timestamp=%v, Type=%v, AdditionalProperties=%v]",
    	w.Ap, w.ApName, w.AuditId, w.DeviceName, w.DeviceType, w.EvType, w.Mac, w.OrgId, w.Reason, w.SiteId, w.SiteName, w.Text, w.Timestamp, w.Type, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookDeviceEventsEvent.
// It customizes the JSON marshaling process for WebhookDeviceEventsEvent objects.
func (w WebhookDeviceEventsEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "ap", "ap_name", "audit_id", "device_name", "device_type", "ev_type", "mac", "org_id", "reason", "site_id", "site_name", "text", "timestamp", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookDeviceEventsEvent object to a map representation for JSON marshaling.
func (w WebhookDeviceEventsEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Ap != nil {
        structMap["ap"] = w.Ap
    }
    if w.ApName != nil {
        structMap["ap_name"] = w.ApName
    }
    if w.AuditId != nil {
        structMap["audit_id"] = w.AuditId
    }
    structMap["device_name"] = w.DeviceName
    structMap["device_type"] = w.DeviceType
    structMap["ev_type"] = w.EvType
    structMap["mac"] = w.Mac
    structMap["org_id"] = w.OrgId
    if w.Reason != nil {
        structMap["reason"] = w.Reason
    }
    if w.SiteId != nil {
        structMap["site_id"] = w.SiteId
    }
    if w.SiteName != nil {
        structMap["site_name"] = w.SiteName
    }
    if w.Text != nil {
        structMap["text"] = w.Text
    }
    structMap["timestamp"] = w.Timestamp
    structMap["type"] = w.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookDeviceEventsEvent.
// It customizes the JSON unmarshaling process for WebhookDeviceEventsEvent objects.
func (w *WebhookDeviceEventsEvent) UnmarshalJSON(input []byte) error {
    var temp tempWebhookDeviceEventsEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "ap_name", "audit_id", "device_name", "device_type", "ev_type", "mac", "org_id", "reason", "site_id", "site_name", "text", "timestamp", "type")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Ap = temp.Ap
    w.ApName = temp.ApName
    w.AuditId = temp.AuditId
    w.DeviceName = *temp.DeviceName
    w.DeviceType = *temp.DeviceType
    w.EvType = *temp.EvType
    w.Mac = *temp.Mac
    w.OrgId = *temp.OrgId
    w.Reason = temp.Reason
    w.SiteId = temp.SiteId
    w.SiteName = temp.SiteName
    w.Text = temp.Text
    w.Timestamp = *temp.Timestamp
    w.Type = *temp.Type
    return nil
}

// tempWebhookDeviceEventsEvent is a temporary struct used for validating the fields of WebhookDeviceEventsEvent.
type tempWebhookDeviceEventsEvent  struct {
    Ap         *string                             `json:"ap,omitempty"`
    ApName     *string                             `json:"ap_name,omitempty"`
    AuditId    *uuid.UUID                          `json:"audit_id,omitempty"`
    DeviceName *string                             `json:"device_name"`
    DeviceType *DeviceTypeEnum                     `json:"device_type"`
    EvType     *WebhookDeviceEventsEventEvTypeEnum `json:"ev_type"`
    Mac        *string                             `json:"mac"`
    OrgId      *uuid.UUID                          `json:"org_id"`
    Reason     *string                             `json:"reason,omitempty"`
    SiteId     *uuid.UUID                          `json:"site_id,omitempty"`
    SiteName   *string                             `json:"site_name,omitempty"`
    Text       *string                             `json:"text,omitempty"`
    Timestamp  *float64                            `json:"timestamp"`
    Type       *string                             `json:"type"`
}

func (w *tempWebhookDeviceEventsEvent) validate() error {
    var errs []string
    if w.DeviceName == nil {
        errs = append(errs, "required field `device_name` is missing for type `webhook_device_events_event`")
    }
    if w.DeviceType == nil {
        errs = append(errs, "required field `device_type` is missing for type `webhook_device_events_event`")
    }
    if w.EvType == nil {
        errs = append(errs, "required field `ev_type` is missing for type `webhook_device_events_event`")
    }
    if w.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `webhook_device_events_event`")
    }
    if w.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `webhook_device_events_event`")
    }
    if w.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `webhook_device_events_event`")
    }
    if w.Type == nil {
        errs = append(errs, "required field `type` is missing for type `webhook_device_events_event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

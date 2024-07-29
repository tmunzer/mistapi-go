package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// WebhookDeviceEventsEvent represents a WebhookDeviceEventsEvent struct.
type WebhookDeviceEventsEvent struct {
    // (will be deprecated soon; please use mac instead) ap mac
    Ap                   *string                                `json:"ap,omitempty"`
    // (will be deprecated soon; please use device_name instead) ap name
    ApName               *string                                `json:"ap_name,omitempty"`
    // (optional) audit id
    AuditId              *uuid.UUID                             `json:"audit_id,omitempty"`
    // device name
    DeviceName           string                                 `json:"device_name"`
    // enum: `ap`, `gateway`, `switch`
    DeviceType           WebhookDeviceEventsEventDeviceTypeEnum `json:"device_type"`
    // (optional) event advisory. enum: `notice`, `warn`
    EvType               WebhookDeviceEventsEventEvTypeEnum     `json:"ev_type"`
    // device mac
    Mac                  string                                 `json:"mac"`
    OrgId                uuid.UUID                              `json:"org_id"`
    // (optional) event reason
    Reason               *string                                `json:"reason,omitempty"`
    SiteId               *uuid.UUID                             `json:"site_id,omitempty"`
    // site name
    SiteName             *string                                `json:"site_name,omitempty"`
    // (optional) event description
    Text                 *string                                `json:"text,omitempty"`
    // time the event occurred e.g. 1565987313
    Timestamp            int                                    `json:"timestamp"`
    // event type
    Type                 string                                 `json:"type"`
    AdditionalProperties map[string]any                         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookDeviceEventsEvent.
// It customizes the JSON marshaling process for WebhookDeviceEventsEvent objects.
func (w WebhookDeviceEventsEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookDeviceEventsEvent object to a map representation for JSON marshaling.
func (w WebhookDeviceEventsEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
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
    var temp webhookDeviceEventsEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap", "ap_name", "audit_id", "device_name", "device_type", "ev_type", "mac", "org_id", "reason", "site_id", "site_name", "text", "timestamp", "type")
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

// webhookDeviceEventsEvent is a temporary struct used for validating the fields of WebhookDeviceEventsEvent.
type webhookDeviceEventsEvent  struct {
    Ap         *string                                 `json:"ap,omitempty"`
    ApName     *string                                 `json:"ap_name,omitempty"`
    AuditId    *uuid.UUID                              `json:"audit_id,omitempty"`
    DeviceName *string                                 `json:"device_name"`
    DeviceType *WebhookDeviceEventsEventDeviceTypeEnum `json:"device_type"`
    EvType     *WebhookDeviceEventsEventEvTypeEnum     `json:"ev_type"`
    Mac        *string                                 `json:"mac"`
    OrgId      *uuid.UUID                              `json:"org_id"`
    Reason     *string                                 `json:"reason,omitempty"`
    SiteId     *uuid.UUID                              `json:"site_id,omitempty"`
    SiteName   *string                                 `json:"site_name,omitempty"`
    Text       *string                                 `json:"text,omitempty"`
    Timestamp  *int                                    `json:"timestamp"`
    Type       *string                                 `json:"type"`
}

func (w *webhookDeviceEventsEvent) validate() error {
    var errs []string
    if w.DeviceName == nil {
        errs = append(errs, "required field `device_name` is missing for type `Webhook_Device_Events_Event`")
    }
    if w.DeviceType == nil {
        errs = append(errs, "required field `device_type` is missing for type `Webhook_Device_Events_Event`")
    }
    if w.EvType == nil {
        errs = append(errs, "required field `ev_type` is missing for type `Webhook_Device_Events_Event`")
    }
    if w.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Webhook_Device_Events_Event`")
    }
    if w.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `Webhook_Device_Events_Event`")
    }
    if w.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Webhook_Device_Events_Event`")
    }
    if w.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Webhook_Device_Events_Event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

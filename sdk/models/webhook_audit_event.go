package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// WebhookAuditEvent represents a WebhookAuditEvent struct.
type WebhookAuditEvent struct {
    AdminName            string         `json:"admin_name"`
    DeviceId             uuid.UUID      `json:"device_id"`
    Id                   uuid.UUID      `json:"id"`
    Message              string         `json:"message"`
    OrgId                uuid.UUID      `json:"org_id"`
    SiteId               uuid.UUID      `json:"site_id"`
    SrcIp                string         `json:"src_ip"`
    Timestamp            float64        `json:"timestamp"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookAuditEvent.
// It customizes the JSON marshaling process for WebhookAuditEvent objects.
func (w WebhookAuditEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookAuditEvent object to a map representation for JSON marshaling.
func (w WebhookAuditEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["admin_name"] = w.AdminName
    structMap["device_id"] = w.DeviceId
    structMap["id"] = w.Id
    structMap["message"] = w.Message
    structMap["org_id"] = w.OrgId
    structMap["site_id"] = w.SiteId
    structMap["src_ip"] = w.SrcIp
    structMap["timestamp"] = w.Timestamp
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookAuditEvent.
// It customizes the JSON unmarshaling process for WebhookAuditEvent objects.
func (w *WebhookAuditEvent) UnmarshalJSON(input []byte) error {
    var temp webhookAuditEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "admin_name", "device_id", "id", "message", "org_id", "site_id", "src_ip", "timestamp")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.AdminName = *temp.AdminName
    w.DeviceId = *temp.DeviceId
    w.Id = *temp.Id
    w.Message = *temp.Message
    w.OrgId = *temp.OrgId
    w.SiteId = *temp.SiteId
    w.SrcIp = *temp.SrcIp
    w.Timestamp = *temp.Timestamp
    return nil
}

// webhookAuditEvent is a temporary struct used for validating the fields of WebhookAuditEvent.
type webhookAuditEvent  struct {
    AdminName *string    `json:"admin_name"`
    DeviceId  *uuid.UUID `json:"device_id"`
    Id        *uuid.UUID `json:"id"`
    Message   *string    `json:"message"`
    OrgId     *uuid.UUID `json:"org_id"`
    SiteId    *uuid.UUID `json:"site_id"`
    SrcIp     *string    `json:"src_ip"`
    Timestamp *float64   `json:"timestamp"`
}

func (w *webhookAuditEvent) validate() error {
    var errs []string
    if w.AdminName == nil {
        errs = append(errs, "required field `admin_name` is missing for type `Webhook_Audit_Event`")
    }
    if w.DeviceId == nil {
        errs = append(errs, "required field `device_id` is missing for type `Webhook_Audit_Event`")
    }
    if w.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Webhook_Audit_Event`")
    }
    if w.Message == nil {
        errs = append(errs, "required field `message` is missing for type `Webhook_Audit_Event`")
    }
    if w.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `Webhook_Audit_Event`")
    }
    if w.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `Webhook_Audit_Event`")
    }
    if w.SrcIp == nil {
        errs = append(errs, "required field `src_ip` is missing for type `Webhook_Audit_Event`")
    }
    if w.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Webhook_Audit_Event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

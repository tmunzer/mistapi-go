package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// WebhookDeviceUpdownsEvent represents a WebhookDeviceUpdownsEvent struct.
type WebhookDeviceUpdownsEvent struct {
    Ap                   string         `json:"ap"`
    ApName               string         `json:"ap_name"`
    ForSite              *bool          `json:"for_site,omitempty"`
    OrgId                uuid.UUID      `json:"org_id"`
    SiteId               uuid.UUID      `json:"site_id"`
    SiteName             string         `json:"site_name"`
    Timestamp            float64        `json:"timestamp"`
    Type                 string         `json:"type"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookDeviceUpdownsEvent.
// It customizes the JSON marshaling process for WebhookDeviceUpdownsEvent objects.
func (w WebhookDeviceUpdownsEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookDeviceUpdownsEvent object to a map representation for JSON marshaling.
func (w WebhookDeviceUpdownsEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["ap"] = w.Ap
    structMap["ap_name"] = w.ApName
    if w.ForSite != nil {
        structMap["for_site"] = w.ForSite
    }
    structMap["org_id"] = w.OrgId
    structMap["site_id"] = w.SiteId
    structMap["site_name"] = w.SiteName
    structMap["timestamp"] = w.Timestamp
    structMap["type"] = w.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookDeviceUpdownsEvent.
// It customizes the JSON unmarshaling process for WebhookDeviceUpdownsEvent objects.
func (w *WebhookDeviceUpdownsEvent) UnmarshalJSON(input []byte) error {
    var temp tempWebhookDeviceUpdownsEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap", "ap_name", "for_site", "org_id", "site_id", "site_name", "timestamp", "type")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Ap = *temp.Ap
    w.ApName = *temp.ApName
    w.ForSite = temp.ForSite
    w.OrgId = *temp.OrgId
    w.SiteId = *temp.SiteId
    w.SiteName = *temp.SiteName
    w.Timestamp = *temp.Timestamp
    w.Type = *temp.Type
    return nil
}

// tempWebhookDeviceUpdownsEvent is a temporary struct used for validating the fields of WebhookDeviceUpdownsEvent.
type tempWebhookDeviceUpdownsEvent  struct {
    Ap        *string    `json:"ap"`
    ApName    *string    `json:"ap_name"`
    ForSite   *bool      `json:"for_site,omitempty"`
    OrgId     *uuid.UUID `json:"org_id"`
    SiteId    *uuid.UUID `json:"site_id"`
    SiteName  *string    `json:"site_name"`
    Timestamp *float64   `json:"timestamp"`
    Type      *string    `json:"type"`
}

func (w *tempWebhookDeviceUpdownsEvent) validate() error {
    var errs []string
    if w.Ap == nil {
        errs = append(errs, "required field `ap` is missing for type `webhook_device_updowns_event`")
    }
    if w.ApName == nil {
        errs = append(errs, "required field `ap_name` is missing for type `webhook_device_updowns_event`")
    }
    if w.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `webhook_device_updowns_event`")
    }
    if w.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `webhook_device_updowns_event`")
    }
    if w.SiteName == nil {
        errs = append(errs, "required field `site_name` is missing for type `webhook_device_updowns_event`")
    }
    if w.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `webhook_device_updowns_event`")
    }
    if w.Type == nil {
        errs = append(errs, "required field `type` is missing for type `webhook_device_updowns_event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

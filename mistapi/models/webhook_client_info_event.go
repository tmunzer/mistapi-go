// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// WebhookClientInfoEvent represents a WebhookClientInfoEvent struct.
type WebhookClientInfoEvent struct {
    // Hostname of client
    Hostname             *string                `json:"hostname,omitempty"`
    // IP address of client
    Ip                   *string                `json:"ip,omitempty"`
    // client's MAC Address
    Mac                  *string                `json:"mac,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    // Epoch (seconds)
    Timestamp            *float64               `json:"timestamp,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookClientInfoEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookClientInfoEvent) String() string {
    return fmt.Sprintf(
    	"WebhookClientInfoEvent[Hostname=%v, Ip=%v, Mac=%v, OrgId=%v, SiteId=%v, Timestamp=%v, AdditionalProperties=%v]",
    	w.Hostname, w.Ip, w.Mac, w.OrgId, w.SiteId, w.Timestamp, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookClientInfoEvent.
// It customizes the JSON marshaling process for WebhookClientInfoEvent objects.
func (w WebhookClientInfoEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "hostname", "ip", "mac", "org_id", "site_id", "timestamp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookClientInfoEvent object to a map representation for JSON marshaling.
func (w WebhookClientInfoEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Hostname != nil {
        structMap["hostname"] = w.Hostname
    }
    if w.Ip != nil {
        structMap["ip"] = w.Ip
    }
    if w.Mac != nil {
        structMap["mac"] = w.Mac
    }
    if w.OrgId != nil {
        structMap["org_id"] = w.OrgId
    }
    if w.SiteId != nil {
        structMap["site_id"] = w.SiteId
    }
    if w.Timestamp != nil {
        structMap["timestamp"] = w.Timestamp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookClientInfoEvent.
// It customizes the JSON unmarshaling process for WebhookClientInfoEvent objects.
func (w *WebhookClientInfoEvent) UnmarshalJSON(input []byte) error {
    var temp tempWebhookClientInfoEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "hostname", "ip", "mac", "org_id", "site_id", "timestamp")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Hostname = temp.Hostname
    w.Ip = temp.Ip
    w.Mac = temp.Mac
    w.OrgId = temp.OrgId
    w.SiteId = temp.SiteId
    w.Timestamp = temp.Timestamp
    return nil
}

// tempWebhookClientInfoEvent is a temporary struct used for validating the fields of WebhookClientInfoEvent.
type tempWebhookClientInfoEvent  struct {
    Hostname  *string    `json:"hostname,omitempty"`
    Ip        *string    `json:"ip,omitempty"`
    Mac       *string    `json:"mac,omitempty"`
    OrgId     *uuid.UUID `json:"org_id,omitempty"`
    SiteId    *uuid.UUID `json:"site_id,omitempty"`
    Timestamp *float64   `json:"timestamp,omitempty"`
}

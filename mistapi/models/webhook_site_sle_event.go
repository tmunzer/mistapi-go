// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// WebhookSiteSleEvent represents a WebhookSiteSleEvent struct.
type WebhookSiteSleEvent struct {
    OrgId                *uuid.UUID              `json:"org_id,omitempty"`
    SiteId               *uuid.UUID              `json:"site_id,omitempty"`
    Sle                  *WebhookSiteSleEventSle `json:"sle,omitempty"`
    // Epoch (seconds)
    Timestamp            *float64                `json:"timestamp,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookSiteSleEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookSiteSleEvent) String() string {
    return fmt.Sprintf(
    	"WebhookSiteSleEvent[OrgId=%v, SiteId=%v, Sle=%v, Timestamp=%v, AdditionalProperties=%v]",
    	w.OrgId, w.SiteId, w.Sle, w.Timestamp, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookSiteSleEvent.
// It customizes the JSON marshaling process for WebhookSiteSleEvent objects.
func (w WebhookSiteSleEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "org_id", "site_id", "sle", "timestamp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookSiteSleEvent object to a map representation for JSON marshaling.
func (w WebhookSiteSleEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.OrgId != nil {
        structMap["org_id"] = w.OrgId
    }
    if w.SiteId != nil {
        structMap["site_id"] = w.SiteId
    }
    if w.Sle != nil {
        structMap["sle"] = w.Sle.toMap()
    }
    if w.Timestamp != nil {
        structMap["timestamp"] = w.Timestamp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookSiteSleEvent.
// It customizes the JSON unmarshaling process for WebhookSiteSleEvent objects.
func (w *WebhookSiteSleEvent) UnmarshalJSON(input []byte) error {
    var temp tempWebhookSiteSleEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "org_id", "site_id", "sle", "timestamp")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.OrgId = temp.OrgId
    w.SiteId = temp.SiteId
    w.Sle = temp.Sle
    w.Timestamp = temp.Timestamp
    return nil
}

// tempWebhookSiteSleEvent is a temporary struct used for validating the fields of WebhookSiteSleEvent.
type tempWebhookSiteSleEvent  struct {
    OrgId     *uuid.UUID              `json:"org_id,omitempty"`
    SiteId    *uuid.UUID              `json:"site_id,omitempty"`
    Sle       *WebhookSiteSleEventSle `json:"sle,omitempty"`
    Timestamp *float64                `json:"timestamp,omitempty"`
}

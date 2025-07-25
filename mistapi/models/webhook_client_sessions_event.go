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

// WebhookClientSessionsEvent represents a WebhookClientSessionsEvent struct.
type WebhookClientSessionsEvent struct {
    // MAC address of the AP the client roamed or disconnected from
    Ap                   string                 `json:"ap"`
    // user-friendly name of the AP the client roamed or disconnected from.
    ApName               string                 `json:"ap_name"`
    // 5GHz or 2.4GHz band
    Band                 string                 `json:"band"`
    Bssid                string                 `json:"bssid"`
    // Device family E.g. "Mac", "iPhone", "Apple watch"
    ClientFamily         string                 `json:"client_family"`
    // Device manufacturer E.g. "Apple"
    ClientManufacture    string                 `json:"client_manufacture"`
    // Device model E.g. "8+", "XS"
    ClientModel          string                 `json:"client_model"`
    // Device operating system E.g. "Mojave", "Windows 10", "Linux"
    ClientOs             string                 `json:"client_os"`
    // Time when the user connects
    Connect              int                    `json:"connect"`
    // floating point connect timestamp with millisecond precision
    ConnectFloat         float64                `json:"connect_float"`
    // Time when the user disconnects
    Disconnect           int                    `json:"disconnect"`
    // floating point disconnect timestamp with millisecond precision
    DisconnectFloat      float64                `json:"disconnect_float"`
    // Duration of the roamed or complete session indicated by termination_reason field.
    Duration             int                    `json:"duration"`
    // Client's MAC Address'
    Mac                  string                 `json:"mac"`
    // the AP the client has roamed to.
    NextAp               string                 `json:"next_ap"`
    OrgId                uuid.UUID              `json:"org_id"`
    // Latest average RSSI before the user disconnects
    Rssi                 float64                `json:"rssi"`
    SiteId               uuid.UUID              `json:"site_id"`
    SiteName             string                 `json:"site_name"`
    Ssid                 string                 `json:"ssid"`
    // 1 disassociate - when the client disassociates. 2 inactive - when the client is timeout. 3 roamed - when the client is roamed between APs
    TerminationReason    int                    `json:"termination_reason"`
    // Epoch (seconds)
    Timestamp            float64                `json:"timestamp"`
    // schema version of this message
    Version              float64                `json:"version"`
    WlanId               uuid.UUID              `json:"wlan_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookClientSessionsEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookClientSessionsEvent) String() string {
    return fmt.Sprintf(
    	"WebhookClientSessionsEvent[Ap=%v, ApName=%v, Band=%v, Bssid=%v, ClientFamily=%v, ClientManufacture=%v, ClientModel=%v, ClientOs=%v, Connect=%v, ConnectFloat=%v, Disconnect=%v, DisconnectFloat=%v, Duration=%v, Mac=%v, NextAp=%v, OrgId=%v, Rssi=%v, SiteId=%v, SiteName=%v, Ssid=%v, TerminationReason=%v, Timestamp=%v, Version=%v, WlanId=%v, AdditionalProperties=%v]",
    	w.Ap, w.ApName, w.Band, w.Bssid, w.ClientFamily, w.ClientManufacture, w.ClientModel, w.ClientOs, w.Connect, w.ConnectFloat, w.Disconnect, w.DisconnectFloat, w.Duration, w.Mac, w.NextAp, w.OrgId, w.Rssi, w.SiteId, w.SiteName, w.Ssid, w.TerminationReason, w.Timestamp, w.Version, w.WlanId, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookClientSessionsEvent.
// It customizes the JSON marshaling process for WebhookClientSessionsEvent objects.
func (w WebhookClientSessionsEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "ap", "ap_name", "band", "bssid", "client_family", "client_manufacture", "client_model", "client_os", "connect", "connect_float", "disconnect", "disconnect_float", "duration", "mac", "next_ap", "org_id", "rssi", "site_id", "site_name", "ssid", "termination_reason", "timestamp", "version", "wlan_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookClientSessionsEvent object to a map representation for JSON marshaling.
func (w WebhookClientSessionsEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["ap"] = w.Ap
    structMap["ap_name"] = w.ApName
    structMap["band"] = w.Band
    structMap["bssid"] = w.Bssid
    structMap["client_family"] = w.ClientFamily
    structMap["client_manufacture"] = w.ClientManufacture
    structMap["client_model"] = w.ClientModel
    structMap["client_os"] = w.ClientOs
    structMap["connect"] = w.Connect
    structMap["connect_float"] = w.ConnectFloat
    structMap["disconnect"] = w.Disconnect
    structMap["disconnect_float"] = w.DisconnectFloat
    structMap["duration"] = w.Duration
    structMap["mac"] = w.Mac
    structMap["next_ap"] = w.NextAp
    structMap["org_id"] = w.OrgId
    structMap["rssi"] = w.Rssi
    structMap["site_id"] = w.SiteId
    structMap["site_name"] = w.SiteName
    structMap["ssid"] = w.Ssid
    structMap["termination_reason"] = w.TerminationReason
    structMap["timestamp"] = w.Timestamp
    structMap["version"] = w.Version
    structMap["wlan_id"] = w.WlanId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookClientSessionsEvent.
// It customizes the JSON unmarshaling process for WebhookClientSessionsEvent objects.
func (w *WebhookClientSessionsEvent) UnmarshalJSON(input []byte) error {
    var temp tempWebhookClientSessionsEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "ap_name", "band", "bssid", "client_family", "client_manufacture", "client_model", "client_os", "connect", "connect_float", "disconnect", "disconnect_float", "duration", "mac", "next_ap", "org_id", "rssi", "site_id", "site_name", "ssid", "termination_reason", "timestamp", "version", "wlan_id")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Ap = *temp.Ap
    w.ApName = *temp.ApName
    w.Band = *temp.Band
    w.Bssid = *temp.Bssid
    w.ClientFamily = *temp.ClientFamily
    w.ClientManufacture = *temp.ClientManufacture
    w.ClientModel = *temp.ClientModel
    w.ClientOs = *temp.ClientOs
    w.Connect = *temp.Connect
    w.ConnectFloat = *temp.ConnectFloat
    w.Disconnect = *temp.Disconnect
    w.DisconnectFloat = *temp.DisconnectFloat
    w.Duration = *temp.Duration
    w.Mac = *temp.Mac
    w.NextAp = *temp.NextAp
    w.OrgId = *temp.OrgId
    w.Rssi = *temp.Rssi
    w.SiteId = *temp.SiteId
    w.SiteName = *temp.SiteName
    w.Ssid = *temp.Ssid
    w.TerminationReason = *temp.TerminationReason
    w.Timestamp = *temp.Timestamp
    w.Version = *temp.Version
    w.WlanId = *temp.WlanId
    return nil
}

// tempWebhookClientSessionsEvent is a temporary struct used for validating the fields of WebhookClientSessionsEvent.
type tempWebhookClientSessionsEvent  struct {
    Ap                *string    `json:"ap"`
    ApName            *string    `json:"ap_name"`
    Band              *string    `json:"band"`
    Bssid             *string    `json:"bssid"`
    ClientFamily      *string    `json:"client_family"`
    ClientManufacture *string    `json:"client_manufacture"`
    ClientModel       *string    `json:"client_model"`
    ClientOs          *string    `json:"client_os"`
    Connect           *int       `json:"connect"`
    ConnectFloat      *float64   `json:"connect_float"`
    Disconnect        *int       `json:"disconnect"`
    DisconnectFloat   *float64   `json:"disconnect_float"`
    Duration          *int       `json:"duration"`
    Mac               *string    `json:"mac"`
    NextAp            *string    `json:"next_ap"`
    OrgId             *uuid.UUID `json:"org_id"`
    Rssi              *float64   `json:"rssi"`
    SiteId            *uuid.UUID `json:"site_id"`
    SiteName          *string    `json:"site_name"`
    Ssid              *string    `json:"ssid"`
    TerminationReason *int       `json:"termination_reason"`
    Timestamp         *float64   `json:"timestamp"`
    Version           *float64   `json:"version"`
    WlanId            *uuid.UUID `json:"wlan_id"`
}

func (w *tempWebhookClientSessionsEvent) validate() error {
    var errs []string
    if w.Ap == nil {
        errs = append(errs, "required field `ap` is missing for type `webhook_client_sessions_event`")
    }
    if w.ApName == nil {
        errs = append(errs, "required field `ap_name` is missing for type `webhook_client_sessions_event`")
    }
    if w.Band == nil {
        errs = append(errs, "required field `band` is missing for type `webhook_client_sessions_event`")
    }
    if w.Bssid == nil {
        errs = append(errs, "required field `bssid` is missing for type `webhook_client_sessions_event`")
    }
    if w.ClientFamily == nil {
        errs = append(errs, "required field `client_family` is missing for type `webhook_client_sessions_event`")
    }
    if w.ClientManufacture == nil {
        errs = append(errs, "required field `client_manufacture` is missing for type `webhook_client_sessions_event`")
    }
    if w.ClientModel == nil {
        errs = append(errs, "required field `client_model` is missing for type `webhook_client_sessions_event`")
    }
    if w.ClientOs == nil {
        errs = append(errs, "required field `client_os` is missing for type `webhook_client_sessions_event`")
    }
    if w.Connect == nil {
        errs = append(errs, "required field `connect` is missing for type `webhook_client_sessions_event`")
    }
    if w.ConnectFloat == nil {
        errs = append(errs, "required field `connect_float` is missing for type `webhook_client_sessions_event`")
    }
    if w.Disconnect == nil {
        errs = append(errs, "required field `disconnect` is missing for type `webhook_client_sessions_event`")
    }
    if w.DisconnectFloat == nil {
        errs = append(errs, "required field `disconnect_float` is missing for type `webhook_client_sessions_event`")
    }
    if w.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `webhook_client_sessions_event`")
    }
    if w.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `webhook_client_sessions_event`")
    }
    if w.NextAp == nil {
        errs = append(errs, "required field `next_ap` is missing for type `webhook_client_sessions_event`")
    }
    if w.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `webhook_client_sessions_event`")
    }
    if w.Rssi == nil {
        errs = append(errs, "required field `rssi` is missing for type `webhook_client_sessions_event`")
    }
    if w.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `webhook_client_sessions_event`")
    }
    if w.SiteName == nil {
        errs = append(errs, "required field `site_name` is missing for type `webhook_client_sessions_event`")
    }
    if w.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `webhook_client_sessions_event`")
    }
    if w.TerminationReason == nil {
        errs = append(errs, "required field `termination_reason` is missing for type `webhook_client_sessions_event`")
    }
    if w.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `webhook_client_sessions_event`")
    }
    if w.Version == nil {
        errs = append(errs, "required field `version` is missing for type `webhook_client_sessions_event`")
    }
    if w.WlanId == nil {
        errs = append(errs, "required field `wlan_id` is missing for type `webhook_client_sessions_event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

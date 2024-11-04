package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// WirelssClientSession represents a WirelssClientSession struct.
type WirelssClientSession struct {
    Ap                   string         `json:"ap"`
    Band                 string         `json:"band"`
    ClientManufacture    *string        `json:"client_manufacture"`
    Connect              int            `json:"connect"`
    Disconnect           int            `json:"disconnect"`
    Duration             float64        `json:"duration"`
    ForSite              *bool          `json:"for_site,omitempty"`
    Mac                  string         `json:"mac"`
    OrgId                uuid.UUID      `json:"org_id"`
    SiteId               uuid.UUID      `json:"site_id"`
    Ssid                 string         `json:"ssid"`
    Tags                 []string       `json:"tags,omitempty"`
    Timestamp            float64        `json:"timestamp"`
    WlanId               uuid.UUID      `json:"wlan_id"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WirelssClientSession.
// It customizes the JSON marshaling process for WirelssClientSession objects.
func (w WirelssClientSession) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WirelssClientSession object to a map representation for JSON marshaling.
func (w WirelssClientSession) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["ap"] = w.Ap
    structMap["band"] = w.Band
    if w.ClientManufacture != nil {
        structMap["client_manufacture"] = w.ClientManufacture
    } else {
        structMap["client_manufacture"] = nil
    }
    structMap["connect"] = w.Connect
    structMap["disconnect"] = w.Disconnect
    structMap["duration"] = w.Duration
    if w.ForSite != nil {
        structMap["for_site"] = w.ForSite
    }
    structMap["mac"] = w.Mac
    structMap["org_id"] = w.OrgId
    structMap["site_id"] = w.SiteId
    structMap["ssid"] = w.Ssid
    if w.Tags != nil {
        structMap["tags"] = w.Tags
    }
    structMap["timestamp"] = w.Timestamp
    structMap["wlan_id"] = w.WlanId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WirelssClientSession.
// It customizes the JSON unmarshaling process for WirelssClientSession objects.
func (w *WirelssClientSession) UnmarshalJSON(input []byte) error {
    var temp tempWirelssClientSession
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap", "band", "client_manufacture", "connect", "disconnect", "duration", "for_site", "mac", "org_id", "site_id", "ssid", "tags", "timestamp", "wlan_id")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Ap = *temp.Ap
    w.Band = *temp.Band
    w.ClientManufacture = temp.ClientManufacture
    w.Connect = *temp.Connect
    w.Disconnect = *temp.Disconnect
    w.Duration = *temp.Duration
    w.ForSite = temp.ForSite
    w.Mac = *temp.Mac
    w.OrgId = *temp.OrgId
    w.SiteId = *temp.SiteId
    w.Ssid = *temp.Ssid
    w.Tags = temp.Tags
    w.Timestamp = *temp.Timestamp
    w.WlanId = *temp.WlanId
    return nil
}

// tempWirelssClientSession is a temporary struct used for validating the fields of WirelssClientSession.
type tempWirelssClientSession  struct {
    Ap                *string    `json:"ap"`
    Band              *string    `json:"band"`
    ClientManufacture *string    `json:"client_manufacture"`
    Connect           *int       `json:"connect"`
    Disconnect        *int       `json:"disconnect"`
    Duration          *float64   `json:"duration"`
    ForSite           *bool      `json:"for_site,omitempty"`
    Mac               *string    `json:"mac"`
    OrgId             *uuid.UUID `json:"org_id"`
    SiteId            *uuid.UUID `json:"site_id"`
    Ssid              *string    `json:"ssid"`
    Tags              []string   `json:"tags,omitempty"`
    Timestamp         *float64   `json:"timestamp"`
    WlanId            *uuid.UUID `json:"wlan_id"`
}

func (w *tempWirelssClientSession) validate() error {
    var errs []string
    if w.Ap == nil {
        errs = append(errs, "required field `ap` is missing for type `wirelss_client_session`")
    }
    if w.Band == nil {
        errs = append(errs, "required field `band` is missing for type `wirelss_client_session`")
    }
    if w.Connect == nil {
        errs = append(errs, "required field `connect` is missing for type `wirelss_client_session`")
    }
    if w.Disconnect == nil {
        errs = append(errs, "required field `disconnect` is missing for type `wirelss_client_session`")
    }
    if w.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `wirelss_client_session`")
    }
    if w.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `wirelss_client_session`")
    }
    if w.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `wirelss_client_session`")
    }
    if w.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `wirelss_client_session`")
    }
    if w.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `wirelss_client_session`")
    }
    if w.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `wirelss_client_session`")
    }
    if w.WlanId == nil {
        errs = append(errs, "required field `wlan_id` is missing for type `wirelss_client_session`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

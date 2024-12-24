package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// WirelessClientSession represents a WirelessClientSession struct.
type WirelessClientSession struct {
    Ap                   string                 `json:"ap"`
    Band                 string                 `json:"band"`
    ClientManufacture    *string                `json:"client_manufacture"`
    Connect              int                    `json:"connect"`
    Disconnect           int                    `json:"disconnect"`
    Duration             float64                `json:"duration"`
    ForSite              *bool                  `json:"for_site,omitempty"`
    Mac                  string                 `json:"mac"`
    OrgId                uuid.UUID              `json:"org_id"`
    SiteId               uuid.UUID              `json:"site_id"`
    Ssid                 string                 `json:"ssid"`
    Tags                 []string               `json:"tags,omitempty"`
    Timestamp            float64                `json:"timestamp"`
    WlanId               uuid.UUID              `json:"wlan_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WirelessClientSession,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WirelessClientSession) String() string {
    return fmt.Sprintf(
    	"WirelessClientSession[Ap=%v, Band=%v, ClientManufacture=%v, Connect=%v, Disconnect=%v, Duration=%v, ForSite=%v, Mac=%v, OrgId=%v, SiteId=%v, Ssid=%v, Tags=%v, Timestamp=%v, WlanId=%v, AdditionalProperties=%v]",
    	w.Ap, w.Band, w.ClientManufacture, w.Connect, w.Disconnect, w.Duration, w.ForSite, w.Mac, w.OrgId, w.SiteId, w.Ssid, w.Tags, w.Timestamp, w.WlanId, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WirelessClientSession.
// It customizes the JSON marshaling process for WirelessClientSession objects.
func (w WirelessClientSession) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "ap", "band", "client_manufacture", "connect", "disconnect", "duration", "for_site", "mac", "org_id", "site_id", "ssid", "tags", "timestamp", "wlan_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WirelessClientSession object to a map representation for JSON marshaling.
func (w WirelessClientSession) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for WirelessClientSession.
// It customizes the JSON unmarshaling process for WirelessClientSession objects.
func (w *WirelessClientSession) UnmarshalJSON(input []byte) error {
    var temp tempWirelessClientSession
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "band", "client_manufacture", "connect", "disconnect", "duration", "for_site", "mac", "org_id", "site_id", "ssid", "tags", "timestamp", "wlan_id")
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

// tempWirelessClientSession is a temporary struct used for validating the fields of WirelessClientSession.
type tempWirelessClientSession  struct {
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

func (w *tempWirelessClientSession) validate() error {
    var errs []string
    if w.Ap == nil {
        errs = append(errs, "required field `ap` is missing for type `wireless_client_session`")
    }
    if w.Band == nil {
        errs = append(errs, "required field `band` is missing for type `wireless_client_session`")
    }
    if w.Connect == nil {
        errs = append(errs, "required field `connect` is missing for type `wireless_client_session`")
    }
    if w.Disconnect == nil {
        errs = append(errs, "required field `disconnect` is missing for type `wireless_client_session`")
    }
    if w.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `wireless_client_session`")
    }
    if w.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `wireless_client_session`")
    }
    if w.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `wireless_client_session`")
    }
    if w.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `wireless_client_session`")
    }
    if w.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `wireless_client_session`")
    }
    if w.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `wireless_client_session`")
    }
    if w.WlanId == nil {
        errs = append(errs, "required field `wlan_id` is missing for type `wireless_client_session`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

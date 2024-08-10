package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// ResponseClientSessionsSearchItem represents a ResponseClientSessionsSearchItem struct.
type ResponseClientSessionsSearchItem struct {
    Ap                   string         `json:"ap"`
    Band                 string         `json:"band"`
    ClientManufacture    string         `json:"client_manufacture"`
    Connect              float64        `json:"connect"`
    Disconnect           float64        `json:"disconnect"`
    Duration             float64        `json:"duration"`
    Mac                  string         `json:"mac"`
    OrgId                uuid.UUID      `json:"org_id"`
    SiteId               uuid.UUID      `json:"site_id"`
    Ssid                 string         `json:"ssid"`
    Tags                 []string       `json:"tags,omitempty"`
    Timestamp            float64        `json:"timestamp"`
    WlanId               uuid.UUID      `json:"wlan_id"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseClientSessionsSearchItem.
// It customizes the JSON marshaling process for ResponseClientSessionsSearchItem objects.
func (r ResponseClientSessionsSearchItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseClientSessionsSearchItem object to a map representation for JSON marshaling.
func (r ResponseClientSessionsSearchItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["ap"] = r.Ap
    structMap["band"] = r.Band
    structMap["client_manufacture"] = r.ClientManufacture
    structMap["connect"] = r.Connect
    structMap["disconnect"] = r.Disconnect
    structMap["duration"] = r.Duration
    structMap["mac"] = r.Mac
    structMap["org_id"] = r.OrgId
    structMap["site_id"] = r.SiteId
    structMap["ssid"] = r.Ssid
    if r.Tags != nil {
        structMap["tags"] = r.Tags
    }
    structMap["timestamp"] = r.Timestamp
    structMap["wlan_id"] = r.WlanId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseClientSessionsSearchItem.
// It customizes the JSON unmarshaling process for ResponseClientSessionsSearchItem objects.
func (r *ResponseClientSessionsSearchItem) UnmarshalJSON(input []byte) error {
    var temp tempResponseClientSessionsSearchItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap", "band", "client_manufacture", "connect", "disconnect", "duration", "mac", "org_id", "site_id", "ssid", "tags", "timestamp", "wlan_id")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Ap = *temp.Ap
    r.Band = *temp.Band
    r.ClientManufacture = *temp.ClientManufacture
    r.Connect = *temp.Connect
    r.Disconnect = *temp.Disconnect
    r.Duration = *temp.Duration
    r.Mac = *temp.Mac
    r.OrgId = *temp.OrgId
    r.SiteId = *temp.SiteId
    r.Ssid = *temp.Ssid
    r.Tags = temp.Tags
    r.Timestamp = *temp.Timestamp
    r.WlanId = *temp.WlanId
    return nil
}

// tempResponseClientSessionsSearchItem is a temporary struct used for validating the fields of ResponseClientSessionsSearchItem.
type tempResponseClientSessionsSearchItem  struct {
    Ap                *string    `json:"ap"`
    Band              *string    `json:"band"`
    ClientManufacture *string    `json:"client_manufacture"`
    Connect           *float64   `json:"connect"`
    Disconnect        *float64   `json:"disconnect"`
    Duration          *float64   `json:"duration"`
    Mac               *string    `json:"mac"`
    OrgId             *uuid.UUID `json:"org_id"`
    SiteId            *uuid.UUID `json:"site_id"`
    Ssid              *string    `json:"ssid"`
    Tags              []string   `json:"tags,omitempty"`
    Timestamp         *float64   `json:"timestamp"`
    WlanId            *uuid.UUID `json:"wlan_id"`
}

func (r *tempResponseClientSessionsSearchItem) validate() error {
    var errs []string
    if r.Ap == nil {
        errs = append(errs, "required field `ap` is missing for type `response_client_sessions_search_item`")
    }
    if r.Band == nil {
        errs = append(errs, "required field `band` is missing for type `response_client_sessions_search_item`")
    }
    if r.ClientManufacture == nil {
        errs = append(errs, "required field `client_manufacture` is missing for type `response_client_sessions_search_item`")
    }
    if r.Connect == nil {
        errs = append(errs, "required field `connect` is missing for type `response_client_sessions_search_item`")
    }
    if r.Disconnect == nil {
        errs = append(errs, "required field `disconnect` is missing for type `response_client_sessions_search_item`")
    }
    if r.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `response_client_sessions_search_item`")
    }
    if r.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `response_client_sessions_search_item`")
    }
    if r.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `response_client_sessions_search_item`")
    }
    if r.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `response_client_sessions_search_item`")
    }
    if r.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `response_client_sessions_search_item`")
    }
    if r.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `response_client_sessions_search_item`")
    }
    if r.WlanId == nil {
        errs = append(errs, "required field `wlan_id` is missing for type `response_client_sessions_search_item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

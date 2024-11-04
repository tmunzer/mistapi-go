package models

import (
    "encoding/json"
)

// ResponseStatsCallsSummary represents a ResponseStatsCallsSummary struct.
type ResponseStatsCallsSummary struct {
    BadMinutesClient     *float64       `json:"bad_minutes_client,omitempty"`
    BadMinutesSiteWan    *float64       `json:"bad_minutes_site_wan,omitempty"`
    BadMinutesWireless   *float64       `json:"bad_minutes_wireless,omitempty"`
    NumAps               *int           `json:"num_aps,omitempty"`
    NumUsers             *int           `json:"num_users,omitempty"`
    TotalMinutes         *float64       `json:"total_minutes,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseStatsCallsSummary.
// It customizes the JSON marshaling process for ResponseStatsCallsSummary objects.
func (r ResponseStatsCallsSummary) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseStatsCallsSummary object to a map representation for JSON marshaling.
func (r ResponseStatsCallsSummary) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.BadMinutesClient != nil {
        structMap["bad_minutes_client"] = r.BadMinutesClient
    }
    if r.BadMinutesSiteWan != nil {
        structMap["bad_minutes_site_wan"] = r.BadMinutesSiteWan
    }
    if r.BadMinutesWireless != nil {
        structMap["bad_minutes_wireless"] = r.BadMinutesWireless
    }
    if r.NumAps != nil {
        structMap["num_aps"] = r.NumAps
    }
    if r.NumUsers != nil {
        structMap["num_users"] = r.NumUsers
    }
    if r.TotalMinutes != nil {
        structMap["total_minutes"] = r.TotalMinutes
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseStatsCallsSummary.
// It customizes the JSON unmarshaling process for ResponseStatsCallsSummary objects.
func (r *ResponseStatsCallsSummary) UnmarshalJSON(input []byte) error {
    var temp tempResponseStatsCallsSummary
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "bad_minutes_client", "bad_minutes_site_wan", "bad_minutes_wireless", "num_aps", "num_users", "total_minutes")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.BadMinutesClient = temp.BadMinutesClient
    r.BadMinutesSiteWan = temp.BadMinutesSiteWan
    r.BadMinutesWireless = temp.BadMinutesWireless
    r.NumAps = temp.NumAps
    r.NumUsers = temp.NumUsers
    r.TotalMinutes = temp.TotalMinutes
    return nil
}

// tempResponseStatsCallsSummary is a temporary struct used for validating the fields of ResponseStatsCallsSummary.
type tempResponseStatsCallsSummary  struct {
    BadMinutesClient   *float64 `json:"bad_minutes_client,omitempty"`
    BadMinutesSiteWan  *float64 `json:"bad_minutes_site_wan,omitempty"`
    BadMinutesWireless *float64 `json:"bad_minutes_wireless,omitempty"`
    NumAps             *int     `json:"num_aps,omitempty"`
    NumUsers           *int     `json:"num_users,omitempty"`
    TotalMinutes       *float64 `json:"total_minutes,omitempty"`
}

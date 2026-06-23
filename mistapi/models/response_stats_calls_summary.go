// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseStatsCallsSummary represents a ResponseStatsCallsSummary struct.
// Aggregated site call statistics summary
type ResponseStatsCallsSummary struct {
	// Total call minutes classified as bad across all quality categories
	BadMinutes *float64 `json:"bad_minutes,omitempty"`
	// Call minutes classified as bad due to client-side issues
	BadMinutesClient *float64 `json:"bad_minutes_client,omitempty"`
	// Call minutes classified as bad due to site WAN issues
	BadMinutesSiteWan *float64 `json:"bad_minutes_site_wan,omitempty"`
	// Call minutes classified as bad due to wireless issues
	BadMinutesWireless *float64 `json:"bad_minutes_wireless,omitempty"`
	// Number of APs represented in the call statistics summary
	NumAps *int `json:"num_aps,omitempty"`
	// Number of users represented in the call statistics summary
	NumUsers *int `json:"num_users,omitempty"`
	// Total call minutes represented in the summary
	TotalMinutes         *float64               `json:"total_minutes,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseStatsCallsSummary,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseStatsCallsSummary) String() string {
	return fmt.Sprintf(
		"ResponseStatsCallsSummary[BadMinutes=%v, BadMinutesClient=%v, BadMinutesSiteWan=%v, BadMinutesWireless=%v, NumAps=%v, NumUsers=%v, TotalMinutes=%v, AdditionalProperties=%v]",
		r.BadMinutes, r.BadMinutesClient, r.BadMinutesSiteWan, r.BadMinutesWireless, r.NumAps, r.NumUsers, r.TotalMinutes, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseStatsCallsSummary.
// It customizes the JSON marshaling process for ResponseStatsCallsSummary objects.
func (r ResponseStatsCallsSummary) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"bad_minutes", "bad_minutes_client", "bad_minutes_site_wan", "bad_minutes_wireless", "num_aps", "num_users", "total_minutes"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseStatsCallsSummary object to a map representation for JSON marshaling.
func (r ResponseStatsCallsSummary) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.BadMinutes != nil {
		structMap["bad_minutes"] = r.BadMinutes
	}
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bad_minutes", "bad_minutes_client", "bad_minutes_site_wan", "bad_minutes_wireless", "num_aps", "num_users", "total_minutes")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.BadMinutes = temp.BadMinutes
	r.BadMinutesClient = temp.BadMinutesClient
	r.BadMinutesSiteWan = temp.BadMinutesSiteWan
	r.BadMinutesWireless = temp.BadMinutesWireless
	r.NumAps = temp.NumAps
	r.NumUsers = temp.NumUsers
	r.TotalMinutes = temp.TotalMinutes
	return nil
}

// tempResponseStatsCallsSummary is a temporary struct used for validating the fields of ResponseStatsCallsSummary.
type tempResponseStatsCallsSummary struct {
	BadMinutes         *float64 `json:"bad_minutes,omitempty"`
	BadMinutesClient   *float64 `json:"bad_minutes_client,omitempty"`
	BadMinutesSiteWan  *float64 `json:"bad_minutes_site_wan,omitempty"`
	BadMinutesWireless *float64 `json:"bad_minutes_wireless,omitempty"`
	NumAps             *int     `json:"num_aps,omitempty"`
	NumUsers           *int     `json:"num_users,omitempty"`
	TotalMinutes       *float64 `json:"total_minutes,omitempty"`
}

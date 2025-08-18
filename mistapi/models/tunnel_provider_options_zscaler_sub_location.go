// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// TunnelProviderOptionsZscalerSubLocation represents a TunnelProviderOptionsZscalerSubLocation struct.
type TunnelProviderOptionsZscalerSubLocation struct {
	AupBlockInternetUntilAccepted *bool `json:"aup_block_internet_until_accepted,omitempty"`
	// Can only be `true` when `auth_required`==`false`, display Acceptable Use Policy (AUP)
	AupEnabled *bool `json:"aup_enabled,omitempty"`
	// Proxy HTTPs traffic, requiring Zscaler cert to be installed in browser
	AupForceSslInspection *bool `json:"aup_force_ssl_inspection,omitempty"`
	// Required if `aup_enabled`==`true`. Days before AUP is requested again
	AupTimeoutInDays *int `json:"aup_timeout_in_days,omitempty"`
	// Enable this option to authenticate users
	AuthRequired *bool `json:"auth_required,omitempty"`
	// Can only be `true` when `auth_required`==`false`, display caution notification for non-authenticated users
	CautionEnabled *bool `json:"caution_enabled,omitempty"`
	// Download bandwidth cap of the link, in Mbps. Disabled if not set
	DnBandwidth Optional[float64] `json:"dn_bandwidth"`
	// Required if `surrogate_IP`==`true`, idle Time to Disassociation
	IdleTimeInMinutes *int `json:"idle_time_in_minutes,omitempty"`
	// [network]($h/Orgs%20Networks/_overview) name
	Name *string `json:"name,omitempty"`
	// If `true`, enable the firewall control option
	OfwEnabled *bool `json:"ofw_enabled,omitempty"`
	// Can only be `true` when `auth_required`==`true`. Map a user to a private IP address so it applies the user's policies, instead of the location's policies
	SurrogateIP *bool `json:"surrogate_IP,omitempty"`
	// Can only be `true` when `surrogate_IP`==`true`, enforce surrogate IP for known browsers
	SurrogateIPEnforcedForKnownBrowsers *bool `json:"surrogate_IP_enforced_for_known_browsers,omitempty"`
	// Required if `surrogate_IP_enforced_for_known_browsers`==`true`, must be lower or equal than `idle_time_in_minutes`, refresh Time for re-validation of Surrogacy
	SurrogateRefreshTimeInMinutes *int `json:"surrogate_refresh_time_in_minutes,omitempty"`
	// Download bandwidth cap of the link, in Mbps. Disabled if not set
	UpBandwidth          Optional[float64]      `json:"up_bandwidth"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TunnelProviderOptionsZscalerSubLocation,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TunnelProviderOptionsZscalerSubLocation) String() string {
	return fmt.Sprintf(
		"TunnelProviderOptionsZscalerSubLocation[AupBlockInternetUntilAccepted=%v, AupEnabled=%v, AupForceSslInspection=%v, AupTimeoutInDays=%v, AuthRequired=%v, CautionEnabled=%v, DnBandwidth=%v, IdleTimeInMinutes=%v, Name=%v, OfwEnabled=%v, SurrogateIP=%v, SurrogateIPEnforcedForKnownBrowsers=%v, SurrogateRefreshTimeInMinutes=%v, UpBandwidth=%v, AdditionalProperties=%v]",
		t.AupBlockInternetUntilAccepted, t.AupEnabled, t.AupForceSslInspection, t.AupTimeoutInDays, t.AuthRequired, t.CautionEnabled, t.DnBandwidth, t.IdleTimeInMinutes, t.Name, t.OfwEnabled, t.SurrogateIP, t.SurrogateIPEnforcedForKnownBrowsers, t.SurrogateRefreshTimeInMinutes, t.UpBandwidth, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TunnelProviderOptionsZscalerSubLocation.
// It customizes the JSON marshaling process for TunnelProviderOptionsZscalerSubLocation objects.
func (t TunnelProviderOptionsZscalerSubLocation) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(t.AdditionalProperties,
		"aup_block_internet_until_accepted", "aup_enabled", "aup_force_ssl_inspection", "aup_timeout_in_days", "auth_required", "caution_enabled", "dn_bandwidth", "idle_time_in_minutes", "name", "ofw_enabled", "surrogate_IP", "surrogate_IP_enforced_for_known_browsers", "surrogate_refresh_time_in_minutes", "up_bandwidth"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(t.toMap())
}

// toMap converts the TunnelProviderOptionsZscalerSubLocation object to a map representation for JSON marshaling.
func (t TunnelProviderOptionsZscalerSubLocation) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, t.AdditionalProperties)
	if t.AupBlockInternetUntilAccepted != nil {
		structMap["aup_block_internet_until_accepted"] = t.AupBlockInternetUntilAccepted
	}
	if t.AupEnabled != nil {
		structMap["aup_enabled"] = t.AupEnabled
	}
	if t.AupForceSslInspection != nil {
		structMap["aup_force_ssl_inspection"] = t.AupForceSslInspection
	}
	if t.AupTimeoutInDays != nil {
		structMap["aup_timeout_in_days"] = t.AupTimeoutInDays
	}
	if t.AuthRequired != nil {
		structMap["auth_required"] = t.AuthRequired
	}
	if t.CautionEnabled != nil {
		structMap["caution_enabled"] = t.CautionEnabled
	}
	if t.DnBandwidth.IsValueSet() {
		if t.DnBandwidth.Value() != nil {
			structMap["dn_bandwidth"] = t.DnBandwidth.Value()
		} else {
			structMap["dn_bandwidth"] = nil
		}
	}
	if t.IdleTimeInMinutes != nil {
		structMap["idle_time_in_minutes"] = t.IdleTimeInMinutes
	}
	if t.Name != nil {
		structMap["name"] = t.Name
	}
	if t.OfwEnabled != nil {
		structMap["ofw_enabled"] = t.OfwEnabled
	}
	if t.SurrogateIP != nil {
		structMap["surrogate_IP"] = t.SurrogateIP
	}
	if t.SurrogateIPEnforcedForKnownBrowsers != nil {
		structMap["surrogate_IP_enforced_for_known_browsers"] = t.SurrogateIPEnforcedForKnownBrowsers
	}
	if t.SurrogateRefreshTimeInMinutes != nil {
		structMap["surrogate_refresh_time_in_minutes"] = t.SurrogateRefreshTimeInMinutes
	}
	if t.UpBandwidth.IsValueSet() {
		if t.UpBandwidth.Value() != nil {
			structMap["up_bandwidth"] = t.UpBandwidth.Value()
		} else {
			structMap["up_bandwidth"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TunnelProviderOptionsZscalerSubLocation.
// It customizes the JSON unmarshaling process for TunnelProviderOptionsZscalerSubLocation objects.
func (t *TunnelProviderOptionsZscalerSubLocation) UnmarshalJSON(input []byte) error {
	var temp tempTunnelProviderOptionsZscalerSubLocation
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "aup_block_internet_until_accepted", "aup_enabled", "aup_force_ssl_inspection", "aup_timeout_in_days", "auth_required", "caution_enabled", "dn_bandwidth", "idle_time_in_minutes", "name", "ofw_enabled", "surrogate_IP", "surrogate_IP_enforced_for_known_browsers", "surrogate_refresh_time_in_minutes", "up_bandwidth")
	if err != nil {
		return err
	}
	t.AdditionalProperties = additionalProperties

	t.AupBlockInternetUntilAccepted = temp.AupBlockInternetUntilAccepted
	t.AupEnabled = temp.AupEnabled
	t.AupForceSslInspection = temp.AupForceSslInspection
	t.AupTimeoutInDays = temp.AupTimeoutInDays
	t.AuthRequired = temp.AuthRequired
	t.CautionEnabled = temp.CautionEnabled
	t.DnBandwidth = temp.DnBandwidth
	t.IdleTimeInMinutes = temp.IdleTimeInMinutes
	t.Name = temp.Name
	t.OfwEnabled = temp.OfwEnabled
	t.SurrogateIP = temp.SurrogateIP
	t.SurrogateIPEnforcedForKnownBrowsers = temp.SurrogateIPEnforcedForKnownBrowsers
	t.SurrogateRefreshTimeInMinutes = temp.SurrogateRefreshTimeInMinutes
	t.UpBandwidth = temp.UpBandwidth
	return nil
}

// tempTunnelProviderOptionsZscalerSubLocation is a temporary struct used for validating the fields of TunnelProviderOptionsZscalerSubLocation.
type tempTunnelProviderOptionsZscalerSubLocation struct {
	AupBlockInternetUntilAccepted       *bool             `json:"aup_block_internet_until_accepted,omitempty"`
	AupEnabled                          *bool             `json:"aup_enabled,omitempty"`
	AupForceSslInspection               *bool             `json:"aup_force_ssl_inspection,omitempty"`
	AupTimeoutInDays                    *int              `json:"aup_timeout_in_days,omitempty"`
	AuthRequired                        *bool             `json:"auth_required,omitempty"`
	CautionEnabled                      *bool             `json:"caution_enabled,omitempty"`
	DnBandwidth                         Optional[float64] `json:"dn_bandwidth"`
	IdleTimeInMinutes                   *int              `json:"idle_time_in_minutes,omitempty"`
	Name                                *string           `json:"name,omitempty"`
	OfwEnabled                          *bool             `json:"ofw_enabled,omitempty"`
	SurrogateIP                         *bool             `json:"surrogate_IP,omitempty"`
	SurrogateIPEnforcedForKnownBrowsers *bool             `json:"surrogate_IP_enforced_for_known_browsers,omitempty"`
	SurrogateRefreshTimeInMinutes       *int              `json:"surrogate_refresh_time_in_minutes,omitempty"`
	UpBandwidth                         Optional[float64] `json:"up_bandwidth"`
}

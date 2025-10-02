// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SleImpactedUsersUser represents a SleImpactedUsersUser struct.
type SleImpactedUsersUser struct {
	ApMac                *string                `json:"ap_mac,omitempty"`
	ApName               *string                `json:"ap_name,omitempty"`
	Degraded             *float64               `json:"degraded,omitempty"`
	DeviceOs             *string                `json:"device_os,omitempty"`
	DeviceType           *string                `json:"device_type,omitempty"`
	Duration             *float64               `json:"duration,omitempty"`
	Mac                  *string                `json:"mac,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	Ssid                 *string                `json:"ssid,omitempty"`
	Total                *float64               `json:"total,omitempty"`
	WlanId               *string                `json:"wlan_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactedUsersUser,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactedUsersUser) String() string {
	return fmt.Sprintf(
		"SleImpactedUsersUser[ApMac=%v, ApName=%v, Degraded=%v, DeviceOs=%v, DeviceType=%v, Duration=%v, Mac=%v, Name=%v, Ssid=%v, Total=%v, WlanId=%v, AdditionalProperties=%v]",
		s.ApMac, s.ApName, s.Degraded, s.DeviceOs, s.DeviceType, s.Duration, s.Mac, s.Name, s.Ssid, s.Total, s.WlanId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedUsersUser.
// It customizes the JSON marshaling process for SleImpactedUsersUser objects.
func (s SleImpactedUsersUser) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"ap_mac", "ap_name", "degraded", "device_os", "device_type", "duration", "mac", "name", "ssid", "total", "wlan_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedUsersUser object to a map representation for JSON marshaling.
func (s SleImpactedUsersUser) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.ApMac != nil {
		structMap["ap_mac"] = s.ApMac
	}
	if s.ApName != nil {
		structMap["ap_name"] = s.ApName
	}
	if s.Degraded != nil {
		structMap["degraded"] = s.Degraded
	}
	if s.DeviceOs != nil {
		structMap["device_os"] = s.DeviceOs
	}
	if s.DeviceType != nil {
		structMap["device_type"] = s.DeviceType
	}
	if s.Duration != nil {
		structMap["duration"] = s.Duration
	}
	if s.Mac != nil {
		structMap["mac"] = s.Mac
	}
	if s.Name != nil {
		structMap["name"] = s.Name
	}
	if s.Ssid != nil {
		structMap["ssid"] = s.Ssid
	}
	if s.Total != nil {
		structMap["total"] = s.Total
	}
	if s.WlanId != nil {
		structMap["wlan_id"] = s.WlanId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedUsersUser.
// It customizes the JSON unmarshaling process for SleImpactedUsersUser objects.
func (s *SleImpactedUsersUser) UnmarshalJSON(input []byte) error {
	var temp tempSleImpactedUsersUser
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_mac", "ap_name", "degraded", "device_os", "device_type", "duration", "mac", "name", "ssid", "total", "wlan_id")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.ApMac = temp.ApMac
	s.ApName = temp.ApName
	s.Degraded = temp.Degraded
	s.DeviceOs = temp.DeviceOs
	s.DeviceType = temp.DeviceType
	s.Duration = temp.Duration
	s.Mac = temp.Mac
	s.Name = temp.Name
	s.Ssid = temp.Ssid
	s.Total = temp.Total
	s.WlanId = temp.WlanId
	return nil
}

// tempSleImpactedUsersUser is a temporary struct used for validating the fields of SleImpactedUsersUser.
type tempSleImpactedUsersUser struct {
	ApMac      *string  `json:"ap_mac,omitempty"`
	ApName     *string  `json:"ap_name,omitempty"`
	Degraded   *float64 `json:"degraded,omitempty"`
	DeviceOs   *string  `json:"device_os,omitempty"`
	DeviceType *string  `json:"device_type,omitempty"`
	Duration   *float64 `json:"duration,omitempty"`
	Mac        *string  `json:"mac,omitempty"`
	Name       *string  `json:"name,omitempty"`
	Ssid       *string  `json:"ssid,omitempty"`
	Total      *float64 `json:"total,omitempty"`
	WlanId     *string  `json:"wlan_id,omitempty"`
}

// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SleImpactedUsersUser represents a SleImpactedUsersUser struct.
type SleImpactedUsersUser struct {
    ApMac                string                 `json:"ap_mac"`
    ApName               string                 `json:"ap_name"`
    Degraded             float64                `json:"degraded"`
    DeviceOs             string                 `json:"device_os"`
    DeviceType           string                 `json:"device_type"`
    Duration             float64                `json:"duration"`
    Mac                  string                 `json:"mac"`
    Name                 string                 `json:"name"`
    Ssid                 string                 `json:"ssid"`
    Total                float64                `json:"total"`
    WlanId               string                 `json:"wlan_id"`
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
    structMap["ap_mac"] = s.ApMac
    structMap["ap_name"] = s.ApName
    structMap["degraded"] = s.Degraded
    structMap["device_os"] = s.DeviceOs
    structMap["device_type"] = s.DeviceType
    structMap["duration"] = s.Duration
    structMap["mac"] = s.Mac
    structMap["name"] = s.Name
    structMap["ssid"] = s.Ssid
    structMap["total"] = s.Total
    structMap["wlan_id"] = s.WlanId
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
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_mac", "ap_name", "degraded", "device_os", "device_type", "duration", "mac", "name", "ssid", "total", "wlan_id")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ApMac = *temp.ApMac
    s.ApName = *temp.ApName
    s.Degraded = *temp.Degraded
    s.DeviceOs = *temp.DeviceOs
    s.DeviceType = *temp.DeviceType
    s.Duration = *temp.Duration
    s.Mac = *temp.Mac
    s.Name = *temp.Name
    s.Ssid = *temp.Ssid
    s.Total = *temp.Total
    s.WlanId = *temp.WlanId
    return nil
}

// tempSleImpactedUsersUser is a temporary struct used for validating the fields of SleImpactedUsersUser.
type tempSleImpactedUsersUser  struct {
    ApMac      *string  `json:"ap_mac"`
    ApName     *string  `json:"ap_name"`
    Degraded   *float64 `json:"degraded"`
    DeviceOs   *string  `json:"device_os"`
    DeviceType *string  `json:"device_type"`
    Duration   *float64 `json:"duration"`
    Mac        *string  `json:"mac"`
    Name       *string  `json:"name"`
    Ssid       *string  `json:"ssid"`
    Total      *float64 `json:"total"`
    WlanId     *string  `json:"wlan_id"`
}

func (s *tempSleImpactedUsersUser) validate() error {
    var errs []string
    if s.ApMac == nil {
        errs = append(errs, "required field `ap_mac` is missing for type `sle_impacted_users_user`")
    }
    if s.ApName == nil {
        errs = append(errs, "required field `ap_name` is missing for type `sle_impacted_users_user`")
    }
    if s.Degraded == nil {
        errs = append(errs, "required field `degraded` is missing for type `sle_impacted_users_user`")
    }
    if s.DeviceOs == nil {
        errs = append(errs, "required field `device_os` is missing for type `sle_impacted_users_user`")
    }
    if s.DeviceType == nil {
        errs = append(errs, "required field `device_type` is missing for type `sle_impacted_users_user`")
    }
    if s.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `sle_impacted_users_user`")
    }
    if s.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `sle_impacted_users_user`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `sle_impacted_users_user`")
    }
    if s.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `sle_impacted_users_user`")
    }
    if s.Total == nil {
        errs = append(errs, "required field `total` is missing for type `sle_impacted_users_user`")
    }
    if s.WlanId == nil {
        errs = append(errs, "required field `wlan_id` is missing for type `sle_impacted_users_user`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

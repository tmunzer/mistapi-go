package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SleImpactedUsersUser represents a SleImpactedUsersUser struct.
type SleImpactedUsersUser struct {
    ApMac                string         `json:"ap_mac"`
    ApName               string         `json:"ap_name"`
    Degraded             float64        `json:"degraded"`
    DeviceOs             string         `json:"device_os"`
    DeviceType           string         `json:"device_type"`
    Duration             float64        `json:"duration"`
    Mac                  string         `json:"mac"`
    Name                 string         `json:"name"`
    Ssid                 string         `json:"ssid"`
    Total                float64        `json:"total"`
    WlanId               string         `json:"wlan_id"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedUsersUser.
// It customizes the JSON marshaling process for SleImpactedUsersUser objects.
func (s SleImpactedUsersUser) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedUsersUser object to a map representation for JSON marshaling.
func (s SleImpactedUsersUser) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp sleImpactedUsersUser
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_mac", "ap_name", "degraded", "device_os", "device_type", "duration", "mac", "name", "ssid", "total", "wlan_id")
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

// sleImpactedUsersUser is a temporary struct used for validating the fields of SleImpactedUsersUser.
type sleImpactedUsersUser  struct {
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

func (s *sleImpactedUsersUser) validate() error {
    var errs []string
    if s.ApMac == nil {
        errs = append(errs, "required field `ap_mac` is missing for type `Sle_Impacted_Users_User`")
    }
    if s.ApName == nil {
        errs = append(errs, "required field `ap_name` is missing for type `Sle_Impacted_Users_User`")
    }
    if s.Degraded == nil {
        errs = append(errs, "required field `degraded` is missing for type `Sle_Impacted_Users_User`")
    }
    if s.DeviceOs == nil {
        errs = append(errs, "required field `device_os` is missing for type `Sle_Impacted_Users_User`")
    }
    if s.DeviceType == nil {
        errs = append(errs, "required field `device_type` is missing for type `Sle_Impacted_Users_User`")
    }
    if s.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `Sle_Impacted_Users_User`")
    }
    if s.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Sle_Impacted_Users_User`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Sle_Impacted_Users_User`")
    }
    if s.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `Sle_Impacted_Users_User`")
    }
    if s.Total == nil {
        errs = append(errs, "required field `total` is missing for type `Sle_Impacted_Users_User`")
    }
    if s.WlanId == nil {
        errs = append(errs, "required field `wlan_id` is missing for type `Sle_Impacted_Users_User`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

package models

import (
    "encoding/json"
)

// SiteSettingGatewayMgmtAutoSignatureUpdate represents a SiteSettingGatewayMgmtAutoSignatureUpdate struct.
type SiteSettingGatewayMgmtAutoSignatureUpdate struct {
    DayOfWeek            *DayOfWeekEnum `json:"day_of_week,omitempty"`
    Enable               *bool          `json:"enable,omitempty"`
    // optional, Mist will decide the timing
    TimeOfDay            *string        `json:"time_of_day,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingGatewayMgmtAutoSignatureUpdate.
// It customizes the JSON marshaling process for SiteSettingGatewayMgmtAutoSignatureUpdate objects.
func (s SiteSettingGatewayMgmtAutoSignatureUpdate) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingGatewayMgmtAutoSignatureUpdate object to a map representation for JSON marshaling.
func (s SiteSettingGatewayMgmtAutoSignatureUpdate) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.DayOfWeek != nil {
        structMap["day_of_week"] = s.DayOfWeek
    }
    if s.Enable != nil {
        structMap["enable"] = s.Enable
    }
    if s.TimeOfDay != nil {
        structMap["time_of_day"] = s.TimeOfDay
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingGatewayMgmtAutoSignatureUpdate.
// It customizes the JSON unmarshaling process for SiteSettingGatewayMgmtAutoSignatureUpdate objects.
func (s *SiteSettingGatewayMgmtAutoSignatureUpdate) UnmarshalJSON(input []byte) error {
    var temp siteSettingGatewayMgmtAutoSignatureUpdate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "day_of_week", "enable", "time_of_day")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.DayOfWeek = temp.DayOfWeek
    s.Enable = temp.Enable
    s.TimeOfDay = temp.TimeOfDay
    return nil
}

// siteSettingGatewayMgmtAutoSignatureUpdate is a temporary struct used for validating the fields of SiteSettingGatewayMgmtAutoSignatureUpdate.
type siteSettingGatewayMgmtAutoSignatureUpdate  struct {
    DayOfWeek *DayOfWeekEnum `json:"day_of_week,omitempty"`
    Enable    *bool          `json:"enable,omitempty"`
    TimeOfDay *string        `json:"time_of_day,omitempty"`
}

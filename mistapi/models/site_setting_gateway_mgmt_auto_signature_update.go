package models

import (
    "encoding/json"
    "fmt"
)

// SiteSettingGatewayMgmtAutoSignatureUpdate represents a SiteSettingGatewayMgmtAutoSignatureUpdate struct.
type SiteSettingGatewayMgmtAutoSignatureUpdate struct {
    // enum: `any`, `fri`, `mon`, `sat`, `sun`, `thu`, `tue`, `wed`
    DayOfWeek            *DayOfWeekEnum         `json:"day_of_week,omitempty"`
    Enable               *bool                  `json:"enable,omitempty"`
    // Optional, Mist will decide the timing
    TimeOfDay            *string                `json:"time_of_day,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingGatewayMgmtAutoSignatureUpdate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingGatewayMgmtAutoSignatureUpdate) String() string {
    return fmt.Sprintf(
    	"SiteSettingGatewayMgmtAutoSignatureUpdate[DayOfWeek=%v, Enable=%v, TimeOfDay=%v, AdditionalProperties=%v]",
    	s.DayOfWeek, s.Enable, s.TimeOfDay, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingGatewayMgmtAutoSignatureUpdate.
// It customizes the JSON marshaling process for SiteSettingGatewayMgmtAutoSignatureUpdate objects.
func (s SiteSettingGatewayMgmtAutoSignatureUpdate) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "day_of_week", "enable", "time_of_day"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingGatewayMgmtAutoSignatureUpdate object to a map representation for JSON marshaling.
func (s SiteSettingGatewayMgmtAutoSignatureUpdate) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSiteSettingGatewayMgmtAutoSignatureUpdate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "day_of_week", "enable", "time_of_day")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.DayOfWeek = temp.DayOfWeek
    s.Enable = temp.Enable
    s.TimeOfDay = temp.TimeOfDay
    return nil
}

// tempSiteSettingGatewayMgmtAutoSignatureUpdate is a temporary struct used for validating the fields of SiteSettingGatewayMgmtAutoSignatureUpdate.
type tempSiteSettingGatewayMgmtAutoSignatureUpdate  struct {
    DayOfWeek *DayOfWeekEnum `json:"day_of_week,omitempty"`
    Enable    *bool          `json:"enable,omitempty"`
    TimeOfDay *string        `json:"time_of_day,omitempty"`
}

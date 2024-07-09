package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SleImpactSummaryWlanItem represents a SleImpactSummaryWlanItem struct.
type SleImpactSummaryWlanItem struct {
    Degraded             float64        `json:"degraded"`
    Duration             float64        `json:"duration"`
    Name                 string         `json:"name"`
    Total                float64        `json:"total"`
    WlanId               string         `json:"wlan_id"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleImpactSummaryWlanItem.
// It customizes the JSON marshaling process for SleImpactSummaryWlanItem objects.
func (s SleImpactSummaryWlanItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactSummaryWlanItem object to a map representation for JSON marshaling.
func (s SleImpactSummaryWlanItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["degraded"] = s.Degraded
    structMap["duration"] = s.Duration
    structMap["name"] = s.Name
    structMap["total"] = s.Total
    structMap["wlan_id"] = s.WlanId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactSummaryWlanItem.
// It customizes the JSON unmarshaling process for SleImpactSummaryWlanItem objects.
func (s *SleImpactSummaryWlanItem) UnmarshalJSON(input []byte) error {
    var temp sleImpactSummaryWlanItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "degraded", "duration", "name", "total", "wlan_id")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Degraded = *temp.Degraded
    s.Duration = *temp.Duration
    s.Name = *temp.Name
    s.Total = *temp.Total
    s.WlanId = *temp.WlanId
    return nil
}

// sleImpactSummaryWlanItem is a temporary struct used for validating the fields of SleImpactSummaryWlanItem.
type sleImpactSummaryWlanItem  struct {
    Degraded *float64 `json:"degraded"`
    Duration *float64 `json:"duration"`
    Name     *string  `json:"name"`
    Total    *float64 `json:"total"`
    WlanId   *string  `json:"wlan_id"`
}

func (s *sleImpactSummaryWlanItem) validate() error {
    var errs []string
    if s.Degraded == nil {
        errs = append(errs, "required field `degraded` is missing for type `Sle_Impact_Summary_Wlan_Item`")
    }
    if s.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `Sle_Impact_Summary_Wlan_Item`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Sle_Impact_Summary_Wlan_Item`")
    }
    if s.Total == nil {
        errs = append(errs, "required field `total` is missing for type `Sle_Impact_Summary_Wlan_Item`")
    }
    if s.WlanId == nil {
        errs = append(errs, "required field `wlan_id` is missing for type `Sle_Impact_Summary_Wlan_Item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

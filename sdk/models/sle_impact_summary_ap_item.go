package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SleImpactSummaryApItem represents a SleImpactSummaryApItem struct.
type SleImpactSummaryApItem struct {
    ApMac                string         `json:"ap_mac"`
    Degraded             float64        `json:"degraded"`
    Duration             float64        `json:"duration"`
    Name                 string         `json:"name"`
    Total                float64        `json:"total"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleImpactSummaryApItem.
// It customizes the JSON marshaling process for SleImpactSummaryApItem objects.
func (s SleImpactSummaryApItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactSummaryApItem object to a map representation for JSON marshaling.
func (s SleImpactSummaryApItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["ap_mac"] = s.ApMac
    structMap["degraded"] = s.Degraded
    structMap["duration"] = s.Duration
    structMap["name"] = s.Name
    structMap["total"] = s.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactSummaryApItem.
// It customizes the JSON unmarshaling process for SleImpactSummaryApItem objects.
func (s *SleImpactSummaryApItem) UnmarshalJSON(input []byte) error {
    var temp sleImpactSummaryApItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_mac", "degraded", "duration", "name", "total")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ApMac = *temp.ApMac
    s.Degraded = *temp.Degraded
    s.Duration = *temp.Duration
    s.Name = *temp.Name
    s.Total = *temp.Total
    return nil
}

// sleImpactSummaryApItem is a temporary struct used for validating the fields of SleImpactSummaryApItem.
type sleImpactSummaryApItem  struct {
    ApMac    *string  `json:"ap_mac"`
    Degraded *float64 `json:"degraded"`
    Duration *float64 `json:"duration"`
    Name     *string  `json:"name"`
    Total    *float64 `json:"total"`
}

func (s *sleImpactSummaryApItem) validate() error {
    var errs []string
    if s.ApMac == nil {
        errs = append(errs, "required field `ap_mac` is missing for type `Sle_Impact_Summary_Ap_Item`")
    }
    if s.Degraded == nil {
        errs = append(errs, "required field `degraded` is missing for type `Sle_Impact_Summary_Ap_Item`")
    }
    if s.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `Sle_Impact_Summary_Ap_Item`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Sle_Impact_Summary_Ap_Item`")
    }
    if s.Total == nil {
        errs = append(errs, "required field `total` is missing for type `Sle_Impact_Summary_Ap_Item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

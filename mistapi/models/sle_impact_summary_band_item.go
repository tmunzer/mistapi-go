package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SleImpactSummaryBandItem represents a SleImpactSummaryBandItem struct.
type SleImpactSummaryBandItem struct {
    Band                 string         `json:"band"`
    Degraded             float64        `json:"degraded"`
    Duration             float64        `json:"duration"`
    Name                 string         `json:"name"`
    Total                float64        `json:"total"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleImpactSummaryBandItem.
// It customizes the JSON marshaling process for SleImpactSummaryBandItem objects.
func (s SleImpactSummaryBandItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactSummaryBandItem object to a map representation for JSON marshaling.
func (s SleImpactSummaryBandItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["band"] = s.Band
    structMap["degraded"] = s.Degraded
    structMap["duration"] = s.Duration
    structMap["name"] = s.Name
    structMap["total"] = s.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactSummaryBandItem.
// It customizes the JSON unmarshaling process for SleImpactSummaryBandItem objects.
func (s *SleImpactSummaryBandItem) UnmarshalJSON(input []byte) error {
    var temp sleImpactSummaryBandItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "band", "degraded", "duration", "name", "total")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Band = *temp.Band
    s.Degraded = *temp.Degraded
    s.Duration = *temp.Duration
    s.Name = *temp.Name
    s.Total = *temp.Total
    return nil
}

// sleImpactSummaryBandItem is a temporary struct used for validating the fields of SleImpactSummaryBandItem.
type sleImpactSummaryBandItem  struct {
    Band     *string  `json:"band"`
    Degraded *float64 `json:"degraded"`
    Duration *float64 `json:"duration"`
    Name     *string  `json:"name"`
    Total    *float64 `json:"total"`
}

func (s *sleImpactSummaryBandItem) validate() error {
    var errs []string
    if s.Band == nil {
        errs = append(errs, "required field `band` is missing for type `Sle_Impact_Summary_Band_Item`")
    }
    if s.Degraded == nil {
        errs = append(errs, "required field `degraded` is missing for type `Sle_Impact_Summary_Band_Item`")
    }
    if s.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `Sle_Impact_Summary_Band_Item`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Sle_Impact_Summary_Band_Item`")
    }
    if s.Total == nil {
        errs = append(errs, "required field `total` is missing for type `Sle_Impact_Summary_Band_Item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

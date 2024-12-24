package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SleImpactSummaryApItem represents a SleImpactSummaryApItem struct.
type SleImpactSummaryApItem struct {
    ApMac                string                 `json:"ap_mac"`
    Degraded             float64                `json:"degraded"`
    Duration             float64                `json:"duration"`
    Name                 string                 `json:"name"`
    Total                float64                `json:"total"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactSummaryApItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactSummaryApItem) String() string {
    return fmt.Sprintf(
    	"SleImpactSummaryApItem[ApMac=%v, Degraded=%v, Duration=%v, Name=%v, Total=%v, AdditionalProperties=%v]",
    	s.ApMac, s.Degraded, s.Duration, s.Name, s.Total, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactSummaryApItem.
// It customizes the JSON marshaling process for SleImpactSummaryApItem objects.
func (s SleImpactSummaryApItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "ap_mac", "degraded", "duration", "name", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactSummaryApItem object to a map representation for JSON marshaling.
func (s SleImpactSummaryApItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSleImpactSummaryApItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_mac", "degraded", "duration", "name", "total")
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

// tempSleImpactSummaryApItem is a temporary struct used for validating the fields of SleImpactSummaryApItem.
type tempSleImpactSummaryApItem  struct {
    ApMac    *string  `json:"ap_mac"`
    Degraded *float64 `json:"degraded"`
    Duration *float64 `json:"duration"`
    Name     *string  `json:"name"`
    Total    *float64 `json:"total"`
}

func (s *tempSleImpactSummaryApItem) validate() error {
    var errs []string
    if s.ApMac == nil {
        errs = append(errs, "required field `ap_mac` is missing for type `sle_impact_summary_ap_item`")
    }
    if s.Degraded == nil {
        errs = append(errs, "required field `degraded` is missing for type `sle_impact_summary_ap_item`")
    }
    if s.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `sle_impact_summary_ap_item`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `sle_impact_summary_ap_item`")
    }
    if s.Total == nil {
        errs = append(errs, "required field `total` is missing for type `sle_impact_summary_ap_item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

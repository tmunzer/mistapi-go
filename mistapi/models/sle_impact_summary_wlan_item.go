package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SleImpactSummaryWlanItem represents a SleImpactSummaryWlanItem struct.
type SleImpactSummaryWlanItem struct {
    Degraded             float64                `json:"degraded"`
    Duration             float64                `json:"duration"`
    Name                 string                 `json:"name"`
    Total                float64                `json:"total"`
    WlanId               string                 `json:"wlan_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactSummaryWlanItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactSummaryWlanItem) String() string {
    return fmt.Sprintf(
    	"SleImpactSummaryWlanItem[Degraded=%v, Duration=%v, Name=%v, Total=%v, WlanId=%v, AdditionalProperties=%v]",
    	s.Degraded, s.Duration, s.Name, s.Total, s.WlanId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactSummaryWlanItem.
// It customizes the JSON marshaling process for SleImpactSummaryWlanItem objects.
func (s SleImpactSummaryWlanItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "degraded", "duration", "name", "total", "wlan_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactSummaryWlanItem object to a map representation for JSON marshaling.
func (s SleImpactSummaryWlanItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    var temp tempSleImpactSummaryWlanItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "degraded", "duration", "name", "total", "wlan_id")
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

// tempSleImpactSummaryWlanItem is a temporary struct used for validating the fields of SleImpactSummaryWlanItem.
type tempSleImpactSummaryWlanItem  struct {
    Degraded *float64 `json:"degraded"`
    Duration *float64 `json:"duration"`
    Name     *string  `json:"name"`
    Total    *float64 `json:"total"`
    WlanId   *string  `json:"wlan_id"`
}

func (s *tempSleImpactSummaryWlanItem) validate() error {
    var errs []string
    if s.Degraded == nil {
        errs = append(errs, "required field `degraded` is missing for type `sle_impact_summary_wlan_item`")
    }
    if s.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `sle_impact_summary_wlan_item`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `sle_impact_summary_wlan_item`")
    }
    if s.Total == nil {
        errs = append(errs, "required field `total` is missing for type `sle_impact_summary_wlan_item`")
    }
    if s.WlanId == nil {
        errs = append(errs, "required field `wlan_id` is missing for type `sle_impact_summary_wlan_item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

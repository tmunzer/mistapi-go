package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SleImpactedApsAp represents a SleImpactedApsAp struct.
type SleImpactedApsAp struct {
    ApMac                string                 `json:"ap_mac"`
    Degraded             float64                `json:"degraded"`
    Duration             float64                `json:"duration"`
    Name                 string                 `json:"name"`
    Total                float64                `json:"total"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactedApsAp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactedApsAp) String() string {
    return fmt.Sprintf(
    	"SleImpactedApsAp[ApMac=%v, Degraded=%v, Duration=%v, Name=%v, Total=%v, AdditionalProperties=%v]",
    	s.ApMac, s.Degraded, s.Duration, s.Name, s.Total, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedApsAp.
// It customizes the JSON marshaling process for SleImpactedApsAp objects.
func (s SleImpactedApsAp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "ap_mac", "degraded", "duration", "name", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedApsAp object to a map representation for JSON marshaling.
func (s SleImpactedApsAp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["ap_mac"] = s.ApMac
    structMap["degraded"] = s.Degraded
    structMap["duration"] = s.Duration
    structMap["name"] = s.Name
    structMap["total"] = s.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedApsAp.
// It customizes the JSON unmarshaling process for SleImpactedApsAp objects.
func (s *SleImpactedApsAp) UnmarshalJSON(input []byte) error {
    var temp tempSleImpactedApsAp
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

// tempSleImpactedApsAp is a temporary struct used for validating the fields of SleImpactedApsAp.
type tempSleImpactedApsAp  struct {
    ApMac    *string  `json:"ap_mac"`
    Degraded *float64 `json:"degraded"`
    Duration *float64 `json:"duration"`
    Name     *string  `json:"name"`
    Total    *float64 `json:"total"`
}

func (s *tempSleImpactedApsAp) validate() error {
    var errs []string
    if s.ApMac == nil {
        errs = append(errs, "required field `ap_mac` is missing for type `sle_impacted_aps_ap`")
    }
    if s.Degraded == nil {
        errs = append(errs, "required field `degraded` is missing for type `sle_impacted_aps_ap`")
    }
    if s.Duration == nil {
        errs = append(errs, "required field `duration` is missing for type `sle_impacted_aps_ap`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `sle_impacted_aps_ap`")
    }
    if s.Total == nil {
        errs = append(errs, "required field `total` is missing for type `sle_impacted_aps_ap`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

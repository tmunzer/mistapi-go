package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SleImpactedApsAp represents a SleImpactedApsAp struct.
type SleImpactedApsAp struct {
    ApMac                string         `json:"ap_mac"`
    Degraded             float64        `json:"degraded"`
    Duration             float64        `json:"duration"`
    Name                 string         `json:"name"`
    Total                float64        `json:"total"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedApsAp.
// It customizes the JSON marshaling process for SleImpactedApsAp objects.
func (s SleImpactedApsAp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedApsAp object to a map representation for JSON marshaling.
func (s SleImpactedApsAp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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

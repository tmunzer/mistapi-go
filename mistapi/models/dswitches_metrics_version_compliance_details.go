package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// DswitchesMetricsVersionComplianceDetails represents a DswitchesMetricsVersionComplianceDetails struct.
type DswitchesMetricsVersionComplianceDetails struct {
    MajorVersions        []DswitchesComplianceMajorVersion `json:"major_versions"`
    TotalSwitchCount     int                               `json:"total_switch_count"`
    AdditionalProperties map[string]any                    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DswitchesMetricsVersionComplianceDetails.
// It customizes the JSON marshaling process for DswitchesMetricsVersionComplianceDetails objects.
func (d DswitchesMetricsVersionComplianceDetails) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DswitchesMetricsVersionComplianceDetails object to a map representation for JSON marshaling.
func (d DswitchesMetricsVersionComplianceDetails) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    structMap["major_versions"] = d.MajorVersions
    structMap["total_switch_count"] = d.TotalSwitchCount
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DswitchesMetricsVersionComplianceDetails.
// It customizes the JSON unmarshaling process for DswitchesMetricsVersionComplianceDetails objects.
func (d *DswitchesMetricsVersionComplianceDetails) UnmarshalJSON(input []byte) error {
    var temp tempDswitchesMetricsVersionComplianceDetails
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "major_versions", "total_switch_count")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.MajorVersions = *temp.MajorVersions
    d.TotalSwitchCount = *temp.TotalSwitchCount
    return nil
}

// tempDswitchesMetricsVersionComplianceDetails is a temporary struct used for validating the fields of DswitchesMetricsVersionComplianceDetails.
type tempDswitchesMetricsVersionComplianceDetails  struct {
    MajorVersions    *[]DswitchesComplianceMajorVersion `json:"major_versions"`
    TotalSwitchCount *int                               `json:"total_switch_count"`
}

func (d *tempDswitchesMetricsVersionComplianceDetails) validate() error {
    var errs []string
    if d.MajorVersions == nil {
        errs = append(errs, "required field `major_versions` is missing for type `dswitches_metrics_version_compliance_details`")
    }
    if d.TotalSwitchCount == nil {
        errs = append(errs, "required field `total_switch_count` is missing for type `dswitches_metrics_version_compliance_details`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

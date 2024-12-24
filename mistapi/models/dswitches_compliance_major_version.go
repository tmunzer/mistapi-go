package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// DswitchesComplianceMajorVersion represents a DswitchesComplianceMajorVersion struct.
type DswitchesComplianceMajorVersion struct {
    MajorCount           float64                `json:"major_count"`
    Model                string                 `json:"model"`
    SystemNames          []string               `json:"system_names,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for DswitchesComplianceMajorVersion,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DswitchesComplianceMajorVersion) String() string {
    return fmt.Sprintf(
    	"DswitchesComplianceMajorVersion[MajorCount=%v, Model=%v, SystemNames=%v, AdditionalProperties=%v]",
    	d.MajorCount, d.Model, d.SystemNames, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DswitchesComplianceMajorVersion.
// It customizes the JSON marshaling process for DswitchesComplianceMajorVersion objects.
func (d DswitchesComplianceMajorVersion) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "major_count", "model", "system_names"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DswitchesComplianceMajorVersion object to a map representation for JSON marshaling.
func (d DswitchesComplianceMajorVersion) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    structMap["major_count"] = d.MajorCount
    structMap["model"] = d.Model
    if d.SystemNames != nil {
        structMap["system_names"] = d.SystemNames
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DswitchesComplianceMajorVersion.
// It customizes the JSON unmarshaling process for DswitchesComplianceMajorVersion objects.
func (d *DswitchesComplianceMajorVersion) UnmarshalJSON(input []byte) error {
    var temp tempDswitchesComplianceMajorVersion
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "major_count", "model", "system_names")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.MajorCount = *temp.MajorCount
    d.Model = *temp.Model
    d.SystemNames = temp.SystemNames
    return nil
}

// tempDswitchesComplianceMajorVersion is a temporary struct used for validating the fields of DswitchesComplianceMajorVersion.
type tempDswitchesComplianceMajorVersion  struct {
    MajorCount  *float64 `json:"major_count"`
    Model       *string  `json:"model"`
    SystemNames []string `json:"system_names,omitempty"`
}

func (d *tempDswitchesComplianceMajorVersion) validate() error {
    var errs []string
    if d.MajorCount == nil {
        errs = append(errs, "required field `major_count` is missing for type `dswitches_compliance_major_version`")
    }
    if d.Model == nil {
        errs = append(errs, "required field `model` is missing for type `dswitches_compliance_major_version`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

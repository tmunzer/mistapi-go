package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// MxedgeUpgradeInfoItems represents a MxedgeUpgradeInfoItems struct.
type MxedgeUpgradeInfoItems struct {
    Default              *bool          `json:"default,omitempty"`
    Package              string         `json:"package"`
    Version              string         `json:"version"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeUpgradeInfoItems.
// It customizes the JSON marshaling process for MxedgeUpgradeInfoItems objects.
func (m MxedgeUpgradeInfoItems) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeUpgradeInfoItems object to a map representation for JSON marshaling.
func (m MxedgeUpgradeInfoItems) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Default != nil {
        structMap["default"] = m.Default
    }
    structMap["package"] = m.Package
    structMap["version"] = m.Version
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeUpgradeInfoItems.
// It customizes the JSON unmarshaling process for MxedgeUpgradeInfoItems objects.
func (m *MxedgeUpgradeInfoItems) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeUpgradeInfoItems
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "default", "package", "version")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Default = temp.Default
    m.Package = *temp.Package
    m.Version = *temp.Version
    return nil
}

// tempMxedgeUpgradeInfoItems is a temporary struct used for validating the fields of MxedgeUpgradeInfoItems.
type tempMxedgeUpgradeInfoItems  struct {
    Default *bool   `json:"default,omitempty"`
    Package *string `json:"package"`
    Version *string `json:"version"`
}

func (m *tempMxedgeUpgradeInfoItems) validate() error {
    var errs []string
    if m.Package == nil {
        errs = append(errs, "required field `package` is missing for type `mxedge_upgrade_info_items`")
    }
    if m.Version == nil {
        errs = append(errs, "required field `version` is missing for type `mxedge_upgrade_info_items`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

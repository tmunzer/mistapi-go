// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingSwitch represents a OrgSettingSwitch struct.
type OrgSettingSwitch struct {
    AutoUpgrade          *SwitchAutoUpgrade     `json:"auto_upgrade,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingSwitch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingSwitch) String() string {
    return fmt.Sprintf(
    	"OrgSettingSwitch[AutoUpgrade=%v, AdditionalProperties=%v]",
    	o.AutoUpgrade, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingSwitch.
// It customizes the JSON marshaling process for OrgSettingSwitch objects.
func (o OrgSettingSwitch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "auto_upgrade"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingSwitch object to a map representation for JSON marshaling.
func (o OrgSettingSwitch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.AutoUpgrade != nil {
        structMap["auto_upgrade"] = o.AutoUpgrade.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingSwitch.
// It customizes the JSON unmarshaling process for OrgSettingSwitch objects.
func (o *OrgSettingSwitch) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingSwitch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auto_upgrade")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.AutoUpgrade = temp.AutoUpgrade
    return nil
}

// tempOrgSettingSwitch is a temporary struct used for validating the fields of OrgSettingSwitch.
type tempOrgSettingSwitch  struct {
    AutoUpgrade *SwitchAutoUpgrade `json:"auto_upgrade,omitempty"`
}

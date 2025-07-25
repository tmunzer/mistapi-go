// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingWirelessPma represents a OrgSettingWirelessPma struct.
type OrgSettingWirelessPma struct {
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingWirelessPma,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingWirelessPma) String() string {
    return fmt.Sprintf(
    	"OrgSettingWirelessPma[Enabled=%v, AdditionalProperties=%v]",
    	o.Enabled, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingWirelessPma.
// It customizes the JSON marshaling process for OrgSettingWirelessPma objects.
func (o OrgSettingWirelessPma) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingWirelessPma object to a map representation for JSON marshaling.
func (o OrgSettingWirelessPma) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Enabled != nil {
        structMap["enabled"] = o.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingWirelessPma.
// It customizes the JSON unmarshaling process for OrgSettingWirelessPma objects.
func (o *OrgSettingWirelessPma) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingWirelessPma
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Enabled = temp.Enabled
    return nil
}

// tempOrgSettingWirelessPma is a temporary struct used for validating the fields of OrgSettingWirelessPma.
type tempOrgSettingWirelessPma  struct {
    Enabled *bool `json:"enabled,omitempty"`
}

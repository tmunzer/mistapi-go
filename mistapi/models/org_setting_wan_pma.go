// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingWanPma represents a OrgSettingWanPma struct.
type OrgSettingWanPma struct {
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingWanPma,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingWanPma) String() string {
    return fmt.Sprintf(
    	"OrgSettingWanPma[Enabled=%v, AdditionalProperties=%v]",
    	o.Enabled, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingWanPma.
// It customizes the JSON marshaling process for OrgSettingWanPma objects.
func (o OrgSettingWanPma) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingWanPma object to a map representation for JSON marshaling.
func (o OrgSettingWanPma) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Enabled != nil {
        structMap["enabled"] = o.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingWanPma.
// It customizes the JSON unmarshaling process for OrgSettingWanPma objects.
func (o *OrgSettingWanPma) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingWanPma
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

// tempOrgSettingWanPma is a temporary struct used for validating the fields of OrgSettingWanPma.
type tempOrgSettingWanPma  struct {
    Enabled *bool `json:"enabled,omitempty"`
}

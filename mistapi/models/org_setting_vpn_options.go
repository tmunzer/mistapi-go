// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingVpnOptions represents a OrgSettingVpnOptions struct.
type OrgSettingVpnOptions struct {
    AsBase               *int                   `json:"as_base,omitempty"`
    // requiring /12 or bigger to support 16 private IPs for 65535 gateways
    StSubnet             *string                `json:"st_subnet,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingVpnOptions,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingVpnOptions) String() string {
    return fmt.Sprintf(
    	"OrgSettingVpnOptions[AsBase=%v, StSubnet=%v, AdditionalProperties=%v]",
    	o.AsBase, o.StSubnet, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingVpnOptions.
// It customizes the JSON marshaling process for OrgSettingVpnOptions objects.
func (o OrgSettingVpnOptions) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "as_base", "st_subnet"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingVpnOptions object to a map representation for JSON marshaling.
func (o OrgSettingVpnOptions) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.AsBase != nil {
        structMap["as_base"] = o.AsBase
    }
    if o.StSubnet != nil {
        structMap["st_subnet"] = o.StSubnet
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingVpnOptions.
// It customizes the JSON unmarshaling process for OrgSettingVpnOptions objects.
func (o *OrgSettingVpnOptions) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingVpnOptions
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "as_base", "st_subnet")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.AsBase = temp.AsBase
    o.StSubnet = temp.StSubnet
    return nil
}

// tempOrgSettingVpnOptions is a temporary struct used for validating the fields of OrgSettingVpnOptions.
type tempOrgSettingVpnOptions  struct {
    AsBase   *int    `json:"as_base,omitempty"`
    StSubnet *string `json:"st_subnet,omitempty"`
}

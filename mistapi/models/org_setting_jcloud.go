// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgSettingJcloud represents a OrgSettingJcloud struct.
type OrgSettingJcloud struct {
	// JCloud Org Token
	OrgApitoken *string `json:"org_apitoken,omitempty"`
	// JCloud Org Token Name
	OrgApitokenName *string `json:"org_apitoken_name,omitempty"`
	// JCloud Org ID
	OrgId                *string                `json:"org_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingJcloud,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingJcloud) String() string {
	return fmt.Sprintf(
		"OrgSettingJcloud[OrgApitoken=%v, OrgApitokenName=%v, OrgId=%v, AdditionalProperties=%v]",
		o.OrgApitoken, o.OrgApitokenName, o.OrgId, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingJcloud.
// It customizes the JSON marshaling process for OrgSettingJcloud objects.
func (o OrgSettingJcloud) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"org_apitoken", "org_apitoken_name", "org_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingJcloud object to a map representation for JSON marshaling.
func (o OrgSettingJcloud) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.OrgApitoken != nil {
		structMap["org_apitoken"] = o.OrgApitoken
	}
	if o.OrgApitokenName != nil {
		structMap["org_apitoken_name"] = o.OrgApitokenName
	}
	if o.OrgId != nil {
		structMap["org_id"] = o.OrgId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingJcloud.
// It customizes the JSON unmarshaling process for OrgSettingJcloud objects.
func (o *OrgSettingJcloud) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingJcloud
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "org_apitoken", "org_apitoken_name", "org_id")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.OrgApitoken = temp.OrgApitoken
	o.OrgApitokenName = temp.OrgApitokenName
	o.OrgId = temp.OrgId
	return nil
}

// tempOrgSettingJcloud is a temporary struct used for validating the fields of OrgSettingJcloud.
type tempOrgSettingJcloud struct {
	OrgApitoken     *string `json:"org_apitoken,omitempty"`
	OrgApitokenName *string `json:"org_apitoken_name,omitempty"`
	OrgId           *string `json:"org_id,omitempty"`
}

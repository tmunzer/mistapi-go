// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgSettingCradlepoint represents a OrgSettingCradlepoint struct.
// Read-only Cradlepoint integration settings stored for the organization
type OrgSettingCradlepoint struct {
	// Cradlepoint API ID used by Mist for the integration
	CpApiId *string `json:"cp_api_id,omitempty"`
	// Cradlepoint API key paired with the Cradlepoint API ID
	CpApiKey *string `json:"cp_api_key,omitempty"`
	// Cradlepoint ECM API ID used by Mist for the integration
	EcmApiId *string `json:"ecm_api_id,omitempty"`
	// Cradlepoint ECM API key paired with the ECM API ID
	EcmApiKey *string `json:"ecm_api_key,omitempty"`
	// Whether Mist uses Cradlepoint LLDP data to link routers to Mist sites and devices
	EnableLldp           *bool                  `json:"enable_lldp,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingCradlepoint,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingCradlepoint) String() string {
	return fmt.Sprintf(
		"OrgSettingCradlepoint[CpApiId=%v, CpApiKey=%v, EcmApiId=%v, EcmApiKey=%v, EnableLldp=%v, AdditionalProperties=%v]",
		o.CpApiId, o.CpApiKey, o.EcmApiId, o.EcmApiKey, o.EnableLldp, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingCradlepoint.
// It customizes the JSON marshaling process for OrgSettingCradlepoint objects.
func (o OrgSettingCradlepoint) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"cp_api_id", "cp_api_key", "ecm_api_id", "ecm_api_key", "enable_lldp"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingCradlepoint object to a map representation for JSON marshaling.
func (o OrgSettingCradlepoint) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.CpApiId != nil {
		structMap["cp_api_id"] = o.CpApiId
	}
	if o.CpApiKey != nil {
		structMap["cp_api_key"] = o.CpApiKey
	}
	if o.EcmApiId != nil {
		structMap["ecm_api_id"] = o.EcmApiId
	}
	if o.EcmApiKey != nil {
		structMap["ecm_api_key"] = o.EcmApiKey
	}
	if o.EnableLldp != nil {
		structMap["enable_lldp"] = o.EnableLldp
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingCradlepoint.
// It customizes the JSON unmarshaling process for OrgSettingCradlepoint objects.
func (o *OrgSettingCradlepoint) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingCradlepoint
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cp_api_id", "cp_api_key", "ecm_api_id", "ecm_api_key", "enable_lldp")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.CpApiId = temp.CpApiId
	o.CpApiKey = temp.CpApiKey
	o.EcmApiId = temp.EcmApiId
	o.EcmApiKey = temp.EcmApiKey
	o.EnableLldp = temp.EnableLldp
	return nil
}

// tempOrgSettingCradlepoint is a temporary struct used for validating the fields of OrgSettingCradlepoint.
type tempOrgSettingCradlepoint struct {
	CpApiId    *string `json:"cp_api_id,omitempty"`
	CpApiKey   *string `json:"cp_api_key,omitempty"`
	EcmApiId   *string `json:"ecm_api_id,omitempty"`
	EcmApiKey  *string `json:"ecm_api_key,omitempty"`
	EnableLldp *bool   `json:"enable_lldp,omitempty"`
}

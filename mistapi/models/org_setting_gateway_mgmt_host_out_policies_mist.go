package models

import (
	"encoding/json"
)

// OrgSettingGatewayMgmtHostOutPoliciesMist represents a OrgSettingGatewayMgmtHostOutPoliciesMist struct.
type OrgSettingGatewayMgmtHostOutPoliciesMist struct {
	PathPreference       *string        `json:"path_preference,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingGatewayMgmtHostOutPoliciesMist.
// It customizes the JSON marshaling process for OrgSettingGatewayMgmtHostOutPoliciesMist objects.
func (o OrgSettingGatewayMgmtHostOutPoliciesMist) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingGatewayMgmtHostOutPoliciesMist object to a map representation for JSON marshaling.
func (o OrgSettingGatewayMgmtHostOutPoliciesMist) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, o.AdditionalProperties)
	if o.PathPreference != nil {
		structMap["path_preference"] = o.PathPreference
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingGatewayMgmtHostOutPoliciesMist.
// It customizes the JSON unmarshaling process for OrgSettingGatewayMgmtHostOutPoliciesMist objects.
func (o *OrgSettingGatewayMgmtHostOutPoliciesMist) UnmarshalJSON(input []byte) error {
	var temp tempOrgSettingGatewayMgmtHostOutPoliciesMist
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "path_preference")
	if err != nil {
		return err
	}

	o.AdditionalProperties = additionalProperties
	o.PathPreference = temp.PathPreference
	return nil
}

// tempOrgSettingGatewayMgmtHostOutPoliciesMist is a temporary struct used for validating the fields of OrgSettingGatewayMgmtHostOutPoliciesMist.
type tempOrgSettingGatewayMgmtHostOutPoliciesMist struct {
	PathPreference *string `json:"path_preference,omitempty"`
}

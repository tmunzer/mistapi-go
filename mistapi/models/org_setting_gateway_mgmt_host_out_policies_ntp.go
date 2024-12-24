package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingGatewayMgmtHostOutPoliciesNtp represents a OrgSettingGatewayMgmtHostOutPoliciesNtp struct.
type OrgSettingGatewayMgmtHostOutPoliciesNtp struct {
    PathPreference       *string                `json:"path_preference,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingGatewayMgmtHostOutPoliciesNtp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingGatewayMgmtHostOutPoliciesNtp) String() string {
    return fmt.Sprintf(
    	"OrgSettingGatewayMgmtHostOutPoliciesNtp[PathPreference=%v, AdditionalProperties=%v]",
    	o.PathPreference, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingGatewayMgmtHostOutPoliciesNtp.
// It customizes the JSON marshaling process for OrgSettingGatewayMgmtHostOutPoliciesNtp objects.
func (o OrgSettingGatewayMgmtHostOutPoliciesNtp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "path_preference"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingGatewayMgmtHostOutPoliciesNtp object to a map representation for JSON marshaling.
func (o OrgSettingGatewayMgmtHostOutPoliciesNtp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.PathPreference != nil {
        structMap["path_preference"] = o.PathPreference
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingGatewayMgmtHostOutPoliciesNtp.
// It customizes the JSON unmarshaling process for OrgSettingGatewayMgmtHostOutPoliciesNtp objects.
func (o *OrgSettingGatewayMgmtHostOutPoliciesNtp) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingGatewayMgmtHostOutPoliciesNtp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "path_preference")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.PathPreference = temp.PathPreference
    return nil
}

// tempOrgSettingGatewayMgmtHostOutPoliciesNtp is a temporary struct used for validating the fields of OrgSettingGatewayMgmtHostOutPoliciesNtp.
type tempOrgSettingGatewayMgmtHostOutPoliciesNtp  struct {
    PathPreference *string `json:"path_preference,omitempty"`
}

package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingGatewayMgmtHostOutPoliciesDns represents a OrgSettingGatewayMgmtHostOutPoliciesDns struct.
type OrgSettingGatewayMgmtHostOutPoliciesDns struct {
    PathPreference       *string                `json:"path_preference,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingGatewayMgmtHostOutPoliciesDns,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingGatewayMgmtHostOutPoliciesDns) String() string {
    return fmt.Sprintf(
    	"OrgSettingGatewayMgmtHostOutPoliciesDns[PathPreference=%v, AdditionalProperties=%v]",
    	o.PathPreference, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingGatewayMgmtHostOutPoliciesDns.
// It customizes the JSON marshaling process for OrgSettingGatewayMgmtHostOutPoliciesDns objects.
func (o OrgSettingGatewayMgmtHostOutPoliciesDns) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "path_preference"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingGatewayMgmtHostOutPoliciesDns object to a map representation for JSON marshaling.
func (o OrgSettingGatewayMgmtHostOutPoliciesDns) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.PathPreference != nil {
        structMap["path_preference"] = o.PathPreference
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingGatewayMgmtHostOutPoliciesDns.
// It customizes the JSON unmarshaling process for OrgSettingGatewayMgmtHostOutPoliciesDns objects.
func (o *OrgSettingGatewayMgmtHostOutPoliciesDns) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingGatewayMgmtHostOutPoliciesDns
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

// tempOrgSettingGatewayMgmtHostOutPoliciesDns is a temporary struct used for validating the fields of OrgSettingGatewayMgmtHostOutPoliciesDns.
type tempOrgSettingGatewayMgmtHostOutPoliciesDns  struct {
    PathPreference *string `json:"path_preference,omitempty"`
}

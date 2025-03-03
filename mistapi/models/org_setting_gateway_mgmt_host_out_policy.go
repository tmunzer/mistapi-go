package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingGatewayMgmtHostOutPolicy represents a OrgSettingGatewayMgmtHostOutPolicy struct.
type OrgSettingGatewayMgmtHostOutPolicy struct {
    PathPreference       *string                `json:"path_preference,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingGatewayMgmtHostOutPolicy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingGatewayMgmtHostOutPolicy) String() string {
    return fmt.Sprintf(
    	"OrgSettingGatewayMgmtHostOutPolicy[PathPreference=%v, AdditionalProperties=%v]",
    	o.PathPreference, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingGatewayMgmtHostOutPolicy.
// It customizes the JSON marshaling process for OrgSettingGatewayMgmtHostOutPolicy objects.
func (o OrgSettingGatewayMgmtHostOutPolicy) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "path_preference"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingGatewayMgmtHostOutPolicy object to a map representation for JSON marshaling.
func (o OrgSettingGatewayMgmtHostOutPolicy) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.PathPreference != nil {
        structMap["path_preference"] = o.PathPreference
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingGatewayMgmtHostOutPolicy.
// It customizes the JSON unmarshaling process for OrgSettingGatewayMgmtHostOutPolicy objects.
func (o *OrgSettingGatewayMgmtHostOutPolicy) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingGatewayMgmtHostOutPolicy
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

// tempOrgSettingGatewayMgmtHostOutPolicy is a temporary struct used for validating the fields of OrgSettingGatewayMgmtHostOutPolicy.
type tempOrgSettingGatewayMgmtHostOutPolicy  struct {
    PathPreference *string `json:"path_preference,omitempty"`
}

package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingGatewayMgmtAppProbing represents a OrgSettingGatewayMgmtAppProbing struct.
type OrgSettingGatewayMgmtAppProbing struct {
    // APp-keys from /api/v1/const/applications
    Apps                 []string               `json:"apps,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingGatewayMgmtAppProbing,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingGatewayMgmtAppProbing) String() string {
    return fmt.Sprintf(
    	"OrgSettingGatewayMgmtAppProbing[Apps=%v, AdditionalProperties=%v]",
    	o.Apps, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingGatewayMgmtAppProbing.
// It customizes the JSON marshaling process for OrgSettingGatewayMgmtAppProbing objects.
func (o OrgSettingGatewayMgmtAppProbing) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "apps"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingGatewayMgmtAppProbing object to a map representation for JSON marshaling.
func (o OrgSettingGatewayMgmtAppProbing) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Apps != nil {
        structMap["apps"] = o.Apps
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingGatewayMgmtAppProbing.
// It customizes the JSON unmarshaling process for OrgSettingGatewayMgmtAppProbing objects.
func (o *OrgSettingGatewayMgmtAppProbing) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingGatewayMgmtAppProbing
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "apps")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Apps = temp.Apps
    return nil
}

// tempOrgSettingGatewayMgmtAppProbing is a temporary struct used for validating the fields of OrgSettingGatewayMgmtAppProbing.
type tempOrgSettingGatewayMgmtAppProbing  struct {
    Apps []string `json:"apps,omitempty"`
}

package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingDeviceCert represents a OrgSettingDeviceCert struct.
// common device cert, optional
type OrgSettingDeviceCert struct {
    Cert                 *string                `json:"cert,omitempty"`
    Key                  *string                `json:"key,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingDeviceCert,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingDeviceCert) String() string {
    return fmt.Sprintf(
    	"OrgSettingDeviceCert[Cert=%v, Key=%v, AdditionalProperties=%v]",
    	o.Cert, o.Key, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingDeviceCert.
// It customizes the JSON marshaling process for OrgSettingDeviceCert objects.
func (o OrgSettingDeviceCert) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "cert", "key"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingDeviceCert object to a map representation for JSON marshaling.
func (o OrgSettingDeviceCert) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Cert != nil {
        structMap["cert"] = o.Cert
    }
    if o.Key != nil {
        structMap["key"] = o.Key
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingDeviceCert.
// It customizes the JSON unmarshaling process for OrgSettingDeviceCert objects.
func (o *OrgSettingDeviceCert) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingDeviceCert
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cert", "key")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Cert = temp.Cert
    o.Key = temp.Key
    return nil
}

// tempOrgSettingDeviceCert is a temporary struct used for validating the fields of OrgSettingDeviceCert.
type tempOrgSettingDeviceCert  struct {
    Cert *string `json:"cert,omitempty"`
    Key  *string `json:"key,omitempty"`
}

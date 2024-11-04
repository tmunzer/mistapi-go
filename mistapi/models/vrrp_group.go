package models

import (
    "encoding/json"
)

// VrrpGroup represents a VrrpGroup struct.
// Junos VRRP group
type VrrpGroup struct {
    // if `auth_type`==`md5`
    AuthKey              *string                     `json:"auth_key,omitempty"`
    // if `auth_type`==`simple`
    AuthPassword         *string                     `json:"auth_password,omitempty"`
    // enum: `md5`, `simple`
    AuthType             *VrrpGroupAuthTypeEnum      `json:"auth_type,omitempty"`
    // Property key is the network name
    Networks             map[string]VrrpGroupNetwork `json:"networks,omitempty"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VrrpGroup.
// It customizes the JSON marshaling process for VrrpGroup objects.
func (v VrrpGroup) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VrrpGroup object to a map representation for JSON marshaling.
func (v VrrpGroup) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
    if v.AuthKey != nil {
        structMap["auth_key"] = v.AuthKey
    }
    if v.AuthPassword != nil {
        structMap["auth_password"] = v.AuthPassword
    }
    if v.AuthType != nil {
        structMap["auth_type"] = v.AuthType
    }
    if v.Networks != nil {
        structMap["networks"] = v.Networks
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VrrpGroup.
// It customizes the JSON unmarshaling process for VrrpGroup objects.
func (v *VrrpGroup) UnmarshalJSON(input []byte) error {
    var temp tempVrrpGroup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "auth_key", "auth_password", "auth_type", "networks")
    if err != nil {
    	return err
    }
    
    v.AdditionalProperties = additionalProperties
    v.AuthKey = temp.AuthKey
    v.AuthPassword = temp.AuthPassword
    v.AuthType = temp.AuthType
    v.Networks = temp.Networks
    return nil
}

// tempVrrpGroup is a temporary struct used for validating the fields of VrrpGroup.
type tempVrrpGroup  struct {
    AuthKey      *string                     `json:"auth_key,omitempty"`
    AuthPassword *string                     `json:"auth_password,omitempty"`
    AuthType     *VrrpGroupAuthTypeEnum      `json:"auth_type,omitempty"`
    Networks     map[string]VrrpGroupNetwork `json:"networks,omitempty"`
}

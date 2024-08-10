package models

import (
    "encoding/json"
)

// OrgServicePolicySslProxy represents a OrgServicePolicySslProxy struct.
// for SRX-only
type OrgServicePolicySslProxy struct {
    CiphersCatagory      *SslProxyCiphersCatagoryEnum `json:"ciphers_catagory,omitempty"`
    Enabled              *bool                        `json:"enabled,omitempty"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgServicePolicySslProxy.
// It customizes the JSON marshaling process for OrgServicePolicySslProxy objects.
func (o OrgServicePolicySslProxy) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgServicePolicySslProxy object to a map representation for JSON marshaling.
func (o OrgServicePolicySslProxy) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.CiphersCatagory != nil {
        structMap["ciphers_catagory"] = o.CiphersCatagory
    }
    if o.Enabled != nil {
        structMap["enabled"] = o.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgServicePolicySslProxy.
// It customizes the JSON unmarshaling process for OrgServicePolicySslProxy objects.
func (o *OrgServicePolicySslProxy) UnmarshalJSON(input []byte) error {
    var temp tempOrgServicePolicySslProxy
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ciphers_catagory", "enabled")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.CiphersCatagory = temp.CiphersCatagory
    o.Enabled = temp.Enabled
    return nil
}

// tempOrgServicePolicySslProxy is a temporary struct used for validating the fields of OrgServicePolicySslProxy.
type tempOrgServicePolicySslProxy  struct {
    CiphersCatagory *SslProxyCiphersCatagoryEnum `json:"ciphers_catagory,omitempty"`
    Enabled         *bool                        `json:"enabled,omitempty"`
}

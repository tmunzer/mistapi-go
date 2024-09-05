package models

import (
    "encoding/json"
)

// OrgSettingApiPolicy represents a OrgSettingApiPolicy struct.
type OrgSettingApiPolicy struct {
    // by default, API hides password/secrets when the user doesn't have write access
    // * `true`: API will hide passwords/secrets for all users
    // * `false`: API will hide passwords/secrets for read-only users
    NoReveal             *bool          `json:"no_reveal,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingApiPolicy.
// It customizes the JSON marshaling process for OrgSettingApiPolicy objects.
func (o OrgSettingApiPolicy) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingApiPolicy object to a map representation for JSON marshaling.
func (o OrgSettingApiPolicy) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.NoReveal != nil {
        structMap["no_reveal"] = o.NoReveal
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingApiPolicy.
// It customizes the JSON unmarshaling process for OrgSettingApiPolicy objects.
func (o *OrgSettingApiPolicy) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingApiPolicy
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "no_reveal")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.NoReveal = temp.NoReveal
    return nil
}

// tempOrgSettingApiPolicy is a temporary struct used for validating the fields of OrgSettingApiPolicy.
type tempOrgSettingApiPolicy  struct {
    NoReveal *bool `json:"no_reveal,omitempty"`
}

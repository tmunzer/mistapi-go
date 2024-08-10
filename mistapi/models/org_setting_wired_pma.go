package models

import (
    "encoding/json"
)

// OrgSettingWiredPma represents a OrgSettingWiredPma struct.
type OrgSettingWiredPma struct {
    Enabled              *bool          `json:"enabled,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingWiredPma.
// It customizes the JSON marshaling process for OrgSettingWiredPma objects.
func (o OrgSettingWiredPma) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingWiredPma object to a map representation for JSON marshaling.
func (o OrgSettingWiredPma) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Enabled != nil {
        structMap["enabled"] = o.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingWiredPma.
// It customizes the JSON unmarshaling process for OrgSettingWiredPma objects.
func (o *OrgSettingWiredPma) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingWiredPma
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.Enabled = temp.Enabled
    return nil
}

// tempOrgSettingWiredPma is a temporary struct used for validating the fields of OrgSettingWiredPma.
type tempOrgSettingWiredPma  struct {
    Enabled *bool `json:"enabled,omitempty"`
}

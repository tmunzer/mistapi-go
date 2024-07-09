package models

import (
    "encoding/json"
)

// OrgSettingWanPma represents a OrgSettingWanPma struct.
type OrgSettingWanPma struct {
    Enabled              *bool          `json:"enabled,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingWanPma.
// It customizes the JSON marshaling process for OrgSettingWanPma objects.
func (o OrgSettingWanPma) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingWanPma object to a map representation for JSON marshaling.
func (o OrgSettingWanPma) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Enabled != nil {
        structMap["enabled"] = o.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingWanPma.
// It customizes the JSON unmarshaling process for OrgSettingWanPma objects.
func (o *OrgSettingWanPma) UnmarshalJSON(input []byte) error {
    var temp orgSettingWanPma
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

// orgSettingWanPma is a temporary struct used for validating the fields of OrgSettingWanPma.
type orgSettingWanPma  struct {
    Enabled *bool `json:"enabled,omitempty"`
}

package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingScep represents a OrgSettingScep struct.
type OrgSettingScep struct {
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingScep,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingScep) String() string {
    return fmt.Sprintf(
    	"OrgSettingScep[Enabled=%v, AdditionalProperties=%v]",
    	o.Enabled, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingScep.
// It customizes the JSON marshaling process for OrgSettingScep objects.
func (o OrgSettingScep) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingScep object to a map representation for JSON marshaling.
func (o OrgSettingScep) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Enabled != nil {
        structMap["enabled"] = o.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingScep.
// It customizes the JSON unmarshaling process for OrgSettingScep objects.
func (o *OrgSettingScep) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingScep
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Enabled = temp.Enabled
    return nil
}

// tempOrgSettingScep is a temporary struct used for validating the fields of OrgSettingScep.
type tempOrgSettingScep  struct {
    Enabled *bool `json:"enabled,omitempty"`
}

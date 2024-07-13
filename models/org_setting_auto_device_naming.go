package models

import (
    "encoding/json"
)

// OrgSettingAutoDeviceNaming represents a OrgSettingAutoDeviceNaming struct.
type OrgSettingAutoDeviceNaming struct {
    Enable               *bool                    `json:"enable,omitempty"`
    Rules                Optional[[]OrgAutoRules] `json:"rules"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingAutoDeviceNaming.
// It customizes the JSON marshaling process for OrgSettingAutoDeviceNaming objects.
func (o OrgSettingAutoDeviceNaming) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingAutoDeviceNaming object to a map representation for JSON marshaling.
func (o OrgSettingAutoDeviceNaming) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Enable != nil {
        structMap["enable"] = o.Enable
    }
    if o.Rules.IsValueSet() {
        if o.Rules.Value() != nil {
            structMap["rules"] = o.Rules.Value()
        } else {
            structMap["rules"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingAutoDeviceNaming.
// It customizes the JSON unmarshaling process for OrgSettingAutoDeviceNaming objects.
func (o *OrgSettingAutoDeviceNaming) UnmarshalJSON(input []byte) error {
    var temp orgSettingAutoDeviceNaming
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "enable", "rules")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.Enable = temp.Enable
    o.Rules = temp.Rules
    return nil
}

// orgSettingAutoDeviceNaming is a temporary struct used for validating the fields of OrgSettingAutoDeviceNaming.
type orgSettingAutoDeviceNaming  struct {
    Enable *bool                    `json:"enable,omitempty"`
    Rules  Optional[[]OrgAutoRules] `json:"rules"`
}

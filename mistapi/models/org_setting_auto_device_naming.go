package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingAutoDeviceNaming represents a OrgSettingAutoDeviceNaming struct.
type OrgSettingAutoDeviceNaming struct {
    Enable               *bool                    `json:"enable,omitempty"`
    Rules                Optional[[]OrgAutoRules] `json:"rules"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingAutoDeviceNaming,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingAutoDeviceNaming) String() string {
    return fmt.Sprintf(
    	"OrgSettingAutoDeviceNaming[Enable=%v, Rules=%v, AdditionalProperties=%v]",
    	o.Enable, o.Rules, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingAutoDeviceNaming.
// It customizes the JSON marshaling process for OrgSettingAutoDeviceNaming objects.
func (o OrgSettingAutoDeviceNaming) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "enable", "rules"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingAutoDeviceNaming object to a map representation for JSON marshaling.
func (o OrgSettingAutoDeviceNaming) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
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
    var temp tempOrgSettingAutoDeviceNaming
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enable", "rules")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Enable = temp.Enable
    o.Rules = temp.Rules
    return nil
}

// tempOrgSettingAutoDeviceNaming is a temporary struct used for validating the fields of OrgSettingAutoDeviceNaming.
type tempOrgSettingAutoDeviceNaming  struct {
    Enable *bool                    `json:"enable,omitempty"`
    Rules  Optional[[]OrgAutoRules] `json:"rules"`
}

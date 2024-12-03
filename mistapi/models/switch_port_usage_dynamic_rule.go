package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SwitchPortUsageDynamicRule represents a SwitchPortUsageDynamicRule struct.
type SwitchPortUsageDynamicRule struct {
    Equals               *string                           `json:"equals,omitempty"`
    // use `equals_any` to match any item in a list
    EqualsAny            []string                          `json:"equals_any,omitempty"`
    // "[0:3]":"abcdef" -> "abc"
    // "split(.)[1]": "a.b.c" -> "b"
    // "split(-)[1][0:3]: "a1234-b5678-c90" -> "b56"
    Expression           *string                           `json:"expression,omitempty"`
    // enum: `link_peermac`, `lldp_chassis_id`, `lldp_hardware_revision`, `lldp_manufacturer_name`, `lldp_oui`, `lldp_serial_number`, `lldp_system_name`, `radius_dynamicfilter`, `radius_usermac`, `radius_username`
    Src                  SwitchPortUsageDynamicRuleSrcEnum `json:"src"`
    // `port_usage` name
    Usage                *string                           `json:"usage,omitempty"`
    AdditionalProperties map[string]interface{}            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchPortUsageDynamicRule.
// It customizes the JSON marshaling process for SwitchPortUsageDynamicRule objects.
func (s SwitchPortUsageDynamicRule) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "equals", "equals_any", "expression", "src", "usage"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchPortUsageDynamicRule object to a map representation for JSON marshaling.
func (s SwitchPortUsageDynamicRule) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Equals != nil {
        structMap["equals"] = s.Equals
    }
    if s.EqualsAny != nil {
        structMap["equals_any"] = s.EqualsAny
    }
    if s.Expression != nil {
        structMap["expression"] = s.Expression
    }
    structMap["src"] = s.Src
    if s.Usage != nil {
        structMap["usage"] = s.Usage
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchPortUsageDynamicRule.
// It customizes the JSON unmarshaling process for SwitchPortUsageDynamicRule objects.
func (s *SwitchPortUsageDynamicRule) UnmarshalJSON(input []byte) error {
    var temp tempSwitchPortUsageDynamicRule
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "equals", "equals_any", "expression", "src", "usage")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Equals = temp.Equals
    s.EqualsAny = temp.EqualsAny
    s.Expression = temp.Expression
    s.Src = *temp.Src
    s.Usage = temp.Usage
    return nil
}

// tempSwitchPortUsageDynamicRule is a temporary struct used for validating the fields of SwitchPortUsageDynamicRule.
type tempSwitchPortUsageDynamicRule  struct {
    Equals     *string                            `json:"equals,omitempty"`
    EqualsAny  []string                           `json:"equals_any,omitempty"`
    Expression *string                            `json:"expression,omitempty"`
    Src        *SwitchPortUsageDynamicRuleSrcEnum `json:"src"`
    Usage      *string                            `json:"usage,omitempty"`
}

func (s *tempSwitchPortUsageDynamicRule) validate() error {
    var errs []string
    if s.Src == nil {
        errs = append(errs, "required field `src` is missing for type `switch_port_usage_dynamic_rule`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

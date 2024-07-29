package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// OrgAutoRules represents a OrgAutoRules struct.
// auto_rules in org settings
type OrgAutoRules struct {
    // "[0:3]"            // "abcdef" -> "abc"
    // "split(.)[1]"      // "a.b.c" -> "b"
    // "split(-)[1][0:3]" // "a1234-b5678-c90" -> "b56"
    Expression           Optional[string]                 `json:"expression"`
    // optional/additional filter. enum: `ap`, `gateway`, `other`, `switch`
    MatchDeviceType      *OrgAutoRulesMatchDeviceTypeEnum `json:"match_device_type,omitempty"`
    // optional/additional filter
    MatchModel           *string                          `json:"match_model,omitempty"`
    Model                *string                          `json:"model,omitempty"`
    Prefix               Optional[string]                 `json:"prefix"`
    // enum: `dns_suffix`, `lldp_port_desc`, `lldp_system_name`, `model`, `name`, `subnet`
    Src                  OrgAutoRulesSrcEnum              `json:"src"`
    Subnet               *string                          `json:"subnet,omitempty"`
    Suffix               Optional[string]                 `json:"suffix"`
    Value                *string                          `json:"value,omitempty"`
    AdditionalProperties map[string]any                   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgAutoRules.
// It customizes the JSON marshaling process for OrgAutoRules objects.
func (o OrgAutoRules) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgAutoRules object to a map representation for JSON marshaling.
func (o OrgAutoRules) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Expression.IsValueSet() {
        if o.Expression.Value() != nil {
            structMap["expression"] = o.Expression.Value()
        } else {
            structMap["expression"] = nil
        }
    }
    if o.MatchDeviceType != nil {
        structMap["match_device_type"] = o.MatchDeviceType
    }
    if o.MatchModel != nil {
        structMap["match_model"] = o.MatchModel
    }
    if o.Model != nil {
        structMap["model"] = o.Model
    }
    if o.Prefix.IsValueSet() {
        if o.Prefix.Value() != nil {
            structMap["prefix"] = o.Prefix.Value()
        } else {
            structMap["prefix"] = nil
        }
    }
    structMap["src"] = o.Src
    if o.Subnet != nil {
        structMap["subnet"] = o.Subnet
    }
    if o.Suffix.IsValueSet() {
        if o.Suffix.Value() != nil {
            structMap["suffix"] = o.Suffix.Value()
        } else {
            structMap["suffix"] = nil
        }
    }
    if o.Value != nil {
        structMap["value"] = o.Value
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgAutoRules.
// It customizes the JSON unmarshaling process for OrgAutoRules objects.
func (o *OrgAutoRules) UnmarshalJSON(input []byte) error {
    var temp orgAutoRules
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "expression", "match_device_type", "match_model", "model", "prefix", "src", "subnet", "suffix", "value")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.Expression = temp.Expression
    o.MatchDeviceType = temp.MatchDeviceType
    o.MatchModel = temp.MatchModel
    o.Model = temp.Model
    o.Prefix = temp.Prefix
    o.Src = *temp.Src
    o.Subnet = temp.Subnet
    o.Suffix = temp.Suffix
    o.Value = temp.Value
    return nil
}

// orgAutoRules is a temporary struct used for validating the fields of OrgAutoRules.
type orgAutoRules  struct {
    Expression      Optional[string]                 `json:"expression"`
    MatchDeviceType *OrgAutoRulesMatchDeviceTypeEnum `json:"match_device_type,omitempty"`
    MatchModel      *string                          `json:"match_model,omitempty"`
    Model           *string                          `json:"model,omitempty"`
    Prefix          Optional[string]                 `json:"prefix"`
    Src             *OrgAutoRulesSrcEnum             `json:"src"`
    Subnet          *string                          `json:"subnet,omitempty"`
    Suffix          Optional[string]                 `json:"suffix"`
    Value           *string                          `json:"value,omitempty"`
}

func (o *orgAutoRules) validate() error {
    var errs []string
    if o.Src == nil {
        errs = append(errs, "required field `src` is missing for type `Org_Auto_Rules`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

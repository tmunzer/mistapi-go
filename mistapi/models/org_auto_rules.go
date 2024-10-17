package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// OrgAutoRules represents a OrgAutoRules struct.
// auto_rules in org settings
type OrgAutoRules struct {
	// if `src`==`geoip`. By default, a claimed device only gets assigned if the site exists to auto-create the site, enable this
	CreateNewSiteIfNeeded *bool `json:"create_new_site_if_needed,omitempty"`
	// if `src`==`name`, `src`==`lldp_system_name`,  `src`==`dns_suffix`
	// "[0:3]"            // "abcdef" -> "abc"
	// "split(.)[1]"      // "a.b.c" -> "b"
	// "split(-)[1][0:3]" // "a1234-b5678-c90" -> "b56"'
	Expression Optional[string] `json:"expression"`
	// if `src`==`geoip` and `create_new_site_if_needed`==`true`. If a gateway template is desired for this newly created site
	GatewaytemplateId *string `json:"gatewaytemplate_id,omitempty"`
	// if `src`==`geoip`
	MatchCountry *string `json:"match_country,omitempty"`
	// optional/additional filter. enum: `ap`, `gateway`, `other`, `switch`
	MatchDeviceType *OrgAutoRulesMatchDeviceTypeEnum `json:"match_device_type,omitempty"`
	// optional/additional filter
	MatchModel *string `json:"match_model,omitempty"`
	// if `src`==`model`
	Model *string `json:"model,omitempty"`
	// if `src`==`name`
	Prefix Optional[string] `json:"prefix"`
	// enum: `dns_suffix`, `geoip`, `lldp_port_desc`, `lldp_system_name`, `model`, `name`, `subnet`
	Src OrgAutoRulesSrcEnum `json:"src"`
	// if `src`==`subnet`
	Subnet *string `json:"subnet,omitempty"`
	// if `src`==`name`
	Suffix Optional[string] `json:"suffix"`
	// if
	// * `src`==`model`
	// *  `src`==`geoip: site name for the device to be assigned to ("city" / "city+country" / ...)
	Value                *string        `json:"value,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
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
	if o.CreateNewSiteIfNeeded != nil {
		structMap["create_new_site_if_needed"] = o.CreateNewSiteIfNeeded
	}
	if o.Expression.IsValueSet() {
		if o.Expression.Value() != nil {
			structMap["expression"] = o.Expression.Value()
		} else {
			structMap["expression"] = nil
		}
	}
	if o.GatewaytemplateId != nil {
		structMap["gatewaytemplate_id"] = o.GatewaytemplateId
	}
	if o.MatchCountry != nil {
		structMap["match_country"] = o.MatchCountry
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
	var temp tempOrgAutoRules
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "create_new_site_if_needed", "expression", "gatewaytemplate_id", "match_country", "match_device_type", "match_model", "model", "prefix", "src", "subnet", "suffix", "value")
	if err != nil {
		return err
	}

	o.AdditionalProperties = additionalProperties
	o.CreateNewSiteIfNeeded = temp.CreateNewSiteIfNeeded
	o.Expression = temp.Expression
	o.GatewaytemplateId = temp.GatewaytemplateId
	o.MatchCountry = temp.MatchCountry
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

// tempOrgAutoRules is a temporary struct used for validating the fields of OrgAutoRules.
type tempOrgAutoRules struct {
	CreateNewSiteIfNeeded *bool                            `json:"create_new_site_if_needed,omitempty"`
	Expression            Optional[string]                 `json:"expression"`
	GatewaytemplateId     *string                          `json:"gatewaytemplate_id,omitempty"`
	MatchCountry          *string                          `json:"match_country,omitempty"`
	MatchDeviceType       *OrgAutoRulesMatchDeviceTypeEnum `json:"match_device_type,omitempty"`
	MatchModel            *string                          `json:"match_model,omitempty"`
	Model                 *string                          `json:"model,omitempty"`
	Prefix                Optional[string]                 `json:"prefix"`
	Src                   *OrgAutoRulesSrcEnum             `json:"src"`
	Subnet                *string                          `json:"subnet,omitempty"`
	Suffix                Optional[string]                 `json:"suffix"`
	Value                 *string                          `json:"value,omitempty"`
}

func (o *tempOrgAutoRules) validate() error {
	var errs []string
	if o.Src == nil {
		errs = append(errs, "required field `src` is missing for type `org_auto_rules`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

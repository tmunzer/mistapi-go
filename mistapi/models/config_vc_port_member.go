package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ConfigVcPortMember represents a ConfigVcPortMember struct.
type ConfigVcPortMember struct {
	Member               float64        `json:"member"`
	VcPorts              []string       `json:"vc_ports,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConfigVcPortMember.
// It customizes the JSON marshaling process for ConfigVcPortMember objects.
func (c ConfigVcPortMember) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ConfigVcPortMember object to a map representation for JSON marshaling.
func (c ConfigVcPortMember) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, c.AdditionalProperties)
	structMap["member"] = c.Member
	if c.VcPorts != nil {
		structMap["vc_ports"] = c.VcPorts
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConfigVcPortMember.
// It customizes the JSON unmarshaling process for ConfigVcPortMember objects.
func (c *ConfigVcPortMember) UnmarshalJSON(input []byte) error {
	var temp tempConfigVcPortMember
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "member", "vc_ports")
	if err != nil {
		return err
	}

	c.AdditionalProperties = additionalProperties
	c.Member = *temp.Member
	c.VcPorts = temp.VcPorts
	return nil
}

// tempConfigVcPortMember is a temporary struct used for validating the fields of ConfigVcPortMember.
type tempConfigVcPortMember struct {
	Member  *float64 `json:"member"`
	VcPorts []string `json:"vc_ports,omitempty"`
}

func (c *tempConfigVcPortMember) validate() error {
	var errs []string
	if c.Member == nil {
		errs = append(errs, "required field `member` is missing for type `config_vc_port_member`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

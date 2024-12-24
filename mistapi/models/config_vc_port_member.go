package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ConfigVcPortMember represents a ConfigVcPortMember struct.
type ConfigVcPortMember struct {
    Member               float64                `json:"member"`
    VcPorts              []string               `json:"vc_ports,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConfigVcPortMember,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConfigVcPortMember) String() string {
    return fmt.Sprintf(
    	"ConfigVcPortMember[Member=%v, VcPorts=%v, AdditionalProperties=%v]",
    	c.Member, c.VcPorts, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConfigVcPortMember.
// It customizes the JSON marshaling process for ConfigVcPortMember objects.
func (c ConfigVcPortMember) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "member", "vc_ports"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConfigVcPortMember object to a map representation for JSON marshaling.
func (c ConfigVcPortMember) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "member", "vc_ports")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Member = *temp.Member
    c.VcPorts = temp.VcPorts
    return nil
}

// tempConfigVcPortMember is a temporary struct used for validating the fields of ConfigVcPortMember.
type tempConfigVcPortMember  struct {
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
    return errors.New(strings.Join (errs, "\n"))
}

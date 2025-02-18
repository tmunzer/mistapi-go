package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// MxedgeTuntermOtherIpConfig represents a MxedgeTuntermOtherIpConfig struct.
type MxedgeTuntermOtherIpConfig struct {
    Ip                   string                 `json:"ip"`
    Netmask              string                 `json:"netmask"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeTuntermOtherIpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeTuntermOtherIpConfig) String() string {
    return fmt.Sprintf(
    	"MxedgeTuntermOtherIpConfig[Ip=%v, Netmask=%v, AdditionalProperties=%v]",
    	m.Ip, m.Netmask, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermOtherIpConfig.
// It customizes the JSON marshaling process for MxedgeTuntermOtherIpConfig objects.
func (m MxedgeTuntermOtherIpConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "ip", "netmask"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermOtherIpConfig object to a map representation for JSON marshaling.
func (m MxedgeTuntermOtherIpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["ip"] = m.Ip
    structMap["netmask"] = m.Netmask
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermOtherIpConfig.
// It customizes the JSON unmarshaling process for MxedgeTuntermOtherIpConfig objects.
func (m *MxedgeTuntermOtherIpConfig) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeTuntermOtherIpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ip", "netmask")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Ip = *temp.Ip
    m.Netmask = *temp.Netmask
    return nil
}

// tempMxedgeTuntermOtherIpConfig is a temporary struct used for validating the fields of MxedgeTuntermOtherIpConfig.
type tempMxedgeTuntermOtherIpConfig  struct {
    Ip      *string `json:"ip"`
    Netmask *string `json:"netmask"`
}

func (m *tempMxedgeTuntermOtherIpConfig) validate() error {
    var errs []string
    if m.Ip == nil {
        errs = append(errs, "required field `ip` is missing for type `mxedge_tunterm_other_ip_config`")
    }
    if m.Netmask == nil {
        errs = append(errs, "required field `netmask` is missing for type `mxedge_tunterm_other_ip_config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

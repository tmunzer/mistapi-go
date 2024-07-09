package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// MxedgeTuntermOtherIpConfig represents a MxedgeTuntermOtherIpConfig struct.
type MxedgeTuntermOtherIpConfig struct {
    Ip                   string         `json:"ip"`
    Netmask              string         `json:"netmask"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermOtherIpConfig.
// It customizes the JSON marshaling process for MxedgeTuntermOtherIpConfig objects.
func (m MxedgeTuntermOtherIpConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermOtherIpConfig object to a map representation for JSON marshaling.
func (m MxedgeTuntermOtherIpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["ip"] = m.Ip
    structMap["netmask"] = m.Netmask
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermOtherIpConfig.
// It customizes the JSON unmarshaling process for MxedgeTuntermOtherIpConfig objects.
func (m *MxedgeTuntermOtherIpConfig) UnmarshalJSON(input []byte) error {
    var temp mxedgeTuntermOtherIpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ip", "netmask")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Ip = *temp.Ip
    m.Netmask = *temp.Netmask
    return nil
}

// mxedgeTuntermOtherIpConfig is a temporary struct used for validating the fields of MxedgeTuntermOtherIpConfig.
type mxedgeTuntermOtherIpConfig  struct {
    Ip      *string `json:"ip"`
    Netmask *string `json:"netmask"`
}

func (m *mxedgeTuntermOtherIpConfig) validate() error {
    var errs []string
    if m.Ip == nil {
        errs = append(errs, "required field `ip` is missing for type `Mxedge_Tunterm_Other_Ip_Config`")
    }
    if m.Netmask == nil {
        errs = append(errs, "required field `netmask` is missing for type `Mxedge_Tunterm_Other_Ip_Config`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

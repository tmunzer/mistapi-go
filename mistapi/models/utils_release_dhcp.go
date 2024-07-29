package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UtilsReleaseDhcp represents a UtilsReleaseDhcp struct.
type UtilsReleaseDhcp struct {
    // only for HA. enum: `node0`, `node1`
    Node                 *HaClusterNodeEnum `json:"node,omitempty"`
    // The nework interface on which to release the current DHCP release
    Port                 string             `json:"port"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsReleaseDhcp.
// It customizes the JSON marshaling process for UtilsReleaseDhcp objects.
func (u UtilsReleaseDhcp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsReleaseDhcp object to a map representation for JSON marshaling.
func (u UtilsReleaseDhcp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    structMap["port"] = u.Port
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsReleaseDhcp.
// It customizes the JSON unmarshaling process for UtilsReleaseDhcp objects.
func (u *UtilsReleaseDhcp) UnmarshalJSON(input []byte) error {
    var temp utilsReleaseDhcp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "node", "port")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Node = temp.Node
    u.Port = *temp.Port
    return nil
}

// utilsReleaseDhcp is a temporary struct used for validating the fields of UtilsReleaseDhcp.
type utilsReleaseDhcp  struct {
    Node *HaClusterNodeEnum `json:"node,omitempty"`
    Port *string            `json:"port"`
}

func (u *utilsReleaseDhcp) validate() error {
    var errs []string
    if u.Port == nil {
        errs = append(errs, "required field `port` is missing for type `Utils_Release_Dhcp`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

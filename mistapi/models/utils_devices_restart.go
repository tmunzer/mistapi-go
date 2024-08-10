package models

import (
    "encoding/json"
)

// UtilsDevicesRestart represents a UtilsDevicesRestart struct.
type UtilsDevicesRestart struct {
    // optional for VC member
    Member               *string        `json:"member,omitempty"`
    // only for SSR: if node is not present, both nodes are restarted
    // for other devices: node should not be present
    Node                 *string        `json:"node,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsDevicesRestart.
// It customizes the JSON marshaling process for UtilsDevicesRestart objects.
func (u UtilsDevicesRestart) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsDevicesRestart object to a map representation for JSON marshaling.
func (u UtilsDevicesRestart) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Member != nil {
        structMap["member"] = u.Member
    }
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsDevicesRestart.
// It customizes the JSON unmarshaling process for UtilsDevicesRestart objects.
func (u *UtilsDevicesRestart) UnmarshalJSON(input []byte) error {
    var temp tempUtilsDevicesRestart
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "member", "node")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Member = temp.Member
    u.Node = temp.Node
    return nil
}

// tempUtilsDevicesRestart is a temporary struct used for validating the fields of UtilsDevicesRestart.
type tempUtilsDevicesRestart  struct {
    Member *string `json:"member,omitempty"`
    Node   *string `json:"node,omitempty"`
}

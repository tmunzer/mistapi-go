package models

import (
    "encoding/json"
    "fmt"
)

// UtilsDevicesRestart represents a UtilsDevicesRestart struct.
type UtilsDevicesRestart struct {
    // optional for VC member
    Member               *string                `json:"member,omitempty"`
    // only for SSR: if node is not present, both nodes are restarted
    // for other devices: node should not be present
    Node                 *string                `json:"node,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsDevicesRestart,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsDevicesRestart) String() string {
    return fmt.Sprintf(
    	"UtilsDevicesRestart[Member=%v, Node=%v, AdditionalProperties=%v]",
    	u.Member, u.Node, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsDevicesRestart.
// It customizes the JSON marshaling process for UtilsDevicesRestart objects.
func (u UtilsDevicesRestart) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "member", "node"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsDevicesRestart object to a map representation for JSON marshaling.
func (u UtilsDevicesRestart) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "member", "node")
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

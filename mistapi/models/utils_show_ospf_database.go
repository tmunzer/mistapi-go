package models

import (
    "encoding/json"
    "fmt"
)

// UtilsShowOspfDatabase represents a UtilsShowOspfDatabase struct.
type UtilsShowOspfDatabase struct {
    // only for HA. enum: `node0`, `node1`
    Node                 *HaClusterNodeEnum     `json:"node,omitempty"`
    // Show originating info, default is false
    SelfOriginate        *bool                  `json:"self_originate,omitempty"`
    // VRF name
    Vrf                  *string                `json:"vrf,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsShowOspfDatabase,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsShowOspfDatabase) String() string {
    return fmt.Sprintf(
    	"UtilsShowOspfDatabase[Node=%v, SelfOriginate=%v, Vrf=%v, AdditionalProperties=%v]",
    	u.Node, u.SelfOriginate, u.Vrf, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowOspfDatabase.
// It customizes the JSON marshaling process for UtilsShowOspfDatabase objects.
func (u UtilsShowOspfDatabase) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "node", "self_originate", "vrf"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowOspfDatabase object to a map representation for JSON marshaling.
func (u UtilsShowOspfDatabase) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    if u.SelfOriginate != nil {
        structMap["self_originate"] = u.SelfOriginate
    }
    if u.Vrf != nil {
        structMap["vrf"] = u.Vrf
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowOspfDatabase.
// It customizes the JSON unmarshaling process for UtilsShowOspfDatabase objects.
func (u *UtilsShowOspfDatabase) UnmarshalJSON(input []byte) error {
    var temp tempUtilsShowOspfDatabase
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "node", "self_originate", "vrf")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Node = temp.Node
    u.SelfOriginate = temp.SelfOriginate
    u.Vrf = temp.Vrf
    return nil
}

// tempUtilsShowOspfDatabase is a temporary struct used for validating the fields of UtilsShowOspfDatabase.
type tempUtilsShowOspfDatabase  struct {
    Node          *HaClusterNodeEnum `json:"node,omitempty"`
    SelfOriginate *bool              `json:"self_originate,omitempty"`
    Vrf           *string            `json:"vrf,omitempty"`
}

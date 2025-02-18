package models

import (
    "encoding/json"
    "fmt"
)

// UtilsShowOspfSummary represents a UtilsShowOspfSummary struct.
type UtilsShowOspfSummary struct {
    // only for HA. enum: `node0`, `node1`
    Node                 *HaClusterNodeEnum     `json:"node,omitempty"`
    // VRF name
    Vrf                  *string                `json:"vrf,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsShowOspfSummary,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsShowOspfSummary) String() string {
    return fmt.Sprintf(
    	"UtilsShowOspfSummary[Node=%v, Vrf=%v, AdditionalProperties=%v]",
    	u.Node, u.Vrf, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowOspfSummary.
// It customizes the JSON marshaling process for UtilsShowOspfSummary objects.
func (u UtilsShowOspfSummary) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "node", "vrf"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowOspfSummary object to a map representation for JSON marshaling.
func (u UtilsShowOspfSummary) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    if u.Vrf != nil {
        structMap["vrf"] = u.Vrf
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowOspfSummary.
// It customizes the JSON unmarshaling process for UtilsShowOspfSummary objects.
func (u *UtilsShowOspfSummary) UnmarshalJSON(input []byte) error {
    var temp tempUtilsShowOspfSummary
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "node", "vrf")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Node = temp.Node
    u.Vrf = temp.Vrf
    return nil
}

// tempUtilsShowOspfSummary is a temporary struct used for validating the fields of UtilsShowOspfSummary.
type tempUtilsShowOspfSummary  struct {
    Node *HaClusterNodeEnum `json:"node,omitempty"`
    Vrf  *string            `json:"vrf,omitempty"`
}

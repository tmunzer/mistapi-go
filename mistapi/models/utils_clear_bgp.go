package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// UtilsClearBgp represents a UtilsClearBgp struct.
type UtilsClearBgp struct {
    // neighbor ip-address or 'all'
    Neighbor             string                 `json:"neighbor"`
    // only for HA. enum: `node0`, `node1`
    Node                 *HaClusterNodeEnum     `json:"node,omitempty"`
    // enum: `hard`, `in`, `out`, `soft`
    Type                 UtilsClearBgpTypeEnum  `json:"type"`
    // vrf name
    Vrf                  *string                `json:"vrf,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsClearBgp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsClearBgp) String() string {
    return fmt.Sprintf(
    	"UtilsClearBgp[Neighbor=%v, Node=%v, Type=%v, Vrf=%v, AdditionalProperties=%v]",
    	u.Neighbor, u.Node, u.Type, u.Vrf, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsClearBgp.
// It customizes the JSON marshaling process for UtilsClearBgp objects.
func (u UtilsClearBgp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "neighbor", "node", "type", "vrf"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsClearBgp object to a map representation for JSON marshaling.
func (u UtilsClearBgp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["neighbor"] = u.Neighbor
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    structMap["type"] = u.Type
    if u.Vrf != nil {
        structMap["vrf"] = u.Vrf
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsClearBgp.
// It customizes the JSON unmarshaling process for UtilsClearBgp objects.
func (u *UtilsClearBgp) UnmarshalJSON(input []byte) error {
    var temp tempUtilsClearBgp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "neighbor", "node", "type", "vrf")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Neighbor = *temp.Neighbor
    u.Node = temp.Node
    u.Type = *temp.Type
    u.Vrf = temp.Vrf
    return nil
}

// tempUtilsClearBgp is a temporary struct used for validating the fields of UtilsClearBgp.
type tempUtilsClearBgp  struct {
    Neighbor *string                `json:"neighbor"`
    Node     *HaClusterNodeEnum     `json:"node,omitempty"`
    Type     *UtilsClearBgpTypeEnum `json:"type"`
    Vrf      *string                `json:"vrf,omitempty"`
}

func (u *tempUtilsClearBgp) validate() error {
    var errs []string
    if u.Neighbor == nil {
        errs = append(errs, "required field `neighbor` is missing for type `utils_clear_bgp`")
    }
    if u.Type == nil {
        errs = append(errs, "required field `type` is missing for type `utils_clear_bgp`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

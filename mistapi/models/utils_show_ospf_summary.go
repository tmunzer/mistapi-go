package models

import (
	"encoding/json"
)

// UtilsShowOspfSummary represents a UtilsShowOspfSummary struct.
type UtilsShowOspfSummary struct {
	// only for HA. enum: `node0`, `node1`
	Node *HaClusterNodeEnum `json:"node,omitempty"`
	// VRF name
	Vrf                  *string        `json:"vrf,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowOspfSummary.
// It customizes the JSON marshaling process for UtilsShowOspfSummary objects.
func (u UtilsShowOspfSummary) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowOspfSummary object to a map representation for JSON marshaling.
func (u UtilsShowOspfSummary) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, u.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "node", "vrf")
	if err != nil {
		return err
	}

	u.AdditionalProperties = additionalProperties
	u.Node = temp.Node
	u.Vrf = temp.Vrf
	return nil
}

// tempUtilsShowOspfSummary is a temporary struct used for validating the fields of UtilsShowOspfSummary.
type tempUtilsShowOspfSummary struct {
	Node *HaClusterNodeEnum `json:"node,omitempty"`
	Vrf  *string            `json:"vrf,omitempty"`
}

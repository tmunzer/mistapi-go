package models

import (
	"encoding/json"
)

// CaptureMxedgeMxedges represents a CaptureMxedgeMxedges struct.
// Property key is the Mx Edge ID, currently limited to one mxedge per org capture session
type CaptureMxedgeMxedges struct {
	Interfaces           map[string]CaptureMxedgeMxedgesInterfaces `json:"interfaces,omitempty"`
	AdditionalProperties map[string]any                            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CaptureMxedgeMxedges.
// It customizes the JSON marshaling process for CaptureMxedgeMxedges objects.
func (c CaptureMxedgeMxedges) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CaptureMxedgeMxedges object to a map representation for JSON marshaling.
func (c CaptureMxedgeMxedges) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, c.AdditionalProperties)
	if c.Interfaces != nil {
		structMap["interfaces"] = c.Interfaces
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CaptureMxedgeMxedges.
// It customizes the JSON unmarshaling process for CaptureMxedgeMxedges objects.
func (c *CaptureMxedgeMxedges) UnmarshalJSON(input []byte) error {
	var temp tempCaptureMxedgeMxedges
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "interfaces")
	if err != nil {
		return err
	}

	c.AdditionalProperties = additionalProperties
	c.Interfaces = temp.Interfaces
	return nil
}

// tempCaptureMxedgeMxedges is a temporary struct used for validating the fields of CaptureMxedgeMxedges.
type tempCaptureMxedgeMxedges struct {
	Interfaces map[string]CaptureMxedgeMxedgesInterfaces `json:"interfaces,omitempty"`
}

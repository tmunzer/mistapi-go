package models

import (
	"encoding/json"
)

// DhcpdConfigOption represents a DhcpdConfigOption struct.
type DhcpdConfigOption struct {
	// enum: `boolean`, `hex`, `int16`, `int32`, `ip`, `string`, `uint16`, `uint32`
	Type                 *DhcpdConfigOptionTypeEnum `json:"type,omitempty"`
	Value                *string                    `json:"value,omitempty"`
	AdditionalProperties map[string]any             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DhcpdConfigOption.
// It customizes the JSON marshaling process for DhcpdConfigOption objects.
func (d DhcpdConfigOption) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(d.toMap())
}

// toMap converts the DhcpdConfigOption object to a map representation for JSON marshaling.
func (d DhcpdConfigOption) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, d.AdditionalProperties)
	if d.Type != nil {
		structMap["type"] = d.Type
	}
	if d.Value != nil {
		structMap["value"] = d.Value
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DhcpdConfigOption.
// It customizes the JSON unmarshaling process for DhcpdConfigOption objects.
func (d *DhcpdConfigOption) UnmarshalJSON(input []byte) error {
	var temp tempDhcpdConfigOption
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "type", "value")
	if err != nil {
		return err
	}

	d.AdditionalProperties = additionalProperties
	d.Type = temp.Type
	d.Value = temp.Value
	return nil
}

// tempDhcpdConfigOption is a temporary struct used for validating the fields of DhcpdConfigOption.
type tempDhcpdConfigOption struct {
	Type  *DhcpdConfigOptionTypeEnum `json:"type,omitempty"`
	Value *string                    `json:"value,omitempty"`
}

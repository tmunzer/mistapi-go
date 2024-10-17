package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// PskIdList represents a PskIdList struct.
type PskIdList struct {
	PskIds               []uuid.UUID    `json:"psk_ids,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PskIdList.
// It customizes the JSON marshaling process for PskIdList objects.
func (p PskIdList) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(p.toMap())
}

// toMap converts the PskIdList object to a map representation for JSON marshaling.
func (p PskIdList) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, p.AdditionalProperties)
	if p.PskIds != nil {
		structMap["psk_ids"] = p.PskIds
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PskIdList.
// It customizes the JSON unmarshaling process for PskIdList objects.
func (p *PskIdList) UnmarshalJSON(input []byte) error {
	var temp tempPskIdList
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "psk_ids")
	if err != nil {
		return err
	}

	p.AdditionalProperties = additionalProperties
	p.PskIds = temp.PskIds
	return nil
}

// tempPskIdList is a temporary struct used for validating the fields of PskIdList.
type tempPskIdList struct {
	PskIds []uuid.UUID `json:"psk_ids,omitempty"`
}

package models

import (
	"encoding/json"
)

// MxclusterTuntermExtraRoute represents a MxclusterTuntermExtraRoute struct.
type MxclusterTuntermExtraRoute struct {
	Via                  *string        `json:"via,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxclusterTuntermExtraRoute.
// It customizes the JSON marshaling process for MxclusterTuntermExtraRoute objects.
func (m MxclusterTuntermExtraRoute) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(m.toMap())
}

// toMap converts the MxclusterTuntermExtraRoute object to a map representation for JSON marshaling.
func (m MxclusterTuntermExtraRoute) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Via != nil {
		structMap["via"] = m.Via
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxclusterTuntermExtraRoute.
// It customizes the JSON unmarshaling process for MxclusterTuntermExtraRoute objects.
func (m *MxclusterTuntermExtraRoute) UnmarshalJSON(input []byte) error {
	var temp tempMxclusterTuntermExtraRoute
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "via")
	if err != nil {
		return err
	}

	m.AdditionalProperties = additionalProperties
	m.Via = temp.Via
	return nil
}

// tempMxclusterTuntermExtraRoute is a temporary struct used for validating the fields of MxclusterTuntermExtraRoute.
type tempMxclusterTuntermExtraRoute struct {
	Via *string `json:"via,omitempty"`
}

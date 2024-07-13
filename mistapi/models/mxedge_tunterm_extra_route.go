package models

import (
    "encoding/json"
)

// MxedgeTuntermExtraRoute represents a MxedgeTuntermExtraRoute struct.
type MxedgeTuntermExtraRoute struct {
    Via                  *string        `json:"via,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermExtraRoute.
// It customizes the JSON marshaling process for MxedgeTuntermExtraRoute objects.
func (m MxedgeTuntermExtraRoute) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermExtraRoute object to a map representation for JSON marshaling.
func (m MxedgeTuntermExtraRoute) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Via != nil {
        structMap["via"] = m.Via
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermExtraRoute.
// It customizes the JSON unmarshaling process for MxedgeTuntermExtraRoute objects.
func (m *MxedgeTuntermExtraRoute) UnmarshalJSON(input []byte) error {
    var temp mxedgeTuntermExtraRoute
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

// mxedgeTuntermExtraRoute is a temporary struct used for validating the fields of MxedgeTuntermExtraRoute.
type mxedgeTuntermExtraRoute  struct {
    Via *string `json:"via,omitempty"`
}

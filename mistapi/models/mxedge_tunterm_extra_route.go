package models

import (
    "encoding/json"
    "fmt"
)

// MxedgeTuntermExtraRoute represents a MxedgeTuntermExtraRoute struct.
type MxedgeTuntermExtraRoute struct {
    Via                  *string                `json:"via,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeTuntermExtraRoute,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeTuntermExtraRoute) String() string {
    return fmt.Sprintf(
    	"MxedgeTuntermExtraRoute[Via=%v, AdditionalProperties=%v]",
    	m.Via, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermExtraRoute.
// It customizes the JSON marshaling process for MxedgeTuntermExtraRoute objects.
func (m MxedgeTuntermExtraRoute) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "via"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermExtraRoute object to a map representation for JSON marshaling.
func (m MxedgeTuntermExtraRoute) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Via != nil {
        structMap["via"] = m.Via
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermExtraRoute.
// It customizes the JSON unmarshaling process for MxedgeTuntermExtraRoute objects.
func (m *MxedgeTuntermExtraRoute) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeTuntermExtraRoute
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "via")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Via = temp.Via
    return nil
}

// tempMxedgeTuntermExtraRoute is a temporary struct used for validating the fields of MxedgeTuntermExtraRoute.
type tempMxedgeTuntermExtraRoute  struct {
    Via *string `json:"via,omitempty"`
}

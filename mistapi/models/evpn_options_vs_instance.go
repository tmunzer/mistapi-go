package models

import (
	"encoding/json"
)

// EvpnOptionsVsInstance represents a EvpnOptionsVsInstance struct.
type EvpnOptionsVsInstance struct {
	Networks             []string       `json:"networks,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EvpnOptionsVsInstance.
// It customizes the JSON marshaling process for EvpnOptionsVsInstance objects.
func (e EvpnOptionsVsInstance) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(e.toMap())
}

// toMap converts the EvpnOptionsVsInstance object to a map representation for JSON marshaling.
func (e EvpnOptionsVsInstance) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, e.AdditionalProperties)
	if e.Networks != nil {
		structMap["networks"] = e.Networks
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EvpnOptionsVsInstance.
// It customizes the JSON unmarshaling process for EvpnOptionsVsInstance objects.
func (e *EvpnOptionsVsInstance) UnmarshalJSON(input []byte) error {
	var temp tempEvpnOptionsVsInstance
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "networks")
	if err != nil {
		return err
	}

	e.AdditionalProperties = additionalProperties
	e.Networks = temp.Networks
	return nil
}

// tempEvpnOptionsVsInstance is a temporary struct used for validating the fields of EvpnOptionsVsInstance.
type tempEvpnOptionsVsInstance struct {
	Networks []string `json:"networks,omitempty"`
}

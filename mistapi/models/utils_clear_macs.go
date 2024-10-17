package models

import (
	"encoding/json"
)

// UtilsClearMacs represents a UtilsClearMacs struct.
type UtilsClearMacs struct {
	// a list of ports on which to clear mac addresses. must include logical unit number
	Ports                []string       `json:"ports,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsClearMacs.
// It customizes the JSON marshaling process for UtilsClearMacs objects.
func (u UtilsClearMacs) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UtilsClearMacs object to a map representation for JSON marshaling.
func (u UtilsClearMacs) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, u.AdditionalProperties)
	if u.Ports != nil {
		structMap["ports"] = u.Ports
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsClearMacs.
// It customizes the JSON unmarshaling process for UtilsClearMacs objects.
func (u *UtilsClearMacs) UnmarshalJSON(input []byte) error {
	var temp tempUtilsClearMacs
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "ports")
	if err != nil {
		return err
	}

	u.AdditionalProperties = additionalProperties
	u.Ports = temp.Ports
	return nil
}

// tempUtilsClearMacs is a temporary struct used for validating the fields of UtilsClearMacs.
type tempUtilsClearMacs struct {
	Ports []string `json:"ports,omitempty"`
}

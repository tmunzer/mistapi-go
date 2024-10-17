package models

import (
	"encoding/json"
)

// LatlngTl represents a LatlngTl struct.
// when type=google, latitude / longitude of the top-left corner
type LatlngTl struct {
	Lat                  *string        `json:"lat,omitempty"`
	Lng                  *string        `json:"lng,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for LatlngTl.
// It customizes the JSON marshaling process for LatlngTl objects.
func (l LatlngTl) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the LatlngTl object to a map representation for JSON marshaling.
func (l LatlngTl) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, l.AdditionalProperties)
	if l.Lat != nil {
		structMap["lat"] = l.Lat
	}
	if l.Lng != nil {
		structMap["lng"] = l.Lng
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for LatlngTl.
// It customizes the JSON unmarshaling process for LatlngTl objects.
func (l *LatlngTl) UnmarshalJSON(input []byte) error {
	var temp tempLatlngTl
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "lat", "lng")
	if err != nil {
		return err
	}

	l.AdditionalProperties = additionalProperties
	l.Lat = temp.Lat
	l.Lng = temp.Lng
	return nil
}

// tempLatlngTl is a temporary struct used for validating the fields of LatlngTl.
type tempLatlngTl struct {
	Lat *string `json:"lat,omitempty"`
	Lng *string `json:"lng,omitempty"`
}

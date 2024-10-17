package models

import (
	"encoding/json"
)

// LastTrouble represents a LastTrouble struct.
// last trouble code of switch
type LastTrouble struct {
	// Code definitions list at /api/v1/consts/ap_led_status
	Code                 *string        `json:"code,omitempty"`
	Timestamp            *int           `json:"timestamp,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for LastTrouble.
// It customizes the JSON marshaling process for LastTrouble objects.
func (l LastTrouble) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(l.toMap())
}

// toMap converts the LastTrouble object to a map representation for JSON marshaling.
func (l LastTrouble) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, l.AdditionalProperties)
	if l.Code != nil {
		structMap["code"] = l.Code
	}
	if l.Timestamp != nil {
		structMap["timestamp"] = l.Timestamp
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for LastTrouble.
// It customizes the JSON unmarshaling process for LastTrouble objects.
func (l *LastTrouble) UnmarshalJSON(input []byte) error {
	var temp tempLastTrouble
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "code", "timestamp")
	if err != nil {
		return err
	}

	l.AdditionalProperties = additionalProperties
	l.Code = temp.Code
	l.Timestamp = temp.Timestamp
	return nil
}

// tempLastTrouble is a temporary struct used for validating the fields of LastTrouble.
type tempLastTrouble struct {
	Code      *string `json:"code,omitempty"`
	Timestamp *int    `json:"timestamp,omitempty"`
}

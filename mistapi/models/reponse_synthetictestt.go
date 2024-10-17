package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// ReponseSynthetictest represents a ReponseSynthetictest struct.
type ReponseSynthetictest struct {
	Id                   *uuid.UUID     `json:"id,omitempty"`
	Message              *string        `json:"message,omitempty"`
	Status               *string        `json:"status,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ReponseSynthetictest.
// It customizes the JSON marshaling process for ReponseSynthetictest objects.
func (r ReponseSynthetictest) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ReponseSynthetictest object to a map representation for JSON marshaling.
func (r ReponseSynthetictest) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Id != nil {
		structMap["id"] = r.Id
	}
	if r.Message != nil {
		structMap["message"] = r.Message
	}
	if r.Status != nil {
		structMap["status"] = r.Status
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReponseSynthetictest.
// It customizes the JSON unmarshaling process for ReponseSynthetictest objects.
func (r *ReponseSynthetictest) UnmarshalJSON(input []byte) error {
	var temp tempReponseSynthetictest
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "message", "status")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.Id = temp.Id
	r.Message = temp.Message
	r.Status = temp.Status
	return nil
}

// tempReponseSynthetictest is a temporary struct used for validating the fields of ReponseSynthetictest.
type tempReponseSynthetictest struct {
	Id      *uuid.UUID `json:"id,omitempty"`
	Message *string    `json:"message,omitempty"`
	Status  *string    `json:"status,omitempty"`
}

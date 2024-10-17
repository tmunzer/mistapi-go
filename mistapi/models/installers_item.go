package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// InstallersItem represents a InstallersItem struct.
type InstallersItem struct {
	Id                   *uuid.UUID     `json:"id,omitempty"`
	Name                 *string        `json:"name,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InstallersItem.
// It customizes the JSON marshaling process for InstallersItem objects.
func (i InstallersItem) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the InstallersItem object to a map representation for JSON marshaling.
func (i InstallersItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, i.AdditionalProperties)
	if i.Id != nil {
		structMap["id"] = i.Id
	}
	if i.Name != nil {
		structMap["name"] = i.Name
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InstallersItem.
// It customizes the JSON unmarshaling process for InstallersItem objects.
func (i *InstallersItem) UnmarshalJSON(input []byte) error {
	var temp tempInstallersItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "name")
	if err != nil {
		return err
	}

	i.AdditionalProperties = additionalProperties
	i.Id = temp.Id
	i.Name = temp.Name
	return nil
}

// tempInstallersItem is a temporary struct used for validating the fields of InstallersItem.
type tempInstallersItem struct {
	Id   *uuid.UUID `json:"id,omitempty"`
	Name *string    `json:"name,omitempty"`
}

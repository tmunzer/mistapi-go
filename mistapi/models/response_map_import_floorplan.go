package models

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
)

// ResponseMapImportFloorplan represents a ResponseMapImportFloorplan struct.
type ResponseMapImportFloorplan struct {
	Action               string         `json:"action"`
	Id                   uuid.UUID      `json:"id"`
	MapId                uuid.UUID      `json:"map_id"`
	Name                 string         `json:"name"`
	Reason               *string        `json:"reason,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseMapImportFloorplan.
// It customizes the JSON marshaling process for ResponseMapImportFloorplan objects.
func (r ResponseMapImportFloorplan) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseMapImportFloorplan object to a map representation for JSON marshaling.
func (r ResponseMapImportFloorplan) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["action"] = r.Action
	structMap["id"] = r.Id
	structMap["map_id"] = r.MapId
	structMap["name"] = r.Name
	if r.Reason != nil {
		structMap["reason"] = r.Reason
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseMapImportFloorplan.
// It customizes the JSON unmarshaling process for ResponseMapImportFloorplan objects.
func (r *ResponseMapImportFloorplan) UnmarshalJSON(input []byte) error {
	var temp tempResponseMapImportFloorplan
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "action", "id", "map_id", "name", "reason")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.Action = *temp.Action
	r.Id = *temp.Id
	r.MapId = *temp.MapId
	r.Name = *temp.Name
	r.Reason = temp.Reason
	return nil
}

// tempResponseMapImportFloorplan is a temporary struct used for validating the fields of ResponseMapImportFloorplan.
type tempResponseMapImportFloorplan struct {
	Action *string    `json:"action"`
	Id     *uuid.UUID `json:"id"`
	MapId  *uuid.UUID `json:"map_id"`
	Name   *string    `json:"name"`
	Reason *string    `json:"reason,omitempty"`
}

func (r *tempResponseMapImportFloorplan) validate() error {
	var errs []string
	if r.Action == nil {
		errs = append(errs, "required field `action` is missing for type `response_map_import_floorplan`")
	}
	if r.Id == nil {
		errs = append(errs, "required field `id` is missing for type `response_map_import_floorplan`")
	}
	if r.MapId == nil {
		errs = append(errs, "required field `map_id` is missing for type `response_map_import_floorplan`")
	}
	if r.Name == nil {
		errs = append(errs, "required field `name` is missing for type `response_map_import_floorplan`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// ResponseMapImportFloorplan represents a ResponseMapImportFloorplan struct.
type ResponseMapImportFloorplan struct {
    Action               string                 `json:"action"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   uuid.UUID              `json:"id"`
    MapId                uuid.UUID              `json:"map_id"`
    Name                 string                 `json:"name"`
    Reason               *string                `json:"reason,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseMapImportFloorplan,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseMapImportFloorplan) String() string {
    return fmt.Sprintf(
    	"ResponseMapImportFloorplan[Action=%v, Id=%v, MapId=%v, Name=%v, Reason=%v, AdditionalProperties=%v]",
    	r.Action, r.Id, r.MapId, r.Name, r.Reason, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseMapImportFloorplan.
// It customizes the JSON marshaling process for ResponseMapImportFloorplan objects.
func (r ResponseMapImportFloorplan) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "action", "id", "map_id", "name", "reason"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseMapImportFloorplan object to a map representation for JSON marshaling.
func (r ResponseMapImportFloorplan) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "action", "id", "map_id", "name", "reason")
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
type tempResponseMapImportFloorplan  struct {
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
    return errors.New(strings.Join (errs, "\n"))
}

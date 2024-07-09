package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// ResponseMapImportAp represents a ResponseMapImportAp struct.
type ResponseMapImportAp struct {
    Action               ResponseMapImportApActionEnum `json:"action"`
    FloorplanId          uuid.UUID                     `json:"floorplan_id"`
    Height               *float64                      `json:"height,omitempty"`
    Mac                  string                        `json:"mac"`
    MapId                uuid.UUID                     `json:"map_id"`
    Orientation          int                           `json:"orientation"`
    Reason               *string                       `json:"reason,omitempty"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseMapImportAp.
// It customizes the JSON marshaling process for ResponseMapImportAp objects.
func (r ResponseMapImportAp) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseMapImportAp object to a map representation for JSON marshaling.
func (r ResponseMapImportAp) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["action"] = r.Action
    structMap["floorplan_id"] = r.FloorplanId
    if r.Height != nil {
        structMap["height"] = r.Height
    }
    structMap["mac"] = r.Mac
    structMap["map_id"] = r.MapId
    structMap["orientation"] = r.Orientation
    if r.Reason != nil {
        structMap["reason"] = r.Reason
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseMapImportAp.
// It customizes the JSON unmarshaling process for ResponseMapImportAp objects.
func (r *ResponseMapImportAp) UnmarshalJSON(input []byte) error {
    var temp responseMapImportAp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "action", "floorplan_id", "height", "mac", "map_id", "orientation", "reason")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Action = *temp.Action
    r.FloorplanId = *temp.FloorplanId
    r.Height = temp.Height
    r.Mac = *temp.Mac
    r.MapId = *temp.MapId
    r.Orientation = *temp.Orientation
    r.Reason = temp.Reason
    return nil
}

// responseMapImportAp is a temporary struct used for validating the fields of ResponseMapImportAp.
type responseMapImportAp  struct {
    Action      *ResponseMapImportApActionEnum `json:"action"`
    FloorplanId *uuid.UUID                     `json:"floorplan_id"`
    Height      *float64                       `json:"height,omitempty"`
    Mac         *string                        `json:"mac"`
    MapId       *uuid.UUID                     `json:"map_id"`
    Orientation *int                           `json:"orientation"`
    Reason      *string                        `json:"reason,omitempty"`
}

func (r *responseMapImportAp) validate() error {
    var errs []string
    if r.Action == nil {
        errs = append(errs, "required field `action` is missing for type `Response_Map_Import_Ap`")
    }
    if r.FloorplanId == nil {
        errs = append(errs, "required field `floorplan_id` is missing for type `Response_Map_Import_Ap`")
    }
    if r.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Response_Map_Import_Ap`")
    }
    if r.MapId == nil {
        errs = append(errs, "required field `map_id` is missing for type `Response_Map_Import_Ap`")
    }
    if r.Orientation == nil {
        errs = append(errs, "required field `orientation` is missing for type `Response_Map_Import_Ap`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// ResponseMapImportAp represents a ResponseMapImportAp struct.
type ResponseMapImportAp struct {
    // enum: `assigned-named-placed`, `assigned-placed`, `ignored`, `named-placed`, `placed`
    Action               ResponseMapImportApActionEnum `json:"action"`
    FloorplanId          uuid.UUID                     `json:"floorplan_id"`
    Height               *float64                      `json:"height,omitempty"`
    Mac                  string                        `json:"mac"`
    MapId                uuid.UUID                     `json:"map_id"`
    Orientation          int                           `json:"orientation"`
    Reason               *string                       `json:"reason,omitempty"`
    AdditionalProperties map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseMapImportAp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseMapImportAp) String() string {
    return fmt.Sprintf(
    	"ResponseMapImportAp[Action=%v, FloorplanId=%v, Height=%v, Mac=%v, MapId=%v, Orientation=%v, Reason=%v, AdditionalProperties=%v]",
    	r.Action, r.FloorplanId, r.Height, r.Mac, r.MapId, r.Orientation, r.Reason, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseMapImportAp.
// It customizes the JSON marshaling process for ResponseMapImportAp objects.
func (r ResponseMapImportAp) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "action", "floorplan_id", "height", "mac", "map_id", "orientation", "reason"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseMapImportAp object to a map representation for JSON marshaling.
func (r ResponseMapImportAp) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
    var temp tempResponseMapImportAp
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "action", "floorplan_id", "height", "mac", "map_id", "orientation", "reason")
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

// tempResponseMapImportAp is a temporary struct used for validating the fields of ResponseMapImportAp.
type tempResponseMapImportAp  struct {
    Action      *ResponseMapImportApActionEnum `json:"action"`
    FloorplanId *uuid.UUID                     `json:"floorplan_id"`
    Height      *float64                       `json:"height,omitempty"`
    Mac         *string                        `json:"mac"`
    MapId       *uuid.UUID                     `json:"map_id"`
    Orientation *int                           `json:"orientation"`
    Reason      *string                        `json:"reason,omitempty"`
}

func (r *tempResponseMapImportAp) validate() error {
    var errs []string
    if r.Action == nil {
        errs = append(errs, "required field `action` is missing for type `response_map_import_ap`")
    }
    if r.FloorplanId == nil {
        errs = append(errs, "required field `floorplan_id` is missing for type `response_map_import_ap`")
    }
    if r.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `response_map_import_ap`")
    }
    if r.MapId == nil {
        errs = append(errs, "required field `map_id` is missing for type `response_map_import_ap`")
    }
    if r.Orientation == nil {
        errs = append(errs, "required field `orientation` is missing for type `response_map_import_ap`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

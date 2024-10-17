package models

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
)

// ResponseMapImport represents a ResponseMapImport struct.
type ResponseMapImport struct {
	Aps                  []ResponseMapImportAp        `json:"aps"`
	Floorplans           []ResponseMapImportFloorplan `json:"floorplans"`
	ForSite              *bool                        `json:"for_site,omitempty"`
	SiteId               uuid.UUID                    `json:"site_id"`
	Summary              ResponseMapImportSummary     `json:"summary"`
	AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseMapImport.
// It customizes the JSON marshaling process for ResponseMapImport objects.
func (r ResponseMapImport) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseMapImport object to a map representation for JSON marshaling.
func (r ResponseMapImport) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["aps"] = r.Aps
	structMap["floorplans"] = r.Floorplans
	if r.ForSite != nil {
		structMap["for_site"] = r.ForSite
	}
	structMap["site_id"] = r.SiteId
	structMap["summary"] = r.Summary.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseMapImport.
// It customizes the JSON unmarshaling process for ResponseMapImport objects.
func (r *ResponseMapImport) UnmarshalJSON(input []byte) error {
	var temp tempResponseMapImport
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "aps", "floorplans", "for_site", "site_id", "summary")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.Aps = *temp.Aps
	r.Floorplans = *temp.Floorplans
	r.ForSite = temp.ForSite
	r.SiteId = *temp.SiteId
	r.Summary = *temp.Summary
	return nil
}

// tempResponseMapImport is a temporary struct used for validating the fields of ResponseMapImport.
type tempResponseMapImport struct {
	Aps        *[]ResponseMapImportAp        `json:"aps"`
	Floorplans *[]ResponseMapImportFloorplan `json:"floorplans"`
	ForSite    *bool                         `json:"for_site,omitempty"`
	SiteId     *uuid.UUID                    `json:"site_id"`
	Summary    *ResponseMapImportSummary     `json:"summary"`
}

func (r *tempResponseMapImport) validate() error {
	var errs []string
	if r.Aps == nil {
		errs = append(errs, "required field `aps` is missing for type `response_map_import`")
	}
	if r.Floorplans == nil {
		errs = append(errs, "required field `floorplans` is missing for type `response_map_import`")
	}
	if r.SiteId == nil {
		errs = append(errs, "required field `site_id` is missing for type `response_map_import`")
	}
	if r.Summary == nil {
		errs = append(errs, "required field `summary` is missing for type `response_map_import`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

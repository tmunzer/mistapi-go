package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ResponseMapImportSummary represents a ResponseMapImportSummary struct.
type ResponseMapImportSummary struct {
	NumApAssigned        int            `json:"num_ap_assigned"`
	NumInvAssigned       int            `json:"num_inv_assigned"`
	NumMapAssigned       int            `json:"num_map_assigned"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseMapImportSummary.
// It customizes the JSON marshaling process for ResponseMapImportSummary objects.
func (r ResponseMapImportSummary) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseMapImportSummary object to a map representation for JSON marshaling.
func (r ResponseMapImportSummary) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["num_ap_assigned"] = r.NumApAssigned
	structMap["num_inv_assigned"] = r.NumInvAssigned
	structMap["num_map_assigned"] = r.NumMapAssigned
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseMapImportSummary.
// It customizes the JSON unmarshaling process for ResponseMapImportSummary objects.
func (r *ResponseMapImportSummary) UnmarshalJSON(input []byte) error {
	var temp tempResponseMapImportSummary
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "num_ap_assigned", "num_inv_assigned", "num_map_assigned")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.NumApAssigned = *temp.NumApAssigned
	r.NumInvAssigned = *temp.NumInvAssigned
	r.NumMapAssigned = *temp.NumMapAssigned
	return nil
}

// tempResponseMapImportSummary is a temporary struct used for validating the fields of ResponseMapImportSummary.
type tempResponseMapImportSummary struct {
	NumApAssigned  *int `json:"num_ap_assigned"`
	NumInvAssigned *int `json:"num_inv_assigned"`
	NumMapAssigned *int `json:"num_map_assigned"`
}

func (r *tempResponseMapImportSummary) validate() error {
	var errs []string
	if r.NumApAssigned == nil {
		errs = append(errs, "required field `num_ap_assigned` is missing for type `response_map_import_summary`")
	}
	if r.NumInvAssigned == nil {
		errs = append(errs, "required field `num_inv_assigned` is missing for type `response_map_import_summary`")
	}
	if r.NumMapAssigned == nil {
		errs = append(errs, "required field `num_map_assigned` is missing for type `response_map_import_summary`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

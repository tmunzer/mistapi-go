// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseMapImportSummary represents a ResponseMapImportSummary struct.
type ResponseMapImportSummary struct {
    NumApAssigned        int                    `json:"num_ap_assigned"`
    NumInvAssigned       int                    `json:"num_inv_assigned"`
    NumMapAssigned       int                    `json:"num_map_assigned"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseMapImportSummary,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseMapImportSummary) String() string {
    return fmt.Sprintf(
    	"ResponseMapImportSummary[NumApAssigned=%v, NumInvAssigned=%v, NumMapAssigned=%v, AdditionalProperties=%v]",
    	r.NumApAssigned, r.NumInvAssigned, r.NumMapAssigned, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseMapImportSummary.
// It customizes the JSON marshaling process for ResponseMapImportSummary objects.
func (r ResponseMapImportSummary) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "num_ap_assigned", "num_inv_assigned", "num_map_assigned"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseMapImportSummary object to a map representation for JSON marshaling.
func (r ResponseMapImportSummary) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "num_ap_assigned", "num_inv_assigned", "num_map_assigned")
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
type tempResponseMapImportSummary  struct {
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
    return errors.New(strings.Join (errs, "\n"))
}

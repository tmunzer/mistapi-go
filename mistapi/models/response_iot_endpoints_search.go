// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// ResponseIotEndpointsSearch represents a ResponseIotEndpointsSearch struct.
type ResponseIotEndpointsSearch struct {
	End                  float64                `json:"end"`
	Results              []IotendpointStats     `json:"results"`
	Start                float64                `json:"start"`
	Total                int                    `json:"total"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseIotEndpointsSearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseIotEndpointsSearch) String() string {
	return fmt.Sprintf(
		"ResponseIotEndpointsSearch[End=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		r.End, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseIotEndpointsSearch.
// It customizes the JSON marshaling process for ResponseIotEndpointsSearch objects.
func (r ResponseIotEndpointsSearch) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"end", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseIotEndpointsSearch object to a map representation for JSON marshaling.
func (r ResponseIotEndpointsSearch) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["end"] = r.End
	structMap["results"] = r.Results
	structMap["start"] = r.Start
	structMap["total"] = r.Total
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseIotEndpointsSearch.
// It customizes the JSON unmarshaling process for ResponseIotEndpointsSearch objects.
func (r *ResponseIotEndpointsSearch) UnmarshalJSON(input []byte) error {
	var temp tempResponseIotEndpointsSearch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "results", "start", "total")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.End = *temp.End
	r.Results = *temp.Results
	r.Start = *temp.Start
	r.Total = *temp.Total
	return nil
}

// tempResponseIotEndpointsSearch is a temporary struct used for validating the fields of ResponseIotEndpointsSearch.
type tempResponseIotEndpointsSearch struct {
	End     *float64            `json:"end"`
	Results *[]IotendpointStats `json:"results"`
	Start   *float64            `json:"start"`
	Total   *int                `json:"total"`
}

func (r *tempResponseIotEndpointsSearch) validate() error {
	var errs []string
	if r.End == nil {
		errs = append(errs, "required field `end` is missing for type `response_iot_endpoints_search`")
	}
	if r.Results == nil {
		errs = append(errs, "required field `results` is missing for type `response_iot_endpoints_search`")
	}
	if r.Start == nil {
		errs = append(errs, "required field `start` is missing for type `response_iot_endpoints_search`")
	}
	if r.Total == nil {
		errs = append(errs, "required field `total` is missing for type `response_iot_endpoints_search`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

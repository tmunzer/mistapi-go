// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponsePastSpectrumAnalysis represents a ResponsePastSpectrumAnalysis struct.
// Paginated response for past site spectrum analysis records
type ResponsePastSpectrumAnalysis struct {
	// Epoch timestamp, in seconds, for the end of the spectrum analysis query window
	End *int `json:"end,omitempty"`
	// Maximum number of spectrum analysis records returned in this page
	Limit *int `json:"limit,omitempty"`
	// Current page number returned for the spectrum analysis results
	Page *int `json:"page,omitempty"`
	// Past spectrum analysis records returned for the site
	Results []ResponsePastSpectrumAnalysisResult `json:"results,omitempty"`
	// Epoch timestamp, in seconds, for the start of the spectrum analysis query window
	Start *int `json:"start,omitempty"`
	// Number of spectrum analysis records available for the given time range
	Total                *int                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponsePastSpectrumAnalysis,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponsePastSpectrumAnalysis) String() string {
	return fmt.Sprintf(
		"ResponsePastSpectrumAnalysis[End=%v, Limit=%v, Page=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		r.End, r.Limit, r.Page, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponsePastSpectrumAnalysis.
// It customizes the JSON marshaling process for ResponsePastSpectrumAnalysis objects.
func (r ResponsePastSpectrumAnalysis) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"end", "limit", "page", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponsePastSpectrumAnalysis object to a map representation for JSON marshaling.
func (r ResponsePastSpectrumAnalysis) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.End != nil {
		structMap["end"] = r.End
	}
	if r.Limit != nil {
		structMap["limit"] = r.Limit
	}
	if r.Page != nil {
		structMap["page"] = r.Page
	}
	if r.Results != nil {
		structMap["results"] = r.Results
	}
	if r.Start != nil {
		structMap["start"] = r.Start
	}
	if r.Total != nil {
		structMap["total"] = r.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponsePastSpectrumAnalysis.
// It customizes the JSON unmarshaling process for ResponsePastSpectrumAnalysis objects.
func (r *ResponsePastSpectrumAnalysis) UnmarshalJSON(input []byte) error {
	var temp tempResponsePastSpectrumAnalysis
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "page", "results", "start", "total")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.End = temp.End
	r.Limit = temp.Limit
	r.Page = temp.Page
	r.Results = temp.Results
	r.Start = temp.Start
	r.Total = temp.Total
	return nil
}

// tempResponsePastSpectrumAnalysis is a temporary struct used for validating the fields of ResponsePastSpectrumAnalysis.
type tempResponsePastSpectrumAnalysis struct {
	End     *int                                 `json:"end,omitempty"`
	Limit   *int                                 `json:"limit,omitempty"`
	Page    *int                                 `json:"page,omitempty"`
	Results []ResponsePastSpectrumAnalysisResult `json:"results,omitempty"`
	Start   *int                                 `json:"start,omitempty"`
	Total   *int                                 `json:"total,omitempty"`
}

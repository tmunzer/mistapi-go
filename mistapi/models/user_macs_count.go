// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// UserMacsCount represents a UserMacsCount struct.
// User MACs count response
type UserMacsCount struct {
	// End timestamp
	End *int `json:"end,omitempty"`
	// Number of results to return
	Limit *int `json:"limit,omitempty"`
	// List of user MAC entries
	Results []UserMac `json:"results,omitempty"`
	// Start timestamp
	Start *int `json:"start,omitempty"`
	// Total number of results
	Total                *int                   `json:"total,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UserMacsCount,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UserMacsCount) String() string {
	return fmt.Sprintf(
		"UserMacsCount[End=%v, Limit=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
		u.End, u.Limit, u.Results, u.Start, u.Total, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UserMacsCount.
// It customizes the JSON marshaling process for UserMacsCount objects.
func (u UserMacsCount) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"end", "limit", "results", "start", "total"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UserMacsCount object to a map representation for JSON marshaling.
func (u UserMacsCount) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.End != nil {
		structMap["end"] = u.End
	}
	if u.Limit != nil {
		structMap["limit"] = u.Limit
	}
	if u.Results != nil {
		structMap["results"] = u.Results
	}
	if u.Start != nil {
		structMap["start"] = u.Start
	}
	if u.Total != nil {
		structMap["total"] = u.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UserMacsCount.
// It customizes the JSON unmarshaling process for UserMacsCount objects.
func (u *UserMacsCount) UnmarshalJSON(input []byte) error {
	var temp tempUserMacsCount
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "results", "start", "total")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.End = temp.End
	u.Limit = temp.Limit
	u.Results = temp.Results
	u.Start = temp.Start
	u.Total = temp.Total
	return nil
}

// tempUserMacsCount is a temporary struct used for validating the fields of UserMacsCount.
type tempUserMacsCount struct {
	End     *int      `json:"end,omitempty"`
	Limit   *int      `json:"limit,omitempty"`
	Results []UserMac `json:"results,omitempty"`
	Start   *int      `json:"start,omitempty"`
	Total   *int      `json:"total,omitempty"`
}

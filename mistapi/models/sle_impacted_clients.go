// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SleImpactedClients represents a SleImpactedClients struct.
// Paginated list of clients impacted by an SLE metric
type SleImpactedClients struct {
	// Requested SLE classifier filter applied to the query
	Classifier *string `json:"classifier,omitempty"`
	// Impacted client rows returned for an SLE query
	Clients []SleImpactedClientsClient `json:"clients,omitempty"`
	// Last timestamp in the impacted clients window
	End *int `json:"end,omitempty"`
	// Requested SLE failure filter applied to the query
	Failure *string `json:"failure,omitempty"`
	// Maximum number of impacted client rows returned per page
	Limit *int `json:"limit,omitempty"`
	// SLE metric name used for the impacted clients query
	Metric *string `json:"metric,omitempty"`
	// Current page number for impacted client results
	Page *int `json:"page,omitempty"`
	// First timestamp in the impacted clients window
	Start *int `json:"start,omitempty"`
	// Number of impacted client rows matching the query
	TotalCount           *int                   `json:"total_count,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactedClients,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactedClients) String() string {
	return fmt.Sprintf(
		"SleImpactedClients[Classifier=%v, Clients=%v, End=%v, Failure=%v, Limit=%v, Metric=%v, Page=%v, Start=%v, TotalCount=%v, AdditionalProperties=%v]",
		s.Classifier, s.Clients, s.End, s.Failure, s.Limit, s.Metric, s.Page, s.Start, s.TotalCount, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactedClients.
// It customizes the JSON marshaling process for SleImpactedClients objects.
func (s SleImpactedClients) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"classifier", "clients", "end", "failure", "limit", "metric", "page", "start", "total_count"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SleImpactedClients object to a map representation for JSON marshaling.
func (s SleImpactedClients) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Classifier != nil {
		structMap["classifier"] = s.Classifier
	}
	if s.Clients != nil {
		structMap["clients"] = s.Clients
	}
	if s.End != nil {
		structMap["end"] = s.End
	}
	if s.Failure != nil {
		structMap["failure"] = s.Failure
	}
	if s.Limit != nil {
		structMap["limit"] = s.Limit
	}
	if s.Metric != nil {
		structMap["metric"] = s.Metric
	}
	if s.Page != nil {
		structMap["page"] = s.Page
	}
	if s.Start != nil {
		structMap["start"] = s.Start
	}
	if s.TotalCount != nil {
		structMap["total_count"] = s.TotalCount
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactedClients.
// It customizes the JSON unmarshaling process for SleImpactedClients objects.
func (s *SleImpactedClients) UnmarshalJSON(input []byte) error {
	var temp tempSleImpactedClients
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "classifier", "clients", "end", "failure", "limit", "metric", "page", "start", "total_count")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Classifier = temp.Classifier
	s.Clients = temp.Clients
	s.End = temp.End
	s.Failure = temp.Failure
	s.Limit = temp.Limit
	s.Metric = temp.Metric
	s.Page = temp.Page
	s.Start = temp.Start
	s.TotalCount = temp.TotalCount
	return nil
}

// tempSleImpactedClients is a temporary struct used for validating the fields of SleImpactedClients.
type tempSleImpactedClients struct {
	Classifier *string                    `json:"classifier,omitempty"`
	Clients    []SleImpactedClientsClient `json:"clients,omitempty"`
	End        *int                       `json:"end,omitempty"`
	Failure    *string                    `json:"failure,omitempty"`
	Limit      *int                       `json:"limit,omitempty"`
	Metric     *string                    `json:"metric,omitempty"`
	Page       *int                       `json:"page,omitempty"`
	Start      *int                       `json:"start,omitempty"`
	TotalCount *int                       `json:"total_count,omitempty"`
}

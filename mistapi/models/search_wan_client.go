package models

import (
	"encoding/json"
)

// SearchWanClient represents a SearchWanClient struct.
type SearchWanClient struct {
	End                  *int             `json:"end,omitempty"`
	Limit                *int             `json:"limit,omitempty"`
	Results              []StatsWanClient `json:"results,omitempty"`
	Start                *int             `json:"start,omitempty"`
	Total                *int             `json:"total,omitempty"`
	AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SearchWanClient.
// It customizes the JSON marshaling process for SearchWanClient objects.
func (s SearchWanClient) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SearchWanClient object to a map representation for JSON marshaling.
func (s SearchWanClient) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.End != nil {
		structMap["end"] = s.End
	}
	if s.Limit != nil {
		structMap["limit"] = s.Limit
	}
	if s.Results != nil {
		structMap["results"] = s.Results
	}
	if s.Start != nil {
		structMap["start"] = s.Start
	}
	if s.Total != nil {
		structMap["total"] = s.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SearchWanClient.
// It customizes the JSON unmarshaling process for SearchWanClient objects.
func (s *SearchWanClient) UnmarshalJSON(input []byte) error {
	var temp tempSearchWanClient
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "end", "limit", "results", "start", "total")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.End = temp.End
	s.Limit = temp.Limit
	s.Results = temp.Results
	s.Start = temp.Start
	s.Total = temp.Total
	return nil
}

// tempSearchWanClient is a temporary struct used for validating the fields of SearchWanClient.
type tempSearchWanClient struct {
	End     *int             `json:"end,omitempty"`
	Limit   *int             `json:"limit,omitempty"`
	Results []StatsWanClient `json:"results,omitempty"`
	Start   *int             `json:"start,omitempty"`
	Total   *int             `json:"total,omitempty"`
}

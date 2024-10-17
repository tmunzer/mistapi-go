package models

import (
	"encoding/json"
)

// ResponseMxedgeEventsSearch represents a ResponseMxedgeEventsSearch struct.
type ResponseMxedgeEventsSearch struct {
	End                  *int           `json:"end,omitempty"`
	Limit                *int           `json:"limit,omitempty"`
	Page                 *int           `json:"page,omitempty"`
	Results              []MxedgeEvent  `json:"results,omitempty"`
	Start                *int           `json:"start,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseMxedgeEventsSearch.
// It customizes the JSON marshaling process for ResponseMxedgeEventsSearch objects.
func (r ResponseMxedgeEventsSearch) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseMxedgeEventsSearch object to a map representation for JSON marshaling.
func (r ResponseMxedgeEventsSearch) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
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
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseMxedgeEventsSearch.
// It customizes the JSON unmarshaling process for ResponseMxedgeEventsSearch objects.
func (r *ResponseMxedgeEventsSearch) UnmarshalJSON(input []byte) error {
	var temp tempResponseMxedgeEventsSearch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "end", "limit", "page", "results", "start")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.End = temp.End
	r.Limit = temp.Limit
	r.Page = temp.Page
	r.Results = temp.Results
	r.Start = temp.Start
	return nil
}

// tempResponseMxedgeEventsSearch is a temporary struct used for validating the fields of ResponseMxedgeEventsSearch.
type tempResponseMxedgeEventsSearch struct {
	End     *int          `json:"end,omitempty"`
	Limit   *int          `json:"limit,omitempty"`
	Page    *int          `json:"page,omitempty"`
	Results []MxedgeEvent `json:"results,omitempty"`
	Start   *int          `json:"start,omitempty"`
}

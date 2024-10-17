package models

import (
	"encoding/json"
)

// ResponseSearchBgps represents a ResponseSearchBgps struct.
type ResponseSearchBgps struct {
	End                  *float64       `json:"end,omitempty"`
	Limit                *int           `json:"limit,omitempty"`
	Results              []BgpStats     `json:"results,omitempty"`
	Start                *float64       `json:"start,omitempty"`
	Total                *int           `json:"total,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSearchBgps.
// It customizes the JSON marshaling process for ResponseSearchBgps objects.
func (r ResponseSearchBgps) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSearchBgps object to a map representation for JSON marshaling.
func (r ResponseSearchBgps) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.End != nil {
		structMap["end"] = r.End
	}
	if r.Limit != nil {
		structMap["limit"] = r.Limit
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSearchBgps.
// It customizes the JSON unmarshaling process for ResponseSearchBgps objects.
func (r *ResponseSearchBgps) UnmarshalJSON(input []byte) error {
	var temp tempResponseSearchBgps
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "end", "limit", "results", "start", "total")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.End = temp.End
	r.Limit = temp.Limit
	r.Results = temp.Results
	r.Start = temp.Start
	r.Total = temp.Total
	return nil
}

// tempResponseSearchBgps is a temporary struct used for validating the fields of ResponseSearchBgps.
type tempResponseSearchBgps struct {
	End     *float64   `json:"end,omitempty"`
	Limit   *int       `json:"limit,omitempty"`
	Results []BgpStats `json:"results,omitempty"`
	Start   *float64   `json:"start,omitempty"`
	Total   *int       `json:"total,omitempty"`
}

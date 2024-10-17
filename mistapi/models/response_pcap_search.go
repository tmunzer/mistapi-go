package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ResponsePcapSearch represents a ResponsePcapSearch struct.
type ResponsePcapSearch struct {
	End                  int                      `json:"end"`
	Limit                int                      `json:"limit"`
	Next                 string                   `json:"next"`
	Results              []ResponsePcapSearchItem `json:"results"`
	Start                int                      `json:"start"`
	Total                *int                     `json:"total,omitempty"`
	AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponsePcapSearch.
// It customizes the JSON marshaling process for ResponsePcapSearch objects.
func (r ResponsePcapSearch) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponsePcapSearch object to a map representation for JSON marshaling.
func (r ResponsePcapSearch) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["end"] = r.End
	structMap["limit"] = r.Limit
	structMap["next"] = r.Next
	structMap["results"] = r.Results
	structMap["start"] = r.Start
	if r.Total != nil {
		structMap["total"] = r.Total
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponsePcapSearch.
// It customizes the JSON unmarshaling process for ResponsePcapSearch objects.
func (r *ResponsePcapSearch) UnmarshalJSON(input []byte) error {
	var temp tempResponsePcapSearch
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "end", "limit", "next", "results", "start", "total")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.End = *temp.End
	r.Limit = *temp.Limit
	r.Next = *temp.Next
	r.Results = *temp.Results
	r.Start = *temp.Start
	r.Total = temp.Total
	return nil
}

// tempResponsePcapSearch is a temporary struct used for validating the fields of ResponsePcapSearch.
type tempResponsePcapSearch struct {
	End     *int                      `json:"end"`
	Limit   *int                      `json:"limit"`
	Next    *string                   `json:"next"`
	Results *[]ResponsePcapSearchItem `json:"results"`
	Start   *int                      `json:"start"`
	Total   *int                      `json:"total,omitempty"`
}

func (r *tempResponsePcapSearch) validate() error {
	var errs []string
	if r.End == nil {
		errs = append(errs, "required field `end` is missing for type `response_pcap_search`")
	}
	if r.Limit == nil {
		errs = append(errs, "required field `limit` is missing for type `response_pcap_search`")
	}
	if r.Next == nil {
		errs = append(errs, "required field `next` is missing for type `response_pcap_search`")
	}
	if r.Results == nil {
		errs = append(errs, "required field `results` is missing for type `response_pcap_search`")
	}
	if r.Start == nil {
		errs = append(errs, "required field `start` is missing for type `response_pcap_search`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

package models

import (
	"encoding/json"
)

// SearchWebhookDelivery represents a SearchWebhookDelivery struct.
type SearchWebhookDelivery struct {
	End                  *int              `json:"end,omitempty"`
	Limit                *int              `json:"limit,omitempty"`
	Results              []WebhookDelivery `json:"results,omitempty"`
	Start                *int              `json:"start,omitempty"`
	Total                *int              `json:"total,omitempty"`
	AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SearchWebhookDelivery.
// It customizes the JSON marshaling process for SearchWebhookDelivery objects.
func (s SearchWebhookDelivery) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SearchWebhookDelivery object to a map representation for JSON marshaling.
func (s SearchWebhookDelivery) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SearchWebhookDelivery.
// It customizes the JSON unmarshaling process for SearchWebhookDelivery objects.
func (s *SearchWebhookDelivery) UnmarshalJSON(input []byte) error {
	var temp tempSearchWebhookDelivery
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

// tempSearchWebhookDelivery is a temporary struct used for validating the fields of SearchWebhookDelivery.
type tempSearchWebhookDelivery struct {
	End     *int              `json:"end,omitempty"`
	Limit   *int              `json:"limit,omitempty"`
	Results []WebhookDelivery `json:"results,omitempty"`
	Start   *int              `json:"start,omitempty"`
	Total   *int              `json:"total,omitempty"`
}

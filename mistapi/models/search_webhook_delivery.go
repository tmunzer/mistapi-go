// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SearchWebhookDelivery represents a SearchWebhookDelivery struct.
type SearchWebhookDelivery struct {
    End                  *int                   `json:"end,omitempty"`
    Limit                *int                   `json:"limit,omitempty"`
    Results              []WebhookDelivery      `json:"results,omitempty"`
    Start                *int                   `json:"start,omitempty"`
    Total                *int                   `json:"total,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SearchWebhookDelivery,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SearchWebhookDelivery) String() string {
    return fmt.Sprintf(
    	"SearchWebhookDelivery[End=%v, Limit=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
    	s.End, s.Limit, s.Results, s.Start, s.Total, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SearchWebhookDelivery.
// It customizes the JSON marshaling process for SearchWebhookDelivery objects.
func (s SearchWebhookDelivery) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "end", "limit", "results", "start", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SearchWebhookDelivery object to a map representation for JSON marshaling.
func (s SearchWebhookDelivery) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "results", "start", "total")
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
type tempSearchWebhookDelivery  struct {
    End     *int              `json:"end,omitempty"`
    Limit   *int              `json:"limit,omitempty"`
    Results []WebhookDelivery `json:"results,omitempty"`
    Start   *int              `json:"start,omitempty"`
    Total   *int              `json:"total,omitempty"`
}

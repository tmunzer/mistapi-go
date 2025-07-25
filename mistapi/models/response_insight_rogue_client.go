// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseInsightRogueClient represents a ResponseInsightRogueClient struct.
type ResponseInsightRogueClient struct {
    End                  int                    `json:"end"`
    Limit                int                    `json:"limit"`
    Next                 string                 `json:"next"`
    Results              []InsightRogueClient   `json:"results"`
    Start                int                    `json:"start"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseInsightRogueClient,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseInsightRogueClient) String() string {
    return fmt.Sprintf(
    	"ResponseInsightRogueClient[End=%v, Limit=%v, Next=%v, Results=%v, Start=%v, AdditionalProperties=%v]",
    	r.End, r.Limit, r.Next, r.Results, r.Start, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseInsightRogueClient.
// It customizes the JSON marshaling process for ResponseInsightRogueClient objects.
func (r ResponseInsightRogueClient) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "end", "limit", "next", "results", "start"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseInsightRogueClient object to a map representation for JSON marshaling.
func (r ResponseInsightRogueClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["end"] = r.End
    structMap["limit"] = r.Limit
    structMap["next"] = r.Next
    structMap["results"] = r.Results
    structMap["start"] = r.Start
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseInsightRogueClient.
// It customizes the JSON unmarshaling process for ResponseInsightRogueClient objects.
func (r *ResponseInsightRogueClient) UnmarshalJSON(input []byte) error {
    var temp tempResponseInsightRogueClient
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "end", "limit", "next", "results", "start")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.End = *temp.End
    r.Limit = *temp.Limit
    r.Next = *temp.Next
    r.Results = *temp.Results
    r.Start = *temp.Start
    return nil
}

// tempResponseInsightRogueClient is a temporary struct used for validating the fields of ResponseInsightRogueClient.
type tempResponseInsightRogueClient  struct {
    End     *int                  `json:"end"`
    Limit   *int                  `json:"limit"`
    Next    *string               `json:"next"`
    Results *[]InsightRogueClient `json:"results"`
    Start   *int                  `json:"start"`
}

func (r *tempResponseInsightRogueClient) validate() error {
    var errs []string
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `response_insight_rogue_client`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `response_insight_rogue_client`")
    }
    if r.Next == nil {
        errs = append(errs, "required field `next` is missing for type `response_insight_rogue_client`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_insight_rogue_client`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `response_insight_rogue_client`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

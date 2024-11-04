package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseInsightRogueClient represents a ResponseInsightRogueClient struct.
type ResponseInsightRogueClient struct {
    End                  int                  `json:"end"`
    Limit                int                  `json:"limit"`
    Next                 string               `json:"next"`
    Results              []InsightRogueClient `json:"results"`
    Start                int                  `json:"start"`
    AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseInsightRogueClient.
// It customizes the JSON marshaling process for ResponseInsightRogueClient objects.
func (r ResponseInsightRogueClient) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseInsightRogueClient object to a map representation for JSON marshaling.
func (r ResponseInsightRogueClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "end", "limit", "next", "results", "start")
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

package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// CallTroubleshootSummary represents a CallTroubleshootSummary struct.
type CallTroubleshootSummary struct {
    Mac                  *string                  `json:"mac,omitempty"`
    MeetingId            *uuid.UUID               `json:"meeting_id,omitempty"`
    Results              []CallTroubleshootSummar `json:"results,omitempty"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CallTroubleshootSummary.
// It customizes the JSON marshaling process for CallTroubleshootSummary objects.
func (c CallTroubleshootSummary) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CallTroubleshootSummary object to a map representation for JSON marshaling.
func (c CallTroubleshootSummary) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Mac != nil {
        structMap["mac"] = c.Mac
    }
    if c.MeetingId != nil {
        structMap["meeting_id"] = c.MeetingId
    }
    if c.Results != nil {
        structMap["results"] = c.Results
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CallTroubleshootSummary.
// It customizes the JSON unmarshaling process for CallTroubleshootSummary objects.
func (c *CallTroubleshootSummary) UnmarshalJSON(input []byte) error {
    var temp callTroubleshootSummary
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mac", "meeting_id", "results")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Mac = temp.Mac
    c.MeetingId = temp.MeetingId
    c.Results = temp.Results
    return nil
}

// callTroubleshootSummary is a temporary struct used for validating the fields of CallTroubleshootSummary.
type callTroubleshootSummary  struct {
    Mac       *string                  `json:"mac,omitempty"`
    MeetingId *uuid.UUID               `json:"meeting_id,omitempty"`
    Results   []CallTroubleshootSummar `json:"results,omitempty"`
}

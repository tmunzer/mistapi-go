package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// CallTroubleshootSummary represents a CallTroubleshootSummary struct.
type CallTroubleshootSummary struct {
    Mac                  *string                  `json:"mac,omitempty"`
    MeetingId            *uuid.UUID               `json:"meeting_id,omitempty"`
    Results              []CallTroubleshootSummar `json:"results,omitempty"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for CallTroubleshootSummary,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CallTroubleshootSummary) String() string {
    return fmt.Sprintf(
    	"CallTroubleshootSummary[Mac=%v, MeetingId=%v, Results=%v, AdditionalProperties=%v]",
    	c.Mac, c.MeetingId, c.Results, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CallTroubleshootSummary.
// It customizes the JSON marshaling process for CallTroubleshootSummary objects.
func (c CallTroubleshootSummary) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "mac", "meeting_id", "results"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CallTroubleshootSummary object to a map representation for JSON marshaling.
func (c CallTroubleshootSummary) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
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
    var temp tempCallTroubleshootSummary
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "meeting_id", "results")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Mac = temp.Mac
    c.MeetingId = temp.MeetingId
    c.Results = temp.Results
    return nil
}

// tempCallTroubleshootSummary is a temporary struct used for validating the fields of CallTroubleshootSummary.
type tempCallTroubleshootSummary  struct {
    Mac       *string                  `json:"mac,omitempty"`
    MeetingId *uuid.UUID               `json:"meeting_id,omitempty"`
    Results   []CallTroubleshootSummar `json:"results,omitempty"`
}

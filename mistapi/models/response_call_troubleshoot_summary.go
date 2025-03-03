package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// ResponseCallTroubleshootSummary represents a ResponseCallTroubleshootSummary struct.
type ResponseCallTroubleshootSummary struct {
    Mac                  *string                   `json:"mac,omitempty"`
    MeetingId            *uuid.UUID                `json:"meeting_id,omitempty"`
    Results              []CallTroubleshootSummary `json:"results,omitempty"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseCallTroubleshootSummary,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseCallTroubleshootSummary) String() string {
    return fmt.Sprintf(
    	"ResponseCallTroubleshootSummary[Mac=%v, MeetingId=%v, Results=%v, AdditionalProperties=%v]",
    	r.Mac, r.MeetingId, r.Results, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseCallTroubleshootSummary.
// It customizes the JSON marshaling process for ResponseCallTroubleshootSummary objects.
func (r ResponseCallTroubleshootSummary) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "mac", "meeting_id", "results"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseCallTroubleshootSummary object to a map representation for JSON marshaling.
func (r ResponseCallTroubleshootSummary) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Mac != nil {
        structMap["mac"] = r.Mac
    }
    if r.MeetingId != nil {
        structMap["meeting_id"] = r.MeetingId
    }
    if r.Results != nil {
        structMap["results"] = r.Results
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseCallTroubleshootSummary.
// It customizes the JSON unmarshaling process for ResponseCallTroubleshootSummary objects.
func (r *ResponseCallTroubleshootSummary) UnmarshalJSON(input []byte) error {
    var temp tempResponseCallTroubleshootSummary
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "meeting_id", "results")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Mac = temp.Mac
    r.MeetingId = temp.MeetingId
    r.Results = temp.Results
    return nil
}

// tempResponseCallTroubleshootSummary is a temporary struct used for validating the fields of ResponseCallTroubleshootSummary.
type tempResponseCallTroubleshootSummary  struct {
    Mac       *string                   `json:"mac,omitempty"`
    MeetingId *uuid.UUID                `json:"meeting_id,omitempty"`
    Results   []CallTroubleshootSummary `json:"results,omitempty"`
}

package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// CallTroubleshoot represents a CallTroubleshoot struct.
type CallTroubleshoot struct {
    Mac                  *string                `json:"mac,omitempty"`
    MeetingId            *uuid.UUID             `json:"meeting_id,omitempty"`
    Results              []TroubleshootCallItem `json:"results,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CallTroubleshoot,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CallTroubleshoot) String() string {
    return fmt.Sprintf(
    	"CallTroubleshoot[Mac=%v, MeetingId=%v, Results=%v, AdditionalProperties=%v]",
    	c.Mac, c.MeetingId, c.Results, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CallTroubleshoot.
// It customizes the JSON marshaling process for CallTroubleshoot objects.
func (c CallTroubleshoot) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "mac", "meeting_id", "results"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CallTroubleshoot object to a map representation for JSON marshaling.
func (c CallTroubleshoot) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for CallTroubleshoot.
// It customizes the JSON unmarshaling process for CallTroubleshoot objects.
func (c *CallTroubleshoot) UnmarshalJSON(input []byte) error {
    var temp tempCallTroubleshoot
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

// tempCallTroubleshoot is a temporary struct used for validating the fields of CallTroubleshoot.
type tempCallTroubleshoot  struct {
    Mac       *string                `json:"mac,omitempty"`
    MeetingId *uuid.UUID             `json:"meeting_id,omitempty"`
    Results   []TroubleshootCallItem `json:"results,omitempty"`
}

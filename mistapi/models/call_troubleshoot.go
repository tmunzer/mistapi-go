package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// CallTroubleshoot represents a CallTroubleshoot struct.
type CallTroubleshoot struct {
	Mac                  *string                `json:"mac,omitempty"`
	MeetingId            *uuid.UUID             `json:"meeting_id,omitempty"`
	Results              []TroubleshootCallItem `json:"results,omitempty"`
	AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CallTroubleshoot.
// It customizes the JSON marshaling process for CallTroubleshoot objects.
func (c CallTroubleshoot) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the CallTroubleshoot object to a map representation for JSON marshaling.
func (c CallTroubleshoot) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for CallTroubleshoot.
// It customizes the JSON unmarshaling process for CallTroubleshoot objects.
func (c *CallTroubleshoot) UnmarshalJSON(input []byte) error {
	var temp tempCallTroubleshoot
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

// tempCallTroubleshoot is a temporary struct used for validating the fields of CallTroubleshoot.
type tempCallTroubleshoot struct {
	Mac       *string                `json:"mac,omitempty"`
	MeetingId *uuid.UUID             `json:"meeting_id,omitempty"`
	Results   []TroubleshootCallItem `json:"results,omitempty"`
}

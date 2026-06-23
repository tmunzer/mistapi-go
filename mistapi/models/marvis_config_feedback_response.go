// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// MarvisConfigFeedbackResponse represents a MarvisConfigFeedbackResponse struct.
// Response after submitting feedback on a Marvis config action
type MarvisConfigFeedbackResponse struct {
	// The note provided with the feedback
	FeedbackNote *string `json:"feedback_note,omitempty"`
	// The feedback type that was submitted
	FeedbackType         *string                `json:"feedback_type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisConfigFeedbackResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisConfigFeedbackResponse) String() string {
	return fmt.Sprintf(
		"MarvisConfigFeedbackResponse[FeedbackNote=%v, FeedbackType=%v, AdditionalProperties=%v]",
		m.FeedbackNote, m.FeedbackType, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisConfigFeedbackResponse.
// It customizes the JSON marshaling process for MarvisConfigFeedbackResponse objects.
func (m MarvisConfigFeedbackResponse) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"feedback_note", "feedback_type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MarvisConfigFeedbackResponse object to a map representation for JSON marshaling.
func (m MarvisConfigFeedbackResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.FeedbackNote != nil {
		structMap["feedback_note"] = m.FeedbackNote
	}
	if m.FeedbackType != nil {
		structMap["feedback_type"] = m.FeedbackType
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MarvisConfigFeedbackResponse.
// It customizes the JSON unmarshaling process for MarvisConfigFeedbackResponse objects.
func (m *MarvisConfigFeedbackResponse) UnmarshalJSON(input []byte) error {
	var temp tempMarvisConfigFeedbackResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "feedback_note", "feedback_type")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.FeedbackNote = temp.FeedbackNote
	m.FeedbackType = temp.FeedbackType
	return nil
}

// tempMarvisConfigFeedbackResponse is a temporary struct used for validating the fields of MarvisConfigFeedbackResponse.
type tempMarvisConfigFeedbackResponse struct {
	FeedbackNote *string `json:"feedback_note,omitempty"`
	FeedbackType *string `json:"feedback_type,omitempty"`
}

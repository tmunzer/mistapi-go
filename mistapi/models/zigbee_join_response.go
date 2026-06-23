// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// ZigbeeJoinResponse represents a ZigbeeJoinResponse struct.
// Response containing the session identifier for a Zigbee join operation
type ZigbeeJoinResponse struct {
	// Session ID for the Zigbee join operation
	SessionId            *uuid.UUID             `json:"session_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ZigbeeJoinResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (z ZigbeeJoinResponse) String() string {
	return fmt.Sprintf(
		"ZigbeeJoinResponse[SessionId=%v, AdditionalProperties=%v]",
		z.SessionId, z.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ZigbeeJoinResponse.
// It customizes the JSON marshaling process for ZigbeeJoinResponse objects.
func (z ZigbeeJoinResponse) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(z.AdditionalProperties,
		"session_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(z.toMap())
}

// toMap converts the ZigbeeJoinResponse object to a map representation for JSON marshaling.
func (z ZigbeeJoinResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, z.AdditionalProperties)
	if z.SessionId != nil {
		structMap["session_id"] = z.SessionId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ZigbeeJoinResponse.
// It customizes the JSON unmarshaling process for ZigbeeJoinResponse objects.
func (z *ZigbeeJoinResponse) UnmarshalJSON(input []byte) error {
	var temp tempZigbeeJoinResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "session_id")
	if err != nil {
		return err
	}
	z.AdditionalProperties = additionalProperties

	z.SessionId = temp.SessionId
	return nil
}

// tempZigbeeJoinResponse is a temporary struct used for validating the fields of ZigbeeJoinResponse.
type tempZigbeeJoinResponse struct {
	SessionId *uuid.UUID `json:"session_id,omitempty"`
}

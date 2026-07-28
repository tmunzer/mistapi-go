// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// ZigbeeTrailResponse represents a ZigbeeTrailResponse struct.
// Response containing the session identifier for a Zigbee event or packet trail operation
type ZigbeeTrailResponse struct {
	// Session ID the UI can use to stream trail results
	Session              *uuid.UUID             `json:"session,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ZigbeeTrailResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (z ZigbeeTrailResponse) String() string {
	return fmt.Sprintf(
		"ZigbeeTrailResponse[Session=%v, AdditionalProperties=%v]",
		z.Session, z.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ZigbeeTrailResponse.
// It customizes the JSON marshaling process for ZigbeeTrailResponse objects.
func (z ZigbeeTrailResponse) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(z.AdditionalProperties,
		"session"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(z.toMap())
}

// toMap converts the ZigbeeTrailResponse object to a map representation for JSON marshaling.
func (z ZigbeeTrailResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, z.AdditionalProperties)
	if z.Session != nil {
		structMap["session"] = z.Session
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ZigbeeTrailResponse.
// It customizes the JSON unmarshaling process for ZigbeeTrailResponse objects.
func (z *ZigbeeTrailResponse) UnmarshalJSON(input []byte) error {
	var temp tempZigbeeTrailResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "session")
	if err != nil {
		return err
	}
	z.AdditionalProperties = additionalProperties

	z.Session = temp.Session
	return nil
}

// tempZigbeeTrailResponse is a temporary struct used for validating the fields of ZigbeeTrailResponse.
type tempZigbeeTrailResponse struct {
	Session *uuid.UUID `json:"session,omitempty"`
}

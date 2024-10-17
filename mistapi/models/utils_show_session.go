package models

import (
	"encoding/json"
)

// UtilsShowSession represents a UtilsShowSession struct.
type UtilsShowSession struct {
	// only for HA. enum: `node0`, `node1`
	Node *HaClusterNodeEnum `json:"node,omitempty"`
	// The exact service name for which to display the active sessions
	ServiceName *string `json:"service_name,omitempty"`
	// Show session details by session_id
	SessionId            *string        `json:"session_id,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsShowSession.
// It customizes the JSON marshaling process for UtilsShowSession objects.
func (u UtilsShowSession) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(u.toMap())
}

// toMap converts the UtilsShowSession object to a map representation for JSON marshaling.
func (u UtilsShowSession) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, u.AdditionalProperties)
	if u.Node != nil {
		structMap["node"] = u.Node
	}
	if u.ServiceName != nil {
		structMap["service_name"] = u.ServiceName
	}
	if u.SessionId != nil {
		structMap["session_id"] = u.SessionId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsShowSession.
// It customizes the JSON unmarshaling process for UtilsShowSession objects.
func (u *UtilsShowSession) UnmarshalJSON(input []byte) error {
	var temp tempUtilsShowSession
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "node", "service_name", "session_id")
	if err != nil {
		return err
	}

	u.AdditionalProperties = additionalProperties
	u.Node = temp.Node
	u.ServiceName = temp.ServiceName
	u.SessionId = temp.SessionId
	return nil
}

// tempUtilsShowSession is a temporary struct used for validating the fields of UtilsShowSession.
type tempUtilsShowSession struct {
	Node        *HaClusterNodeEnum `json:"node,omitempty"`
	ServiceName *string            `json:"service_name,omitempty"`
	SessionId   *string            `json:"session_id,omitempty"`
}

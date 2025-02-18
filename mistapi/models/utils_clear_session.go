package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// UtilsClearSession represents a UtilsClearSession struct.
// To use five tuples to lookup the session to be cleared, all must be provided
type UtilsClearSession struct {
    // only for HA. enum: `node0`, `node1`
    Node                 *HaClusterNodeEnum     `json:"node,omitempty"`
    // Service name, only supported in SSR
    ServiceName          *string                `json:"service_name,omitempty"`
    // List of id of the sessions to be cleared
    SessionIds           []uuid.UUID            `json:"session_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsClearSession,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsClearSession) String() string {
    return fmt.Sprintf(
    	"UtilsClearSession[Node=%v, ServiceName=%v, SessionIds=%v, AdditionalProperties=%v]",
    	u.Node, u.ServiceName, u.SessionIds, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsClearSession.
// It customizes the JSON marshaling process for UtilsClearSession objects.
func (u UtilsClearSession) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "node", "service_name", "session_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsClearSession object to a map representation for JSON marshaling.
func (u UtilsClearSession) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    if u.ServiceName != nil {
        structMap["service_name"] = u.ServiceName
    }
    if u.SessionIds != nil {
        structMap["session_ids"] = u.SessionIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsClearSession.
// It customizes the JSON unmarshaling process for UtilsClearSession objects.
func (u *UtilsClearSession) UnmarshalJSON(input []byte) error {
    var temp tempUtilsClearSession
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "node", "service_name", "session_ids")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.Node = temp.Node
    u.ServiceName = temp.ServiceName
    u.SessionIds = temp.SessionIds
    return nil
}

// tempUtilsClearSession is a temporary struct used for validating the fields of UtilsClearSession.
type tempUtilsClearSession  struct {
    Node        *HaClusterNodeEnum `json:"node,omitempty"`
    ServiceName *string            `json:"service_name,omitempty"`
    SessionIds  []uuid.UUID        `json:"session_ids,omitempty"`
}

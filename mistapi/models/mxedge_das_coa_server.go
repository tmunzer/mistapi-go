package models

import (
    "encoding/json"
    "fmt"
)

// MxedgeDasCoaServer represents a MxedgeDasCoaServer struct.
type MxedgeDasCoaServer struct {
    // Whether to disable Event-Timestamp Check
    DisableEventTimestampCheck  *bool                  `json:"disable_event_timestamp_check,omitempty"`
    Enabled                     *bool                  `json:"enabled,omitempty"`
    // This server configured to send CoA|DM to mist edges
    Host                        *string                `json:"host,omitempty"`
    // Mist edges will allow this host on this port
    Port                        *int                   `json:"port,omitempty"`
    // Whether to require Message-Authenticator in requests
    RequireMessageAuthenticator *bool                  `json:"require_message_authenticator,omitempty"`
    Secret                      *string                `json:"secret,omitempty"`
    AdditionalProperties        map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeDasCoaServer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeDasCoaServer) String() string {
    return fmt.Sprintf(
    	"MxedgeDasCoaServer[DisableEventTimestampCheck=%v, Enabled=%v, Host=%v, Port=%v, RequireMessageAuthenticator=%v, Secret=%v, AdditionalProperties=%v]",
    	m.DisableEventTimestampCheck, m.Enabled, m.Host, m.Port, m.RequireMessageAuthenticator, m.Secret, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeDasCoaServer.
// It customizes the JSON marshaling process for MxedgeDasCoaServer objects.
func (m MxedgeDasCoaServer) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "disable_event_timestamp_check", "enabled", "host", "port", "require_message_authenticator", "secret"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeDasCoaServer object to a map representation for JSON marshaling.
func (m MxedgeDasCoaServer) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.DisableEventTimestampCheck != nil {
        structMap["disable_event_timestamp_check"] = m.DisableEventTimestampCheck
    }
    if m.Enabled != nil {
        structMap["enabled"] = m.Enabled
    }
    if m.Host != nil {
        structMap["host"] = m.Host
    }
    if m.Port != nil {
        structMap["port"] = m.Port
    }
    if m.RequireMessageAuthenticator != nil {
        structMap["require_message_authenticator"] = m.RequireMessageAuthenticator
    }
    if m.Secret != nil {
        structMap["secret"] = m.Secret
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeDasCoaServer.
// It customizes the JSON unmarshaling process for MxedgeDasCoaServer objects.
func (m *MxedgeDasCoaServer) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeDasCoaServer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "disable_event_timestamp_check", "enabled", "host", "port", "require_message_authenticator", "secret")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.DisableEventTimestampCheck = temp.DisableEventTimestampCheck
    m.Enabled = temp.Enabled
    m.Host = temp.Host
    m.Port = temp.Port
    m.RequireMessageAuthenticator = temp.RequireMessageAuthenticator
    m.Secret = temp.Secret
    return nil
}

// tempMxedgeDasCoaServer is a temporary struct used for validating the fields of MxedgeDasCoaServer.
type tempMxedgeDasCoaServer  struct {
    DisableEventTimestampCheck  *bool   `json:"disable_event_timestamp_check,omitempty"`
    Enabled                     *bool   `json:"enabled,omitempty"`
    Host                        *string `json:"host,omitempty"`
    Port                        *int    `json:"port,omitempty"`
    RequireMessageAuthenticator *bool   `json:"require_message_authenticator,omitempty"`
    Secret                      *string `json:"secret,omitempty"`
}

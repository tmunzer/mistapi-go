package models

import (
    "encoding/json"
)

// MxedgeDasCoaServer represents a MxedgeDasCoaServer struct.
type MxedgeDasCoaServer struct {
    // whether to disable Event-Timestamp Check
    DisableEventTimestampCheck *bool          `json:"disable_event_timestamp_check,omitempty"`
    Enabled                    *bool          `json:"enabled,omitempty"`
    // this server configured to send CoA|DM to mist edges
    Host                       *string        `json:"host,omitempty"`
    // mist edges will allow this host on this port
    Port                       *int           `json:"port,omitempty"`
    Secret                     *string        `json:"secret,omitempty"`
    AdditionalProperties       map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeDasCoaServer.
// It customizes the JSON marshaling process for MxedgeDasCoaServer objects.
func (m MxedgeDasCoaServer) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeDasCoaServer object to a map representation for JSON marshaling.
func (m MxedgeDasCoaServer) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "disable_event_timestamp_check", "enabled", "host", "port", "secret")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.DisableEventTimestampCheck = temp.DisableEventTimestampCheck
    m.Enabled = temp.Enabled
    m.Host = temp.Host
    m.Port = temp.Port
    m.Secret = temp.Secret
    return nil
}

// tempMxedgeDasCoaServer is a temporary struct used for validating the fields of MxedgeDasCoaServer.
type tempMxedgeDasCoaServer  struct {
    DisableEventTimestampCheck *bool   `json:"disable_event_timestamp_check,omitempty"`
    Enabled                    *bool   `json:"enabled,omitempty"`
    Host                       *string `json:"host,omitempty"`
    Port                       *int    `json:"port,omitempty"`
    Secret                     *string `json:"secret,omitempty"`
}

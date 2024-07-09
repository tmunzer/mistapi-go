package models

import (
    "encoding/json"
)

// ApStatsL2TpStatSession represents a ApStatsL2TpStatSession struct.
type ApStatsL2TpStatSession struct {
    // remote sessions id (dynamically unless Tunnel is said to be static)
    LocalSid             *int           `json:"local_sid,omitempty"`
    // WxlanTunnel Remote ID (user-configured)
    RemoteId             *string        `json:"remote_id,omitempty"`
    // remote sessions id (dynamically unless Tunnel is said to be static)
    RemoteSid            *int           `json:"remote_sid,omitempty"`
    State                *L2TpStateEnum `json:"state,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsL2TpStatSession.
// It customizes the JSON marshaling process for ApStatsL2TpStatSession objects.
func (a ApStatsL2TpStatSession) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsL2TpStatSession object to a map representation for JSON marshaling.
func (a ApStatsL2TpStatSession) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.LocalSid != nil {
        structMap["local_sid"] = a.LocalSid
    }
    if a.RemoteId != nil {
        structMap["remote_id"] = a.RemoteId
    }
    if a.RemoteSid != nil {
        structMap["remote_sid"] = a.RemoteSid
    }
    if a.State != nil {
        structMap["state"] = a.State
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsL2TpStatSession.
// It customizes the JSON unmarshaling process for ApStatsL2TpStatSession objects.
func (a *ApStatsL2TpStatSession) UnmarshalJSON(input []byte) error {
    var temp apStatsL2TpStatSession
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "local_sid", "remote_id", "remote_sid", "state")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.LocalSid = temp.LocalSid
    a.RemoteId = temp.RemoteId
    a.RemoteSid = temp.RemoteSid
    a.State = temp.State
    return nil
}

// apStatsL2TpStatSession is a temporary struct used for validating the fields of ApStatsL2TpStatSession.
type apStatsL2TpStatSession  struct {
    LocalSid  *int           `json:"local_sid,omitempty"`
    RemoteId  *string        `json:"remote_id,omitempty"`
    RemoteSid *int           `json:"remote_sid,omitempty"`
    State     *L2TpStateEnum `json:"state,omitempty"`
}

package models

import (
    "encoding/json"
)

// ApStatsL2tpStatSession represents a ApStatsL2tpStatSession struct.
type ApStatsL2tpStatSession struct {
    // remote sessions id (dynamically unless Tunnel is said to be static)
    LocalSid             Optional[int]    `json:"local_sid"`
    // WxlanTunnel Remote ID (user-configured)
    RemoteId             Optional[string] `json:"remote_id"`
    // remote sessions id (dynamically unless Tunnel is said to be static)
    RemoteSid            Optional[int]    `json:"remote_sid"`
    // enum: `established`, `established_with_session`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply`
    State                *L2tpStateEnum   `json:"state,omitempty"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsL2tpStatSession.
// It customizes the JSON marshaling process for ApStatsL2tpStatSession objects.
func (a ApStatsL2tpStatSession) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsL2tpStatSession object to a map representation for JSON marshaling.
func (a ApStatsL2tpStatSession) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.LocalSid.IsValueSet() {
        if a.LocalSid.Value() != nil {
            structMap["local_sid"] = a.LocalSid.Value()
        } else {
            structMap["local_sid"] = nil
        }
    }
    if a.RemoteId.IsValueSet() {
        if a.RemoteId.Value() != nil {
            structMap["remote_id"] = a.RemoteId.Value()
        } else {
            structMap["remote_id"] = nil
        }
    }
    if a.RemoteSid.IsValueSet() {
        if a.RemoteSid.Value() != nil {
            structMap["remote_sid"] = a.RemoteSid.Value()
        } else {
            structMap["remote_sid"] = nil
        }
    }
    if a.State != nil {
        structMap["state"] = a.State
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsL2tpStatSession.
// It customizes the JSON unmarshaling process for ApStatsL2tpStatSession objects.
func (a *ApStatsL2tpStatSession) UnmarshalJSON(input []byte) error {
    var temp tempApStatsL2tpStatSession
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

// tempApStatsL2tpStatSession is a temporary struct used for validating the fields of ApStatsL2tpStatSession.
type tempApStatsL2tpStatSession  struct {
    LocalSid  Optional[int]    `json:"local_sid"`
    RemoteId  Optional[string] `json:"remote_id"`
    RemoteSid Optional[int]    `json:"remote_sid"`
    State     *L2tpStateEnum   `json:"state,omitempty"`
}

package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// StatsMxtunnelSession represents a StatsMxtunnelSession struct.
type StatsMxtunnelSession struct {
    // remote sessions id (dynamically unless Tunnel is said to be static)
    LocalSid             int            `json:"local_sid"`
    // WxlanTunnel Remote ID
    RemoteId             string         `json:"remote_id"`
    // remote sessions id (dynamically unless Tunnel is said to be static)
    RemoteSid            int            `json:"remote_sid"`
    State                string         `json:"state"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsMxtunnelSession.
// It customizes the JSON marshaling process for StatsMxtunnelSession objects.
func (s StatsMxtunnelSession) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsMxtunnelSession object to a map representation for JSON marshaling.
func (s StatsMxtunnelSession) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["local_sid"] = s.LocalSid
    structMap["remote_id"] = s.RemoteId
    structMap["remote_sid"] = s.RemoteSid
    structMap["state"] = s.State
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsMxtunnelSession.
// It customizes the JSON unmarshaling process for StatsMxtunnelSession objects.
func (s *StatsMxtunnelSession) UnmarshalJSON(input []byte) error {
    var temp tempStatsMxtunnelSession
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "local_sid", "remote_id", "remote_sid", "state")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.LocalSid = *temp.LocalSid
    s.RemoteId = *temp.RemoteId
    s.RemoteSid = *temp.RemoteSid
    s.State = *temp.State
    return nil
}

// tempStatsMxtunnelSession is a temporary struct used for validating the fields of StatsMxtunnelSession.
type tempStatsMxtunnelSession  struct {
    LocalSid  *int    `json:"local_sid"`
    RemoteId  *string `json:"remote_id"`
    RemoteSid *int    `json:"remote_sid"`
    State     *string `json:"state"`
}

func (s *tempStatsMxtunnelSession) validate() error {
    var errs []string
    if s.LocalSid == nil {
        errs = append(errs, "required field `local_sid` is missing for type `stats_mxtunnel_session`")
    }
    if s.RemoteId == nil {
        errs = append(errs, "required field `remote_id` is missing for type `stats_mxtunnel_session`")
    }
    if s.RemoteSid == nil {
        errs = append(errs, "required field `remote_sid` is missing for type `stats_mxtunnel_session`")
    }
    if s.State == nil {
        errs = append(errs, "required field `state` is missing for type `stats_mxtunnel_session`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// MxtunnelStatsSession represents a MxtunnelStatsSession struct.
type MxtunnelStatsSession struct {
    // remote sessions id (dynamically unless Tunnel is said to be static)
    LocalSid             int            `json:"local_sid"`
    // WxlanTunnel Remote ID
    RemoteId             string         `json:"remote_id"`
    // remote sessions id (dynamically unless Tunnel is said to be static)
    RemoteSid            int            `json:"remote_sid"`
    State                string         `json:"state"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxtunnelStatsSession.
// It customizes the JSON marshaling process for MxtunnelStatsSession objects.
func (m MxtunnelStatsSession) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxtunnelStatsSession object to a map representation for JSON marshaling.
func (m MxtunnelStatsSession) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["local_sid"] = m.LocalSid
    structMap["remote_id"] = m.RemoteId
    structMap["remote_sid"] = m.RemoteSid
    structMap["state"] = m.State
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxtunnelStatsSession.
// It customizes the JSON unmarshaling process for MxtunnelStatsSession objects.
func (m *MxtunnelStatsSession) UnmarshalJSON(input []byte) error {
    var temp tempMxtunnelStatsSession
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
    
    m.AdditionalProperties = additionalProperties
    m.LocalSid = *temp.LocalSid
    m.RemoteId = *temp.RemoteId
    m.RemoteSid = *temp.RemoteSid
    m.State = *temp.State
    return nil
}

// tempMxtunnelStatsSession is a temporary struct used for validating the fields of MxtunnelStatsSession.
type tempMxtunnelStatsSession  struct {
    LocalSid  *int    `json:"local_sid"`
    RemoteId  *string `json:"remote_id"`
    RemoteSid *int    `json:"remote_sid"`
    State     *string `json:"state"`
}

func (m *tempMxtunnelStatsSession) validate() error {
    var errs []string
    if m.LocalSid == nil {
        errs = append(errs, "required field `local_sid` is missing for type `mxtunnel_stats_session`")
    }
    if m.RemoteId == nil {
        errs = append(errs, "required field `remote_id` is missing for type `mxtunnel_stats_session`")
    }
    if m.RemoteSid == nil {
        errs = append(errs, "required field `remote_sid` is missing for type `mxtunnel_stats_session`")
    }
    if m.State == nil {
        errs = append(errs, "required field `state` is missing for type `mxtunnel_stats_session`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

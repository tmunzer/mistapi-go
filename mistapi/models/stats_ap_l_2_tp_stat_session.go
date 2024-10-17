package models

import (
	"encoding/json"
)

// StatsApL2tpStatSession represents a StatsApL2tpStatSession struct.
type StatsApL2tpStatSession struct {
	// remote sessions id (dynamically unless Tunnel is said to be static)
	LocalSid Optional[int] `json:"local_sid"`
	// WxlanTunnel Remote ID (user-configured)
	RemoteId Optional[string] `json:"remote_id"`
	// remote sessions id (dynamically unless Tunnel is said to be static)
	RemoteSid Optional[int] `json:"remote_sid"`
	// enum: `established`, `established_with_session`, `idle`, `wait-ctrl-conn`, `wait-ctrl-reply`
	State                *L2tpStateEnum `json:"state,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsApL2tpStatSession.
// It customizes the JSON marshaling process for StatsApL2tpStatSession objects.
func (s StatsApL2tpStatSession) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the StatsApL2tpStatSession object to a map representation for JSON marshaling.
func (s StatsApL2tpStatSession) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.LocalSid.IsValueSet() {
		if s.LocalSid.Value() != nil {
			structMap["local_sid"] = s.LocalSid.Value()
		} else {
			structMap["local_sid"] = nil
		}
	}
	if s.RemoteId.IsValueSet() {
		if s.RemoteId.Value() != nil {
			structMap["remote_id"] = s.RemoteId.Value()
		} else {
			structMap["remote_id"] = nil
		}
	}
	if s.RemoteSid.IsValueSet() {
		if s.RemoteSid.Value() != nil {
			structMap["remote_sid"] = s.RemoteSid.Value()
		} else {
			structMap["remote_sid"] = nil
		}
	}
	if s.State != nil {
		structMap["state"] = s.State
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApL2tpStatSession.
// It customizes the JSON unmarshaling process for StatsApL2tpStatSession objects.
func (s *StatsApL2tpStatSession) UnmarshalJSON(input []byte) error {
	var temp tempStatsApL2tpStatSession
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "local_sid", "remote_id", "remote_sid", "state")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.LocalSid = temp.LocalSid
	s.RemoteId = temp.RemoteId
	s.RemoteSid = temp.RemoteSid
	s.State = temp.State
	return nil
}

// tempStatsApL2tpStatSession is a temporary struct used for validating the fields of StatsApL2tpStatSession.
type tempStatsApL2tpStatSession struct {
	LocalSid  Optional[int]    `json:"local_sid"`
	RemoteId  Optional[string] `json:"remote_id"`
	RemoteSid Optional[int]    `json:"remote_sid"`
	State     *L2tpStateEnum   `json:"state,omitempty"`
}

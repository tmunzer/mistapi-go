package models

import (
	"encoding/json"
)

// StatsApUsbStat represents a StatsApUsbStat struct.
type StatsApUsbStat struct {
	Channel              Optional[int]    `json:"channel"`
	Connected            Optional[bool]   `json:"connected"`
	LastActivity         Optional[int]    `json:"last_activity"`
	Type                 Optional[string] `json:"type"`
	Up                   Optional[bool]   `json:"up"`
	AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsApUsbStat.
// It customizes the JSON marshaling process for StatsApUsbStat objects.
func (s StatsApUsbStat) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the StatsApUsbStat object to a map representation for JSON marshaling.
func (s StatsApUsbStat) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Channel.IsValueSet() {
		if s.Channel.Value() != nil {
			structMap["channel"] = s.Channel.Value()
		} else {
			structMap["channel"] = nil
		}
	}
	if s.Connected.IsValueSet() {
		if s.Connected.Value() != nil {
			structMap["connected"] = s.Connected.Value()
		} else {
			structMap["connected"] = nil
		}
	}
	if s.LastActivity.IsValueSet() {
		if s.LastActivity.Value() != nil {
			structMap["last_activity"] = s.LastActivity.Value()
		} else {
			structMap["last_activity"] = nil
		}
	}
	if s.Type.IsValueSet() {
		if s.Type.Value() != nil {
			structMap["type"] = s.Type.Value()
		} else {
			structMap["type"] = nil
		}
	}
	if s.Up.IsValueSet() {
		if s.Up.Value() != nil {
			structMap["up"] = s.Up.Value()
		} else {
			structMap["up"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApUsbStat.
// It customizes the JSON unmarshaling process for StatsApUsbStat objects.
func (s *StatsApUsbStat) UnmarshalJSON(input []byte) error {
	var temp tempStatsApUsbStat
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "channel", "connected", "last_activity", "type", "up")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Channel = temp.Channel
	s.Connected = temp.Connected
	s.LastActivity = temp.LastActivity
	s.Type = temp.Type
	s.Up = temp.Up
	return nil
}

// tempStatsApUsbStat is a temporary struct used for validating the fields of StatsApUsbStat.
type tempStatsApUsbStat struct {
	Channel      Optional[int]    `json:"channel"`
	Connected    Optional[bool]   `json:"connected"`
	LastActivity Optional[int]    `json:"last_activity"`
	Type         Optional[string] `json:"type"`
	Up           Optional[bool]   `json:"up"`
}

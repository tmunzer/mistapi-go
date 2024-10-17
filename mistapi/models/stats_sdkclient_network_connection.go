package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// StatsSdkclientNetworkConnection represents a StatsSdkclientNetworkConnection struct.
// various network connection info for the SDK client (if known, else omitted), with RSSI in dBm, and signal level as
type StatsSdkclientNetworkConnection struct {
	Mac                  string         `json:"mac"`
	Rssi                 float64        `json:"rssi"`
	SignalLevel          float64        `json:"signal_level"`
	Type                 string         `json:"type"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsSdkclientNetworkConnection.
// It customizes the JSON marshaling process for StatsSdkclientNetworkConnection objects.
func (s StatsSdkclientNetworkConnection) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the StatsSdkclientNetworkConnection object to a map representation for JSON marshaling.
func (s StatsSdkclientNetworkConnection) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["mac"] = s.Mac
	structMap["rssi"] = s.Rssi
	structMap["signal_level"] = s.SignalLevel
	structMap["type"] = s.Type
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsSdkclientNetworkConnection.
// It customizes the JSON unmarshaling process for StatsSdkclientNetworkConnection objects.
func (s *StatsSdkclientNetworkConnection) UnmarshalJSON(input []byte) error {
	var temp tempStatsSdkclientNetworkConnection
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "mac", "rssi", "signal_level", "type")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Mac = *temp.Mac
	s.Rssi = *temp.Rssi
	s.SignalLevel = *temp.SignalLevel
	s.Type = *temp.Type
	return nil
}

// tempStatsSdkclientNetworkConnection is a temporary struct used for validating the fields of StatsSdkclientNetworkConnection.
type tempStatsSdkclientNetworkConnection struct {
	Mac         *string  `json:"mac"`
	Rssi        *float64 `json:"rssi"`
	SignalLevel *float64 `json:"signal_level"`
	Type        *string  `json:"type"`
}

func (s *tempStatsSdkclientNetworkConnection) validate() error {
	var errs []string
	if s.Mac == nil {
		errs = append(errs, "required field `mac` is missing for type `stats_sdkclient_network_connection`")
	}
	if s.Rssi == nil {
		errs = append(errs, "required field `rssi` is missing for type `stats_sdkclient_network_connection`")
	}
	if s.SignalLevel == nil {
		errs = append(errs, "required field `signal_level` is missing for type `stats_sdkclient_network_connection`")
	}
	if s.Type == nil {
		errs = append(errs, "required field `type` is missing for type `stats_sdkclient_network_connection`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

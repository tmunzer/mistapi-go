package models

import (
    "encoding/json"
)

// ApStatsRadioConfigBand represents a ApStatsRadioConfigBand struct.
type ApStatsRadioConfigBand struct {
    Bandwidth              *float64       `json:"bandwidth,omitempty"`
    Channel                *int           `json:"channel,omitempty"`
    DynamicChainingEnabled *bool          `json:"dynamic_chaining_enabled,omitempty"`
    Power                  *float64       `json:"power,omitempty"`
    RxChain                *int           `json:"rx_chain,omitempty"`
    TxChain                *int           `json:"tx_chain,omitempty"`
    AdditionalProperties   map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsRadioConfigBand.
// It customizes the JSON marshaling process for ApStatsRadioConfigBand objects.
func (a ApStatsRadioConfigBand) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsRadioConfigBand object to a map representation for JSON marshaling.
func (a ApStatsRadioConfigBand) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Bandwidth != nil {
        structMap["bandwidth"] = a.Bandwidth
    }
    if a.Channel != nil {
        structMap["channel"] = a.Channel
    }
    if a.DynamicChainingEnabled != nil {
        structMap["dynamic_chaining_enabled"] = a.DynamicChainingEnabled
    }
    if a.Power != nil {
        structMap["power"] = a.Power
    }
    if a.RxChain != nil {
        structMap["rx_chain"] = a.RxChain
    }
    if a.TxChain != nil {
        structMap["tx_chain"] = a.TxChain
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsRadioConfigBand.
// It customizes the JSON unmarshaling process for ApStatsRadioConfigBand objects.
func (a *ApStatsRadioConfigBand) UnmarshalJSON(input []byte) error {
    var temp apStatsRadioConfigBand
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "bandwidth", "channel", "dynamic_chaining_enabled", "power", "rx_chain", "tx_chain")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Bandwidth = temp.Bandwidth
    a.Channel = temp.Channel
    a.DynamicChainingEnabled = temp.DynamicChainingEnabled
    a.Power = temp.Power
    a.RxChain = temp.RxChain
    a.TxChain = temp.TxChain
    return nil
}

// apStatsRadioConfigBand is a temporary struct used for validating the fields of ApStatsRadioConfigBand.
type apStatsRadioConfigBand  struct {
    Bandwidth              *float64 `json:"bandwidth,omitempty"`
    Channel                *int     `json:"channel,omitempty"`
    DynamicChainingEnabled *bool    `json:"dynamic_chaining_enabled,omitempty"`
    Power                  *float64 `json:"power,omitempty"`
    RxChain                *int     `json:"rx_chain,omitempty"`
    TxChain                *int     `json:"tx_chain,omitempty"`
}

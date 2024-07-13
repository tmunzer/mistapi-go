package models

import (
    "encoding/json"
)

// ApStatsRadioConfigBand represents a ApStatsRadioConfigBand struct.
type ApStatsRadioConfigBand struct {
    AllowRrmDisable        Optional[bool]    `json:"allow_rrm_disable"`
    Bandwidth              Optional[float64] `json:"bandwidth"`
    Channel                *int              `json:"channel,omitempty"`
    Disabled               Optional[bool]    `json:"disabled"`
    DynamicChainingEnabled Optional[bool]    `json:"dynamic_chaining_enabled"`
    Power                  Optional[float64] `json:"power"`
    PowerMax               Optional[float64] `json:"power_max"`
    PowerMin               Optional[float64] `json:"power_min"`
    RxChain                Optional[int]     `json:"rx_chain"`
    TxChain                Optional[int]     `json:"tx_chain"`
    AdditionalProperties   map[string]any    `json:"_"`
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
    if a.AllowRrmDisable.IsValueSet() {
        if a.AllowRrmDisable.Value() != nil {
            structMap["allow_rrm_disable"] = a.AllowRrmDisable.Value()
        } else {
            structMap["allow_rrm_disable"] = nil
        }
    }
    if a.Bandwidth.IsValueSet() {
        if a.Bandwidth.Value() != nil {
            structMap["bandwidth"] = a.Bandwidth.Value()
        } else {
            structMap["bandwidth"] = nil
        }
    }
    if a.Channel != nil {
        structMap["channel"] = a.Channel
    }
    if a.Disabled.IsValueSet() {
        if a.Disabled.Value() != nil {
            structMap["disabled"] = a.Disabled.Value()
        } else {
            structMap["disabled"] = nil
        }
    }
    if a.DynamicChainingEnabled.IsValueSet() {
        if a.DynamicChainingEnabled.Value() != nil {
            structMap["dynamic_chaining_enabled"] = a.DynamicChainingEnabled.Value()
        } else {
            structMap["dynamic_chaining_enabled"] = nil
        }
    }
    if a.Power.IsValueSet() {
        if a.Power.Value() != nil {
            structMap["power"] = a.Power.Value()
        } else {
            structMap["power"] = nil
        }
    }
    if a.PowerMax.IsValueSet() {
        if a.PowerMax.Value() != nil {
            structMap["power_max"] = a.PowerMax.Value()
        } else {
            structMap["power_max"] = nil
        }
    }
    if a.PowerMin.IsValueSet() {
        if a.PowerMin.Value() != nil {
            structMap["power_min"] = a.PowerMin.Value()
        } else {
            structMap["power_min"] = nil
        }
    }
    if a.RxChain.IsValueSet() {
        if a.RxChain.Value() != nil {
            structMap["rx_chain"] = a.RxChain.Value()
        } else {
            structMap["rx_chain"] = nil
        }
    }
    if a.TxChain.IsValueSet() {
        if a.TxChain.Value() != nil {
            structMap["tx_chain"] = a.TxChain.Value()
        } else {
            structMap["tx_chain"] = nil
        }
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "allow_rrm_disable", "bandwidth", "channel", "disabled", "dynamic_chaining_enabled", "power", "power_max", "power_min", "rx_chain", "tx_chain")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.AllowRrmDisable = temp.AllowRrmDisable
    a.Bandwidth = temp.Bandwidth
    a.Channel = temp.Channel
    a.Disabled = temp.Disabled
    a.DynamicChainingEnabled = temp.DynamicChainingEnabled
    a.Power = temp.Power
    a.PowerMax = temp.PowerMax
    a.PowerMin = temp.PowerMin
    a.RxChain = temp.RxChain
    a.TxChain = temp.TxChain
    return nil
}

// apStatsRadioConfigBand is a temporary struct used for validating the fields of ApStatsRadioConfigBand.
type apStatsRadioConfigBand  struct {
    AllowRrmDisable        Optional[bool]    `json:"allow_rrm_disable"`
    Bandwidth              Optional[float64] `json:"bandwidth"`
    Channel                *int              `json:"channel,omitempty"`
    Disabled               Optional[bool]    `json:"disabled"`
    DynamicChainingEnabled Optional[bool]    `json:"dynamic_chaining_enabled"`
    Power                  Optional[float64] `json:"power"`
    PowerMax               Optional[float64] `json:"power_max"`
    PowerMin               Optional[float64] `json:"power_min"`
    RxChain                Optional[int]     `json:"rx_chain"`
    TxChain                Optional[int]     `json:"tx_chain"`
}

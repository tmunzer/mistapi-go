package models

import (
    "encoding/json"
)

// MxedgeStatsLagStat represents a MxedgeStatsLagStat struct.
type MxedgeStatsLagStat struct {
    // list of ports active on the LAG defined by the LACP
    ActivePorts          []string       `json:"active_ports,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeStatsLagStat.
// It customizes the JSON marshaling process for MxedgeStatsLagStat objects.
func (m MxedgeStatsLagStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeStatsLagStat object to a map representation for JSON marshaling.
func (m MxedgeStatsLagStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.ActivePorts != nil {
        structMap["active_ports"] = m.ActivePorts
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeStatsLagStat.
// It customizes the JSON unmarshaling process for MxedgeStatsLagStat objects.
func (m *MxedgeStatsLagStat) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeStatsLagStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "active_ports")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.ActivePorts = temp.ActivePorts
    return nil
}

// tempMxedgeStatsLagStat is a temporary struct used for validating the fields of MxedgeStatsLagStat.
type tempMxedgeStatsLagStat  struct {
    ActivePorts []string `json:"active_ports,omitempty"`
}

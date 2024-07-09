package models

import (
    "encoding/json"
)

// MxedgeStatsTuntermStat represents a MxedgeStatsTuntermStat struct.
type MxedgeStatsTuntermStat struct {
    MonitoringFailed     *bool          `json:"monitoring_failed,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeStatsTuntermStat.
// It customizes the JSON marshaling process for MxedgeStatsTuntermStat objects.
func (m MxedgeStatsTuntermStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeStatsTuntermStat object to a map representation for JSON marshaling.
func (m MxedgeStatsTuntermStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.MonitoringFailed != nil {
        structMap["monitoring_failed"] = m.MonitoringFailed
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeStatsTuntermStat.
// It customizes the JSON unmarshaling process for MxedgeStatsTuntermStat objects.
func (m *MxedgeStatsTuntermStat) UnmarshalJSON(input []byte) error {
    var temp mxedgeStatsTuntermStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "monitoring_failed")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.MonitoringFailed = temp.MonitoringFailed
    return nil
}

// mxedgeStatsTuntermStat is a temporary struct used for validating the fields of MxedgeStatsTuntermStat.
type mxedgeStatsTuntermStat  struct {
    MonitoringFailed *bool `json:"monitoring_failed,omitempty"`
}

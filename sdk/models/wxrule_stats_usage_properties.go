package models

import (
    "encoding/json"
)

// WxruleStatsUsageProperties represents a WxruleStatsUsageProperties struct.
type WxruleStatsUsageProperties struct {
    NumFlows             *int           `json:"num_flows,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WxruleStatsUsageProperties.
// It customizes the JSON marshaling process for WxruleStatsUsageProperties objects.
func (w WxruleStatsUsageProperties) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WxruleStatsUsageProperties object to a map representation for JSON marshaling.
func (w WxruleStatsUsageProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.NumFlows != nil {
        structMap["num_flows"] = w.NumFlows
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WxruleStatsUsageProperties.
// It customizes the JSON unmarshaling process for WxruleStatsUsageProperties objects.
func (w *WxruleStatsUsageProperties) UnmarshalJSON(input []byte) error {
    var temp wxruleStatsUsageProperties
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "num_flows")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.NumFlows = temp.NumFlows
    return nil
}

// wxruleStatsUsageProperties is a temporary struct used for validating the fields of WxruleStatsUsageProperties.
type wxruleStatsUsageProperties  struct {
    NumFlows *int `json:"num_flows,omitempty"`
}

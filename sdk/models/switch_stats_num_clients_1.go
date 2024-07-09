package models

import (
    "encoding/json"
)

// SwitchStatsNumClients1 represents a SwitchStatsNumClients1 struct.
// how many wireless clients are currently connected
type SwitchStatsNumClients1 struct {
    Total                *SwitchStatsNumClientsTotal `json:"total,omitempty"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchStatsNumClients1.
// It customizes the JSON marshaling process for SwitchStatsNumClients1 objects.
func (s SwitchStatsNumClients1) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchStatsNumClients1 object to a map representation for JSON marshaling.
func (s SwitchStatsNumClients1) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Total != nil {
        structMap["total"] = s.Total.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchStatsNumClients1.
// It customizes the JSON unmarshaling process for SwitchStatsNumClients1 objects.
func (s *SwitchStatsNumClients1) UnmarshalJSON(input []byte) error {
    var temp switchStatsNumClients1
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "total")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Total = temp.Total
    return nil
}

// switchStatsNumClients1 is a temporary struct used for validating the fields of SwitchStatsNumClients1.
type switchStatsNumClients1  struct {
    Total *SwitchStatsNumClientsTotal `json:"total,omitempty"`
}

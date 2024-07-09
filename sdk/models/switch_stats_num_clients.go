package models

import (
    "encoding/json"
)

// SwitchStatsNumClients represents a SwitchStatsNumClients struct.
type SwitchStatsNumClients struct {
    Total                *SwitchStatsNumClientsTotal `json:"total,omitempty"`
    AdditionalProperties map[string]any              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchStatsNumClients.
// It customizes the JSON marshaling process for SwitchStatsNumClients objects.
func (s SwitchStatsNumClients) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchStatsNumClients object to a map representation for JSON marshaling.
func (s SwitchStatsNumClients) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Total != nil {
        structMap["total"] = s.Total.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchStatsNumClients.
// It customizes the JSON unmarshaling process for SwitchStatsNumClients objects.
func (s *SwitchStatsNumClients) UnmarshalJSON(input []byte) error {
    var temp switchStatsNumClients
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

// switchStatsNumClients is a temporary struct used for validating the fields of SwitchStatsNumClients.
type switchStatsNumClients  struct {
    Total *SwitchStatsNumClientsTotal `json:"total,omitempty"`
}

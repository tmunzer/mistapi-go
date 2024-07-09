package models

import (
    "encoding/json"
)

// SwitchStatsNumClientsTotalProperties represents a SwitchStatsNumClientsTotalProperties struct.
// propery key is the FPC number
type SwitchStatsNumClientsTotalProperties struct {
    NumAps               *string        `json:"num_aps,omitempty"`
    NumClients           *string        `json:"num_clients,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SwitchStatsNumClientsTotalProperties.
// It customizes the JSON marshaling process for SwitchStatsNumClientsTotalProperties objects.
func (s SwitchStatsNumClientsTotalProperties) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SwitchStatsNumClientsTotalProperties object to a map representation for JSON marshaling.
func (s SwitchStatsNumClientsTotalProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.NumAps != nil {
        structMap["num_aps"] = s.NumAps
    }
    if s.NumClients != nil {
        structMap["num_clients"] = s.NumClients
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchStatsNumClientsTotalProperties.
// It customizes the JSON unmarshaling process for SwitchStatsNumClientsTotalProperties objects.
func (s *SwitchStatsNumClientsTotalProperties) UnmarshalJSON(input []byte) error {
    var temp switchStatsNumClientsTotalProperties
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "num_aps", "num_clients")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.NumAps = temp.NumAps
    s.NumClients = temp.NumClients
    return nil
}

// switchStatsNumClientsTotalProperties is a temporary struct used for validating the fields of SwitchStatsNumClientsTotalProperties.
type switchStatsNumClientsTotalProperties  struct {
    NumAps     *string `json:"num_aps,omitempty"`
    NumClients *string `json:"num_clients,omitempty"`
}

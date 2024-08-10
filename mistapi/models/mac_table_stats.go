package models

import (
    "encoding/json"
)

// MacTableStats represents a MacTableStats struct.
type MacTableStats struct {
    MacTableCount          *int           `json:"mac_table_count,omitempty"`
    MaxMacEntriesSupported *int           `json:"max_mac_entries_supported,omitempty"`
    AdditionalProperties   map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MacTableStats.
// It customizes the JSON marshaling process for MacTableStats objects.
func (m MacTableStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MacTableStats object to a map representation for JSON marshaling.
func (m MacTableStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.MacTableCount != nil {
        structMap["mac_table_count"] = m.MacTableCount
    }
    if m.MaxMacEntriesSupported != nil {
        structMap["max_mac_entries_supported"] = m.MaxMacEntriesSupported
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MacTableStats.
// It customizes the JSON unmarshaling process for MacTableStats objects.
func (m *MacTableStats) UnmarshalJSON(input []byte) error {
    var temp tempMacTableStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mac_table_count", "max_mac_entries_supported")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.MacTableCount = temp.MacTableCount
    m.MaxMacEntriesSupported = temp.MaxMacEntriesSupported
    return nil
}

// tempMacTableStats is a temporary struct used for validating the fields of MacTableStats.
type tempMacTableStats  struct {
    MacTableCount          *int `json:"mac_table_count,omitempty"`
    MaxMacEntriesSupported *int `json:"max_mac_entries_supported,omitempty"`
}

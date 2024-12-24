package models

import (
    "encoding/json"
    "fmt"
)

// MacTableStats represents a MacTableStats struct.
type MacTableStats struct {
    MacTableCount          *int                   `json:"mac_table_count,omitempty"`
    MaxMacEntriesSupported *int                   `json:"max_mac_entries_supported,omitempty"`
    AdditionalProperties   map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MacTableStats,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MacTableStats) String() string {
    return fmt.Sprintf(
    	"MacTableStats[MacTableCount=%v, MaxMacEntriesSupported=%v, AdditionalProperties=%v]",
    	m.MacTableCount, m.MaxMacEntriesSupported, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MacTableStats.
// It customizes the JSON marshaling process for MacTableStats objects.
func (m MacTableStats) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "mac_table_count", "max_mac_entries_supported"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MacTableStats object to a map representation for JSON marshaling.
func (m MacTableStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac_table_count", "max_mac_entries_supported")
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

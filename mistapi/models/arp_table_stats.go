package models

import (
    "encoding/json"
    "fmt"
)

// ArpTableStats represents a ArpTableStats struct.
type ArpTableStats struct {
    ArpTableCount        *int                   `json:"arp_table_count,omitempty"`
    MaxEntriesSupported  *int                   `json:"max_entries_supported,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ArpTableStats,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ArpTableStats) String() string {
    return fmt.Sprintf(
    	"ArpTableStats[ArpTableCount=%v, MaxEntriesSupported=%v, AdditionalProperties=%v]",
    	a.ArpTableCount, a.MaxEntriesSupported, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ArpTableStats.
// It customizes the JSON marshaling process for ArpTableStats objects.
func (a ArpTableStats) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "arp_table_count", "max_entries_supported"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ArpTableStats object to a map representation for JSON marshaling.
func (a ArpTableStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.ArpTableCount != nil {
        structMap["arp_table_count"] = a.ArpTableCount
    }
    if a.MaxEntriesSupported != nil {
        structMap["max_entries_supported"] = a.MaxEntriesSupported
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ArpTableStats.
// It customizes the JSON unmarshaling process for ArpTableStats objects.
func (a *ArpTableStats) UnmarshalJSON(input []byte) error {
    var temp tempArpTableStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "arp_table_count", "max_entries_supported")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.ArpTableCount = temp.ArpTableCount
    a.MaxEntriesSupported = temp.MaxEntriesSupported
    return nil
}

// tempArpTableStats is a temporary struct used for validating the fields of ArpTableStats.
type tempArpTableStats  struct {
    ArpTableCount       *int `json:"arp_table_count,omitempty"`
    MaxEntriesSupported *int `json:"max_entries_supported,omitempty"`
}

package models

import (
    "encoding/json"
    "fmt"
)

// SiteSettingSsr represents a SiteSettingSsr struct.
type SiteSettingSsr struct {
    ConductorHosts       []string               `json:"conductor_hosts,omitempty"`
    DisableStats         *bool                  `json:"disable_stats,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingSsr,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingSsr) String() string {
    return fmt.Sprintf(
    	"SiteSettingSsr[ConductorHosts=%v, DisableStats=%v, AdditionalProperties=%v]",
    	s.ConductorHosts, s.DisableStats, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingSsr.
// It customizes the JSON marshaling process for SiteSettingSsr objects.
func (s SiteSettingSsr) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "conductor_hosts", "disable_stats"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingSsr object to a map representation for JSON marshaling.
func (s SiteSettingSsr) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ConductorHosts != nil {
        structMap["conductor_hosts"] = s.ConductorHosts
    }
    if s.DisableStats != nil {
        structMap["disable_stats"] = s.DisableStats
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingSsr.
// It customizes the JSON unmarshaling process for SiteSettingSsr objects.
func (s *SiteSettingSsr) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingSsr
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "conductor_hosts", "disable_stats")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ConductorHosts = temp.ConductorHosts
    s.DisableStats = temp.DisableStats
    return nil
}

// tempSiteSettingSsr is a temporary struct used for validating the fields of SiteSettingSsr.
type tempSiteSettingSsr  struct {
    ConductorHosts []string `json:"conductor_hosts,omitempty"`
    DisableStats   *bool    `json:"disable_stats,omitempty"`
}

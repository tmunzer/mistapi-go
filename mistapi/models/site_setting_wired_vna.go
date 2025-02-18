package models

import (
    "encoding/json"
    "fmt"
)

// SiteSettingWiredVna represents a SiteSettingWiredVna struct.
type SiteSettingWiredVna struct {
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingWiredVna,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingWiredVna) String() string {
    return fmt.Sprintf(
    	"SiteSettingWiredVna[Enabled=%v, AdditionalProperties=%v]",
    	s.Enabled, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingWiredVna.
// It customizes the JSON marshaling process for SiteSettingWiredVna objects.
func (s SiteSettingWiredVna) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingWiredVna object to a map representation for JSON marshaling.
func (s SiteSettingWiredVna) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Enabled != nil {
        structMap["enabled"] = s.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingWiredVna.
// It customizes the JSON unmarshaling process for SiteSettingWiredVna objects.
func (s *SiteSettingWiredVna) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingWiredVna
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enabled")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Enabled = temp.Enabled
    return nil
}

// tempSiteSettingWiredVna is a temporary struct used for validating the fields of SiteSettingWiredVna.
type tempSiteSettingWiredVna  struct {
    Enabled *bool `json:"enabled,omitempty"`
}

package models

import (
    "encoding/json"
    "fmt"
)

// SiteSettingApPortConfig represents a SiteSettingApPortConfig struct.
type SiteSettingApPortConfig struct {
    // Property key is the AP model (e.g "AP32")
    ModelSpecific        map[string]ApPortConfig `json:"model_specific,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingApPortConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingApPortConfig) String() string {
    return fmt.Sprintf(
    	"SiteSettingApPortConfig[ModelSpecific=%v, AdditionalProperties=%v]",
    	s.ModelSpecific, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingApPortConfig.
// It customizes the JSON marshaling process for SiteSettingApPortConfig objects.
func (s SiteSettingApPortConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "model_specific"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingApPortConfig object to a map representation for JSON marshaling.
func (s SiteSettingApPortConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.ModelSpecific != nil {
        structMap["model_specific"] = s.ModelSpecific
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingApPortConfig.
// It customizes the JSON unmarshaling process for SiteSettingApPortConfig objects.
func (s *SiteSettingApPortConfig) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingApPortConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "model_specific")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.ModelSpecific = temp.ModelSpecific
    return nil
}

// tempSiteSettingApPortConfig is a temporary struct used for validating the fields of SiteSettingApPortConfig.
type tempSiteSettingApPortConfig  struct {
    ModelSpecific map[string]ApPortConfig `json:"model_specific,omitempty"`
}

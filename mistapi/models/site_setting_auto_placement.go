package models

import (
    "encoding/json"
)

// SiteSettingAutoPlacement represents a SiteSettingAutoPlacement struct.
// if we're able to determine its x/y/orientation, this will be populated
type SiteSettingAutoPlacement struct {
    Orientation          *int                   `json:"orientation,omitempty"`
    X                    *float64               `json:"x,omitempty"`
    Y                    *float64               `json:"y,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingAutoPlacement.
// It customizes the JSON marshaling process for SiteSettingAutoPlacement objects.
func (s SiteSettingAutoPlacement) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "orientation", "x", "y"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingAutoPlacement object to a map representation for JSON marshaling.
func (s SiteSettingAutoPlacement) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Orientation != nil {
        structMap["orientation"] = s.Orientation
    }
    if s.X != nil {
        structMap["x"] = s.X
    }
    if s.Y != nil {
        structMap["y"] = s.Y
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingAutoPlacement.
// It customizes the JSON unmarshaling process for SiteSettingAutoPlacement objects.
func (s *SiteSettingAutoPlacement) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingAutoPlacement
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "orientation", "x", "y")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Orientation = temp.Orientation
    s.X = temp.X
    s.Y = temp.Y
    return nil
}

// tempSiteSettingAutoPlacement is a temporary struct used for validating the fields of SiteSettingAutoPlacement.
type tempSiteSettingAutoPlacement  struct {
    Orientation *int     `json:"orientation,omitempty"`
    X           *float64 `json:"x,omitempty"`
    Y           *float64 `json:"y,omitempty"`
}

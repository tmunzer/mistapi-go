package models

import (
    "encoding/json"
)

// AppProbing represents a AppProbing struct.
type AppProbing struct {
    Apps                 []string              `json:"apps,omitempty"`
    CustomApps           []AppProbingCustomApp `json:"custom_apps,omitempty"`
    Enabled              *bool                 `json:"enabled,omitempty"`
    AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AppProbing.
// It customizes the JSON marshaling process for AppProbing objects.
func (a AppProbing) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AppProbing object to a map representation for JSON marshaling.
func (a AppProbing) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Apps != nil {
        structMap["apps"] = a.Apps
    }
    if a.CustomApps != nil {
        structMap["custom_apps"] = a.CustomApps
    }
    if a.Enabled != nil {
        structMap["enabled"] = a.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AppProbing.
// It customizes the JSON unmarshaling process for AppProbing objects.
func (a *AppProbing) UnmarshalJSON(input []byte) error {
    var temp appProbing
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "apps", "custom_apps", "enabled")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Apps = temp.Apps
    a.CustomApps = temp.CustomApps
    a.Enabled = temp.Enabled
    return nil
}

// appProbing is a temporary struct used for validating the fields of AppProbing.
type appProbing  struct {
    Apps       []string              `json:"apps,omitempty"`
    CustomApps []AppProbingCustomApp `json:"custom_apps,omitempty"`
    Enabled    *bool                 `json:"enabled,omitempty"`
}

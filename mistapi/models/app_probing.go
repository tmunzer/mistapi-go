package models

import (
    "encoding/json"
    "fmt"
)

// AppProbing represents a AppProbing struct.
type AppProbing struct {
    // app-keys from /api/v1/const/applications
    Apps                 []string               `json:"apps,omitempty"`
    CustomApps           []AppProbingCustomApp  `json:"custom_apps,omitempty"`
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AppProbing,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AppProbing) String() string {
    return fmt.Sprintf(
    	"AppProbing[Apps=%v, CustomApps=%v, Enabled=%v, AdditionalProperties=%v]",
    	a.Apps, a.CustomApps, a.Enabled, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AppProbing.
// It customizes the JSON marshaling process for AppProbing objects.
func (a AppProbing) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "apps", "custom_apps", "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AppProbing object to a map representation for JSON marshaling.
func (a AppProbing) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
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
    var temp tempAppProbing
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "apps", "custom_apps", "enabled")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Apps = temp.Apps
    a.CustomApps = temp.CustomApps
    a.Enabled = temp.Enabled
    return nil
}

// tempAppProbing is a temporary struct used for validating the fields of AppProbing.
type tempAppProbing  struct {
    Apps       []string              `json:"apps,omitempty"`
    CustomApps []AppProbingCustomApp `json:"custom_apps,omitempty"`
    Enabled    *bool                 `json:"enabled,omitempty"`
}

// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SiteSettingConfigPushPolicy represents a SiteSettingConfigPushPolicy struct.
// Mist also uses some heuristic rules to prevent destructive configs from being pushed
type SiteSettingConfigPushPolicy struct {
    // Stop any new config from being pushed to the device
    NoPush               *bool                  `json:"no_push,omitempty"`
    // If enabled, new config will only be pushed to device within the specified time window
    PushWindow           *PushPolicyPushWindow  `json:"push_window,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingConfigPushPolicy,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingConfigPushPolicy) String() string {
    return fmt.Sprintf(
    	"SiteSettingConfigPushPolicy[NoPush=%v, PushWindow=%v, AdditionalProperties=%v]",
    	s.NoPush, s.PushWindow, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingConfigPushPolicy.
// It customizes the JSON marshaling process for SiteSettingConfigPushPolicy objects.
func (s SiteSettingConfigPushPolicy) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "no_push", "push_window"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingConfigPushPolicy object to a map representation for JSON marshaling.
func (s SiteSettingConfigPushPolicy) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.NoPush != nil {
        structMap["no_push"] = s.NoPush
    }
    if s.PushWindow != nil {
        structMap["push_window"] = s.PushWindow.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingConfigPushPolicy.
// It customizes the JSON unmarshaling process for SiteSettingConfigPushPolicy objects.
func (s *SiteSettingConfigPushPolicy) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingConfigPushPolicy
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "no_push", "push_window")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.NoPush = temp.NoPush
    s.PushWindow = temp.PushWindow
    return nil
}

// tempSiteSettingConfigPushPolicy is a temporary struct used for validating the fields of SiteSettingConfigPushPolicy.
type tempSiteSettingConfigPushPolicy  struct {
    NoPush     *bool                 `json:"no_push,omitempty"`
    PushWindow *PushPolicyPushWindow `json:"push_window,omitempty"`
}

// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// WlanAppLimit represents a WlanAppLimit struct.
// Bandwidth limiting for apps (applies to up/down)
type WlanAppLimit struct {
    // Map from app key to bandwidth in kbps.
    // Property key is the app key, defined in Get Application List
    Apps                 map[string]int         `json:"apps,omitempty"`
    Enabled              *bool                  `json:"enabled,omitempty"`
    // Map from wxtag_id of Hostname Wxlan Tags to bandwidth in kbps. Property key is the `wxtag_id`
    WxtagIds             map[string]int         `json:"wxtag_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WlanAppLimit,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanAppLimit) String() string {
    return fmt.Sprintf(
    	"WlanAppLimit[Apps=%v, Enabled=%v, WxtagIds=%v, AdditionalProperties=%v]",
    	w.Apps, w.Enabled, w.WxtagIds, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanAppLimit.
// It customizes the JSON marshaling process for WlanAppLimit objects.
func (w WlanAppLimit) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "apps", "enabled", "wxtag_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanAppLimit object to a map representation for JSON marshaling.
func (w WlanAppLimit) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Apps != nil {
        structMap["apps"] = w.Apps
    }
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    if w.WxtagIds != nil {
        structMap["wxtag_ids"] = w.WxtagIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanAppLimit.
// It customizes the JSON unmarshaling process for WlanAppLimit objects.
func (w *WlanAppLimit) UnmarshalJSON(input []byte) error {
    var temp tempWlanAppLimit
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "apps", "enabled", "wxtag_ids")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Apps = temp.Apps
    w.Enabled = temp.Enabled
    w.WxtagIds = temp.WxtagIds
    return nil
}

// tempWlanAppLimit is a temporary struct used for validating the fields of WlanAppLimit.
type tempWlanAppLimit  struct {
    Apps     map[string]int `json:"apps,omitempty"`
    Enabled  *bool          `json:"enabled,omitempty"`
    WxtagIds map[string]int `json:"wxtag_ids,omitempty"`
}

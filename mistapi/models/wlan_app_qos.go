// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// WlanAppQos represents a WlanAppQos struct.
// APp qos wlan settings
type WlanAppQos struct {
    Apps                 map[string]WlanAppQosAppsProperties `json:"apps,omitempty"`
    Enabled              *bool                               `json:"enabled,omitempty"`
    Others               []WlanAppQosOthersItem              `json:"others,omitempty"`
    AdditionalProperties map[string]interface{}              `json:"_"`
}

// String implements the fmt.Stringer interface for WlanAppQos,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanAppQos) String() string {
    return fmt.Sprintf(
    	"WlanAppQos[Apps=%v, Enabled=%v, Others=%v, AdditionalProperties=%v]",
    	w.Apps, w.Enabled, w.Others, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanAppQos.
// It customizes the JSON marshaling process for WlanAppQos objects.
func (w WlanAppQos) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "apps", "enabled", "others"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanAppQos object to a map representation for JSON marshaling.
func (w WlanAppQos) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Apps != nil {
        structMap["apps"] = w.Apps
    }
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    if w.Others != nil {
        structMap["others"] = w.Others
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanAppQos.
// It customizes the JSON unmarshaling process for WlanAppQos objects.
func (w *WlanAppQos) UnmarshalJSON(input []byte) error {
    var temp tempWlanAppQos
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "apps", "enabled", "others")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Apps = temp.Apps
    w.Enabled = temp.Enabled
    w.Others = temp.Others
    return nil
}

// tempWlanAppQos is a temporary struct used for validating the fields of WlanAppQos.
type tempWlanAppQos  struct {
    Apps    map[string]WlanAppQosAppsProperties `json:"apps,omitempty"`
    Enabled *bool                               `json:"enabled,omitempty"`
    Others  []WlanAppQosOthersItem              `json:"others,omitempty"`
}

package models

import (
    "encoding/json"
)

// WlanAppQos represents a WlanAppQos struct.
// app qos wlan settings
type WlanAppQos struct {
    Apps                 map[string]WlanAppQosAppsProperties `json:"apps,omitempty"`
    Enabled              *bool                               `json:"enabled,omitempty"`
    Others               []WlanAppQosOthersItem              `json:"others,omitempty"`
    AdditionalProperties map[string]any                      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanAppQos.
// It customizes the JSON marshaling process for WlanAppQos objects.
func (w WlanAppQos) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WlanAppQos object to a map representation for JSON marshaling.
func (w WlanAppQos) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "apps", "enabled", "others")
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

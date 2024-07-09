package models

import (
    "encoding/json"
)

// WlanBonjourServiceProperties represents a WlanBonjourServiceProperties struct.
type WlanBonjourServiceProperties struct {
    // whether to prevent wireless clients to discover bonjour devices on the same WLAN
    DisableLocal         *bool                                  `json:"disable_local,omitempty"`
    // optional, if the service is further restricted for certain RADIUS groups
    RadiusGroups         []string                               `json:"radius_groups,omitempty"`
    // how bonjour services should be discovered for the same WLAN
    Scope                *WlanBonjourServicePropertiesScopeEnum `json:"scope,omitempty"`
    AdditionalProperties map[string]any                         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanBonjourServiceProperties.
// It customizes the JSON marshaling process for WlanBonjourServiceProperties objects.
func (w WlanBonjourServiceProperties) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WlanBonjourServiceProperties object to a map representation for JSON marshaling.
func (w WlanBonjourServiceProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.DisableLocal != nil {
        structMap["disable_local"] = w.DisableLocal
    }
    if w.RadiusGroups != nil {
        structMap["radius_groups"] = w.RadiusGroups
    }
    if w.Scope != nil {
        structMap["scope"] = w.Scope
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanBonjourServiceProperties.
// It customizes the JSON unmarshaling process for WlanBonjourServiceProperties objects.
func (w *WlanBonjourServiceProperties) UnmarshalJSON(input []byte) error {
    var temp wlanBonjourServiceProperties
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "disable_local", "radius_groups", "scope")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.DisableLocal = temp.DisableLocal
    w.RadiusGroups = temp.RadiusGroups
    w.Scope = temp.Scope
    return nil
}

// wlanBonjourServiceProperties is a temporary struct used for validating the fields of WlanBonjourServiceProperties.
type wlanBonjourServiceProperties  struct {
    DisableLocal *bool                                  `json:"disable_local,omitempty"`
    RadiusGroups []string                               `json:"radius_groups,omitempty"`
    Scope        *WlanBonjourServicePropertiesScopeEnum `json:"scope,omitempty"`
}
package models

import (
    "encoding/json"
    "fmt"
)

// WlanBonjourServiceProperties represents a WlanBonjourServiceProperties struct.
type WlanBonjourServiceProperties struct {
    // Whether to prevent wireless clients to discover bonjour devices on the same WLAN
    DisableLocal         *bool                                  `json:"disable_local,omitempty"`
    // Optional, if the service is further restricted for certain RADIUS groups
    RadiusGroups         []string                               `json:"radius_groups,omitempty"`
    // how bonjour services should be discovered for the same WLAN. enum: `same_ap`, `same_map`, `same_site`
    Scope                *WlanBonjourServicePropertiesScopeEnum `json:"scope,omitempty"`
    AdditionalProperties map[string]interface{}                 `json:"_"`
}

// String implements the fmt.Stringer interface for WlanBonjourServiceProperties,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanBonjourServiceProperties) String() string {
    return fmt.Sprintf(
    	"WlanBonjourServiceProperties[DisableLocal=%v, RadiusGroups=%v, Scope=%v, AdditionalProperties=%v]",
    	w.DisableLocal, w.RadiusGroups, w.Scope, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanBonjourServiceProperties.
// It customizes the JSON marshaling process for WlanBonjourServiceProperties objects.
func (w WlanBonjourServiceProperties) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "disable_local", "radius_groups", "scope"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanBonjourServiceProperties object to a map representation for JSON marshaling.
func (w WlanBonjourServiceProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
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
    var temp tempWlanBonjourServiceProperties
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "disable_local", "radius_groups", "scope")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.DisableLocal = temp.DisableLocal
    w.RadiusGroups = temp.RadiusGroups
    w.Scope = temp.Scope
    return nil
}

// tempWlanBonjourServiceProperties is a temporary struct used for validating the fields of WlanBonjourServiceProperties.
type tempWlanBonjourServiceProperties  struct {
    DisableLocal *bool                                  `json:"disable_local,omitempty"`
    RadiusGroups []string                               `json:"radius_groups,omitempty"`
    Scope        *WlanBonjourServicePropertiesScopeEnum `json:"scope,omitempty"`
}

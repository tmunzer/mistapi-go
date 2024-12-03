package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WlanBonjour represents a WlanBonjour struct.
// bonjour gateway wlan settings
type WlanBonjour struct {
    // comma sperated list of additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses
    AdditionalVlanIds    string                                  `json:"additional_vlan_ids"`
    // whether to enable bonjour for this WLAN. Once enabled, limit_bcast is assumed true, allow_mdns is assumed false
    Enabled              *bool                                   `json:"enabled,omitempty"`
    // what services are allowed.
    // Property key is the service name
    Services             map[string]WlanBonjourServiceProperties `json:"services"`
    AdditionalProperties map[string]interface{}                  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanBonjour.
// It customizes the JSON marshaling process for WlanBonjour objects.
func (w WlanBonjour) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "additional_vlan_ids", "enabled", "services"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanBonjour object to a map representation for JSON marshaling.
func (w WlanBonjour) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["additional_vlan_ids"] = w.AdditionalVlanIds
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    structMap["services"] = w.Services
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanBonjour.
// It customizes the JSON unmarshaling process for WlanBonjour objects.
func (w *WlanBonjour) UnmarshalJSON(input []byte) error {
    var temp tempWlanBonjour
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "additional_vlan_ids", "enabled", "services")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.AdditionalVlanIds = *temp.AdditionalVlanIds
    w.Enabled = temp.Enabled
    w.Services = *temp.Services
    return nil
}

// tempWlanBonjour is a temporary struct used for validating the fields of WlanBonjour.
type tempWlanBonjour  struct {
    AdditionalVlanIds *string                                  `json:"additional_vlan_ids"`
    Enabled           *bool                                    `json:"enabled,omitempty"`
    Services          *map[string]WlanBonjourServiceProperties `json:"services"`
}

func (w *tempWlanBonjour) validate() error {
    var errs []string
    if w.AdditionalVlanIds == nil {
        errs = append(errs, "required field `additional_vlan_ids` is missing for type `wlan_bonjour`")
    }
    if w.Services == nil {
        errs = append(errs, "required field `services` is missing for type `wlan_bonjour`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

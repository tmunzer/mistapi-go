package models

import (
    "encoding/json"
)

// ApUplinkPortConfig represents a ApUplinkPortConfig struct.
type ApUplinkPortConfig struct {
    // Whether to do 802.1x against uplink switch. When enaled, AP cert will be used to do EAP-TLS and the Org's CA Cert has to be provisioned at the switch
    Dot1x                *bool          `json:"dot1x,omitempty"`
    // by default, WLANs are disabled when uplink is down. In some scenario, like SiteSurvey, one would want the AP to keep sending beacons.
    KeepWlansUpIfDown    *bool          `json:"keep_wlans_up_if_down,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApUplinkPortConfig.
// It customizes the JSON marshaling process for ApUplinkPortConfig objects.
func (a ApUplinkPortConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApUplinkPortConfig object to a map representation for JSON marshaling.
func (a ApUplinkPortConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Dot1x != nil {
        structMap["dot1x"] = a.Dot1x
    }
    if a.KeepWlansUpIfDown != nil {
        structMap["keep_wlans_up_if_down"] = a.KeepWlansUpIfDown
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApUplinkPortConfig.
// It customizes the JSON unmarshaling process for ApUplinkPortConfig objects.
func (a *ApUplinkPortConfig) UnmarshalJSON(input []byte) error {
    var temp apUplinkPortConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "dot1x", "keep_wlans_up_if_down")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Dot1x = temp.Dot1x
    a.KeepWlansUpIfDown = temp.KeepWlansUpIfDown
    return nil
}

// apUplinkPortConfig is a temporary struct used for validating the fields of ApUplinkPortConfig.
type apUplinkPortConfig  struct {
    Dot1x             *bool `json:"dot1x,omitempty"`
    KeepWlansUpIfDown *bool `json:"keep_wlans_up_if_down,omitempty"`
}
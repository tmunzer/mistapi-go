package models

import (
    "encoding/json"
    "fmt"
)

// ApUplinkPortConfig represents a ApUplinkPortConfig struct.
// AP Uplink port configuration
type ApUplinkPortConfig struct {
    // Whether to do 802.1x against uplink switch. When enaled, AP cert will be used to do EAP-TLS and the Org's CA Cert has to be provisioned at the switch
    Dot1x                *bool                  `json:"dot1x,omitempty"`
    // By default, WLANs are disabled when uplink is down. In some scenario, like SiteSurvey, one would want the AP to keep sending beacons.
    KeepWlansUpIfDown    *bool                  `json:"keep_wlans_up_if_down,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApUplinkPortConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApUplinkPortConfig) String() string {
    return fmt.Sprintf(
    	"ApUplinkPortConfig[Dot1x=%v, KeepWlansUpIfDown=%v, AdditionalProperties=%v]",
    	a.Dot1x, a.KeepWlansUpIfDown, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApUplinkPortConfig.
// It customizes the JSON marshaling process for ApUplinkPortConfig objects.
func (a ApUplinkPortConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "dot1x", "keep_wlans_up_if_down"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApUplinkPortConfig object to a map representation for JSON marshaling.
func (a ApUplinkPortConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
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
    var temp tempApUplinkPortConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dot1x", "keep_wlans_up_if_down")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Dot1x = temp.Dot1x
    a.KeepWlansUpIfDown = temp.KeepWlansUpIfDown
    return nil
}

// tempApUplinkPortConfig is a temporary struct used for validating the fields of ApUplinkPortConfig.
type tempApUplinkPortConfig  struct {
    Dot1x             *bool `json:"dot1x,omitempty"`
    KeepWlansUpIfDown *bool `json:"keep_wlans_up_if_down,omitempty"`
}

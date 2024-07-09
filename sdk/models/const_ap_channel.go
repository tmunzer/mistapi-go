package models

import (
    "encoding/json"
)

// ConstApChannel represents a ConstApChannel struct.
type ConstApChannel struct {
    Band2440mhzAllowed   *bool                  `json:"band24_40mhz_allowed,omitempty"`
    // Property key is the channel width
    Band24Channels       map[string]interface{} `json:"band24_channels,omitempty"`
    Band24Enabled        *bool                  `json:"band24_enabled,omitempty"`
    // Property key is the channel width
    Band5Channels        map[string]interface{} `json:"band5_channels,omitempty"`
    Band5Enabled         *bool                  `json:"band5_enabled,omitempty"`
    // Property key is the channel width
    Band6Channels        map[string]interface{} `json:"band6_channels,omitempty"`
    Band6Enabled         *bool                  `json:"band6_enabled,omitempty"`
    Certified            *bool                  `json:"certified,omitempty"`
    // country code, ISO 3166-1 numeric
    Code                 *int                   `json:"code,omitempty"`
    DfsOk                *bool                  `json:"dfs_ok,omitempty"`
    // country code, in two-character
    Key                  *string                `json:"key,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    Uses                 *string                `json:"uses,omitempty"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstApChannel.
// It customizes the JSON marshaling process for ConstApChannel objects.
func (c ConstApChannel) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConstApChannel object to a map representation for JSON marshaling.
func (c ConstApChannel) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Band2440mhzAllowed != nil {
        structMap["band24_40mhz_allowed"] = c.Band2440mhzAllowed
    }
    if c.Band24Channels != nil {
        structMap["band24_channels"] = c.Band24Channels
    }
    if c.Band24Enabled != nil {
        structMap["band24_enabled"] = c.Band24Enabled
    }
    if c.Band5Channels != nil {
        structMap["band5_channels"] = c.Band5Channels
    }
    if c.Band5Enabled != nil {
        structMap["band5_enabled"] = c.Band5Enabled
    }
    if c.Band6Channels != nil {
        structMap["band6_channels"] = c.Band6Channels
    }
    if c.Band6Enabled != nil {
        structMap["band6_enabled"] = c.Band6Enabled
    }
    if c.Certified != nil {
        structMap["certified"] = c.Certified
    }
    if c.Code != nil {
        structMap["code"] = c.Code
    }
    if c.DfsOk != nil {
        structMap["dfs_ok"] = c.DfsOk
    }
    if c.Key != nil {
        structMap["key"] = c.Key
    }
    if c.Name != nil {
        structMap["name"] = c.Name
    }
    if c.Uses != nil {
        structMap["uses"] = c.Uses
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstApChannel.
// It customizes the JSON unmarshaling process for ConstApChannel objects.
func (c *ConstApChannel) UnmarshalJSON(input []byte) error {
    var temp constApChannel
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "band24_40mhz_allowed", "band24_channels", "band24_enabled", "band5_channels", "band5_enabled", "band6_channels", "band6_enabled", "certified", "code", "dfs_ok", "key", "name", "uses")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Band2440mhzAllowed = temp.Band2440mhzAllowed
    c.Band24Channels = temp.Band24Channels
    c.Band24Enabled = temp.Band24Enabled
    c.Band5Channels = temp.Band5Channels
    c.Band5Enabled = temp.Band5Enabled
    c.Band6Channels = temp.Band6Channels
    c.Band6Enabled = temp.Band6Enabled
    c.Certified = temp.Certified
    c.Code = temp.Code
    c.DfsOk = temp.DfsOk
    c.Key = temp.Key
    c.Name = temp.Name
    c.Uses = temp.Uses
    return nil
}

// constApChannel is a temporary struct used for validating the fields of ConstApChannel.
type constApChannel  struct {
    Band2440mhzAllowed *bool                  `json:"band24_40mhz_allowed,omitempty"`
    Band24Channels     map[string]interface{} `json:"band24_channels,omitempty"`
    Band24Enabled      *bool                  `json:"band24_enabled,omitempty"`
    Band5Channels      map[string]interface{} `json:"band5_channels,omitempty"`
    Band5Enabled       *bool                  `json:"band5_enabled,omitempty"`
    Band6Channels      map[string]interface{} `json:"band6_channels,omitempty"`
    Band6Enabled       *bool                  `json:"band6_enabled,omitempty"`
    Certified          *bool                  `json:"certified,omitempty"`
    Code               *int                   `json:"code,omitempty"`
    DfsOk              *bool                  `json:"dfs_ok,omitempty"`
    Key                *string                `json:"key,omitempty"`
    Name               *string                `json:"name,omitempty"`
    Uses               *string                `json:"uses,omitempty"`
}

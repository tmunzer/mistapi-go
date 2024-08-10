package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseDeviceRadioChannels represents a ResponseDeviceRadioChannels struct.
type ResponseDeviceRadioChannels struct {
    Band2440mhzAllowed   bool                   `json:"band24_40mhz_allowed"`
    // Property key is the channel width
    Band24Channels       map[string]interface{} `json:"band24_channels"`
    Band24Enabled        bool                   `json:"band24_enabled"`
    // Property key is the channel width
    Band5Channels        map[string]interface{} `json:"band5_channels"`
    Band5Enabled         bool                   `json:"band5_enabled"`
    // Property key is the channel width
    Band6Channels        map[string]interface{} `json:"band6_channels,omitempty"`
    Band6Enabled         *bool                  `json:"band6_enabled,omitempty"`
    Certified            bool                   `json:"certified"`
    Code                 int                    `json:"code"`
    DfsOk                bool                   `json:"dfs_ok"`
    Key                  string                 `json:"key"`
    Name                 string                 `json:"name"`
    Uses                 string                 `json:"uses"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceRadioChannels.
// It customizes the JSON marshaling process for ResponseDeviceRadioChannels objects.
func (r ResponseDeviceRadioChannels) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceRadioChannels object to a map representation for JSON marshaling.
func (r ResponseDeviceRadioChannels) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["band24_40mhz_allowed"] = r.Band2440mhzAllowed
    structMap["band24_channels"] = r.Band24Channels
    structMap["band24_enabled"] = r.Band24Enabled
    structMap["band5_channels"] = r.Band5Channels
    structMap["band5_enabled"] = r.Band5Enabled
    if r.Band6Channels != nil {
        structMap["band6_channels"] = r.Band6Channels
    }
    if r.Band6Enabled != nil {
        structMap["band6_enabled"] = r.Band6Enabled
    }
    structMap["certified"] = r.Certified
    structMap["code"] = r.Code
    structMap["dfs_ok"] = r.DfsOk
    structMap["key"] = r.Key
    structMap["name"] = r.Name
    structMap["uses"] = r.Uses
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceRadioChannels.
// It customizes the JSON unmarshaling process for ResponseDeviceRadioChannels objects.
func (r *ResponseDeviceRadioChannels) UnmarshalJSON(input []byte) error {
    var temp tempResponseDeviceRadioChannels
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "band24_40mhz_allowed", "band24_channels", "band24_enabled", "band5_channels", "band5_enabled", "band6_channels", "band6_enabled", "certified", "code", "dfs_ok", "key", "name", "uses")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Band2440mhzAllowed = *temp.Band2440mhzAllowed
    r.Band24Channels = *temp.Band24Channels
    r.Band24Enabled = *temp.Band24Enabled
    r.Band5Channels = *temp.Band5Channels
    r.Band5Enabled = *temp.Band5Enabled
    r.Band6Channels = temp.Band6Channels
    r.Band6Enabled = temp.Band6Enabled
    r.Certified = *temp.Certified
    r.Code = *temp.Code
    r.DfsOk = *temp.DfsOk
    r.Key = *temp.Key
    r.Name = *temp.Name
    r.Uses = *temp.Uses
    return nil
}

// tempResponseDeviceRadioChannels is a temporary struct used for validating the fields of ResponseDeviceRadioChannels.
type tempResponseDeviceRadioChannels  struct {
    Band2440mhzAllowed *bool                   `json:"band24_40mhz_allowed"`
    Band24Channels     *map[string]interface{} `json:"band24_channels"`
    Band24Enabled      *bool                   `json:"band24_enabled"`
    Band5Channels      *map[string]interface{} `json:"band5_channels"`
    Band5Enabled       *bool                   `json:"band5_enabled"`
    Band6Channels      map[string]interface{}  `json:"band6_channels,omitempty"`
    Band6Enabled       *bool                   `json:"band6_enabled,omitempty"`
    Certified          *bool                   `json:"certified"`
    Code               *int                    `json:"code"`
    DfsOk              *bool                   `json:"dfs_ok"`
    Key                *string                 `json:"key"`
    Name               *string                 `json:"name"`
    Uses               *string                 `json:"uses"`
}

func (r *tempResponseDeviceRadioChannels) validate() error {
    var errs []string
    if r.Band2440mhzAllowed == nil {
        errs = append(errs, "required field `band24_40mhz_allowed` is missing for type `response_device_radio_channels`")
    }
    if r.Band24Channels == nil {
        errs = append(errs, "required field `band24_channels` is missing for type `response_device_radio_channels`")
    }
    if r.Band24Enabled == nil {
        errs = append(errs, "required field `band24_enabled` is missing for type `response_device_radio_channels`")
    }
    if r.Band5Channels == nil {
        errs = append(errs, "required field `band5_channels` is missing for type `response_device_radio_channels`")
    }
    if r.Band5Enabled == nil {
        errs = append(errs, "required field `band5_enabled` is missing for type `response_device_radio_channels`")
    }
    if r.Certified == nil {
        errs = append(errs, "required field `certified` is missing for type `response_device_radio_channels`")
    }
    if r.Code == nil {
        errs = append(errs, "required field `code` is missing for type `response_device_radio_channels`")
    }
    if r.DfsOk == nil {
        errs = append(errs, "required field `dfs_ok` is missing for type `response_device_radio_channels`")
    }
    if r.Key == nil {
        errs = append(errs, "required field `key` is missing for type `response_device_radio_channels`")
    }
    if r.Name == nil {
        errs = append(errs, "required field `name` is missing for type `response_device_radio_channels`")
    }
    if r.Uses == nil {
        errs = append(errs, "required field `uses` is missing for type `response_device_radio_channels`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

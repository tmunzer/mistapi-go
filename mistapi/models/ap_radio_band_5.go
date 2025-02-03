package models

import (
    "encoding/json"
    "fmt"
)

// ApRadioBand5 represents a ApRadioBand5 struct.
// Radio Band AP settings
type ApRadioBand5 struct {
    AllowRrmDisable      *bool                     `json:"allow_rrm_disable,omitempty"`
    AntGain              Optional[int]             `json:"ant_gain"`
    // enum: `1x1`, `2x2`, `3x3`, `4x4`, `default`
    AntennaMode          *RadioBandAntennaModeEnum `json:"antenna_mode,omitempty"`
    // channel width for the 5GHz band. enum: `20`, `40`, `80`
    Bandwidth            *Dot11Bandwidth5Enum      `json:"bandwidth,omitempty"`
    // For Device. (primary) channel for the band, 0 means using the Site Setting
    Channel              Optional[int]             `json:"channel"`
    // For RFTemplates. List of channels, null or empty array means auto
    Channels             Optional[[]int]           `json:"channels"`
    // Whether to disable the radio
    Disabled             *bool                     `json:"disabled,omitempty"`
    // TX power of the radio. For Devices, 0 means auto. -1 / -2 / -3 / …: treated as 0 / -1 / -2 / …
    Power                Optional[int]             `json:"power"`
    // When power=0, max tx power to use, HW-specific values will be used if not set
    PowerMax             Optional[int]             `json:"power_max"`
    // When power=0, min tx power to use, HW-specific values will be used if not set
    PowerMin             Optional[int]             `json:"power_min"`
    // enum: `auto`, `long`, `short`
    Preamble             *RadioBandPreambleEnum    `json:"preamble,omitempty"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for ApRadioBand5,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApRadioBand5) String() string {
    return fmt.Sprintf(
    	"ApRadioBand5[AllowRrmDisable=%v, AntGain=%v, AntennaMode=%v, Bandwidth=%v, Channel=%v, Channels=%v, Disabled=%v, Power=%v, PowerMax=%v, PowerMin=%v, Preamble=%v, AdditionalProperties=%v]",
    	a.AllowRrmDisable, a.AntGain, a.AntennaMode, a.Bandwidth, a.Channel, a.Channels, a.Disabled, a.Power, a.PowerMax, a.PowerMin, a.Preamble, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApRadioBand5.
// It customizes the JSON marshaling process for ApRadioBand5 objects.
func (a ApRadioBand5) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "allow_rrm_disable", "ant_gain", "antenna_mode", "bandwidth", "channel", "channels", "disabled", "power", "power_max", "power_min", "preamble"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the ApRadioBand5 object to a map representation for JSON marshaling.
func (a ApRadioBand5) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.AllowRrmDisable != nil {
        structMap["allow_rrm_disable"] = a.AllowRrmDisable
    }
    if a.AntGain.IsValueSet() {
        if a.AntGain.Value() != nil {
            structMap["ant_gain"] = a.AntGain.Value()
        } else {
            structMap["ant_gain"] = nil
        }
    }
    if a.AntennaMode != nil {
        structMap["antenna_mode"] = a.AntennaMode
    }
    if a.Bandwidth != nil {
        structMap["bandwidth"] = a.Bandwidth
    }
    if a.Channel.IsValueSet() {
        if a.Channel.Value() != nil {
            structMap["channel"] = a.Channel.Value()
        } else {
            structMap["channel"] = nil
        }
    }
    if a.Channels.IsValueSet() {
        if a.Channels.Value() != nil {
            structMap["channels"] = a.Channels.Value()
        } else {
            structMap["channels"] = nil
        }
    }
    if a.Disabled != nil {
        structMap["disabled"] = a.Disabled
    }
    if a.Power.IsValueSet() {
        if a.Power.Value() != nil {
            structMap["power"] = a.Power.Value()
        } else {
            structMap["power"] = nil
        }
    }
    if a.PowerMax.IsValueSet() {
        if a.PowerMax.Value() != nil {
            structMap["power_max"] = a.PowerMax.Value()
        } else {
            structMap["power_max"] = nil
        }
    }
    if a.PowerMin.IsValueSet() {
        if a.PowerMin.Value() != nil {
            structMap["power_min"] = a.PowerMin.Value()
        } else {
            structMap["power_min"] = nil
        }
    }
    if a.Preamble != nil {
        structMap["preamble"] = a.Preamble
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApRadioBand5.
// It customizes the JSON unmarshaling process for ApRadioBand5 objects.
func (a *ApRadioBand5) UnmarshalJSON(input []byte) error {
    var temp tempApRadioBand5
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "allow_rrm_disable", "ant_gain", "antenna_mode", "bandwidth", "channel", "channels", "disabled", "power", "power_max", "power_min", "preamble")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.AllowRrmDisable = temp.AllowRrmDisable
    a.AntGain = temp.AntGain
    a.AntennaMode = temp.AntennaMode
    a.Bandwidth = temp.Bandwidth
    a.Channel = temp.Channel
    a.Channels = temp.Channels
    a.Disabled = temp.Disabled
    a.Power = temp.Power
    a.PowerMax = temp.PowerMax
    a.PowerMin = temp.PowerMin
    a.Preamble = temp.Preamble
    return nil
}

// tempApRadioBand5 is a temporary struct used for validating the fields of ApRadioBand5.
type tempApRadioBand5  struct {
    AllowRrmDisable *bool                     `json:"allow_rrm_disable,omitempty"`
    AntGain         Optional[int]             `json:"ant_gain"`
    AntennaMode     *RadioBandAntennaModeEnum `json:"antenna_mode,omitempty"`
    Bandwidth       *Dot11Bandwidth5Enum      `json:"bandwidth,omitempty"`
    Channel         Optional[int]             `json:"channel"`
    Channels        Optional[[]int]           `json:"channels"`
    Disabled        *bool                     `json:"disabled,omitempty"`
    Power           Optional[int]             `json:"power"`
    PowerMax        Optional[int]             `json:"power_max"`
    PowerMin        Optional[int]             `json:"power_min"`
    Preamble        *RadioBandPreambleEnum    `json:"preamble,omitempty"`
}

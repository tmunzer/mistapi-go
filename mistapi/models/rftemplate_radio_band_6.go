package models

import (
    "encoding/json"
)

// RftemplateRadioBand6 represents a RftemplateRadioBand6 struct.
// Radio Band AP settings
type RftemplateRadioBand6 struct {
    AllowRrmDisable      *bool                     `json:"allow_rrm_disable,omitempty"`
    AntGain              Optional[int]             `json:"ant_gain"`
    // enum: `1x1`, `2x2`, `3x3`, `4x4`, `default`
    AntennaMode          *RadioBandAntennaModeEnum `json:"antenna_mode,omitempty"`
    // channel width for the 6GHz band. enum: `20`, `40`, `80`, `160`
    Bandwidth            *Dot11Bandwidth6Enum      `json:"bandwidth,omitempty"`
    // For RFTemplates. List of channels, null or empty array means auto
    Channels             Optional[[]int]           `json:"channels"`
    // whether to disable the radio
    Disabled             *bool                     `json:"disabled,omitempty"`
    // TX power of the radio. For Devices, 0 means auto. -1 / -2 / -3 / …: treated as 0 / -1 / -2 / …
    Power                Optional[int]             `json:"power"`
    // when power=0, max tx power to use, HW-specific values will be used if not set
    PowerMax             Optional[int]             `json:"power_max"`
    // when power=0, min tx power to use, HW-specific values will be used if not set
    PowerMin             Optional[int]             `json:"power_min"`
    // enum: `auto`, `long`, `short`
    Preamble             *RadioBandPreambleEnum    `json:"preamble,omitempty"`
    // for 6GHz Only, standard-power operation, AFC (Automatic Frequency Coordination) will be performed and we'll fallback to Low Power Indoor if AFC failed
    StandardPower        *bool                     `json:"standard_power,omitempty"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RftemplateRadioBand6.
// It customizes the JSON marshaling process for RftemplateRadioBand6 objects.
func (r RftemplateRadioBand6) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RftemplateRadioBand6 object to a map representation for JSON marshaling.
func (r RftemplateRadioBand6) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.AllowRrmDisable != nil {
        structMap["allow_rrm_disable"] = r.AllowRrmDisable
    }
    if r.AntGain.IsValueSet() {
        if r.AntGain.Value() != nil {
            structMap["ant_gain"] = r.AntGain.Value()
        } else {
            structMap["ant_gain"] = nil
        }
    }
    if r.AntennaMode != nil {
        structMap["antenna_mode"] = r.AntennaMode
    }
    if r.Bandwidth != nil {
        structMap["bandwidth"] = r.Bandwidth
    }
    if r.Channels.IsValueSet() {
        if r.Channels.Value() != nil {
            structMap["channels"] = r.Channels.Value()
        } else {
            structMap["channels"] = nil
        }
    }
    if r.Disabled != nil {
        structMap["disabled"] = r.Disabled
    }
    if r.Power.IsValueSet() {
        if r.Power.Value() != nil {
            structMap["power"] = r.Power.Value()
        } else {
            structMap["power"] = nil
        }
    }
    if r.PowerMax.IsValueSet() {
        if r.PowerMax.Value() != nil {
            structMap["power_max"] = r.PowerMax.Value()
        } else {
            structMap["power_max"] = nil
        }
    }
    if r.PowerMin.IsValueSet() {
        if r.PowerMin.Value() != nil {
            structMap["power_min"] = r.PowerMin.Value()
        } else {
            structMap["power_min"] = nil
        }
    }
    if r.Preamble != nil {
        structMap["preamble"] = r.Preamble
    }
    if r.StandardPower != nil {
        structMap["standard_power"] = r.StandardPower
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RftemplateRadioBand6.
// It customizes the JSON unmarshaling process for RftemplateRadioBand6 objects.
func (r *RftemplateRadioBand6) UnmarshalJSON(input []byte) error {
    var temp rftemplateRadioBand6
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "allow_rrm_disable", "ant_gain", "antenna_mode", "bandwidth", "channels", "disabled", "power", "power_max", "power_min", "preamble", "standard_power")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.AllowRrmDisable = temp.AllowRrmDisable
    r.AntGain = temp.AntGain
    r.AntennaMode = temp.AntennaMode
    r.Bandwidth = temp.Bandwidth
    r.Channels = temp.Channels
    r.Disabled = temp.Disabled
    r.Power = temp.Power
    r.PowerMax = temp.PowerMax
    r.PowerMin = temp.PowerMin
    r.Preamble = temp.Preamble
    r.StandardPower = temp.StandardPower
    return nil
}

// rftemplateRadioBand6 is a temporary struct used for validating the fields of RftemplateRadioBand6.
type rftemplateRadioBand6  struct {
    AllowRrmDisable *bool                     `json:"allow_rrm_disable,omitempty"`
    AntGain         Optional[int]             `json:"ant_gain"`
    AntennaMode     *RadioBandAntennaModeEnum `json:"antenna_mode,omitempty"`
    Bandwidth       *Dot11Bandwidth6Enum      `json:"bandwidth,omitempty"`
    Channels        Optional[[]int]           `json:"channels"`
    Disabled        *bool                     `json:"disabled,omitempty"`
    Power           Optional[int]             `json:"power"`
    PowerMax        Optional[int]             `json:"power_max"`
    PowerMin        Optional[int]             `json:"power_min"`
    Preamble        *RadioBandPreambleEnum    `json:"preamble,omitempty"`
    StandardPower   *bool                     `json:"standard_power,omitempty"`
}

package models

import (
    "encoding/json"
)

// RrmBand represents a RrmBand struct.
type RrmBand struct {
    // channel width for the band * `80` is only applicable for band_5 and band_6 * `160` is only for band_6
    Bandwidth            *Dot11BandwidthEnum `json:"bandwidth,omitempty"`
    // proposed channel
    Channel              *int                `json:"channel,omitempty"`
    // channel width for the band * `80` is only applicable for band_5 and band_6 * `160` is only for band_6
    CurrBandwidht        *Dot11BandwidthEnum `json:"curr_bandwidht,omitempty"`
    // current channel
    CurrChannel          *int                `json:"curr_channel,omitempty"`
    // current tx power
    CurrPower            *int                `json:"curr_power,omitempty"`
    // current radio band
    CurrUsage            *string             `json:"curr_usage,omitempty"`
    // proposed tx power
    Power                *int                `json:"power,omitempty"`
    // proposed radio band
    Usage                *string             `json:"usage,omitempty"`
    AdditionalProperties map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RrmBand.
// It customizes the JSON marshaling process for RrmBand objects.
func (r RrmBand) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RrmBand object to a map representation for JSON marshaling.
func (r RrmBand) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Bandwidth != nil {
        structMap["bandwidth"] = r.Bandwidth
    }
    if r.Channel != nil {
        structMap["channel"] = r.Channel
    }
    if r.CurrBandwidht != nil {
        structMap["curr_bandwidht"] = r.CurrBandwidht
    }
    if r.CurrChannel != nil {
        structMap["curr_channel"] = r.CurrChannel
    }
    if r.CurrPower != nil {
        structMap["curr_power"] = r.CurrPower
    }
    if r.CurrUsage != nil {
        structMap["curr_usage"] = r.CurrUsage
    }
    if r.Power != nil {
        structMap["power"] = r.Power
    }
    if r.Usage != nil {
        structMap["usage"] = r.Usage
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RrmBand.
// It customizes the JSON unmarshaling process for RrmBand objects.
func (r *RrmBand) UnmarshalJSON(input []byte) error {
    var temp rrmBand
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "bandwidth", "channel", "curr_bandwidht", "curr_channel", "curr_power", "curr_usage", "power", "usage")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Bandwidth = temp.Bandwidth
    r.Channel = temp.Channel
    r.CurrBandwidht = temp.CurrBandwidht
    r.CurrChannel = temp.CurrChannel
    r.CurrPower = temp.CurrPower
    r.CurrUsage = temp.CurrUsage
    r.Power = temp.Power
    r.Usage = temp.Usage
    return nil
}

// rrmBand is a temporary struct used for validating the fields of RrmBand.
type rrmBand  struct {
    Bandwidth     *Dot11BandwidthEnum `json:"bandwidth,omitempty"`
    Channel       *int                `json:"channel,omitempty"`
    CurrBandwidht *Dot11BandwidthEnum `json:"curr_bandwidht,omitempty"`
    CurrChannel   *int                `json:"curr_channel,omitempty"`
    CurrPower     *int                `json:"curr_power,omitempty"`
    CurrUsage     *string             `json:"curr_usage,omitempty"`
    Power         *int                `json:"power,omitempty"`
    Usage         *string             `json:"usage,omitempty"`
}

package models

import (
    "encoding/json"
    "fmt"
)

// RrmBand represents a RrmBand struct.
type RrmBand struct {
    // channel width for the band.enum: `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6)
    Bandwidth            *Dot11BandwidthEnum    `json:"bandwidth,omitempty"`
    // proposed channel
    Channel              *int                   `json:"channel,omitempty"`
    // channel width for the band.enum: `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6)
    CurrBandwidht        *Dot11BandwidthEnum    `json:"curr_bandwidht,omitempty"`
    // current channel
    CurrChannel          *int                   `json:"curr_channel,omitempty"`
    // current tx power
    CurrPower            *int                   `json:"curr_power,omitempty"`
    // current radio band
    CurrUsage            *string                `json:"curr_usage,omitempty"`
    // proposed tx power
    Power                *int                   `json:"power,omitempty"`
    // proposed radio band
    Usage                *string                `json:"usage,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RrmBand,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RrmBand) String() string {
    return fmt.Sprintf(
    	"RrmBand[Bandwidth=%v, Channel=%v, CurrBandwidht=%v, CurrChannel=%v, CurrPower=%v, CurrUsage=%v, Power=%v, Usage=%v, AdditionalProperties=%v]",
    	r.Bandwidth, r.Channel, r.CurrBandwidht, r.CurrChannel, r.CurrPower, r.CurrUsage, r.Power, r.Usage, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RrmBand.
// It customizes the JSON marshaling process for RrmBand objects.
func (r RrmBand) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "bandwidth", "channel", "curr_bandwidht", "curr_channel", "curr_power", "curr_usage", "power", "usage"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RrmBand object to a map representation for JSON marshaling.
func (r RrmBand) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
    var temp tempRrmBand
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bandwidth", "channel", "curr_bandwidht", "curr_channel", "curr_power", "curr_usage", "power", "usage")
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

// tempRrmBand is a temporary struct used for validating the fields of RrmBand.
type tempRrmBand  struct {
    Bandwidth     *Dot11BandwidthEnum `json:"bandwidth,omitempty"`
    Channel       *int                `json:"channel,omitempty"`
    CurrBandwidht *Dot11BandwidthEnum `json:"curr_bandwidht,omitempty"`
    CurrChannel   *int                `json:"curr_channel,omitempty"`
    CurrPower     *int                `json:"curr_power,omitempty"`
    CurrUsage     *string             `json:"curr_usage,omitempty"`
    Power         *int                `json:"power,omitempty"`
    Usage         *string             `json:"usage,omitempty"`
}

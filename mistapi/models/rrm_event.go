package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// RrmEvent represents a RrmEvent struct.
type RrmEvent struct {
    ApId                 uuid.UUID                `json:"ap_id"`
    // enum: `24`, `5`, `6`
    Band                 Dot11BandEnum            `json:"band"`
    // channel width for the band.enum: `20`, `40`, `80` (only applicable for band_5 and band_6), `160` (only for band_6)
    Bandwidth            Dot11BandwidthEnum       `json:"bandwidth"`
    // channel for the band from rrm
    Channel              int                      `json:"channel"`
    // enum: `interference-ap-co-channel`, `interference-ap-non-wifi`, `neighbor-ap-down`, `neighbor-ap-recovered`, `radar-detected`, `rrm-radar`, `scheduled-site_rrm`, `triggered-site_rrm`
    Event                RrmEventTypeEnum         `json:"event"`
    // tx power of the radio
    Power                int                      `json:"power"`
    // (previously) channel width for the band , 0 means no previously available. enum: `0`, `20`, `40`, `80`, `160`
    PreBandwidth         RrmEventPreBandwidthEnum `json:"pre_bandwidth"`
    // (previously) channel for the band, 0 means no previously available
    PreChannel           int                      `json:"pre_channel"`
    // (previously) tx power of the radio, 0 means no previously available
    PrePower             float64                  `json:"pre_power"`
    PreUsage             string                   `json:"pre_usage"`
    // timestamp of the event
    Timestamp            float64                  `json:"timestamp"`
    Usage                string                   `json:"usage"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RrmEvent.
// It customizes the JSON marshaling process for RrmEvent objects.
func (r RrmEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RrmEvent object to a map representation for JSON marshaling.
func (r RrmEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["ap_id"] = r.ApId
    structMap["band"] = r.Band
    structMap["bandwidth"] = r.Bandwidth
    structMap["channel"] = r.Channel
    structMap["event"] = r.Event
    structMap["power"] = r.Power
    structMap["pre_bandwidth"] = r.PreBandwidth
    structMap["pre_channel"] = r.PreChannel
    structMap["pre_power"] = r.PrePower
    structMap["pre_usage"] = r.PreUsage
    structMap["timestamp"] = r.Timestamp
    structMap["usage"] = r.Usage
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RrmEvent.
// It customizes the JSON unmarshaling process for RrmEvent objects.
func (r *RrmEvent) UnmarshalJSON(input []byte) error {
    var temp tempRrmEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_id", "band", "bandwidth", "channel", "event", "power", "pre_bandwidth", "pre_channel", "pre_power", "pre_usage", "timestamp", "usage")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.ApId = *temp.ApId
    r.Band = *temp.Band
    r.Bandwidth = *temp.Bandwidth
    r.Channel = *temp.Channel
    r.Event = *temp.Event
    r.Power = *temp.Power
    r.PreBandwidth = *temp.PreBandwidth
    r.PreChannel = *temp.PreChannel
    r.PrePower = *temp.PrePower
    r.PreUsage = *temp.PreUsage
    r.Timestamp = *temp.Timestamp
    r.Usage = *temp.Usage
    return nil
}

// tempRrmEvent is a temporary struct used for validating the fields of RrmEvent.
type tempRrmEvent  struct {
    ApId         *uuid.UUID                `json:"ap_id"`
    Band         *Dot11BandEnum            `json:"band"`
    Bandwidth    *Dot11BandwidthEnum       `json:"bandwidth"`
    Channel      *int                      `json:"channel"`
    Event        *RrmEventTypeEnum         `json:"event"`
    Power        *int                      `json:"power"`
    PreBandwidth *RrmEventPreBandwidthEnum `json:"pre_bandwidth"`
    PreChannel   *int                      `json:"pre_channel"`
    PrePower     *float64                  `json:"pre_power"`
    PreUsage     *string                   `json:"pre_usage"`
    Timestamp    *float64                  `json:"timestamp"`
    Usage        *string                   `json:"usage"`
}

func (r *tempRrmEvent) validate() error {
    var errs []string
    if r.ApId == nil {
        errs = append(errs, "required field `ap_id` is missing for type `rrm_event`")
    }
    if r.Band == nil {
        errs = append(errs, "required field `band` is missing for type `rrm_event`")
    }
    if r.Bandwidth == nil {
        errs = append(errs, "required field `bandwidth` is missing for type `rrm_event`")
    }
    if r.Channel == nil {
        errs = append(errs, "required field `channel` is missing for type `rrm_event`")
    }
    if r.Event == nil {
        errs = append(errs, "required field `event` is missing for type `rrm_event`")
    }
    if r.Power == nil {
        errs = append(errs, "required field `power` is missing for type `rrm_event`")
    }
    if r.PreBandwidth == nil {
        errs = append(errs, "required field `pre_bandwidth` is missing for type `rrm_event`")
    }
    if r.PreChannel == nil {
        errs = append(errs, "required field `pre_channel` is missing for type `rrm_event`")
    }
    if r.PrePower == nil {
        errs = append(errs, "required field `pre_power` is missing for type `rrm_event`")
    }
    if r.PreUsage == nil {
        errs = append(errs, "required field `pre_usage` is missing for type `rrm_event`")
    }
    if r.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `rrm_event`")
    }
    if r.Usage == nil {
        errs = append(errs, "required field `usage` is missing for type `rrm_event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

package models

import (
    "encoding/json"
)

// UtilsSendBleBeacon represents a UtilsSendBleBeacon struct.
type UtilsSendBleBeacon struct {
    BeaconFrame          *string                `json:"beacon_frame,omitempty"`
    BeaconFreq           *int                   `json:"beacon_freq,omitempty"`
    Duration             *int                   `json:"duration,omitempty"`
    Macs                 []string               `json:"macs,omitempty"`
    MapIds               []string               `json:"map_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsSendBleBeacon.
// It customizes the JSON marshaling process for UtilsSendBleBeacon objects.
func (u UtilsSendBleBeacon) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(u.AdditionalProperties,
        "beacon_frame", "beacon_freq", "duration", "macs", "map_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsSendBleBeacon object to a map representation for JSON marshaling.
func (u UtilsSendBleBeacon) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, u.AdditionalProperties)
    if u.BeaconFrame != nil {
        structMap["beacon_frame"] = u.BeaconFrame
    }
    if u.BeaconFreq != nil {
        structMap["beacon_freq"] = u.BeaconFreq
    }
    if u.Duration != nil {
        structMap["duration"] = u.Duration
    }
    if u.Macs != nil {
        structMap["macs"] = u.Macs
    }
    if u.MapIds != nil {
        structMap["map_ids"] = u.MapIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsSendBleBeacon.
// It customizes the JSON unmarshaling process for UtilsSendBleBeacon objects.
func (u *UtilsSendBleBeacon) UnmarshalJSON(input []byte) error {
    var temp tempUtilsSendBleBeacon
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "beacon_frame", "beacon_freq", "duration", "macs", "map_ids")
    if err != nil {
    	return err
    }
    u.AdditionalProperties = additionalProperties
    
    u.BeaconFrame = temp.BeaconFrame
    u.BeaconFreq = temp.BeaconFreq
    u.Duration = temp.Duration
    u.Macs = temp.Macs
    u.MapIds = temp.MapIds
    return nil
}

// tempUtilsSendBleBeacon is a temporary struct used for validating the fields of UtilsSendBleBeacon.
type tempUtilsSendBleBeacon  struct {
    BeaconFrame *string  `json:"beacon_frame,omitempty"`
    BeaconFreq  *int     `json:"beacon_freq,omitempty"`
    Duration    *int     `json:"duration,omitempty"`
    Macs        []string `json:"macs,omitempty"`
    MapIds      []string `json:"map_ids,omitempty"`
}

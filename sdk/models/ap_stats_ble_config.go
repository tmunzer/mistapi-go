package models

import (
    "encoding/json"
)

// ApStatsBleConfig represents a ApStatsBleConfig struct.
type ApStatsBleConfig struct {
    BeaconRate           *int           `json:"beacon_rate,omitempty"`
    BeaconRateModel      *string        `json:"beacon_rate_model,omitempty"`
    BeamDisabled         []int          `json:"beam_disabled,omitempty"`
    Power                *int           `json:"power,omitempty"`
    PowerMode            *string        `json:"power_mode,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsBleConfig.
// It customizes the JSON marshaling process for ApStatsBleConfig objects.
func (a ApStatsBleConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsBleConfig object to a map representation for JSON marshaling.
func (a ApStatsBleConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.BeaconRate != nil {
        structMap["beacon_rate"] = a.BeaconRate
    }
    if a.BeaconRateModel != nil {
        structMap["beacon_rate_model"] = a.BeaconRateModel
    }
    if a.BeamDisabled != nil {
        structMap["beam_disabled"] = a.BeamDisabled
    }
    if a.Power != nil {
        structMap["power"] = a.Power
    }
    if a.PowerMode != nil {
        structMap["power_mode"] = a.PowerMode
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsBleConfig.
// It customizes the JSON unmarshaling process for ApStatsBleConfig objects.
func (a *ApStatsBleConfig) UnmarshalJSON(input []byte) error {
    var temp apStatsBleConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "beacon_rate", "beacon_rate_model", "beam_disabled", "power", "power_mode")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.BeaconRate = temp.BeaconRate
    a.BeaconRateModel = temp.BeaconRateModel
    a.BeamDisabled = temp.BeamDisabled
    a.Power = temp.Power
    a.PowerMode = temp.PowerMode
    return nil
}

// apStatsBleConfig is a temporary struct used for validating the fields of ApStatsBleConfig.
type apStatsBleConfig  struct {
    BeaconRate      *int    `json:"beacon_rate,omitempty"`
    BeaconRateModel *string `json:"beacon_rate_model,omitempty"`
    BeamDisabled    []int   `json:"beam_disabled,omitempty"`
    Power           *int    `json:"power,omitempty"`
    PowerMode       *string `json:"power_mode,omitempty"`
}

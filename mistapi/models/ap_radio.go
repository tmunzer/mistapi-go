package models

import (
    "encoding/json"
)

// ApRadio represents a ApRadio struct.
// Radio AP settings
type ApRadio struct {
    AllowRrmDisable      *bool                   `json:"allow_rrm_disable,omitempty"`
    // antenna gain for 2.4G - for models with external antenna only
    AntGain24            *int                    `json:"ant_gain_24,omitempty"`
    // antenna gain for 5G - for models with external antenna only
    AntGain5             *int                    `json:"ant_gain_5,omitempty"`
    // antenna gain for 6G - for models with external antenna only
    AntGain6             *int                    `json:"ant_gain_6,omitempty"`
    AntennaMode          *ApRadioAntennaModeEnum `json:"antenna_mode,omitempty"`
    // Radio Band AP settings
    Band24               *ApRadioBand24          `json:"band_24,omitempty"`
    Band24Usage          *RadioBand24UsageEnum   `json:"band_24_usage,omitempty"`
    // Radio Band AP settings
    Band5                *ApRadioBand5           `json:"band_5,omitempty"`
    // Radio Band AP settings
    Band5On24Radio       *ApRadioBand5           `json:"band_5_on_24_radio,omitempty"`
    // Radio Band AP settings
    Band6                *ApRadioBand6           `json:"band_6,omitempty"`
    // to make an outdoor operate indoor.
    // for an outdoor-ap, some channels are disallowed by default, this allows the user to use it as an indoor-ap
    IndoorUse            *bool                   `json:"indoor_use,omitempty"`
    // whether scanning radio is enabled
    ScanningEnabled      *bool                   `json:"scanning_enabled,omitempty"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApRadio.
// It customizes the JSON marshaling process for ApRadio objects.
func (a ApRadio) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApRadio object to a map representation for JSON marshaling.
func (a ApRadio) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.AllowRrmDisable != nil {
        structMap["allow_rrm_disable"] = a.AllowRrmDisable
    }
    if a.AntGain24 != nil {
        structMap["ant_gain_24"] = a.AntGain24
    }
    if a.AntGain5 != nil {
        structMap["ant_gain_5"] = a.AntGain5
    }
    if a.AntGain6 != nil {
        structMap["ant_gain_6"] = a.AntGain6
    }
    if a.AntennaMode != nil {
        structMap["antenna_mode"] = a.AntennaMode
    }
    if a.Band24 != nil {
        structMap["band_24"] = a.Band24.toMap()
    }
    if a.Band24Usage != nil {
        structMap["band_24_usage"] = a.Band24Usage
    }
    if a.Band5 != nil {
        structMap["band_5"] = a.Band5.toMap()
    }
    if a.Band5On24Radio != nil {
        structMap["band_5_on_24_radio"] = a.Band5On24Radio.toMap()
    }
    if a.Band6 != nil {
        structMap["band_6"] = a.Band6.toMap()
    }
    if a.IndoorUse != nil {
        structMap["indoor_use"] = a.IndoorUse
    }
    if a.ScanningEnabled != nil {
        structMap["scanning_enabled"] = a.ScanningEnabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApRadio.
// It customizes the JSON unmarshaling process for ApRadio objects.
func (a *ApRadio) UnmarshalJSON(input []byte) error {
    var temp apRadio
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "allow_rrm_disable", "ant_gain_24", "ant_gain_5", "ant_gain_6", "antenna_mode", "band_24", "band_24_usage", "band_5", "band_5_on_24_radio", "band_6", "indoor_use", "scanning_enabled")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.AllowRrmDisable = temp.AllowRrmDisable
    a.AntGain24 = temp.AntGain24
    a.AntGain5 = temp.AntGain5
    a.AntGain6 = temp.AntGain6
    a.AntennaMode = temp.AntennaMode
    a.Band24 = temp.Band24
    a.Band24Usage = temp.Band24Usage
    a.Band5 = temp.Band5
    a.Band5On24Radio = temp.Band5On24Radio
    a.Band6 = temp.Band6
    a.IndoorUse = temp.IndoorUse
    a.ScanningEnabled = temp.ScanningEnabled
    return nil
}

// apRadio is a temporary struct used for validating the fields of ApRadio.
type apRadio  struct {
    AllowRrmDisable *bool                   `json:"allow_rrm_disable,omitempty"`
    AntGain24       *int                    `json:"ant_gain_24,omitempty"`
    AntGain5        *int                    `json:"ant_gain_5,omitempty"`
    AntGain6        *int                    `json:"ant_gain_6,omitempty"`
    AntennaMode     *ApRadioAntennaModeEnum `json:"antenna_mode,omitempty"`
    Band24          *ApRadioBand24          `json:"band_24,omitempty"`
    Band24Usage     *RadioBand24UsageEnum   `json:"band_24_usage,omitempty"`
    Band5           *ApRadioBand5           `json:"band_5,omitempty"`
    Band5On24Radio  *ApRadioBand5           `json:"band_5_on_24_radio,omitempty"`
    Band6           *ApRadioBand6           `json:"band_6,omitempty"`
    IndoorUse       *bool                   `json:"indoor_use,omitempty"`
    ScanningEnabled *bool                   `json:"scanning_enabled,omitempty"`
}
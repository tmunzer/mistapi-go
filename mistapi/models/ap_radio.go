// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ApRadio represents a ApRadio struct.
// Radio AP settings
type ApRadio struct {
	AllowRrmDisable *bool `json:"allow_rrm_disable,omitempty"`
	// Antenna gain for 2.4G - for models with external antenna only
	AntGain24 *int `json:"ant_gain_24,omitempty"`
	// Antenna gain for 5G - for models with external antenna only
	AntGain5 *int `json:"ant_gain_5,omitempty"`
	// Antenna gain for 6G - for models with external antenna only
	AntGain6 *int `json:"ant_gain_6,omitempty"`
	// enum: `1x1`, `2x2`, `3x3`, `4x4`, `default`
	AntennaMode *ApRadioAntennaModeEnum `json:"antenna_mode,omitempty"`
	// Antenna Mode for AP which supports selectable antennas. enum: `""` (default), `external`, `internal`
	AntennaSelect *AntennaSelectEnum `json:"antenna_select,omitempty"`
	// Radio Band AP settings
	Band24 *ApRadioBand24 `json:"band_24,omitempty"`
	// enum: `24`, `5`, `6`, `auto`
	Band24Usage *RadioBand24UsageEnum `json:"band_24_usage,omitempty"`
	// Radio Band AP settings
	Band5 *ApRadioBand5 `json:"band_5,omitempty"`
	// Radio Band AP settings
	Band5On24Radio *ApRadioBand5 `json:"band_5_on_24_radio,omitempty"`
	// Radio Band AP settings
	Band6 *ApRadioBand6 `json:"band_6,omitempty"`
	// Let RRM control everything, only the `channels` and `ant_gain` will be honored (i.e. disabled/bandwidth/power/band_24_usage are all controlled by RRM)
	FullAutomaticRrm *bool `json:"full_automatic_rrm,omitempty"`
	// To make an outdoor operate indoor. For an outdoor-ap, some channels are disallowed by default, this allows the user to use it as an indoor-ap
	IndoorUse *bool `json:"indoor_use,omitempty"`
	// Enable RRM to manage all radio settings (ignores all band_xxx configs)
	RrmManaged *bool `json:"rrm_managed,omitempty"`
	// Whether scanning radio is enabled
	ScanningEnabled      *bool                  `json:"scanning_enabled,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApRadio,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApRadio) String() string {
	return fmt.Sprintf(
		"ApRadio[AllowRrmDisable=%v, AntGain24=%v, AntGain5=%v, AntGain6=%v, AntennaMode=%v, AntennaSelect=%v, Band24=%v, Band24Usage=%v, Band5=%v, Band5On24Radio=%v, Band6=%v, FullAutomaticRrm=%v, IndoorUse=%v, RrmManaged=%v, ScanningEnabled=%v, AdditionalProperties=%v]",
		a.AllowRrmDisable, a.AntGain24, a.AntGain5, a.AntGain6, a.AntennaMode, a.AntennaSelect, a.Band24, a.Band24Usage, a.Band5, a.Band5On24Radio, a.Band6, a.FullAutomaticRrm, a.IndoorUse, a.RrmManaged, a.ScanningEnabled, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApRadio.
// It customizes the JSON marshaling process for ApRadio objects.
func (a ApRadio) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"allow_rrm_disable", "ant_gain_24", "ant_gain_5", "ant_gain_6", "antenna_mode", "antenna_select", "band_24", "band_24_usage", "band_5", "band_5_on_24_radio", "band_6", "full_automatic_rrm", "indoor_use", "rrm_managed", "scanning_enabled"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the ApRadio object to a map representation for JSON marshaling.
func (a ApRadio) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
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
	if a.AntennaSelect != nil {
		structMap["antenna_select"] = a.AntennaSelect
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
	if a.FullAutomaticRrm != nil {
		structMap["full_automatic_rrm"] = a.FullAutomaticRrm
	}
	if a.IndoorUse != nil {
		structMap["indoor_use"] = a.IndoorUse
	}
	if a.RrmManaged != nil {
		structMap["rrm_managed"] = a.RrmManaged
	}
	if a.ScanningEnabled != nil {
		structMap["scanning_enabled"] = a.ScanningEnabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApRadio.
// It customizes the JSON unmarshaling process for ApRadio objects.
func (a *ApRadio) UnmarshalJSON(input []byte) error {
	var temp tempApRadio
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "allow_rrm_disable", "ant_gain_24", "ant_gain_5", "ant_gain_6", "antenna_mode", "antenna_select", "band_24", "band_24_usage", "band_5", "band_5_on_24_radio", "band_6", "full_automatic_rrm", "indoor_use", "rrm_managed", "scanning_enabled")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.AllowRrmDisable = temp.AllowRrmDisable
	a.AntGain24 = temp.AntGain24
	a.AntGain5 = temp.AntGain5
	a.AntGain6 = temp.AntGain6
	a.AntennaMode = temp.AntennaMode
	a.AntennaSelect = temp.AntennaSelect
	a.Band24 = temp.Band24
	a.Band24Usage = temp.Band24Usage
	a.Band5 = temp.Band5
	a.Band5On24Radio = temp.Band5On24Radio
	a.Band6 = temp.Band6
	a.FullAutomaticRrm = temp.FullAutomaticRrm
	a.IndoorUse = temp.IndoorUse
	a.RrmManaged = temp.RrmManaged
	a.ScanningEnabled = temp.ScanningEnabled
	return nil
}

// tempApRadio is a temporary struct used for validating the fields of ApRadio.
type tempApRadio struct {
	AllowRrmDisable  *bool                   `json:"allow_rrm_disable,omitempty"`
	AntGain24        *int                    `json:"ant_gain_24,omitempty"`
	AntGain5         *int                    `json:"ant_gain_5,omitempty"`
	AntGain6         *int                    `json:"ant_gain_6,omitempty"`
	AntennaMode      *ApRadioAntennaModeEnum `json:"antenna_mode,omitempty"`
	AntennaSelect    *AntennaSelectEnum      `json:"antenna_select,omitempty"`
	Band24           *ApRadioBand24          `json:"band_24,omitempty"`
	Band24Usage      *RadioBand24UsageEnum   `json:"band_24_usage,omitempty"`
	Band5            *ApRadioBand5           `json:"band_5,omitempty"`
	Band5On24Radio   *ApRadioBand5           `json:"band_5_on_24_radio,omitempty"`
	Band6            *ApRadioBand6           `json:"band_6,omitempty"`
	FullAutomaticRrm *bool                   `json:"full_automatic_rrm,omitempty"`
	IndoorUse        *bool                   `json:"indoor_use,omitempty"`
	RrmManaged       *bool                   `json:"rrm_managed,omitempty"`
	ScanningEnabled  *bool                   `json:"scanning_enabled,omitempty"`
}

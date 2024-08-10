package models

import (
    "encoding/json"
)

// RfTemplateModelSpecificProperty represents a RfTemplateModelSpecificProperty struct.
type RfTemplateModelSpecificProperty struct {
    AntGain24            *int                   `json:"ant_gain_24,omitempty"`
    AntGain5             *int                   `json:"ant_gain_5,omitempty"`
    AntGain6             *int                   `json:"ant_gain_6,omitempty"`
    // Radio Band AP settings
    Band24               *RftemplateRadioBand24 `json:"band_24,omitempty"`
    // enum: `24`, `5`, `6`, `auto`
    Band24Usage          *RadioBand24UsageEnum  `json:"band_24_usage,omitempty"`
    // Radio Band AP settings
    Band5                *RftemplateRadioBand5  `json:"band_5,omitempty"`
    // Radio Band AP settings
    Band5On24Radio       *RftemplateRadioBand5  `json:"band_5_on_24_radio,omitempty"`
    // Radio Band AP settings
    Band6                *RftemplateRadioBand6  `json:"band_6,omitempty"`
    AdditionalProperties map[string]any         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RfTemplateModelSpecificProperty.
// It customizes the JSON marshaling process for RfTemplateModelSpecificProperty objects.
func (r RfTemplateModelSpecificProperty) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RfTemplateModelSpecificProperty object to a map representation for JSON marshaling.
func (r RfTemplateModelSpecificProperty) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.AntGain24 != nil {
        structMap["ant_gain_24"] = r.AntGain24
    }
    if r.AntGain5 != nil {
        structMap["ant_gain_5"] = r.AntGain5
    }
    if r.AntGain6 != nil {
        structMap["ant_gain_6"] = r.AntGain6
    }
    if r.Band24 != nil {
        structMap["band_24"] = r.Band24.toMap()
    }
    if r.Band24Usage != nil {
        structMap["band_24_usage"] = r.Band24Usage
    }
    if r.Band5 != nil {
        structMap["band_5"] = r.Band5.toMap()
    }
    if r.Band5On24Radio != nil {
        structMap["band_5_on_24_radio"] = r.Band5On24Radio.toMap()
    }
    if r.Band6 != nil {
        structMap["band_6"] = r.Band6.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RfTemplateModelSpecificProperty.
// It customizes the JSON unmarshaling process for RfTemplateModelSpecificProperty objects.
func (r *RfTemplateModelSpecificProperty) UnmarshalJSON(input []byte) error {
    var temp tempRfTemplateModelSpecificProperty
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ant_gain_24", "ant_gain_5", "ant_gain_6", "band_24", "band_24_usage", "band_5", "band_5_on_24_radio", "band_6")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.AntGain24 = temp.AntGain24
    r.AntGain5 = temp.AntGain5
    r.AntGain6 = temp.AntGain6
    r.Band24 = temp.Band24
    r.Band24Usage = temp.Band24Usage
    r.Band5 = temp.Band5
    r.Band5On24Radio = temp.Band5On24Radio
    r.Band6 = temp.Band6
    return nil
}

// tempRfTemplateModelSpecificProperty is a temporary struct used for validating the fields of RfTemplateModelSpecificProperty.
type tempRfTemplateModelSpecificProperty  struct {
    AntGain24      *int                   `json:"ant_gain_24,omitempty"`
    AntGain5       *int                   `json:"ant_gain_5,omitempty"`
    AntGain6       *int                   `json:"ant_gain_6,omitempty"`
    Band24         *RftemplateRadioBand24 `json:"band_24,omitempty"`
    Band24Usage    *RadioBand24UsageEnum  `json:"band_24_usage,omitempty"`
    Band5          *RftemplateRadioBand5  `json:"band_5,omitempty"`
    Band5On24Radio *RftemplateRadioBand5  `json:"band_5_on_24_radio,omitempty"`
    Band6          *RftemplateRadioBand6  `json:"band_6,omitempty"`
}

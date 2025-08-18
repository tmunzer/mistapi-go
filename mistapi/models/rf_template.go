// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// RfTemplate represents a RfTemplate struct.
// RF Template
type RfTemplate struct {
	AntGain24 *int `json:"ant_gain_24,omitempty"`
	AntGain5  *int `json:"ant_gain_5,omitempty"`
	AntGain6  *int `json:"ant_gain_6,omitempty"`
	// Radio Band AP settings
	Band24 *RftemplateRadioBand24 `json:"band_24,omitempty"`
	// enum: `24`, `5`, `6`, `auto`
	Band24Usage *RadioBand24UsageEnum `json:"band_24_usage,omitempty"`
	// Radio Band AP settings
	Band5 *RftemplateRadioBand5 `json:"band_5,omitempty"`
	// Radio Band AP settings
	Band5On24Radio *RftemplateRadioBand5 `json:"band_5_on_24_radio,omitempty"`
	// Radio Band AP settings
	Band6 *RftemplateRadioBand6 `json:"band_6,omitempty"`
	// Optional, country code to use. If specified, this gets applied to all sites using the RF Template
	CountryCode *string `json:"country_code,omitempty"`
	// When the object has been created, in epoch
	CreatedTime *float64 `json:"created_time,omitempty"`
	ForSite     *bool    `json:"for_site,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id *uuid.UUID `json:"id,omitempty"`
	// overwrites for a specific model. If a band is specified, it will shadow the default. Property key is the model name (e.g. "AP63")
	ModelSpecific map[string]RfTemplateModelSpecificProperty `json:"model_specific,omitempty"`
	// When the object has been modified for the last time, in epoch
	ModifiedTime *float64 `json:"modified_time,omitempty"`
	// The name of the RF template
	Name  string     `json:"name"`
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// Whether scanning radio is enabled
	ScanningEnabled      *bool                  `json:"scanning_enabled,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RfTemplate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RfTemplate) String() string {
	return fmt.Sprintf(
		"RfTemplate[AntGain24=%v, AntGain5=%v, AntGain6=%v, Band24=%v, Band24Usage=%v, Band5=%v, Band5On24Radio=%v, Band6=%v, CountryCode=%v, CreatedTime=%v, ForSite=%v, Id=%v, ModelSpecific=%v, ModifiedTime=%v, Name=%v, OrgId=%v, ScanningEnabled=%v, AdditionalProperties=%v]",
		r.AntGain24, r.AntGain5, r.AntGain6, r.Band24, r.Band24Usage, r.Band5, r.Band5On24Radio, r.Band6, r.CountryCode, r.CreatedTime, r.ForSite, r.Id, r.ModelSpecific, r.ModifiedTime, r.Name, r.OrgId, r.ScanningEnabled, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RfTemplate.
// It customizes the JSON marshaling process for RfTemplate objects.
func (r RfTemplate) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"ant_gain_24", "ant_gain_5", "ant_gain_6", "band_24", "band_24_usage", "band_5", "band_5_on_24_radio", "band_6", "country_code", "created_time", "for_site", "id", "model_specific", "modified_time", "name", "org_id", "scanning_enabled"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the RfTemplate object to a map representation for JSON marshaling.
func (r RfTemplate) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
	if r.CountryCode != nil {
		structMap["country_code"] = r.CountryCode
	}
	if r.CreatedTime != nil {
		structMap["created_time"] = r.CreatedTime
	}
	if r.ForSite != nil {
		structMap["for_site"] = r.ForSite
	}
	if r.Id != nil {
		structMap["id"] = r.Id
	}
	if r.ModelSpecific != nil {
		structMap["model_specific"] = r.ModelSpecific
	}
	if r.ModifiedTime != nil {
		structMap["modified_time"] = r.ModifiedTime
	}
	structMap["name"] = r.Name
	if r.OrgId != nil {
		structMap["org_id"] = r.OrgId
	}
	if r.ScanningEnabled != nil {
		structMap["scanning_enabled"] = r.ScanningEnabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RfTemplate.
// It customizes the JSON unmarshaling process for RfTemplate objects.
func (r *RfTemplate) UnmarshalJSON(input []byte) error {
	var temp tempRfTemplate
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ant_gain_24", "ant_gain_5", "ant_gain_6", "band_24", "band_24_usage", "band_5", "band_5_on_24_radio", "band_6", "country_code", "created_time", "for_site", "id", "model_specific", "modified_time", "name", "org_id", "scanning_enabled")
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
	r.CountryCode = temp.CountryCode
	r.CreatedTime = temp.CreatedTime
	r.ForSite = temp.ForSite
	r.Id = temp.Id
	r.ModelSpecific = temp.ModelSpecific
	r.ModifiedTime = temp.ModifiedTime
	r.Name = *temp.Name
	r.OrgId = temp.OrgId
	r.ScanningEnabled = temp.ScanningEnabled
	return nil
}

// tempRfTemplate is a temporary struct used for validating the fields of RfTemplate.
type tempRfTemplate struct {
	AntGain24       *int                                       `json:"ant_gain_24,omitempty"`
	AntGain5        *int                                       `json:"ant_gain_5,omitempty"`
	AntGain6        *int                                       `json:"ant_gain_6,omitempty"`
	Band24          *RftemplateRadioBand24                     `json:"band_24,omitempty"`
	Band24Usage     *RadioBand24UsageEnum                      `json:"band_24_usage,omitempty"`
	Band5           *RftemplateRadioBand5                      `json:"band_5,omitempty"`
	Band5On24Radio  *RftemplateRadioBand5                      `json:"band_5_on_24_radio,omitempty"`
	Band6           *RftemplateRadioBand6                      `json:"band_6,omitempty"`
	CountryCode     *string                                    `json:"country_code,omitempty"`
	CreatedTime     *float64                                   `json:"created_time,omitempty"`
	ForSite         *bool                                      `json:"for_site,omitempty"`
	Id              *uuid.UUID                                 `json:"id,omitempty"`
	ModelSpecific   map[string]RfTemplateModelSpecificProperty `json:"model_specific,omitempty"`
	ModifiedTime    *float64                                   `json:"modified_time,omitempty"`
	Name            *string                                    `json:"name"`
	OrgId           *uuid.UUID                                 `json:"org_id,omitempty"`
	ScanningEnabled *bool                                      `json:"scanning_enabled,omitempty"`
}

func (r *tempRfTemplate) validate() error {
	var errs []string
	if r.Name == nil {
		errs = append(errs, "required field `name` is missing for type `rf_template`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SiteSettingMxedge represents a SiteSettingMxedge struct.
// Site Mist Edges form a cluster of RadSec Proxy servers
type SiteSettingMxedge struct {
	// Configure cloud-assisted dynamic authorization service on this cluster of mist edges
	MistDas     *MxedgeDas    `json:"mist_das,omitempty"`
	MistNac     *MxclusterNac `json:"mist_nac,omitempty"`
	MistNacedge *MistNacedge  `json:"mist_nacedge,omitempty"`
	// MxEdge RadSec Configuration
	Radsec               *MxclusterRadsec       `json:"radsec,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingMxedge,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingMxedge) String() string {
	return fmt.Sprintf(
		"SiteSettingMxedge[MistDas=%v, MistNac=%v, MistNacedge=%v, Radsec=%v, AdditionalProperties=%v]",
		s.MistDas, s.MistNac, s.MistNacedge, s.Radsec, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingMxedge.
// It customizes the JSON marshaling process for SiteSettingMxedge objects.
func (s SiteSettingMxedge) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"mist_das", "mist_nac", "mist_nacedge", "radsec"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingMxedge object to a map representation for JSON marshaling.
func (s SiteSettingMxedge) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.MistDas != nil {
		structMap["mist_das"] = s.MistDas.toMap()
	}
	if s.MistNac != nil {
		structMap["mist_nac"] = s.MistNac.toMap()
	}
	if s.MistNacedge != nil {
		structMap["mist_nacedge"] = s.MistNacedge.toMap()
	}
	if s.Radsec != nil {
		structMap["radsec"] = s.Radsec.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingMxedge.
// It customizes the JSON unmarshaling process for SiteSettingMxedge objects.
func (s *SiteSettingMxedge) UnmarshalJSON(input []byte) error {
	var temp tempSiteSettingMxedge
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mist_das", "mist_nac", "mist_nacedge", "radsec")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.MistDas = temp.MistDas
	s.MistNac = temp.MistNac
	s.MistNacedge = temp.MistNacedge
	s.Radsec = temp.Radsec
	return nil
}

// tempSiteSettingMxedge is a temporary struct used for validating the fields of SiteSettingMxedge.
type tempSiteSettingMxedge struct {
	MistDas     *MxedgeDas       `json:"mist_das,omitempty"`
	MistNac     *MxclusterNac    `json:"mist_nac,omitempty"`
	MistNacedge *MistNacedge     `json:"mist_nacedge,omitempty"`
	Radsec      *MxclusterRadsec `json:"radsec,omitempty"`
}

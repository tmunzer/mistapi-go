// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SiteEngagementDwellTagNames represents a SiteEngagementDwellTagNames struct.
// Name associated to each tag
type SiteEngagementDwellTagNames struct {
	Bounce               *string                `json:"bounce,omitempty"`
	Engaged              *string                `json:"engaged,omitempty"`
	Passerby             *string                `json:"passerby,omitempty"`
	Stationed            *string                `json:"stationed,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SiteEngagementDwellTagNames,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteEngagementDwellTagNames) String() string {
	return fmt.Sprintf(
		"SiteEngagementDwellTagNames[Bounce=%v, Engaged=%v, Passerby=%v, Stationed=%v, AdditionalProperties=%v]",
		s.Bounce, s.Engaged, s.Passerby, s.Stationed, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteEngagementDwellTagNames.
// It customizes the JSON marshaling process for SiteEngagementDwellTagNames objects.
func (s SiteEngagementDwellTagNames) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"bounce", "engaged", "passerby", "stationed"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SiteEngagementDwellTagNames object to a map representation for JSON marshaling.
func (s SiteEngagementDwellTagNames) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.Bounce != nil {
		structMap["bounce"] = s.Bounce
	}
	if s.Engaged != nil {
		structMap["engaged"] = s.Engaged
	}
	if s.Passerby != nil {
		structMap["passerby"] = s.Passerby
	}
	if s.Stationed != nil {
		structMap["stationed"] = s.Stationed
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteEngagementDwellTagNames.
// It customizes the JSON unmarshaling process for SiteEngagementDwellTagNames objects.
func (s *SiteEngagementDwellTagNames) UnmarshalJSON(input []byte) error {
	var temp tempSiteEngagementDwellTagNames
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bounce", "engaged", "passerby", "stationed")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.Bounce = temp.Bounce
	s.Engaged = temp.Engaged
	s.Passerby = temp.Passerby
	s.Stationed = temp.Stationed
	return nil
}

// tempSiteEngagementDwellTagNames is a temporary struct used for validating the fields of SiteEngagementDwellTagNames.
type tempSiteEngagementDwellTagNames struct {
	Bounce    *string `json:"bounce,omitempty"`
	Engaged   *string `json:"engaged,omitempty"`
	Passerby  *string `json:"passerby,omitempty"`
	Stationed *string `json:"stationed,omitempty"`
}

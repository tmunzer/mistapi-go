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

// Sdktemplate represents a Sdktemplate struct.
// SDK Template
type Sdktemplate struct {
	BgImage       *string `json:"bg_image,omitempty"`
	BtnFlrBgcolor *string `json:"btn_flr_bgcolor,omitempty"`
	// When the object has been created, in epoch
	CreatedTime *float64 `json:"created_time,omitempty"`
	// Whether this is the default template when there are multiple templates
	Default   *bool   `json:"default,omitempty"`
	ForSite   *bool   `json:"for_site,omitempty"`
	HeaderTxt *string `json:"header_txt,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id *uuid.UUID `json:"id,omitempty"`
	// When the object has been modified for the last time, in epoch
	ModifiedTime *float64 `json:"modified_time,omitempty"`
	// Name for identification purpose
	Name                 string                 `json:"name"`
	OrgId                *uuid.UUID             `json:"org_id,omitempty"`
	SearchTxtcolor       *string                `json:"search_txtcolor,omitempty"`
	SiteId               *uuid.UUID             `json:"site_id,omitempty"`
	WelcomeMsg           *string                `json:"welcome_msg,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Sdktemplate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s Sdktemplate) String() string {
	return fmt.Sprintf(
		"Sdktemplate[BgImage=%v, BtnFlrBgcolor=%v, CreatedTime=%v, Default=%v, ForSite=%v, HeaderTxt=%v, Id=%v, ModifiedTime=%v, Name=%v, OrgId=%v, SearchTxtcolor=%v, SiteId=%v, WelcomeMsg=%v, AdditionalProperties=%v]",
		s.BgImage, s.BtnFlrBgcolor, s.CreatedTime, s.Default, s.ForSite, s.HeaderTxt, s.Id, s.ModifiedTime, s.Name, s.OrgId, s.SearchTxtcolor, s.SiteId, s.WelcomeMsg, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Sdktemplate.
// It customizes the JSON marshaling process for Sdktemplate objects.
func (s Sdktemplate) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"bg_image", "btn_flr_bgcolor", "created_time", "default", "for_site", "header_txt", "id", "modified_time", "name", "org_id", "search_txtcolor", "site_id", "welcome_msg"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the Sdktemplate object to a map representation for JSON marshaling.
func (s Sdktemplate) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.BgImage != nil {
		structMap["bg_image"] = s.BgImage
	}
	if s.BtnFlrBgcolor != nil {
		structMap["btn_flr_bgcolor"] = s.BtnFlrBgcolor
	}
	if s.CreatedTime != nil {
		structMap["created_time"] = s.CreatedTime
	}
	if s.Default != nil {
		structMap["default"] = s.Default
	}
	if s.ForSite != nil {
		structMap["for_site"] = s.ForSite
	}
	if s.HeaderTxt != nil {
		structMap["header_txt"] = s.HeaderTxt
	}
	if s.Id != nil {
		structMap["id"] = s.Id
	}
	if s.ModifiedTime != nil {
		structMap["modified_time"] = s.ModifiedTime
	}
	structMap["name"] = s.Name
	if s.OrgId != nil {
		structMap["org_id"] = s.OrgId
	}
	if s.SearchTxtcolor != nil {
		structMap["search_txtcolor"] = s.SearchTxtcolor
	}
	if s.SiteId != nil {
		structMap["site_id"] = s.SiteId
	}
	if s.WelcomeMsg != nil {
		structMap["welcome_msg"] = s.WelcomeMsg
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Sdktemplate.
// It customizes the JSON unmarshaling process for Sdktemplate objects.
func (s *Sdktemplate) UnmarshalJSON(input []byte) error {
	var temp tempSdktemplate
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "bg_image", "btn_flr_bgcolor", "created_time", "default", "for_site", "header_txt", "id", "modified_time", "name", "org_id", "search_txtcolor", "site_id", "welcome_msg")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.BgImage = temp.BgImage
	s.BtnFlrBgcolor = temp.BtnFlrBgcolor
	s.CreatedTime = temp.CreatedTime
	s.Default = temp.Default
	s.ForSite = temp.ForSite
	s.HeaderTxt = temp.HeaderTxt
	s.Id = temp.Id
	s.ModifiedTime = temp.ModifiedTime
	s.Name = *temp.Name
	s.OrgId = temp.OrgId
	s.SearchTxtcolor = temp.SearchTxtcolor
	s.SiteId = temp.SiteId
	s.WelcomeMsg = temp.WelcomeMsg
	return nil
}

// tempSdktemplate is a temporary struct used for validating the fields of Sdktemplate.
type tempSdktemplate struct {
	BgImage        *string    `json:"bg_image,omitempty"`
	BtnFlrBgcolor  *string    `json:"btn_flr_bgcolor,omitempty"`
	CreatedTime    *float64   `json:"created_time,omitempty"`
	Default        *bool      `json:"default,omitempty"`
	ForSite        *bool      `json:"for_site,omitempty"`
	HeaderTxt      *string    `json:"header_txt,omitempty"`
	Id             *uuid.UUID `json:"id,omitempty"`
	ModifiedTime   *float64   `json:"modified_time,omitempty"`
	Name           *string    `json:"name"`
	OrgId          *uuid.UUID `json:"org_id,omitempty"`
	SearchTxtcolor *string    `json:"search_txtcolor,omitempty"`
	SiteId         *uuid.UUID `json:"site_id,omitempty"`
	WelcomeMsg     *string    `json:"welcome_msg,omitempty"`
}

func (s *tempSdktemplate) validate() error {
	var errs []string
	if s.Name == nil {
		errs = append(errs, "required field `name` is missing for type `sdktemplate`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

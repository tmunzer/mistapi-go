// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// TemplateExceptions represents a TemplateExceptions struct.
// Where this template should not be applied to (takes precedence)
type TemplateExceptions struct {
	// List of site ids
	SiteIds []uuid.UUID `json:"site_ids,omitempty"`
	// List of sitegroup ids
	SitegroupIds         []uuid.UUID            `json:"sitegroup_ids,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for TemplateExceptions,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (t TemplateExceptions) String() string {
	return fmt.Sprintf(
		"TemplateExceptions[SiteIds=%v, SitegroupIds=%v, AdditionalProperties=%v]",
		t.SiteIds, t.SitegroupIds, t.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for TemplateExceptions.
// It customizes the JSON marshaling process for TemplateExceptions objects.
func (t TemplateExceptions) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(t.AdditionalProperties,
		"site_ids", "sitegroup_ids"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(t.toMap())
}

// toMap converts the TemplateExceptions object to a map representation for JSON marshaling.
func (t TemplateExceptions) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, t.AdditionalProperties)
	if t.SiteIds != nil {
		structMap["site_ids"] = t.SiteIds
	}
	if t.SitegroupIds != nil {
		structMap["sitegroup_ids"] = t.SitegroupIds
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TemplateExceptions.
// It customizes the JSON unmarshaling process for TemplateExceptions objects.
func (t *TemplateExceptions) UnmarshalJSON(input []byte) error {
	var temp tempTemplateExceptions
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "site_ids", "sitegroup_ids")
	if err != nil {
		return err
	}
	t.AdditionalProperties = additionalProperties

	t.SiteIds = temp.SiteIds
	t.SitegroupIds = temp.SitegroupIds
	return nil
}

// tempTemplateExceptions is a temporary struct used for validating the fields of TemplateExceptions.
type tempTemplateExceptions struct {
	SiteIds      []uuid.UUID `json:"site_ids,omitempty"`
	SitegroupIds []uuid.UUID `json:"sitegroup_ids,omitempty"`
}

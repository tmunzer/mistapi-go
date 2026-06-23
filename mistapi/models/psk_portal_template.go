// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// PskPortalTemplate represents a PskPortalTemplate struct.
// Portal UI customization payload
type PskPortalTemplate struct {
	// Custom UI settings for the PSK portal template
	PortalTemplate       *PskPortalTemplateSetting `json:"portal_template,omitempty"`
	AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for PskPortalTemplate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PskPortalTemplate) String() string {
	return fmt.Sprintf(
		"PskPortalTemplate[PortalTemplate=%v, AdditionalProperties=%v]",
		p.PortalTemplate, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PskPortalTemplate.
// It customizes the JSON marshaling process for PskPortalTemplate objects.
func (p PskPortalTemplate) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(p.AdditionalProperties,
		"portal_template"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(p.toMap())
}

// toMap converts the PskPortalTemplate object to a map representation for JSON marshaling.
func (p PskPortalTemplate) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, p.AdditionalProperties)
	if p.PortalTemplate != nil {
		structMap["portal_template"] = p.PortalTemplate.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PskPortalTemplate.
// It customizes the JSON unmarshaling process for PskPortalTemplate objects.
func (p *PskPortalTemplate) UnmarshalJSON(input []byte) error {
	var temp tempPskPortalTemplate
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "portal_template")
	if err != nil {
		return err
	}
	p.AdditionalProperties = additionalProperties

	p.PortalTemplate = temp.PortalTemplate
	return nil
}

// tempPskPortalTemplate is a temporary struct used for validating the fields of PskPortalTemplate.
type tempPskPortalTemplate struct {
	PortalTemplate *PskPortalTemplateSetting `json:"portal_template,omitempty"`
}

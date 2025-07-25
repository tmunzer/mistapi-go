// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// WlanPortalTemplate represents a WlanPortalTemplate struct.
type WlanPortalTemplate struct {
    // Portal template wlan settings
    PortalTemplate       *WlanPortalTemplateSetting `json:"portal_template,omitempty"`
    AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for WlanPortalTemplate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanPortalTemplate) String() string {
    return fmt.Sprintf(
    	"WlanPortalTemplate[PortalTemplate=%v, AdditionalProperties=%v]",
    	w.PortalTemplate, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanPortalTemplate.
// It customizes the JSON marshaling process for WlanPortalTemplate objects.
func (w WlanPortalTemplate) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "portal_template"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanPortalTemplate object to a map representation for JSON marshaling.
func (w WlanPortalTemplate) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.PortalTemplate != nil {
        structMap["portal_template"] = w.PortalTemplate.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanPortalTemplate.
// It customizes the JSON unmarshaling process for WlanPortalTemplate objects.
func (w *WlanPortalTemplate) UnmarshalJSON(input []byte) error {
    var temp tempWlanPortalTemplate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "portal_template")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.PortalTemplate = temp.PortalTemplate
    return nil
}

// tempWlanPortalTemplate is a temporary struct used for validating the fields of WlanPortalTemplate.
type tempWlanPortalTemplate  struct {
    PortalTemplate *WlanPortalTemplateSetting `json:"portal_template,omitempty"`
}

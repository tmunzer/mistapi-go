// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// PskPortalTemplate represents a PskPortalTemplate struct.
type PskPortalTemplate struct {
    // defines alignment on portal. enum: `center`, `left`, `right`
    Alignment            *PortalTemplateAlignmentEnum `json:"alignment,omitempty"`
    Color                *string                      `json:"color,omitempty"`
    // Custom logo with "data:image/png;base64," format.  default null, uses Juniper Mist Logo
    Logo                 Optional[string]             `json:"logo"`
    // Whether to hide "Powered by Juniper Mist" and email footers
    PoweredBy            *bool                        `json:"poweredBy,omitempty"`
    // Whether to show Terms of Service
    Tos                  *bool                        `json:"tos,omitempty"`
    // Terms of Service accept button label
    TosAcceptLabel       *string                      `json:"tosAcceptLabel,omitempty"`
    // Terror message for not accepting tos
    TosError             *string                      `json:"tosError,omitempty"`
    TosLink              *string                      `json:"tosLink,omitempty"`
    // terms and service text displayed in footer if tos is enabled
    TosText              *string                      `json:"tosText,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for PskPortalTemplate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PskPortalTemplate) String() string {
    return fmt.Sprintf(
    	"PskPortalTemplate[Alignment=%v, Color=%v, Logo=%v, PoweredBy=%v, Tos=%v, TosAcceptLabel=%v, TosError=%v, TosLink=%v, TosText=%v, AdditionalProperties=%v]",
    	p.Alignment, p.Color, p.Logo, p.PoweredBy, p.Tos, p.TosAcceptLabel, p.TosError, p.TosLink, p.TosText, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PskPortalTemplate.
// It customizes the JSON marshaling process for PskPortalTemplate objects.
func (p PskPortalTemplate) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "alignment", "color", "logo", "poweredBy", "tos", "tosAcceptLabel", "tosError", "tosLink", "tosText"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PskPortalTemplate object to a map representation for JSON marshaling.
func (p PskPortalTemplate) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.Alignment != nil {
        structMap["alignment"] = p.Alignment
    }
    if p.Color != nil {
        structMap["color"] = p.Color
    }
    if p.Logo.IsValueSet() {
        if p.Logo.Value() != nil {
            structMap["logo"] = p.Logo.Value()
        } else {
            structMap["logo"] = nil
        }
    }
    if p.PoweredBy != nil {
        structMap["poweredBy"] = p.PoweredBy
    }
    if p.Tos != nil {
        structMap["tos"] = p.Tos
    }
    if p.TosAcceptLabel != nil {
        structMap["tosAcceptLabel"] = p.TosAcceptLabel
    }
    if p.TosError != nil {
        structMap["tosError"] = p.TosError
    }
    if p.TosLink != nil {
        structMap["tosLink"] = p.TosLink
    }
    if p.TosText != nil {
        structMap["tosText"] = p.TosText
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "alignment", "color", "logo", "poweredBy", "tos", "tosAcceptLabel", "tosError", "tosLink", "tosText")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.Alignment = temp.Alignment
    p.Color = temp.Color
    p.Logo = temp.Logo
    p.PoweredBy = temp.PoweredBy
    p.Tos = temp.Tos
    p.TosAcceptLabel = temp.TosAcceptLabel
    p.TosError = temp.TosError
    p.TosLink = temp.TosLink
    p.TosText = temp.TosText
    return nil
}

// tempPskPortalTemplate is a temporary struct used for validating the fields of PskPortalTemplate.
type tempPskPortalTemplate  struct {
    Alignment      *PortalTemplateAlignmentEnum `json:"alignment,omitempty"`
    Color          *string                      `json:"color,omitempty"`
    Logo           Optional[string]             `json:"logo"`
    PoweredBy      *bool                        `json:"poweredBy,omitempty"`
    Tos            *bool                        `json:"tos,omitempty"`
    TosAcceptLabel *string                      `json:"tosAcceptLabel,omitempty"`
    TosError       *string                      `json:"tosError,omitempty"`
    TosLink        *string                      `json:"tosLink,omitempty"`
    TosText        *string                      `json:"tosText,omitempty"`
}

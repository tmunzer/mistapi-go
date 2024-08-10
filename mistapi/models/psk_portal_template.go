package models

import (
    "encoding/json"
)

// PskPortalTemplate represents a PskPortalTemplate struct.
type PskPortalTemplate struct {
    // defines alignment on portal. enum: `center`, `left`, `right`
    Alignment            *PortalTemplateAlignmentEnum `json:"alignment,omitempty"`
    Color                *string                      `json:"color,omitempty"`
    // custom logo with “data:image/png;base64,” format.  default null, uses Juniper Mist Logo
    Logo                 Optional[string]             `json:"logo"`
    // whether to hide "Powered by Juniper Mist" and email footers
    PoweredBy            *bool                        `json:"poweredBy,omitempty"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for PskPortalTemplate.
// It customizes the JSON marshaling process for PskPortalTemplate objects.
func (p PskPortalTemplate) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(p.toMap())
}

// toMap converts the PskPortalTemplate object to a map representation for JSON marshaling.
func (p PskPortalTemplate) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, p.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "alignment", "color", "logo", "poweredBy")
    if err != nil {
    	return err
    }
    
    p.AdditionalProperties = additionalProperties
    p.Alignment = temp.Alignment
    p.Color = temp.Color
    p.Logo = temp.Logo
    p.PoweredBy = temp.PoweredBy
    return nil
}

// tempPskPortalTemplate is a temporary struct used for validating the fields of PskPortalTemplate.
type tempPskPortalTemplate  struct {
    Alignment *PortalTemplateAlignmentEnum `json:"alignment,omitempty"`
    Color     *string                      `json:"color,omitempty"`
    Logo      Optional[string]             `json:"logo"`
    PoweredBy *bool                        `json:"poweredBy,omitempty"`
}

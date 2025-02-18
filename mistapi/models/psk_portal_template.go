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
    // Custom logo with “data:image/png;base64,” format.  default null, uses Juniper Mist Logo
    Logo                 Optional[string]             `json:"logo"`
    // Whether to hide "Powered by Juniper Mist" and email footers
    PoweredBy            *bool                        `json:"poweredBy,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for PskPortalTemplate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PskPortalTemplate) String() string {
    return fmt.Sprintf(
    	"PskPortalTemplate[Alignment=%v, Color=%v, Logo=%v, PoweredBy=%v, AdditionalProperties=%v]",
    	p.Alignment, p.Color, p.Logo, p.PoweredBy, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PskPortalTemplate.
// It customizes the JSON marshaling process for PskPortalTemplate objects.
func (p PskPortalTemplate) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "alignment", "color", "logo", "poweredBy"); err != nil {
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "alignment", "color", "logo", "poweredBy")
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

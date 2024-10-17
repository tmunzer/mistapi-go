package models

import (
	"encoding/json"
)

// NacPortalTemplate represents a NacPortalTemplate struct.
type NacPortalTemplate struct {
	// defines alignment on portal. enum: `center`, `left`, `right`
	Alignment *PortalTemplateAlignmentEnum `json:"alignment,omitempty"`
	Color     *string                      `json:"color,omitempty"`
	// custom logo custom logo with “data:image/png;base64,” format. default null, uses Juniper Mist Logo.
	Logo *string `json:"logo,omitempty"`
	// whether to hide “Powered by Juniper Mist” and email footers
	PoweredBy            *bool          `json:"poweredBy,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NacPortalTemplate.
// It customizes the JSON marshaling process for NacPortalTemplate objects.
func (n NacPortalTemplate) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(n.toMap())
}

// toMap converts the NacPortalTemplate object to a map representation for JSON marshaling.
func (n NacPortalTemplate) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, n.AdditionalProperties)
	if n.Alignment != nil {
		structMap["alignment"] = n.Alignment
	}
	if n.Color != nil {
		structMap["color"] = n.Color
	}
	if n.Logo != nil {
		structMap["logo"] = n.Logo
	}
	if n.PoweredBy != nil {
		structMap["poweredBy"] = n.PoweredBy
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NacPortalTemplate.
// It customizes the JSON unmarshaling process for NacPortalTemplate objects.
func (n *NacPortalTemplate) UnmarshalJSON(input []byte) error {
	var temp tempNacPortalTemplate
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "alignment", "color", "logo", "poweredBy")
	if err != nil {
		return err
	}

	n.AdditionalProperties = additionalProperties
	n.Alignment = temp.Alignment
	n.Color = temp.Color
	n.Logo = temp.Logo
	n.PoweredBy = temp.PoweredBy
	return nil
}

// tempNacPortalTemplate is a temporary struct used for validating the fields of NacPortalTemplate.
type tempNacPortalTemplate struct {
	Alignment *PortalTemplateAlignmentEnum `json:"alignment,omitempty"`
	Color     *string                      `json:"color,omitempty"`
	Logo      *string                      `json:"logo,omitempty"`
	PoweredBy *bool                        `json:"poweredBy,omitempty"`
}

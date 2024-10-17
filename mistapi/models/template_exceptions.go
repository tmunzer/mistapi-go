package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// TemplateExceptions represents a TemplateExceptions struct.
// where this template should not be applied to (takes precedence)
type TemplateExceptions struct {
	// list of site ids
	SiteIds []uuid.UUID `json:"site_ids,omitempty"`
	// list of sitegroup ids
	SitegroupIds         []uuid.UUID    `json:"sitegroup_ids,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TemplateExceptions.
// It customizes the JSON marshaling process for TemplateExceptions objects.
func (t TemplateExceptions) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(t.toMap())
}

// toMap converts the TemplateExceptions object to a map representation for JSON marshaling.
func (t TemplateExceptions) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, t.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "site_ids", "sitegroup_ids")
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

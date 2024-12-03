package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// TemplateApplies represents a TemplateApplies struct.
// where this template should be applied to, can be org_id, site_ids, sitegroup_ids
type TemplateApplies struct {
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    // list of site ids
    SiteIds              []uuid.UUID            `json:"site_ids,omitempty"`
    // list of sitegroup ids
    SitegroupIds         []uuid.UUID            `json:"sitegroup_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for TemplateApplies.
// It customizes the JSON marshaling process for TemplateApplies objects.
func (t TemplateApplies) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(t.AdditionalProperties,
        "org_id", "site_ids", "sitegroup_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(t.toMap())
}

// toMap converts the TemplateApplies object to a map representation for JSON marshaling.
func (t TemplateApplies) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, t.AdditionalProperties)
    if t.OrgId != nil {
        structMap["org_id"] = t.OrgId
    }
    if t.SiteIds != nil {
        structMap["site_ids"] = t.SiteIds
    }
    if t.SitegroupIds != nil {
        structMap["sitegroup_ids"] = t.SitegroupIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for TemplateApplies.
// It customizes the JSON unmarshaling process for TemplateApplies objects.
func (t *TemplateApplies) UnmarshalJSON(input []byte) error {
    var temp tempTemplateApplies
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "org_id", "site_ids", "sitegroup_ids")
    if err != nil {
    	return err
    }
    t.AdditionalProperties = additionalProperties
    
    t.OrgId = temp.OrgId
    t.SiteIds = temp.SiteIds
    t.SitegroupIds = temp.SitegroupIds
    return nil
}

// tempTemplateApplies is a temporary struct used for validating the fields of TemplateApplies.
type tempTemplateApplies  struct {
    OrgId        *uuid.UUID  `json:"org_id,omitempty"`
    SiteIds      []uuid.UUID `json:"site_ids,omitempty"`
    SitegroupIds []uuid.UUID `json:"sitegroup_ids,omitempty"`
}

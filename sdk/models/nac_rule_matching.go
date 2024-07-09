package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// NacRuleMatching represents a NacRuleMatching struct.
type NacRuleMatching struct {
    AuthType             *NacRuleMatchingAuthTypeEnum  `json:"auth_type,omitempty"`
    Nactags              []string                      `json:"nactags,omitempty"`
    PortTypes            []NacRuleMatchingPortTypeEnum `json:"port_types,omitempty"`
    // list of site ids to match
    SiteIds              []uuid.UUID                   `json:"site_ids,omitempty"`
    // list of sitegroup ids to match
    SitegroupIds         []uuid.UUID                   `json:"sitegroup_ids,omitempty"`
    // list of vendors to match
    Vendor               []string                      `json:"vendor,omitempty"`
    AdditionalProperties map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NacRuleMatching.
// It customizes the JSON marshaling process for NacRuleMatching objects.
func (n NacRuleMatching) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(n.toMap())
}

// toMap converts the NacRuleMatching object to a map representation for JSON marshaling.
func (n NacRuleMatching) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, n.AdditionalProperties)
    if n.AuthType != nil {
        structMap["auth_type"] = n.AuthType
    }
    if n.Nactags != nil {
        structMap["nactags"] = n.Nactags
    }
    if n.PortTypes != nil {
        structMap["port_types"] = n.PortTypes
    }
    if n.SiteIds != nil {
        structMap["site_ids"] = n.SiteIds
    }
    if n.SitegroupIds != nil {
        structMap["sitegroup_ids"] = n.SitegroupIds
    }
    if n.Vendor != nil {
        structMap["vendor"] = n.Vendor
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NacRuleMatching.
// It customizes the JSON unmarshaling process for NacRuleMatching objects.
func (n *NacRuleMatching) UnmarshalJSON(input []byte) error {
    var temp nacRuleMatching
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "auth_type", "nactags", "port_types", "site_ids", "sitegroup_ids", "vendor")
    if err != nil {
    	return err
    }
    
    n.AdditionalProperties = additionalProperties
    n.AuthType = temp.AuthType
    n.Nactags = temp.Nactags
    n.PortTypes = temp.PortTypes
    n.SiteIds = temp.SiteIds
    n.SitegroupIds = temp.SitegroupIds
    n.Vendor = temp.Vendor
    return nil
}

// nacRuleMatching is a temporary struct used for validating the fields of NacRuleMatching.
type nacRuleMatching  struct {
    AuthType     *NacRuleMatchingAuthTypeEnum  `json:"auth_type,omitempty"`
    Nactags      []string                      `json:"nactags,omitempty"`
    PortTypes    []NacRuleMatchingPortTypeEnum `json:"port_types,omitempty"`
    SiteIds      []uuid.UUID                   `json:"site_ids,omitempty"`
    SitegroupIds []uuid.UUID                   `json:"sitegroup_ids,omitempty"`
    Vendor       []string                      `json:"vendor,omitempty"`
}

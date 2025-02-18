package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// SuppressedAlarmApplies represents a SuppressedAlarmApplies struct.
// If `scope`==`site`. Object defines the scope (within the org e.g. whole org, and/or some site_groups, and/or some sites) for which the alarm service has to be suppressed for some `duration`
type SuppressedAlarmApplies struct {
    // UUID of the current org (if provided, the alarms will be suppressed at org level)
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    // List of UUID of the sites within the org (if provided, the alarms will be suppressed for all the mentioned sites under the org)
    SiteIds              []uuid.UUID            `json:"site_ids,omitempty"`
    // List of UUID of the site groups within the org (if provided, the alarms will be suppressed for all the sites under those site groups)
    SitegroupIds         []uuid.UUID            `json:"sitegroup_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SuppressedAlarmApplies,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SuppressedAlarmApplies) String() string {
    return fmt.Sprintf(
    	"SuppressedAlarmApplies[OrgId=%v, SiteIds=%v, SitegroupIds=%v, AdditionalProperties=%v]",
    	s.OrgId, s.SiteIds, s.SitegroupIds, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SuppressedAlarmApplies.
// It customizes the JSON marshaling process for SuppressedAlarmApplies objects.
func (s SuppressedAlarmApplies) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "org_id", "site_ids", "sitegroup_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SuppressedAlarmApplies object to a map representation for JSON marshaling.
func (s SuppressedAlarmApplies) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    if s.SiteIds != nil {
        structMap["site_ids"] = s.SiteIds
    }
    if s.SitegroupIds != nil {
        structMap["sitegroup_ids"] = s.SitegroupIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SuppressedAlarmApplies.
// It customizes the JSON unmarshaling process for SuppressedAlarmApplies objects.
func (s *SuppressedAlarmApplies) UnmarshalJSON(input []byte) error {
    var temp tempSuppressedAlarmApplies
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "org_id", "site_ids", "sitegroup_ids")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.OrgId = temp.OrgId
    s.SiteIds = temp.SiteIds
    s.SitegroupIds = temp.SitegroupIds
    return nil
}

// tempSuppressedAlarmApplies is a temporary struct used for validating the fields of SuppressedAlarmApplies.
type tempSuppressedAlarmApplies  struct {
    OrgId        *uuid.UUID  `json:"org_id,omitempty"`
    SiteIds      []uuid.UUID `json:"site_ids,omitempty"`
    SitegroupIds []uuid.UUID `json:"sitegroup_ids,omitempty"`
}

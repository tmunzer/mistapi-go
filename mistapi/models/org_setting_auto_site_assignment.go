// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingAutoSiteAssignment represents a OrgSettingAutoSiteAssignment struct.
type OrgSettingAutoSiteAssignment struct {
    Enable               *bool                                    `json:"enable,omitempty"`
    Rules                Optional[[]OrgSettingAutoAssignmentRule] `json:"rules"`
    AdditionalProperties map[string]interface{}                   `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingAutoSiteAssignment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingAutoSiteAssignment) String() string {
    return fmt.Sprintf(
    	"OrgSettingAutoSiteAssignment[Enable=%v, Rules=%v, AdditionalProperties=%v]",
    	o.Enable, o.Rules, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingAutoSiteAssignment.
// It customizes the JSON marshaling process for OrgSettingAutoSiteAssignment objects.
func (o OrgSettingAutoSiteAssignment) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "enable", "rules"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingAutoSiteAssignment object to a map representation for JSON marshaling.
func (o OrgSettingAutoSiteAssignment) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.Enable != nil {
        structMap["enable"] = o.Enable
    }
    if o.Rules.IsValueSet() {
        if o.Rules.Value() != nil {
            structMap["rules"] = o.Rules.Value()
        } else {
            structMap["rules"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingAutoSiteAssignment.
// It customizes the JSON unmarshaling process for OrgSettingAutoSiteAssignment objects.
func (o *OrgSettingAutoSiteAssignment) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingAutoSiteAssignment
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "enable", "rules")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.Enable = temp.Enable
    o.Rules = temp.Rules
    return nil
}

// tempOrgSettingAutoSiteAssignment is a temporary struct used for validating the fields of OrgSettingAutoSiteAssignment.
type tempOrgSettingAutoSiteAssignment  struct {
    Enable *bool                                    `json:"enable,omitempty"`
    Rules  Optional[[]OrgSettingAutoAssignmentRule] `json:"rules"`
}

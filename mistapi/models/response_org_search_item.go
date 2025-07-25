// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// ResponseOrgSearchItem represents a ResponseOrgSearchItem struct.
type ResponseOrgSearchItem struct {
    MspId                *uuid.UUID             `json:"msp_id,omitempty"`
    // org name
    Name                 *string                `json:"name,omitempty"`
    NumAps               *int                   `json:"num_aps,omitempty"`
    NumGateways          *int                   `json:"num_gateways,omitempty"`
    NumSites             *int                   `json:"num_sites,omitempty"`
    NumSwitches          *int                   `json:"num_switches,omitempty"`
    NumUnassignedAps     *int                   `json:"num_unassigned_aps,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    SubAnaEntitled       *int                   `json:"sub_ana_entitled,omitempty"`
    SubAnaRequired       *int                   `json:"sub_ana_required,omitempty"`
    SubAstEntitled       *int                   `json:"sub_ast_entitled,omitempty"`
    SubAstRequired       *int                   `json:"sub_ast_required,omitempty"`
    SubEngEntitled       *int                   `json:"sub_eng_entitled,omitempty"`
    SubEngRequired       *int                   `json:"sub_eng_required,omitempty"`
    SubEx12Required      *int                   `json:"sub_ex12_required,omitempty"`
    // If this org has sufficient subscription
    SubInsufficient      *bool                  `json:"sub_insufficient,omitempty"`
    SubManEntitled       *int                   `json:"sub_man_entitled,omitempty"`
    SubManRequired       *int                   `json:"sub_man_required,omitempty"`
    SubMeEntitled        *int                   `json:"sub_me_entitled,omitempty"`
    SubVnaEntitled       *int                   `json:"sub_vna_entitled,omitempty"`
    SubVnaRequired       *int                   `json:"sub_vna_required,omitempty"`
    // Epoch (seconds)
    Timestamp            *float64               `json:"timestamp,omitempty"`
    // If this org is under trial period
    TrialEnabled         *bool                  `json:"trial_enabled,omitempty"`
    // a list of types that enabled by usage
    UsageTypes           []string               `json:"usage_types,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseOrgSearchItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseOrgSearchItem) String() string {
    return fmt.Sprintf(
    	"ResponseOrgSearchItem[MspId=%v, Name=%v, NumAps=%v, NumGateways=%v, NumSites=%v, NumSwitches=%v, NumUnassignedAps=%v, OrgId=%v, SubAnaEntitled=%v, SubAnaRequired=%v, SubAstEntitled=%v, SubAstRequired=%v, SubEngEntitled=%v, SubEngRequired=%v, SubEx12Required=%v, SubInsufficient=%v, SubManEntitled=%v, SubManRequired=%v, SubMeEntitled=%v, SubVnaEntitled=%v, SubVnaRequired=%v, Timestamp=%v, TrialEnabled=%v, UsageTypes=%v, AdditionalProperties=%v]",
    	r.MspId, r.Name, r.NumAps, r.NumGateways, r.NumSites, r.NumSwitches, r.NumUnassignedAps, r.OrgId, r.SubAnaEntitled, r.SubAnaRequired, r.SubAstEntitled, r.SubAstRequired, r.SubEngEntitled, r.SubEngRequired, r.SubEx12Required, r.SubInsufficient, r.SubManEntitled, r.SubManRequired, r.SubMeEntitled, r.SubVnaEntitled, r.SubVnaRequired, r.Timestamp, r.TrialEnabled, r.UsageTypes, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseOrgSearchItem.
// It customizes the JSON marshaling process for ResponseOrgSearchItem objects.
func (r ResponseOrgSearchItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "msp_id", "name", "num_aps", "num_gateways", "num_sites", "num_switches", "num_unassigned_aps", "org_id", "sub_ana_entitled", "sub_ana_required", "sub_ast_entitled", "sub_ast_required", "sub_eng_entitled", "sub_eng_required", "sub_ex12_required", "sub_insufficient", "sub_man_entitled", "sub_man_required", "sub_me_entitled", "sub_vna_entitled", "sub_vna_required", "timestamp", "trial_enabled", "usage_types"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseOrgSearchItem object to a map representation for JSON marshaling.
func (r ResponseOrgSearchItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.MspId != nil {
        structMap["msp_id"] = r.MspId
    }
    if r.Name != nil {
        structMap["name"] = r.Name
    }
    if r.NumAps != nil {
        structMap["num_aps"] = r.NumAps
    }
    if r.NumGateways != nil {
        structMap["num_gateways"] = r.NumGateways
    }
    if r.NumSites != nil {
        structMap["num_sites"] = r.NumSites
    }
    if r.NumSwitches != nil {
        structMap["num_switches"] = r.NumSwitches
    }
    if r.NumUnassignedAps != nil {
        structMap["num_unassigned_aps"] = r.NumUnassignedAps
    }
    if r.OrgId != nil {
        structMap["org_id"] = r.OrgId
    }
    if r.SubAnaEntitled != nil {
        structMap["sub_ana_entitled"] = r.SubAnaEntitled
    }
    if r.SubAnaRequired != nil {
        structMap["sub_ana_required"] = r.SubAnaRequired
    }
    if r.SubAstEntitled != nil {
        structMap["sub_ast_entitled"] = r.SubAstEntitled
    }
    if r.SubAstRequired != nil {
        structMap["sub_ast_required"] = r.SubAstRequired
    }
    if r.SubEngEntitled != nil {
        structMap["sub_eng_entitled"] = r.SubEngEntitled
    }
    if r.SubEngRequired != nil {
        structMap["sub_eng_required"] = r.SubEngRequired
    }
    if r.SubEx12Required != nil {
        structMap["sub_ex12_required"] = r.SubEx12Required
    }
    if r.SubInsufficient != nil {
        structMap["sub_insufficient"] = r.SubInsufficient
    }
    if r.SubManEntitled != nil {
        structMap["sub_man_entitled"] = r.SubManEntitled
    }
    if r.SubManRequired != nil {
        structMap["sub_man_required"] = r.SubManRequired
    }
    if r.SubMeEntitled != nil {
        structMap["sub_me_entitled"] = r.SubMeEntitled
    }
    if r.SubVnaEntitled != nil {
        structMap["sub_vna_entitled"] = r.SubVnaEntitled
    }
    if r.SubVnaRequired != nil {
        structMap["sub_vna_required"] = r.SubVnaRequired
    }
    if r.Timestamp != nil {
        structMap["timestamp"] = r.Timestamp
    }
    if r.TrialEnabled != nil {
        structMap["trial_enabled"] = r.TrialEnabled
    }
    if r.UsageTypes != nil {
        structMap["usage_types"] = r.UsageTypes
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOrgSearchItem.
// It customizes the JSON unmarshaling process for ResponseOrgSearchItem objects.
func (r *ResponseOrgSearchItem) UnmarshalJSON(input []byte) error {
    var temp tempResponseOrgSearchItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "msp_id", "name", "num_aps", "num_gateways", "num_sites", "num_switches", "num_unassigned_aps", "org_id", "sub_ana_entitled", "sub_ana_required", "sub_ast_entitled", "sub_ast_required", "sub_eng_entitled", "sub_eng_required", "sub_ex12_required", "sub_insufficient", "sub_man_entitled", "sub_man_required", "sub_me_entitled", "sub_vna_entitled", "sub_vna_required", "timestamp", "trial_enabled", "usage_types")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.MspId = temp.MspId
    r.Name = temp.Name
    r.NumAps = temp.NumAps
    r.NumGateways = temp.NumGateways
    r.NumSites = temp.NumSites
    r.NumSwitches = temp.NumSwitches
    r.NumUnassignedAps = temp.NumUnassignedAps
    r.OrgId = temp.OrgId
    r.SubAnaEntitled = temp.SubAnaEntitled
    r.SubAnaRequired = temp.SubAnaRequired
    r.SubAstEntitled = temp.SubAstEntitled
    r.SubAstRequired = temp.SubAstRequired
    r.SubEngEntitled = temp.SubEngEntitled
    r.SubEngRequired = temp.SubEngRequired
    r.SubEx12Required = temp.SubEx12Required
    r.SubInsufficient = temp.SubInsufficient
    r.SubManEntitled = temp.SubManEntitled
    r.SubManRequired = temp.SubManRequired
    r.SubMeEntitled = temp.SubMeEntitled
    r.SubVnaEntitled = temp.SubVnaEntitled
    r.SubVnaRequired = temp.SubVnaRequired
    r.Timestamp = temp.Timestamp
    r.TrialEnabled = temp.TrialEnabled
    r.UsageTypes = temp.UsageTypes
    return nil
}

// tempResponseOrgSearchItem is a temporary struct used for validating the fields of ResponseOrgSearchItem.
type tempResponseOrgSearchItem  struct {
    MspId            *uuid.UUID `json:"msp_id,omitempty"`
    Name             *string    `json:"name,omitempty"`
    NumAps           *int       `json:"num_aps,omitempty"`
    NumGateways      *int       `json:"num_gateways,omitempty"`
    NumSites         *int       `json:"num_sites,omitempty"`
    NumSwitches      *int       `json:"num_switches,omitempty"`
    NumUnassignedAps *int       `json:"num_unassigned_aps,omitempty"`
    OrgId            *uuid.UUID `json:"org_id,omitempty"`
    SubAnaEntitled   *int       `json:"sub_ana_entitled,omitempty"`
    SubAnaRequired   *int       `json:"sub_ana_required,omitempty"`
    SubAstEntitled   *int       `json:"sub_ast_entitled,omitempty"`
    SubAstRequired   *int       `json:"sub_ast_required,omitempty"`
    SubEngEntitled   *int       `json:"sub_eng_entitled,omitempty"`
    SubEngRequired   *int       `json:"sub_eng_required,omitempty"`
    SubEx12Required  *int       `json:"sub_ex12_required,omitempty"`
    SubInsufficient  *bool      `json:"sub_insufficient,omitempty"`
    SubManEntitled   *int       `json:"sub_man_entitled,omitempty"`
    SubManRequired   *int       `json:"sub_man_required,omitempty"`
    SubMeEntitled    *int       `json:"sub_me_entitled,omitempty"`
    SubVnaEntitled   *int       `json:"sub_vna_entitled,omitempty"`
    SubVnaRequired   *int       `json:"sub_vna_required,omitempty"`
    Timestamp        *float64   `json:"timestamp,omitempty"`
    TrialEnabled     *bool      `json:"trial_enabled,omitempty"`
    UsageTypes       []string   `json:"usage_types,omitempty"`
}

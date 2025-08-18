// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

// NacRule represents a NacRule struct.
type NacRule struct {
	// enum: `allow`, `block`
	Action NacRuleActionEnum `json:"action"`
	// All optional, this goes into Access-Accept
	ApplyTags []string `json:"apply_tags,omitempty"`
	// When the object has been created, in epoch
	CreatedTime *float64 `json:"created_time,omitempty"`
	// Enabled or not
	Enabled *bool `json:"enabled,omitempty"`
	// Guest portal authorization state. enum: `authorized`, `unknown`
	GuestAuthState *NacRuleGuestAuthStateEnum `json:"guest_auth_state,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id       *uuid.UUID       `json:"id,omitempty"`
	Matching *NacRuleMatching `json:"matching,omitempty"`
	// When the object has been modified for the last time, in epoch
	ModifiedTime *float64         `json:"modified_time,omitempty"`
	Name         string           `json:"name"`
	NotMatching  *NacRuleMatching `json:"not_matching,omitempty"`
	// Order of the rule, lower value implies higher priority
	Order                *int                   `json:"order,omitempty"`
	OrgId                *uuid.UUID             `json:"org_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for NacRule,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (n NacRule) String() string {
	return fmt.Sprintf(
		"NacRule[Action=%v, ApplyTags=%v, CreatedTime=%v, Enabled=%v, GuestAuthState=%v, Id=%v, Matching=%v, ModifiedTime=%v, Name=%v, NotMatching=%v, Order=%v, OrgId=%v, AdditionalProperties=%v]",
		n.Action, n.ApplyTags, n.CreatedTime, n.Enabled, n.GuestAuthState, n.Id, n.Matching, n.ModifiedTime, n.Name, n.NotMatching, n.Order, n.OrgId, n.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for NacRule.
// It customizes the JSON marshaling process for NacRule objects.
func (n NacRule) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(n.AdditionalProperties,
		"action", "apply_tags", "created_time", "enabled", "guest_auth_state", "id", "matching", "modified_time", "name", "not_matching", "order", "org_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(n.toMap())
}

// toMap converts the NacRule object to a map representation for JSON marshaling.
func (n NacRule) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, n.AdditionalProperties)
	structMap["action"] = n.Action
	if n.ApplyTags != nil {
		structMap["apply_tags"] = n.ApplyTags
	}
	if n.CreatedTime != nil {
		structMap["created_time"] = n.CreatedTime
	}
	if n.Enabled != nil {
		structMap["enabled"] = n.Enabled
	}
	if n.GuestAuthState != nil {
		structMap["guest_auth_state"] = n.GuestAuthState
	}
	if n.Id != nil {
		structMap["id"] = n.Id
	}
	if n.Matching != nil {
		structMap["matching"] = n.Matching.toMap()
	}
	if n.ModifiedTime != nil {
		structMap["modified_time"] = n.ModifiedTime
	}
	structMap["name"] = n.Name
	if n.NotMatching != nil {
		structMap["not_matching"] = n.NotMatching.toMap()
	}
	if n.Order != nil {
		structMap["order"] = n.Order
	}
	if n.OrgId != nil {
		structMap["org_id"] = n.OrgId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NacRule.
// It customizes the JSON unmarshaling process for NacRule objects.
func (n *NacRule) UnmarshalJSON(input []byte) error {
	var temp tempNacRule
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "action", "apply_tags", "created_time", "enabled", "guest_auth_state", "id", "matching", "modified_time", "name", "not_matching", "order", "org_id")
	if err != nil {
		return err
	}
	n.AdditionalProperties = additionalProperties

	n.Action = *temp.Action
	n.ApplyTags = temp.ApplyTags
	n.CreatedTime = temp.CreatedTime
	n.Enabled = temp.Enabled
	n.GuestAuthState = temp.GuestAuthState
	n.Id = temp.Id
	n.Matching = temp.Matching
	n.ModifiedTime = temp.ModifiedTime
	n.Name = *temp.Name
	n.NotMatching = temp.NotMatching
	n.Order = temp.Order
	n.OrgId = temp.OrgId
	return nil
}

// tempNacRule is a temporary struct used for validating the fields of NacRule.
type tempNacRule struct {
	Action         *NacRuleActionEnum         `json:"action"`
	ApplyTags      []string                   `json:"apply_tags,omitempty"`
	CreatedTime    *float64                   `json:"created_time,omitempty"`
	Enabled        *bool                      `json:"enabled,omitempty"`
	GuestAuthState *NacRuleGuestAuthStateEnum `json:"guest_auth_state,omitempty"`
	Id             *uuid.UUID                 `json:"id,omitempty"`
	Matching       *NacRuleMatching           `json:"matching,omitempty"`
	ModifiedTime   *float64                   `json:"modified_time,omitempty"`
	Name           *string                    `json:"name"`
	NotMatching    *NacRuleMatching           `json:"not_matching,omitempty"`
	Order          *int                       `json:"order,omitempty"`
	OrgId          *uuid.UUID                 `json:"org_id,omitempty"`
}

func (n *tempNacRule) validate() error {
	var errs []string
	if n.Action == nil {
		errs = append(errs, "required field `action` is missing for type `nac_rule`")
	}
	if n.Name == nil {
		errs = append(errs, "required field `name` is missing for type `nac_rule`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

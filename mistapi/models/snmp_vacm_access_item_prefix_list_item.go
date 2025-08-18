// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// SnmpVacmAccessItemPrefixListItem represents a SnmpVacmAccessItemPrefixListItem struct.
type SnmpVacmAccessItemPrefixListItem struct {
	// Only required if `type`==`context_prefix`
	ContextPrefix *string `json:"context_prefix,omitempty"`
	// Refer to view name
	NotifyView *string `json:"notify_view,omitempty"`
	// Refer to view name
	ReadView *string `json:"read_view,omitempty"`
	// enum: `authentication`, `none`, `privacy`
	SecurityLevel *SnmpVacmAccessItemPrefixListItemLevelEnum `json:"security_level,omitempty"`
	// enum: `any`, `usm`, `v1`, `v2c`
	SecurityModel *SnmpVacmAccessItemPrefixListItemModelEnum `json:"security_model,omitempty"`
	// enum: `context_prefix`, `default_context_prefix`
	Type *SnmpVacmAccessItemTypeEnum `json:"type,omitempty"`
	// Refer to view name
	WriteView            *string                `json:"write_view,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SnmpVacmAccessItemPrefixListItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SnmpVacmAccessItemPrefixListItem) String() string {
	return fmt.Sprintf(
		"SnmpVacmAccessItemPrefixListItem[ContextPrefix=%v, NotifyView=%v, ReadView=%v, SecurityLevel=%v, SecurityModel=%v, Type=%v, WriteView=%v, AdditionalProperties=%v]",
		s.ContextPrefix, s.NotifyView, s.ReadView, s.SecurityLevel, s.SecurityModel, s.Type, s.WriteView, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SnmpVacmAccessItemPrefixListItem.
// It customizes the JSON marshaling process for SnmpVacmAccessItemPrefixListItem objects.
func (s SnmpVacmAccessItemPrefixListItem) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"context_prefix", "notify_view", "read_view", "security_level", "security_model", "type", "write_view"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SnmpVacmAccessItemPrefixListItem object to a map representation for JSON marshaling.
func (s SnmpVacmAccessItemPrefixListItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.ContextPrefix != nil {
		structMap["context_prefix"] = s.ContextPrefix
	}
	if s.NotifyView != nil {
		structMap["notify_view"] = s.NotifyView
	}
	if s.ReadView != nil {
		structMap["read_view"] = s.ReadView
	}
	if s.SecurityLevel != nil {
		structMap["security_level"] = s.SecurityLevel
	}
	if s.SecurityModel != nil {
		structMap["security_model"] = s.SecurityModel
	}
	if s.Type != nil {
		structMap["type"] = s.Type
	}
	if s.WriteView != nil {
		structMap["write_view"] = s.WriteView
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SnmpVacmAccessItemPrefixListItem.
// It customizes the JSON unmarshaling process for SnmpVacmAccessItemPrefixListItem objects.
func (s *SnmpVacmAccessItemPrefixListItem) UnmarshalJSON(input []byte) error {
	var temp tempSnmpVacmAccessItemPrefixListItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "context_prefix", "notify_view", "read_view", "security_level", "security_model", "type", "write_view")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.ContextPrefix = temp.ContextPrefix
	s.NotifyView = temp.NotifyView
	s.ReadView = temp.ReadView
	s.SecurityLevel = temp.SecurityLevel
	s.SecurityModel = temp.SecurityModel
	s.Type = temp.Type
	s.WriteView = temp.WriteView
	return nil
}

// tempSnmpVacmAccessItemPrefixListItem is a temporary struct used for validating the fields of SnmpVacmAccessItemPrefixListItem.
type tempSnmpVacmAccessItemPrefixListItem struct {
	ContextPrefix *string                                    `json:"context_prefix,omitempty"`
	NotifyView    *string                                    `json:"notify_view,omitempty"`
	ReadView      *string                                    `json:"read_view,omitempty"`
	SecurityLevel *SnmpVacmAccessItemPrefixListItemLevelEnum `json:"security_level,omitempty"`
	SecurityModel *SnmpVacmAccessItemPrefixListItemModelEnum `json:"security_model,omitempty"`
	Type          *SnmpVacmAccessItemTypeEnum                `json:"type,omitempty"`
	WriteView     *string                                    `json:"write_view,omitempty"`
}

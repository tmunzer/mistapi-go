// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// AamwProfileCategory represents a AamwProfileCategory struct.
type AamwProfileCategory struct {
	// enum: `archive`, `document`, `pdf`, `executable`, `rich_application`, `library`, `os_package`, `mobile`, `java`, `configuration`, `script`
	Category             *AamwProfileCategoryCategoryEnum `json:"category,omitempty"`
	HashLookupOnly       *bool                            `json:"hash_lookup_only,omitempty"`
	AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for AamwProfileCategory,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AamwProfileCategory) String() string {
	return fmt.Sprintf(
		"AamwProfileCategory[Category=%v, HashLookupOnly=%v, AdditionalProperties=%v]",
		a.Category, a.HashLookupOnly, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AamwProfileCategory.
// It customizes the JSON marshaling process for AamwProfileCategory objects.
func (a AamwProfileCategory) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"category", "hash_lookup_only"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the AamwProfileCategory object to a map representation for JSON marshaling.
func (a AamwProfileCategory) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.Category != nil {
		structMap["category"] = a.Category
	}
	if a.HashLookupOnly != nil {
		structMap["hash_lookup_only"] = a.HashLookupOnly
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AamwProfileCategory.
// It customizes the JSON unmarshaling process for AamwProfileCategory objects.
func (a *AamwProfileCategory) UnmarshalJSON(input []byte) error {
	var temp tempAamwProfileCategory
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "category", "hash_lookup_only")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.Category = temp.Category
	a.HashLookupOnly = temp.HashLookupOnly
	return nil
}

// tempAamwProfileCategory is a temporary struct used for validating the fields of AamwProfileCategory.
type tempAamwProfileCategory struct {
	Category       *AamwProfileCategoryCategoryEnum `json:"category,omitempty"`
	HashLookupOnly *bool                            `json:"hash_lookup_only,omitempty"`
}

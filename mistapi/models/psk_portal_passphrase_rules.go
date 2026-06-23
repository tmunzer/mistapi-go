// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// PskPortalPassphraseRules represents a PskPortalPassphraseRules struct.
// Passphrase generation rules for PSKs created through a portal
type PskPortalPassphraseRules struct {
	// Whether generated passphrases may include alphabetic characters
	AlphabetsEnabled *bool `json:"alphabets_enabled,omitempty"`
	// Fixed generated passphrase length used when min and max length are not both valid
	Length *int `json:"length,omitempty"`
	// Maximum generated passphrase length when paired with a valid `min_length`. If `max_length` or `min_length` is invalid, the portal uses `length`; if `length` is unset or invalid, it uses 8. Valid values are integers from 8 through 63, and `max_length` must be greater than `min_length`
	MaxLength *int `json:"max_length,omitempty"`
	// Minimum generated passphrase length when paired with a valid `max_length`. If `max_length` or `min_length` is invalid, the portal uses `length`; if `length` is unset or invalid, it uses 8. Valid values are integers from 8 through 63, and `max_length` must be greater than `min_length`
	MinLength *int `json:"min_length,omitempty"`
	// Whether generated passphrases may include numeric characters
	NumericsEnabled *bool `json:"numerics_enabled,omitempty"`
	// Allowed symbol characters for generated passphrases
	Symbols *string `json:"symbols,omitempty"`
	// Whether generated passphrases may include symbols
	SymbolsEnabled       *bool                  `json:"symbols_enabled,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PskPortalPassphraseRules,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PskPortalPassphraseRules) String() string {
	return fmt.Sprintf(
		"PskPortalPassphraseRules[AlphabetsEnabled=%v, Length=%v, MaxLength=%v, MinLength=%v, NumericsEnabled=%v, Symbols=%v, SymbolsEnabled=%v, AdditionalProperties=%v]",
		p.AlphabetsEnabled, p.Length, p.MaxLength, p.MinLength, p.NumericsEnabled, p.Symbols, p.SymbolsEnabled, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PskPortalPassphraseRules.
// It customizes the JSON marshaling process for PskPortalPassphraseRules objects.
func (p PskPortalPassphraseRules) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(p.AdditionalProperties,
		"alphabets_enabled", "length", "max_length", "min_length", "numerics_enabled", "symbols", "symbols_enabled"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(p.toMap())
}

// toMap converts the PskPortalPassphraseRules object to a map representation for JSON marshaling.
func (p PskPortalPassphraseRules) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, p.AdditionalProperties)
	if p.AlphabetsEnabled != nil {
		structMap["alphabets_enabled"] = p.AlphabetsEnabled
	}
	if p.Length != nil {
		structMap["length"] = p.Length
	}
	if p.MaxLength != nil {
		structMap["max_length"] = p.MaxLength
	}
	if p.MinLength != nil {
		structMap["min_length"] = p.MinLength
	}
	if p.NumericsEnabled != nil {
		structMap["numerics_enabled"] = p.NumericsEnabled
	}
	if p.Symbols != nil {
		structMap["symbols"] = p.Symbols
	}
	if p.SymbolsEnabled != nil {
		structMap["symbols_enabled"] = p.SymbolsEnabled
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for PskPortalPassphraseRules.
// It customizes the JSON unmarshaling process for PskPortalPassphraseRules objects.
func (p *PskPortalPassphraseRules) UnmarshalJSON(input []byte) error {
	var temp tempPskPortalPassphraseRules
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "alphabets_enabled", "length", "max_length", "min_length", "numerics_enabled", "symbols", "symbols_enabled")
	if err != nil {
		return err
	}
	p.AdditionalProperties = additionalProperties

	p.AlphabetsEnabled = temp.AlphabetsEnabled
	p.Length = temp.Length
	p.MaxLength = temp.MaxLength
	p.MinLength = temp.MinLength
	p.NumericsEnabled = temp.NumericsEnabled
	p.Symbols = temp.Symbols
	p.SymbolsEnabled = temp.SymbolsEnabled
	return nil
}

// tempPskPortalPassphraseRules is a temporary struct used for validating the fields of PskPortalPassphraseRules.
type tempPskPortalPassphraseRules struct {
	AlphabetsEnabled *bool   `json:"alphabets_enabled,omitempty"`
	Length           *int    `json:"length,omitempty"`
	MaxLength        *int    `json:"max_length,omitempty"`
	MinLength        *int    `json:"min_length,omitempty"`
	NumericsEnabled  *bool   `json:"numerics_enabled,omitempty"`
	Symbols          *string `json:"symbols,omitempty"`
	SymbolsEnabled   *bool   `json:"symbols_enabled,omitempty"`
}

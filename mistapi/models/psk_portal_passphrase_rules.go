// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// PskPortalPassphraseRules represents a PskPortalPassphraseRules struct.
type PskPortalPassphraseRules struct {
    AlphabertsEnabled    *bool                  `json:"alphaberts_enabled,omitempty"`
    Length               *int                   `json:"length,omitempty"`
    // For valid `max_length` and `min_length`, passphrase size is set randomly from that range.
    // - if `max_length` and/or `min_length` are invalid, passphrase size is equal to `length` parameter
    // - if `length` is not set or is invalid, default passphrase size is 8.
    // - valid `max_length`, `min_length`, `length` should be an integer between 8 to 63. Also, `max_length` > `min_length`
    MaxLength            *int                   `json:"max_length,omitempty"`
    // Ror valid `max_length` and `min_length`, passphrase size is set randomly from that range.
    // - if `max_length` and/or `min_length` are invalid, passphrase size is equal to `length` parameter
    // - if `length` is not set or is invalid, default passphrase size is 8.
    // - valid `max_length`, `min_length`, `length` should be an integer between 8 to 63. Also, `max_length` > `min_length`
    MinLength            *int                   `json:"min_length,omitempty"`
    NumericsEnabled      *bool                  `json:"numerics_enabled,omitempty"`
    Symbols              *string                `json:"symbols,omitempty"`
    SymbolsEnabled       *bool                  `json:"symbols_enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for PskPortalPassphraseRules,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (p PskPortalPassphraseRules) String() string {
    return fmt.Sprintf(
    	"PskPortalPassphraseRules[AlphabertsEnabled=%v, Length=%v, MaxLength=%v, MinLength=%v, NumericsEnabled=%v, Symbols=%v, SymbolsEnabled=%v, AdditionalProperties=%v]",
    	p.AlphabertsEnabled, p.Length, p.MaxLength, p.MinLength, p.NumericsEnabled, p.Symbols, p.SymbolsEnabled, p.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for PskPortalPassphraseRules.
// It customizes the JSON marshaling process for PskPortalPassphraseRules objects.
func (p PskPortalPassphraseRules) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(p.AdditionalProperties,
        "alphaberts_enabled", "length", "max_length", "min_length", "numerics_enabled", "symbols", "symbols_enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(p.toMap())
}

// toMap converts the PskPortalPassphraseRules object to a map representation for JSON marshaling.
func (p PskPortalPassphraseRules) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, p.AdditionalProperties)
    if p.AlphabertsEnabled != nil {
        structMap["alphaberts_enabled"] = p.AlphabertsEnabled
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "alphaberts_enabled", "length", "max_length", "min_length", "numerics_enabled", "symbols", "symbols_enabled")
    if err != nil {
    	return err
    }
    p.AdditionalProperties = additionalProperties
    
    p.AlphabertsEnabled = temp.AlphabertsEnabled
    p.Length = temp.Length
    p.MaxLength = temp.MaxLength
    p.MinLength = temp.MinLength
    p.NumericsEnabled = temp.NumericsEnabled
    p.Symbols = temp.Symbols
    p.SymbolsEnabled = temp.SymbolsEnabled
    return nil
}

// tempPskPortalPassphraseRules is a temporary struct used for validating the fields of PskPortalPassphraseRules.
type tempPskPortalPassphraseRules  struct {
    AlphabertsEnabled *bool   `json:"alphaberts_enabled,omitempty"`
    Length            *int    `json:"length,omitempty"`
    MaxLength         *int    `json:"max_length,omitempty"`
    MinLength         *int    `json:"min_length,omitempty"`
    NumericsEnabled   *bool   `json:"numerics_enabled,omitempty"`
    Symbols           *string `json:"symbols,omitempty"`
    SymbolsEnabled    *bool   `json:"symbols_enabled,omitempty"`
}

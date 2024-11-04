package models

import (
    "encoding/json"
)

// AccountJuniperInfo represents a AccountJuniperInfo struct.
type AccountJuniperInfo struct {
    Accounts             []JuniperAccount `json:"accounts,omitempty"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AccountJuniperInfo.
// It customizes the JSON marshaling process for AccountJuniperInfo objects.
func (a AccountJuniperInfo) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AccountJuniperInfo object to a map representation for JSON marshaling.
func (a AccountJuniperInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Accounts != nil {
        structMap["accounts"] = a.Accounts
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountJuniperInfo.
// It customizes the JSON unmarshaling process for AccountJuniperInfo objects.
func (a *AccountJuniperInfo) UnmarshalJSON(input []byte) error {
    var temp tempAccountJuniperInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "accounts")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Accounts = temp.Accounts
    return nil
}

// tempAccountJuniperInfo is a temporary struct used for validating the fields of AccountJuniperInfo.
type tempAccountJuniperInfo  struct {
    Accounts []JuniperAccount `json:"accounts,omitempty"`
}

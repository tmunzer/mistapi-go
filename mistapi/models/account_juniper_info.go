package models

import (
    "encoding/json"
    "fmt"
)

// AccountJuniperInfo represents a AccountJuniperInfo struct.
type AccountJuniperInfo struct {
    Accounts             []JuniperAccount       `json:"accounts,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AccountJuniperInfo,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AccountJuniperInfo) String() string {
    return fmt.Sprintf(
    	"AccountJuniperInfo[Accounts=%v, AdditionalProperties=%v]",
    	a.Accounts, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AccountJuniperInfo.
// It customizes the JSON marshaling process for AccountJuniperInfo objects.
func (a AccountJuniperInfo) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "accounts"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountJuniperInfo object to a map representation for JSON marshaling.
func (a AccountJuniperInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "accounts")
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

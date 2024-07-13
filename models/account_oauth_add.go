package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// AccountOauthAdd represents a AccountOauthAdd struct.
type AccountOauthAdd struct {
    value                 any
    isAccountJamfConfig   bool
    isAccountVmwareConfig bool
}

// String converts the AccountOauthAdd object to a string representation.
func (a AccountOauthAdd) String() string {
    if bytes, err := json.Marshal(a.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for AccountOauthAdd.
// It customizes the JSON marshaling process for AccountOauthAdd objects.
func (a AccountOauthAdd) MarshalJSON() (
    []byte,
    error) {
    if a.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.AccountOauthAddContainer.From*` functions to initialize the AccountOauthAdd object.")
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountOauthAdd object to a map representation for JSON marshaling.
func (a *AccountOauthAdd) toMap() any {
    switch obj := a.value.(type) {
    case *AccountJamfConfig:
        return obj.toMap()
    case *AccountVmwareConfig:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountOauthAdd.
// It customizes the JSON unmarshaling process for AccountOauthAdd objects.
func (a *AccountOauthAdd) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&AccountJamfConfig{}, false, &a.isAccountJamfConfig),
        NewTypeHolder(&AccountVmwareConfig{}, false, &a.isAccountVmwareConfig),
    )
    
    a.value = result
    return err
}

func (a *AccountOauthAdd) AsAccountJamfConfig() (
    *AccountJamfConfig,
    bool) {
    if !a.isAccountJamfConfig {
        return nil, false
    }
    return a.value.(*AccountJamfConfig), true
}

func (a *AccountOauthAdd) AsAccountVmwareConfig() (
    *AccountVmwareConfig,
    bool) {
    if !a.isAccountVmwareConfig {
        return nil, false
    }
    return a.value.(*AccountVmwareConfig), true
}

// internalAccountOauthAdd represents a accountOauthAdd struct.
type internalAccountOauthAdd struct {}

var AccountOauthAddContainer internalAccountOauthAdd

// The internalAccountOauthAdd instance, wrapping the provided AccountJamfConfig value.
func (a *internalAccountOauthAdd) FromAccountJamfConfig(val AccountJamfConfig) AccountOauthAdd {
    return AccountOauthAdd{value: &val}
}

// The internalAccountOauthAdd instance, wrapping the provided AccountVmwareConfig value.
func (a *internalAccountOauthAdd) FromAccountVmwareConfig(val AccountVmwareConfig) AccountOauthAdd {
    return AccountOauthAdd{value: &val}
}
package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// AccountOauthAdd represents a AccountOauthAdd struct.
type AccountOauthAdd struct {
    value                      any
    isAccountJamfConfig        bool
    isAccountVmwareConfig      bool
    isAccountMobicontrolConfig bool
    isAccountZdxConfig         bool
    isAccountCrowdstrikeConfig bool
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
    case *AccountMobicontrolConfig:
        return obj.toMap()
    case *AccountZdxConfig:
        return obj.toMap()
    case *AccountCrowdstrikeConfig:
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
        NewTypeHolder(&AccountMobicontrolConfig{}, false, &a.isAccountMobicontrolConfig),
        NewTypeHolder(&AccountZdxConfig{}, false, &a.isAccountZdxConfig),
        NewTypeHolder(&AccountCrowdstrikeConfig{}, false, &a.isAccountCrowdstrikeConfig),
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

func (a *AccountOauthAdd) AsAccountMobicontrolConfig() (
    *AccountMobicontrolConfig,
    bool) {
    if !a.isAccountMobicontrolConfig {
        return nil, false
    }
    return a.value.(*AccountMobicontrolConfig), true
}

func (a *AccountOauthAdd) AsAccountZdxConfig() (
    *AccountZdxConfig,
    bool) {
    if !a.isAccountZdxConfig {
        return nil, false
    }
    return a.value.(*AccountZdxConfig), true
}

func (a *AccountOauthAdd) AsAccountCrowdstrikeConfig() (
    *AccountCrowdstrikeConfig,
    bool) {
    if !a.isAccountCrowdstrikeConfig {
        return nil, false
    }
    return a.value.(*AccountCrowdstrikeConfig), true
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

// The internalAccountOauthAdd instance, wrapping the provided AccountMobicontrolConfig value.
func (a *internalAccountOauthAdd) FromAccountMobicontrolConfig(val AccountMobicontrolConfig) AccountOauthAdd {
    return AccountOauthAdd{value: &val}
}

// The internalAccountOauthAdd instance, wrapping the provided AccountZdxConfig value.
func (a *internalAccountOauthAdd) FromAccountZdxConfig(val AccountZdxConfig) AccountOauthAdd {
    return AccountOauthAdd{value: &val}
}

// The internalAccountOauthAdd instance, wrapping the provided AccountCrowdstrikeConfig value.
func (a *internalAccountOauthAdd) FromAccountCrowdstrikeConfig(val AccountCrowdstrikeConfig) AccountOauthAdd {
    return AccountOauthAdd{value: &val}
}

package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseOauthAppLinkAccounts represents a ResponseOauthAppLinkAccounts struct.
// This is Array of a container for any-of cases.
type ResponseOauthAppLinkAccounts struct {
    value               any
    isAccountJamfInfo   bool
    isAccountOauthInfo  bool
    isAccountVmwareInfo bool
}

// String converts the ResponseOauthAppLinkAccounts object to a string representation.
func (r ResponseOauthAppLinkAccounts) String() string {
    if bytes, err := json.Marshal(r.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ResponseOauthAppLinkAccounts.
// It customizes the JSON marshaling process for ResponseOauthAppLinkAccounts objects.
func (r ResponseOauthAppLinkAccounts) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ResponseOauthAppLinkAccountsContainer.From*` functions to initialize the ResponseOauthAppLinkAccounts object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseOauthAppLinkAccounts object to a map representation for JSON marshaling.
func (r *ResponseOauthAppLinkAccounts) toMap() any {
    switch obj := r.value.(type) {
    case *AccountJamfInfo:
        return obj.toMap()
    case *AccountOauthInfo:
        return obj.toMap()
    case *AccountVmwareInfo:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOauthAppLinkAccounts.
// It customizes the JSON unmarshaling process for ResponseOauthAppLinkAccounts objects.
func (r *ResponseOauthAppLinkAccounts) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&AccountJamfInfo{}, false, &r.isAccountJamfInfo),
        NewTypeHolder(&AccountOauthInfo{}, false, &r.isAccountOauthInfo),
        NewTypeHolder(&AccountVmwareInfo{}, false, &r.isAccountVmwareInfo),
    )
    
    r.value = result
    return err
}

func (r *ResponseOauthAppLinkAccounts) AsAccountJamfInfo() (
    *AccountJamfInfo,
    bool) {
    if !r.isAccountJamfInfo {
        return nil, false
    }
    return r.value.(*AccountJamfInfo), true
}

func (r *ResponseOauthAppLinkAccounts) AsAccountOauthInfo() (
    *AccountOauthInfo,
    bool) {
    if !r.isAccountOauthInfo {
        return nil, false
    }
    return r.value.(*AccountOauthInfo), true
}

func (r *ResponseOauthAppLinkAccounts) AsAccountVmwareInfo() (
    *AccountVmwareInfo,
    bool) {
    if !r.isAccountVmwareInfo {
        return nil, false
    }
    return r.value.(*AccountVmwareInfo), true
}

// internalResponseOauthAppLinkAccounts represents a responseOauthAppLinkAccounts struct.
// This is Array of a container for any-of cases.
type internalResponseOauthAppLinkAccounts struct {}

var ResponseOauthAppLinkAccountsContainer internalResponseOauthAppLinkAccounts

// The internalResponseOauthAppLinkAccounts instance, wrapping the provided AccountJamfInfo value.
func (r *internalResponseOauthAppLinkAccounts) FromAccountJamfInfo(val AccountJamfInfo) ResponseOauthAppLinkAccounts {
    return ResponseOauthAppLinkAccounts{value: &val}
}

// The internalResponseOauthAppLinkAccounts instance, wrapping the provided AccountOauthInfo value.
func (r *internalResponseOauthAppLinkAccounts) FromAccountOauthInfo(val AccountOauthInfo) ResponseOauthAppLinkAccounts {
    return ResponseOauthAppLinkAccounts{value: &val}
}

// The internalResponseOauthAppLinkAccounts instance, wrapping the provided AccountVmwareInfo value.
func (r *internalResponseOauthAppLinkAccounts) FromAccountVmwareInfo(val AccountVmwareInfo) ResponseOauthAppLinkAccounts {
    return ResponseOauthAppLinkAccounts{value: &val}
}

package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseOauthAppLinkItem2 represents a ResponseOauthAppLinkItem2 struct.
type ResponseOauthAppLinkItem2 struct {
    value                    any
    isAccountJamfInfo        bool
    isAccountOauthInfo       bool
    isAccountVmwareInfo      bool
    isAccountMobicontrolInfo bool
}

// String converts the ResponseOauthAppLinkItem2 object to a string representation.
func (r ResponseOauthAppLinkItem2) String() string {
    if bytes, err := json.Marshal(r.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ResponseOauthAppLinkItem2.
// It customizes the JSON marshaling process for ResponseOauthAppLinkItem2 objects.
func (r ResponseOauthAppLinkItem2) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ResponseOauthAppLinkItemContainer.From*` functions to initialize the ResponseOauthAppLinkItem2 object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseOauthAppLinkItem2 object to a map representation for JSON marshaling.
func (r *ResponseOauthAppLinkItem2) toMap() any {
    switch obj := r.value.(type) {
    case *AccountJamfInfo:
        return obj.toMap()
    case *AccountOauthInfo:
        return obj.toMap()
    case *AccountVmwareInfo:
        return obj.toMap()
    case *AccountMobicontrolInfo:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOauthAppLinkItem2.
// It customizes the JSON unmarshaling process for ResponseOauthAppLinkItem2 objects.
func (r *ResponseOauthAppLinkItem2) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallAnyOf(input,
        NewTypeHolder(&AccountJamfInfo{}, false, &r.isAccountJamfInfo),
        NewTypeHolder(&AccountOauthInfo{}, false, &r.isAccountOauthInfo),
        NewTypeHolder(&AccountVmwareInfo{}, false, &r.isAccountVmwareInfo),
        NewTypeHolder(&AccountMobicontrolInfo{}, false, &r.isAccountMobicontrolInfo),
    )
    
    r.value = result
    return err
}

func (r *ResponseOauthAppLinkItem2) AsAccountJamfInfo() (
    *AccountJamfInfo,
    bool) {
    if !r.isAccountJamfInfo {
        return nil, false
    }
    return r.value.(*AccountJamfInfo), true
}

func (r *ResponseOauthAppLinkItem2) AsAccountOauthInfo() (
    *AccountOauthInfo,
    bool) {
    if !r.isAccountOauthInfo {
        return nil, false
    }
    return r.value.(*AccountOauthInfo), true
}

func (r *ResponseOauthAppLinkItem2) AsAccountVmwareInfo() (
    *AccountVmwareInfo,
    bool) {
    if !r.isAccountVmwareInfo {
        return nil, false
    }
    return r.value.(*AccountVmwareInfo), true
}

func (r *ResponseOauthAppLinkItem2) AsAccountMobicontrolInfo() (
    *AccountMobicontrolInfo,
    bool) {
    if !r.isAccountMobicontrolInfo {
        return nil, false
    }
    return r.value.(*AccountMobicontrolInfo), true
}

// internalResponseOauthAppLinkItem2 represents a responseOauthAppLinkItem2 struct.
type internalResponseOauthAppLinkItem2 struct {}

var ResponseOauthAppLinkItemContainer internalResponseOauthAppLinkItem2

// The internalResponseOauthAppLinkItem2 instance, wrapping the provided AccountJamfInfo value.
func (r *internalResponseOauthAppLinkItem2) FromAccountJamfInfo(val AccountJamfInfo) ResponseOauthAppLinkItem2 {
    return ResponseOauthAppLinkItem2{value: &val}
}

// The internalResponseOauthAppLinkItem2 instance, wrapping the provided AccountOauthInfo value.
func (r *internalResponseOauthAppLinkItem2) FromAccountOauthInfo(val AccountOauthInfo) ResponseOauthAppLinkItem2 {
    return ResponseOauthAppLinkItem2{value: &val}
}

// The internalResponseOauthAppLinkItem2 instance, wrapping the provided AccountVmwareInfo value.
func (r *internalResponseOauthAppLinkItem2) FromAccountVmwareInfo(val AccountVmwareInfo) ResponseOauthAppLinkItem2 {
    return ResponseOauthAppLinkItem2{value: &val}
}

// The internalResponseOauthAppLinkItem2 instance, wrapping the provided AccountMobicontrolInfo value.
func (r *internalResponseOauthAppLinkItem2) FromAccountMobicontrolInfo(val AccountMobicontrolInfo) ResponseOauthAppLinkItem2 {
    return ResponseOauthAppLinkItem2{value: &val}
}

package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseOauthAppLink represents a ResponseOauthAppLink struct.
type ResponseOauthAppLink struct {
    value              any
    isAccountOauthInfo bool
    isAccountZdxInfo   bool
}

// String converts the ResponseOauthAppLink object to a string representation.
func (r ResponseOauthAppLink) String() string {
    if bytes, err := json.Marshal(r.value); err == nil {
         return strings.Trim(string(bytes), "\"")
    }
    return ""
}

// MarshalJSON implements the json.Marshaler interface for ResponseOauthAppLink.
// It customizes the JSON marshaling process for ResponseOauthAppLink objects.
func (r ResponseOauthAppLink) MarshalJSON() (
    []byte,
    error) {
    if r.value == nil {
        return nil, errors.New("No underlying type is set. Please use any of the `models.ResponseOauthAppLinkContainer.From*` functions to initialize the ResponseOauthAppLink object.")
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseOauthAppLink object to a map representation for JSON marshaling.
func (r *ResponseOauthAppLink) toMap() any {
    switch obj := r.value.(type) {
    case *AccountOauthInfo:
        return obj.toMap()
    case *AccountZdxInfo:
        return obj.toMap()
    }
    return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOauthAppLink.
// It customizes the JSON unmarshaling process for ResponseOauthAppLink objects.
func (r *ResponseOauthAppLink) UnmarshalJSON(input []byte) error {
    result, err := UnmarshallOneOf(input,
        NewTypeHolder(&AccountOauthInfo{}, false, &r.isAccountOauthInfo),
        NewTypeHolder(&AccountZdxInfo{}, false, &r.isAccountZdxInfo),
    )
    
    r.value = result
    return err
}

func (r *ResponseOauthAppLink) AsAccountOauthInfo() (
    *AccountOauthInfo,
    bool) {
    if !r.isAccountOauthInfo {
        return nil, false
    }
    return r.value.(*AccountOauthInfo), true
}

func (r *ResponseOauthAppLink) AsAccountZdxInfo() (
    *AccountZdxInfo,
    bool) {
    if !r.isAccountZdxInfo {
        return nil, false
    }
    return r.value.(*AccountZdxInfo), true
}

// internalResponseOauthAppLink represents a responseOauthAppLink struct.
type internalResponseOauthAppLink struct {}

var ResponseOauthAppLinkContainer internalResponseOauthAppLink

// The internalResponseOauthAppLink instance, wrapping the provided AccountOauthInfo value.
func (r *internalResponseOauthAppLink) FromAccountOauthInfo(val AccountOauthInfo) ResponseOauthAppLink {
    return ResponseOauthAppLink{value: &val}
}

// The internalResponseOauthAppLink instance, wrapping the provided AccountZdxInfo value.
func (r *internalResponseOauthAppLink) FromAccountZdxInfo(val AccountZdxInfo) ResponseOauthAppLink {
    return ResponseOauthAppLink{value: &val}
}

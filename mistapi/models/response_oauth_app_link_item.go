package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ResponseOauthAppLinkItem represents a ResponseOauthAppLinkItem struct.
type ResponseOauthAppLinkItem struct {
	value                    any
	isAccountJamfInfo        bool
	isAccountOauthInfo       bool
	isAccountVmwareInfo      bool
	isAccountMobicontrolInfo bool
}

// String converts the ResponseOauthAppLinkItem object to a string representation.
func (r ResponseOauthAppLinkItem) String() string {
	if bytes, err := json.Marshal(r.value); err == nil {
		return strings.Trim(string(bytes), "\"")
	}
	return ""
}

// MarshalJSON implements the json.Marshaler interface for ResponseOauthAppLinkItem.
// It customizes the JSON marshaling process for ResponseOauthAppLinkItem objects.
func (r ResponseOauthAppLinkItem) MarshalJSON() (
	[]byte,
	error) {
	if r.value == nil {
		return nil, errors.New("No underlying type is set. Please use any of the `models.ResponseOauthAppLinkItemContainer.From*` functions to initialize the ResponseOauthAppLinkItem object.")
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseOauthAppLinkItem object to a map representation for JSON marshaling.
func (r *ResponseOauthAppLinkItem) toMap() any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOauthAppLinkItem.
// It customizes the JSON unmarshaling process for ResponseOauthAppLinkItem objects.
func (r *ResponseOauthAppLinkItem) UnmarshalJSON(input []byte) error {
	result, err := UnmarshallOneOf(input,
		NewTypeHolder(&AccountJamfInfo{}, false, &r.isAccountJamfInfo),
		NewTypeHolder(&AccountOauthInfo{}, false, &r.isAccountOauthInfo),
		NewTypeHolder(&AccountVmwareInfo{}, false, &r.isAccountVmwareInfo),
		NewTypeHolder(&AccountMobicontrolInfo{}, false, &r.isAccountMobicontrolInfo),
	)

	r.value = result
	return err
}

func (r *ResponseOauthAppLinkItem) AsAccountJamfInfo() (
	*AccountJamfInfo,
	bool) {
	if !r.isAccountJamfInfo {
		return nil, false
	}
	return r.value.(*AccountJamfInfo), true
}

func (r *ResponseOauthAppLinkItem) AsAccountOauthInfo() (
	*AccountOauthInfo,
	bool) {
	if !r.isAccountOauthInfo {
		return nil, false
	}
	return r.value.(*AccountOauthInfo), true
}

func (r *ResponseOauthAppLinkItem) AsAccountVmwareInfo() (
	*AccountVmwareInfo,
	bool) {
	if !r.isAccountVmwareInfo {
		return nil, false
	}
	return r.value.(*AccountVmwareInfo), true
}

func (r *ResponseOauthAppLinkItem) AsAccountMobicontrolInfo() (
	*AccountMobicontrolInfo,
	bool) {
	if !r.isAccountMobicontrolInfo {
		return nil, false
	}
	return r.value.(*AccountMobicontrolInfo), true
}

// internalResponseOauthAppLinkItem represents a responseOauthAppLinkItem struct.
type internalResponseOauthAppLinkItem struct{}

var ResponseOauthAppLinkItemContainer internalResponseOauthAppLinkItem

// The internalResponseOauthAppLinkItem instance, wrapping the provided AccountJamfInfo value.
func (r *internalResponseOauthAppLinkItem) FromAccountJamfInfo(val AccountJamfInfo) ResponseOauthAppLinkItem {
	return ResponseOauthAppLinkItem{value: &val}
}

// The internalResponseOauthAppLinkItem instance, wrapping the provided AccountOauthInfo value.
func (r *internalResponseOauthAppLinkItem) FromAccountOauthInfo(val AccountOauthInfo) ResponseOauthAppLinkItem {
	return ResponseOauthAppLinkItem{value: &val}
}

// The internalResponseOauthAppLinkItem instance, wrapping the provided AccountVmwareInfo value.
func (r *internalResponseOauthAppLinkItem) FromAccountVmwareInfo(val AccountVmwareInfo) ResponseOauthAppLinkItem {
	return ResponseOauthAppLinkItem{value: &val}
}

// The internalResponseOauthAppLinkItem instance, wrapping the provided AccountMobicontrolInfo value.
func (r *internalResponseOauthAppLinkItem) FromAccountMobicontrolInfo(val AccountMobicontrolInfo) ResponseOauthAppLinkItem {
	return ResponseOauthAppLinkItem{value: &val}
}

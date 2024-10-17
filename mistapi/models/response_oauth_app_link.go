package models

import (
	"encoding/json"
)

// ResponseOauthAppLink represents a ResponseOauthAppLink struct.
type ResponseOauthAppLink struct {
	// List of linked account details
	Accounts []ResponseOauthAppLinkItem `json:"accounts,omitempty"`
	// Basic Auth application linked status in mist portal enabled for VMware
	Linked               *bool          `json:"linked,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseOauthAppLink.
// It customizes the JSON marshaling process for ResponseOauthAppLink objects.
func (r ResponseOauthAppLink) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseOauthAppLink object to a map representation for JSON marshaling.
func (r ResponseOauthAppLink) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Accounts != nil {
		structMap["accounts"] = r.Accounts
	}
	if r.Linked != nil {
		structMap["linked"] = r.Linked
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOauthAppLink.
// It customizes the JSON unmarshaling process for ResponseOauthAppLink objects.
func (r *ResponseOauthAppLink) UnmarshalJSON(input []byte) error {
	var temp tempResponseOauthAppLink
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "accounts", "linked")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.Accounts = temp.Accounts
	r.Linked = temp.Linked
	return nil
}

// tempResponseOauthAppLink is a temporary struct used for validating the fields of ResponseOauthAppLink.
type tempResponseOauthAppLink struct {
	Accounts []ResponseOauthAppLinkItem `json:"accounts,omitempty"`
	Linked   *bool                      `json:"linked,omitempty"`
}

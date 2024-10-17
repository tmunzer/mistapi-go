package models

import (
	"encoding/json"
)

// ResponseLogout represents a ResponseLogout struct.
type ResponseLogout struct {
	// if configured in SSO as custom_logout_url
	ForwardUrl           *string        `json:"forward_url,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseLogout.
// It customizes the JSON marshaling process for ResponseLogout objects.
func (r ResponseLogout) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseLogout object to a map representation for JSON marshaling.
func (r ResponseLogout) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.ForwardUrl != nil {
		structMap["forward_url"] = r.ForwardUrl
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseLogout.
// It customizes the JSON unmarshaling process for ResponseLogout objects.
func (r *ResponseLogout) UnmarshalJSON(input []byte) error {
	var temp tempResponseLogout
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "forward_url")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.ForwardUrl = temp.ForwardUrl
	return nil
}

// tempResponseLogout is a temporary struct used for validating the fields of ResponseLogout.
type tempResponseLogout struct {
	ForwardUrl *string `json:"forward_url,omitempty"`
}

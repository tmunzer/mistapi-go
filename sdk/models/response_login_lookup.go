package models

import (
    "encoding/json"
)

// ResponseLoginLookup represents a ResponseLoginLookup struct.
type ResponseLoginLookup struct {
    SsoUrl               *string        `json:"sso_url,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseLoginLookup.
// It customizes the JSON marshaling process for ResponseLoginLookup objects.
func (r ResponseLoginLookup) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseLoginLookup object to a map representation for JSON marshaling.
func (r ResponseLoginLookup) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.SsoUrl != nil {
        structMap["sso_url"] = r.SsoUrl
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseLoginLookup.
// It customizes the JSON unmarshaling process for ResponseLoginLookup objects.
func (r *ResponseLoginLookup) UnmarshalJSON(input []byte) error {
    var temp responseLoginLookup
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "sso_url")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.SsoUrl = temp.SsoUrl
    return nil
}

// responseLoginLookup is a temporary struct used for validating the fields of ResponseLoginLookup.
type responseLoginLookup  struct {
    SsoUrl *string `json:"sso_url,omitempty"`
}

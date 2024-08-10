package models

import (
    "encoding/json"
)

// ResponseTwoFactorJson represents a ResponseTwoFactorJson struct.
type ResponseTwoFactorJson struct {
    TwoFactorSecret      *string        `json:"two_factor_secret,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseTwoFactorJson.
// It customizes the JSON marshaling process for ResponseTwoFactorJson objects.
func (r ResponseTwoFactorJson) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseTwoFactorJson object to a map representation for JSON marshaling.
func (r ResponseTwoFactorJson) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.TwoFactorSecret != nil {
        structMap["two_factor_secret"] = r.TwoFactorSecret
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseTwoFactorJson.
// It customizes the JSON unmarshaling process for ResponseTwoFactorJson objects.
func (r *ResponseTwoFactorJson) UnmarshalJSON(input []byte) error {
    var temp tempResponseTwoFactorJson
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "two_factor_secret")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.TwoFactorSecret = temp.TwoFactorSecret
    return nil
}

// tempResponseTwoFactorJson is a temporary struct used for validating the fields of ResponseTwoFactorJson.
type tempResponseTwoFactorJson  struct {
    TwoFactorSecret *string `json:"two_factor_secret,omitempty"`
}

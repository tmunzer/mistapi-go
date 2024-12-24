package models

import (
    "encoding/json"
    "fmt"
)

// ResponseTwoFactorJson represents a ResponseTwoFactorJson struct.
type ResponseTwoFactorJson struct {
    TwoFactorSecret      *string                `json:"two_factor_secret,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseTwoFactorJson,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseTwoFactorJson) String() string {
    return fmt.Sprintf(
    	"ResponseTwoFactorJson[TwoFactorSecret=%v, AdditionalProperties=%v]",
    	r.TwoFactorSecret, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseTwoFactorJson.
// It customizes the JSON marshaling process for ResponseTwoFactorJson objects.
func (r ResponseTwoFactorJson) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "two_factor_secret"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseTwoFactorJson object to a map representation for JSON marshaling.
func (r ResponseTwoFactorJson) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "two_factor_secret")
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

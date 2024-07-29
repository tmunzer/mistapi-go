package models

import (
    "encoding/json"
)

// Recaptcha represents a Recaptcha struct.
type Recaptcha struct {
    // flavor of the captcha. enum: `google`, `hcaptcha`
    Flavor               *RecaptchaFlavorEnum `json:"flavor,omitempty"`
    Required             *bool                `json:"required,omitempty"`
    Sitekey              *string              `json:"sitekey,omitempty"`
    AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Recaptcha.
// It customizes the JSON marshaling process for Recaptcha objects.
func (r Recaptcha) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the Recaptcha object to a map representation for JSON marshaling.
func (r Recaptcha) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Flavor != nil {
        structMap["flavor"] = r.Flavor
    }
    if r.Required != nil {
        structMap["required"] = r.Required
    }
    if r.Sitekey != nil {
        structMap["sitekey"] = r.Sitekey
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Recaptcha.
// It customizes the JSON unmarshaling process for Recaptcha objects.
func (r *Recaptcha) UnmarshalJSON(input []byte) error {
    var temp recaptcha
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "flavor", "required", "sitekey")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Flavor = temp.Flavor
    r.Required = temp.Required
    r.Sitekey = temp.Sitekey
    return nil
}

// recaptcha is a temporary struct used for validating the fields of Recaptcha.
type recaptcha  struct {
    Flavor   *RecaptchaFlavorEnum `json:"flavor,omitempty"`
    Required *bool                `json:"required,omitempty"`
    Sitekey  *string              `json:"sitekey,omitempty"`
}

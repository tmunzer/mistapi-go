package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// Recover represents a Recover struct.
type Recover struct {
    Email                string               `json:"email"`
    // see https://www.google.com/recaptcha/
    Recaptcha            *string              `json:"recaptcha,omitempty"`
    // flavor of the captcha. enum: `google`, `hcaptcha`
    RecaptchaFlavor      *RecaptchaFlavorEnum `json:"recaptcha_flavor,omitempty"`
    AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Recover.
// It customizes the JSON marshaling process for Recover objects.
func (r Recover) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the Recover object to a map representation for JSON marshaling.
func (r Recover) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["email"] = r.Email
    if r.Recaptcha != nil {
        structMap["recaptcha"] = r.Recaptcha
    }
    if r.RecaptchaFlavor != nil {
        structMap["recaptcha_flavor"] = r.RecaptchaFlavor
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Recover.
// It customizes the JSON unmarshaling process for Recover objects.
func (r *Recover) UnmarshalJSON(input []byte) error {
    var temp recover
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "email", "recaptcha", "recaptcha_flavor")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Email = *temp.Email
    r.Recaptcha = temp.Recaptcha
    r.RecaptchaFlavor = temp.RecaptchaFlavor
    return nil
}

// recover is a temporary struct used for validating the fields of Recover.
type recover  struct {
    Email           *string              `json:"email"`
    Recaptcha       *string              `json:"recaptcha,omitempty"`
    RecaptchaFlavor *RecaptchaFlavorEnum `json:"recaptcha_flavor,omitempty"`
}

func (r *recover) validate() error {
    var errs []string
    if r.Email == nil {
        errs = append(errs, "required field `email` is missing for type `Recover`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}

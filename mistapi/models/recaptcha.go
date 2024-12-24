package models

import (
    "encoding/json"
    "fmt"
)

// Recaptcha represents a Recaptcha struct.
type Recaptcha struct {
    // flavor of the captcha. enum: `google`, `hcaptcha`
    Flavor               *RecaptchaFlavorEnum   `json:"flavor,omitempty"`
    Required             *bool                  `json:"required,omitempty"`
    Sitekey              *string                `json:"sitekey,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Recaptcha,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r Recaptcha) String() string {
    return fmt.Sprintf(
    	"Recaptcha[Flavor=%v, Required=%v, Sitekey=%v, AdditionalProperties=%v]",
    	r.Flavor, r.Required, r.Sitekey, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Recaptcha.
// It customizes the JSON marshaling process for Recaptcha objects.
func (r Recaptcha) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "flavor", "required", "sitekey"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the Recaptcha object to a map representation for JSON marshaling.
func (r Recaptcha) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
    var temp tempRecaptcha
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "flavor", "required", "sitekey")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Flavor = temp.Flavor
    r.Required = temp.Required
    r.Sitekey = temp.Sitekey
    return nil
}

// tempRecaptcha is a temporary struct used for validating the fields of Recaptcha.
type tempRecaptcha  struct {
    Flavor   *RecaptchaFlavorEnum `json:"flavor,omitempty"`
    Required *bool                `json:"required,omitempty"`
    Sitekey  *string              `json:"sitekey,omitempty"`
}

package models

import (
    "encoding/json"
    "fmt"
)

// ResponseVerifyTokenSuccess represents a ResponseVerifyTokenSuccess struct.
type ResponseVerifyTokenSuccess struct {
    Detail               *string                `json:"detail,omitempty"`
    InviteNotApplied     *bool                  `json:"invite_not_applied,omitempty"`
    MinLength            *int                   `json:"min_length,omitempty"`
    ReturnTo             *string                `json:"return_to,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseVerifyTokenSuccess,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseVerifyTokenSuccess) String() string {
    return fmt.Sprintf(
    	"ResponseVerifyTokenSuccess[Detail=%v, InviteNotApplied=%v, MinLength=%v, ReturnTo=%v, AdditionalProperties=%v]",
    	r.Detail, r.InviteNotApplied, r.MinLength, r.ReturnTo, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseVerifyTokenSuccess.
// It customizes the JSON marshaling process for ResponseVerifyTokenSuccess objects.
func (r ResponseVerifyTokenSuccess) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "detail", "invite_not_applied", "min_length", "return_to"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseVerifyTokenSuccess object to a map representation for JSON marshaling.
func (r ResponseVerifyTokenSuccess) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Detail != nil {
        structMap["detail"] = r.Detail
    }
    if r.InviteNotApplied != nil {
        structMap["invite_not_applied"] = r.InviteNotApplied
    }
    if r.MinLength != nil {
        structMap["min_length"] = r.MinLength
    }
    if r.ReturnTo != nil {
        structMap["return_to"] = r.ReturnTo
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseVerifyTokenSuccess.
// It customizes the JSON unmarshaling process for ResponseVerifyTokenSuccess objects.
func (r *ResponseVerifyTokenSuccess) UnmarshalJSON(input []byte) error {
    var temp tempResponseVerifyTokenSuccess
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "detail", "invite_not_applied", "min_length", "return_to")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Detail = temp.Detail
    r.InviteNotApplied = temp.InviteNotApplied
    r.MinLength = temp.MinLength
    r.ReturnTo = temp.ReturnTo
    return nil
}

// tempResponseVerifyTokenSuccess is a temporary struct used for validating the fields of ResponseVerifyTokenSuccess.
type tempResponseVerifyTokenSuccess  struct {
    Detail           *string `json:"detail,omitempty"`
    InviteNotApplied *bool   `json:"invite_not_applied,omitempty"`
    MinLength        *int    `json:"min_length,omitempty"`
    ReturnTo         *string `json:"return_to,omitempty"`
}

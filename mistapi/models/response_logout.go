// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ResponseLogout represents a ResponseLogout struct.
type ResponseLogout struct {
    // If configured in SSO as custom_logout_url
    ForwardUrl           *string                `json:"forward_url,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseLogout,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseLogout) String() string {
    return fmt.Sprintf(
    	"ResponseLogout[ForwardUrl=%v, AdditionalProperties=%v]",
    	r.ForwardUrl, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseLogout.
// It customizes the JSON marshaling process for ResponseLogout objects.
func (r ResponseLogout) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "forward_url"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseLogout object to a map representation for JSON marshaling.
func (r ResponseLogout) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "forward_url")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.ForwardUrl = temp.ForwardUrl
    return nil
}

// tempResponseLogout is a temporary struct used for validating the fields of ResponseLogout.
type tempResponseLogout  struct {
    ForwardUrl *string `json:"forward_url,omitempty"`
}

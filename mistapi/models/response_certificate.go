// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseCertificate represents a ResponseCertificate struct.
// If the current Org CA certificate is set to expire within 30 days, a pending certificate will be returned along with the expected auto-renewal timestamp.
type ResponseCertificate struct {
    Cert                 string                 `json:"cert"`
    PendingCert          *string                `json:"pending_cert,omitempty"`
    PendingCertExpiry    *int                   `json:"pending_cert_expiry,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseCertificate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseCertificate) String() string {
    return fmt.Sprintf(
    	"ResponseCertificate[Cert=%v, PendingCert=%v, PendingCertExpiry=%v, AdditionalProperties=%v]",
    	r.Cert, r.PendingCert, r.PendingCertExpiry, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseCertificate.
// It customizes the JSON marshaling process for ResponseCertificate objects.
func (r ResponseCertificate) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "cert", "pending_cert", "pending_cert_expiry"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseCertificate object to a map representation for JSON marshaling.
func (r ResponseCertificate) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["cert"] = r.Cert
    if r.PendingCert != nil {
        structMap["pending_cert"] = r.PendingCert
    }
    if r.PendingCertExpiry != nil {
        structMap["pending_cert_expiry"] = r.PendingCertExpiry
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseCertificate.
// It customizes the JSON unmarshaling process for ResponseCertificate objects.
func (r *ResponseCertificate) UnmarshalJSON(input []byte) error {
    var temp tempResponseCertificate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cert", "pending_cert", "pending_cert_expiry")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Cert = *temp.Cert
    r.PendingCert = temp.PendingCert
    r.PendingCertExpiry = temp.PendingCertExpiry
    return nil
}

// tempResponseCertificate is a temporary struct used for validating the fields of ResponseCertificate.
type tempResponseCertificate  struct {
    Cert              *string `json:"cert"`
    PendingCert       *string `json:"pending_cert,omitempty"`
    PendingCertExpiry *int    `json:"pending_cert_expiry,omitempty"`
}

func (r *tempResponseCertificate) validate() error {
    var errs []string
    if r.Cert == nil {
        errs = append(errs, "required field `cert` is missing for type `response_certificate`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}

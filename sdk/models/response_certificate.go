package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseCertificate represents a ResponseCertificate struct.
type ResponseCertificate struct {
    Cert                 string         `json:"cert"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseCertificate.
// It customizes the JSON marshaling process for ResponseCertificate objects.
func (r ResponseCertificate) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseCertificate object to a map representation for JSON marshaling.
func (r ResponseCertificate) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["cert"] = r.Cert
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseCertificate.
// It customizes the JSON unmarshaling process for ResponseCertificate objects.
func (r *ResponseCertificate) UnmarshalJSON(input []byte) error {
    var temp responseCertificate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "cert")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Cert = *temp.Cert
    return nil
}

// responseCertificate is a temporary struct used for validating the fields of ResponseCertificate.
type responseCertificate  struct {
    Cert *string `json:"cert"`
}

func (r *responseCertificate) validate() error {
    var errs []string
    if r.Cert == nil {
        errs = append(errs, "required field `cert` is missing for type `Response_Certificate`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
